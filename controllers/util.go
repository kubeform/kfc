/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"k8s.io/client-go/kubernetes"

	"k8s.io/apimachinery/pkg/runtime"

	"k8s.io/klog"

	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/fatih/structs"
	"kmodules.xyz/client-go/meta"

	jsoniter "github.com/json-iterator/go"

	"github.com/gobuffalo/flect"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

const KFCFinalizer = "kfc.io"

var (
	homePath = os.Getenv("HOME")
	basePath = filepath.Join(homePath, ".kfc")
)

func secretToTFProvider(secret *corev1.Secret, providerName, providerFile string) error {
	d1 := []byte(`{ "provider": { "` + providerName + `":`)
	providerJson, err := json.Marshal(secret.Data)
	if err != nil {
		return err
	}
	d1 = append(d1, providerJson...)
	d1 = append(d1, []byte("} }")...)

	prettyData, err := prettyJSON(d1)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(providerFile, prettyData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func crdToTFResource(gv schema.GroupVersion, kind, namespace, providerName string, kubeclient kubernetes.Interface, obj *unstructured.Unstructured, mainFile string) error {
	resourceName := providerName + "_" + flect.Underscore(kind)

	data, err := meta.MarshalToJson(obj, gv)
	if err != nil {
		klog.Error(err)
	}
	typedObj, err := meta.UnmarshalFromJSON(data, gv)
	if err != nil {
		klog.Error(err)
	}

	typedStruct := structs.New(typedObj)
	spec := typedStruct.Field("Spec")
	specValue := spec.Value()
	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
	}.Froze()

	tfStr, err := jsonit.Marshal(specValue)
	if err != nil {
		klog.Error(err)
	}

	var u1 map[string]interface{}
	err = json.Unmarshal(tfStr, &u1)
	if err != nil {
		return err
	}

	fields := typedStruct.Field("Spec").Fields()

	var ok bool
	for _, field := range fields {
		if field.Name() == "Secret" {
			ok = true
			break
		}
	}

	if ok {
		secretName := typedStruct.Field("Spec").Field("Secret").Field("Name").Value()
		if secretName != nil {
			secret, err := kubeclient.CoreV1().Secrets(namespace).Get(secretName.(string), v1.GetOptions{})
			if err != nil {
				return err
			}

			for key := range secret.Data {
				value := secret.Data[key]

				var tempMap = make(map[string]interface{}, 0)
				filedName := strings.Split(key, ".")

				buffer := new(bytes.Buffer)
				if err := json.Compact(buffer, value); err != nil {
					d := strings.ReplaceAll(string(value), "\n", "")
					err = unstructured.SetNestedField(u1, d, filedName...)
					if err != nil {
						return err
					}
				} else {
					err = json.Unmarshal(buffer.Bytes(), &tempMap)
					if err != nil {
						return err
					}

					err = unstructured.SetNestedMap(u1, tempMap, filedName...)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	str, err := json.Marshal(u1)
	if err != nil {
		return err
	}

	d1 := []byte(`{"resource":{ "` + resourceName + `":{"` + obj.GetName() + `":`)

	d1 = append(d1, str...)
	d1 = append(d1, []byte("} } }")...)
	prettyData, err := prettyJSON(d1)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(mainFile, prettyData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func hasFinalizer(finalizers []string, finalizer string) bool {
	for _, f := range finalizers {
		if f == finalizer {
			return true
		}
	}

	return false
}

func addFinalizer(u *unstructured.Unstructured, finalizer string) error {
	finalizers := u.GetFinalizers()
	for _, v := range finalizers {
		if v == finalizer {
			return nil
		}
	}

	finalizers = append(finalizers, finalizer)
	err := unstructured.SetNestedStringSlice(u.Object, finalizers, "metadata", "finalizers")
	if err != nil {
		return err
	}

	return nil
}

func removeFinalizer(u *unstructured.Unstructured, finalizer string) error {
	finalizers := u.GetFinalizers()
	for i, v := range finalizers {
		if v == finalizer {
			finalizers = append(finalizers[:i], finalizers[i+1:]...)
			break
		}
	}

	err := unstructured.SetNestedStringSlice(u.Object, finalizers, "metadata", "finalizers")
	if err != nil {
		return err
	}

	return nil
}

func createFiles(resPath, providerFile, mainFile string) error {
	_, err := os.Stat(resPath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(resPath, 0777)
		if err != nil {
			return err
		}
		_, err = os.Create(providerFile)
		if err != nil {
			return err
		}

		_, err = os.Create(mainFile)
		if err != nil {
			return err
		}
	}

	return nil
}

func deleteFiles(resPath string) error {
	err := os.RemoveAll(resPath)
	if err != nil {
		return err
	}

	return nil
}

func prettyJSON(byteJson []byte) ([]byte, error) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, byteJson, "", "  ")
	if err != nil {
		return nil, err
	}

	return prettyJSON.Bytes(), err
}

func createTFState(filePath string, gv schema.GroupVersion, u *unstructured.Unstructured) {
	_, existErr := os.Stat(filePath)

	data, err := meta.MarshalToJson(u, gv)
	if err != nil {
		klog.Error(err)
	}

	typedObj, err := meta.UnmarshalFromJSON(data, gv)
	if err != nil {
		klog.Error(err)
	}

	typedStruct := structs.New(typedObj)
	spec := typedStruct.Field("Status").Field("TFState")
	value := spec.Value()

	if os.IsNotExist(existErr) && value.(*runtime.RawExtension) != nil {
		err = ioutil.WriteFile(filePath, value.(*runtime.RawExtension).Raw, 0644)
		if err != nil {
			klog.Errorf("failed to write file hash : %s", err.Error())
		}
	}
}

func updateTFState(filePath string, gv schema.GroupVersion, u *unstructured.Unstructured) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		klog.Errorf("failed to read tfstate file : %s", err.Error())
	}

	jsonData, err := meta.MarshalToJson(u, gv)
	if err != nil {
		klog.Error(err)
	}
	typedObj, err := meta.UnmarshalFromJSON(jsonData, gv)
	if err != nil {
		klog.Error(err)
	}

	typedStruct := structs.New(typedObj)
	spec := typedStruct.Field("Status").Field("TFState")
	value := spec.Value()

	rawData := &runtime.RawExtension{
		Raw: data,
	}

	if value.(*runtime.RawExtension) == nil {
		err = setNestedFieldNoCopy(u.Object, rawData, "status", "tfState")
		if err != nil {
			klog.Errorf("failed to update tfstate : %s", err.Error())
		}
		return
	}

	if bytes.Compare(data, value.(*runtime.RawExtension).Raw) != 0 {
		err = setNestedFieldNoCopy(u.Object, rawData, "status", "tfState")
		if err != nil {
			klog.Errorf("failed to update tfstate : %s", err.Error())
		}
	}
}

func setNestedFieldNoCopy(obj map[string]interface{}, value interface{}, fields ...string) error {
	m := obj

	for i, field := range fields[:len(fields)-1] {
		if val, ok := m[field]; ok {
			if valMap, ok := val.(map[string]interface{}); ok {
				m = valMap
			} else {
				return fmt.Errorf("value cannot be set because %v is not a map[string]interface{}", jsonPath(fields[:i+1]))
			}
		} else {
			newVal := make(map[string]interface{})
			m[field] = newVal
			m = newVal
		}
	}
	m[fields[len(fields)-1]] = value
	return nil
}

func jsonPath(fields []string) string {
	return "." + strings.Join(fields, ".")
}

//var sensitiveData = map[string][]string{
//	"linode_instance": {"Spec.RootPass", ""},
//}
