package controllers

import (
	"fmt"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/appscode/go/log"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/dynamic/dynamiclister"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog"
	"kmodules.xyz/client-go/tools/queue"
)

const controllerAgentName = "kfc"

var SecretKey string

const (
	// SuccessSynced is used as part of the Event 'reason' when a Resource is synced
	SuccessSynced = "Synced"

	// MessageResourceSynced is the message used for an Event fired when a Resource
	// is synced successfully
	MessageResourceSynced = "Resource synced successfully"
)

// Controller is the controller implementation for KubeForm resources
type Controller struct {
	sync.Mutex

	kubeclientset kubernetes.Interface
	dynamicclient dynamic.Interface

	crdListers map[schema.GroupVersionResource]dynamiclister.Lister
	crdWorkers map[schema.GroupVersionResource]*queue.Worker
	syncedFns  map[schema.GroupVersionResource]cache.InformerSynced

	// recorder is an event recorder for recording Event resources to the
	// Kubernetes API.
	recorder record.EventRecorder
}

// NewController returns a new sample controller
func NewController(
	kubeclientset kubernetes.Interface,
	dynamicclient dynamic.Interface) *Controller {
	// Create event broadcaster
	// Add sample-controller types to the default Kubernetes Scheme so Events can be
	// logged for sample-controller types.
	//utilruntime.Must(kubeformscheme.AddToScheme(scheme.Scheme))
	klog.V(4).Info("Creating event broadcaster")
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(klog.Infof)
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{Interface: kubeclientset.CoreV1().Events("")})
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: controllerAgentName})

	controller := &Controller{
		kubeclientset: kubeclientset,
		dynamicclient: dynamicclient,
		crdListers:    make(map[schema.GroupVersionResource]dynamiclister.Lister),
		crdWorkers:    make(map[schema.GroupVersionResource]*queue.Worker),
		syncedFns:     make(map[schema.GroupVersionResource]cache.InformerSynced),
		recorder:      recorder,
	}
	klog.Info("Setting up event handlers")

	return controller
}

func (c *Controller) AddNewCRD(gvr schema.GroupVersionResource, dynamicClient dynamic.Interface, stopCh <-chan struct{}) error {
	factory := dynamicinformer.NewDynamicSharedInformerFactory(dynamicClient, time.Second*30)

	maxNumRequeues := 5
	numThreads := 5
	i := factory.ForResource(gvr)
	q := queue.New(gvr.String(), maxNumRequeues, numThreads, func(key string) error {
		return c.reconcile(gvr, key)
	})

	i.Informer().AddEventHandler(queue.DefaultEventHandler(q.GetQueue()))

	c.crdListers[gvr] = dynamiclister.New(i.Informer().GetIndexer(), gvr)
	c.crdWorkers[gvr] = q
	c.syncedFns[gvr] = i.Informer().HasSynced

	factory.Start(stopCh)

	klog.Info("Waiting for informer caches to sync")
	if ok := cache.WaitForCacheSync(stopCh, c.syncedFns[gvr]); !ok {
		return fmt.Errorf("failed to wait for caches to sync")
	}

	klog.Info("Starting worker")

	c.crdWorkers[gvr].Run(stopCh)

	return nil
}

func (c *Controller) GetWorker(gvr schema.GroupVersionResource) *queue.Worker {
	c.Lock()
	defer c.Unlock()
	return c.crdWorkers[gvr]
}

// Run will set up the event handlers for types we are interested in, as well
// as syncing informer caches and starting workers. It will block until stopCh
// is closed, at which point it will shutdown the workqueue and wait for
// workers to finish processing their current work items.
func (c *Controller) Run(stopCh <-chan struct{}) error {
	defer utilruntime.HandleCrash()

	klog.Info("Starting KubeForm controller")

	<-stopCh
	klog.Info("Shutting down workers")

	return nil
}

func (c *Controller) Lister(gvr schema.GroupVersionResource) dynamiclister.Lister {
	c.Lock()
	defer c.Unlock()
	return c.crdListers[gvr]
}

// reconcile compares the actual state with the desired, and attempts to
// converge the two. It then updates the Status block of the KubeForm resource
// with the current status of the resource.
func (c *Controller) reconcile(gvr schema.GroupVersionResource, key string) error {
	// Convert the namespace/name string into a distinct namespace and name
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("invalid resource key: %s", key))
		return nil
	}

	// Get the resource with this namespace/name
	obj, err := c.Lister(gvr).Namespace(namespace).Get(name)
	if err != nil {
		// The resource may no longer exist, in which case we stop
		// processing.
		if errors.IsNotFound(err) {
			utilruntime.HandleError(fmt.Errorf("resource '%s' in work queue no longer exists", key))
			return nil
		}

		return err
	}

	// TODO: make a namer
	resPath := filepath.Join(basePath, gvr.Resource+"."+namespace+"."+name)
	providerFile := filepath.Join(resPath, "provider.tf.json")
	mainFile := filepath.Join(resPath, "main.tf.json")
	stateFile := filepath.Join(resPath, "terraform.tfstate")
	outputFile := filepath.Join(resPath, "output.tf")

	if hasFinalizer(obj.GetFinalizers(), KFCFinalizer) {
		if obj.GetDeletionTimestamp() != nil {
			err := terraformDestroy(resPath)
			if err != nil {
				log.Error("failed to terraform destroy: ", err)
			}

			err = deleteFiles(resPath)
			if err != nil {
				log.Error("failed to delete files: ", err)
			}

			err = removeFinalizer(obj, KFCFinalizer)
			if err != nil {
				log.Error("failed to remove finalizer: ", err)
			}

			c.updateResource(gvr, obj)
			return nil
		}
	} else {
		err := addFinalizer(obj, KFCFinalizer)
		if err != nil {
			return fmt.Errorf("failed to add finalizer : %s", err)
		}
	}

	err = createFiles(resPath, providerFile, mainFile)
	if err != nil {
		return fmt.Errorf("failed to create tf files : %s", err)
	}

	providerRef, _, err := unstructured.NestedString(obj.Object, "spec", "providerRef", "name")
	if err != nil {
		log.Error(err, "failed to get providerRef")
		return nil
	}

	secret, err := c.kubeclientset.CoreV1().Secrets(namespace).Get(providerRef, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("unable to fetch secret : %s", err)
	}

	isModule := isModule(gvr.Group)

	var providerName, source string
	if isModule {
		providerName, source = getModuleProviderAndSource(obj.GetKind())
	} else {
		providerName = strings.Split(gvr.Group, ".")[0]
	}

	err = secretToTFProvider(secret, providerName, providerFile)
	if err != nil {
		return fmt.Errorf("unable to create provider from secret : %s", err)
	}

	if isModule {
		err = crdToModule(c.kubeclientset, gvr.GroupVersion(), obj, source, mainFile, outputFile)
		if err != nil {
			return fmt.Errorf("unable to get crd module : %s", err)
		}
	} else {
		err = crdToTFResource(gvr.GroupVersion(), namespace, providerName, c.kubeclientset, obj, mainFile)
		if err != nil {
			return fmt.Errorf("unable to get crd resource : %s", err)
		}
	}

	err = terraformInit(resPath)
	if err != nil {
		return fmt.Errorf("unable to initialize terraform : %s", err)
	}

	err = createTFState(c.kubeclientset, stateFile, providerName, isModule, gvr.GroupVersion(), obj)
	if err != nil {
		return fmt.Errorf("unable to create tfstate file : %s", err)
	}

	err = terraformApply(resPath)
	if err != nil {
		return fmt.Errorf("unable to apply terraform : %s", err)
	}

	err = updateTFStateFile(stateFile, isModule, gvr.GroupVersion(), obj)
	if err != nil {

		return fmt.Errorf("unable to update TFState : %s", err)
	}

	if !isModule {
		err = updateStateField(c.kubeclientset, namespace, providerName, stateFile, gvr, obj)
		if err != nil {
			return fmt.Errorf("unable to update resource fields from tfstate : %s", err)
		}
	} else {
		err = updateOutputField(c.kubeclientset, resPath, namespace, providerName, obj)
		if err != nil {
			return fmt.Errorf("unable to update output tfstate : %s", err)
		}
	}

	c.updateResource(gvr, obj)

	c.recorder.Event(obj, corev1.EventTypeNormal, SuccessSynced, MessageResourceSynced)
	return nil
}

func (c *Controller) updateResource(gvr schema.GroupVersionResource, u *unstructured.Unstructured) {
	_, err := c.dynamicclient.Resource(gvr).Namespace(u.GetNamespace()).Update(u, metav1.UpdateOptions{})
	if err != nil {
		klog.Errorf("failed to update resource, reason : %s", err.Error())
	}
}
