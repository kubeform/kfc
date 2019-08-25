package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	"github.com/appscode/go/log"
	"github.com/fatih/structs"
	"github.com/gobuffalo/flect"
	jsoniter "github.com/json-iterator/go"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
	"kmodules.xyz/client-go/meta"
	"kubeform.dev/kubeform/apis"
)

const KFCFinalizer = "kfc.io"

var (
	basePath = filepath.Join("/tmp", ".kfc")
)

func secretToTFProvider(secret *corev1.Secret, providerName, providerFile string) error {
	d1 := []byte(`{ "provider": { "` + providerName + `":`)
	tempData := make(map[string]string, 0)
	for key, val := range secret.Data {
		tempData[key] = strings.ReplaceAll(string(val), "\n", "")
	}
	providerJson, err := json.Marshal(tempData)
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

func crdToTFResource(gv schema.GroupVersion, namespace, providerName string, kubeclient kubernetes.Interface, obj *unstructured.Unstructured, mainFile string) error {
	resourceName := providerName + "_" + flect.Underscore(obj.GetKind())

	data, err := meta.MarshalToJson(obj, gv)
	if err != nil {
		return err
	}

	typedObj, err := meta.UnmarshalFromJSON(data, gv)
	if err != nil {
		return err
	}

	typedStruct := structs.New(typedObj)
	spec := reflect.ValueOf(typedStruct.Field("Spec").Value())
	specType := reflect.TypeOf(typedStruct.Field("Spec").Value())
	specValue := reflect.New(specType)
	specValue.Elem().Set(spec)
	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
	}.Froze()

	secretRef, _, err := unstructured.NestedFieldNoCopy(obj.Object, "spec", "secretRef")
	if err != nil {
		return err
	}

	if secretRef != nil {
		secretName := typedStruct.Field("Spec").Field("SecretRef").Field("Name").Value()
		if secretName != nil {
			secret, err := kubeclient.CoreV1().Secrets(namespace).Get(secretName.(string), v1.GetOptions{})
			if err != nil {
				return err
			}

			for key := range secret.Data {
				if strings.Contains(key, "out.") {
					continue
				}
				value := secret.Data[key]

				fieldName := strings.Split(key, ".")
				var tempMap = make(map[string]string, 0)
				buffer := new(bytes.Buffer)
				var secretData interface{}

				if err := json.Compact(buffer, value); err != nil {
					secretData = strings.ReplaceAll(string(value), "\n", "")
				} else {
					err = json.Unmarshal(buffer.Bytes(), &tempMap)
					if err != nil {
						return err
					}
					secretData = tempMap
				}

				field := specValue.Elem()
				for _, f := range fieldName {
					if index, err := strconv.Atoi(f); err == nil {
						field = field.Index(index)
						continue
					}
					field = reflect.Indirect(field).FieldByName(flect.Capitalize(flect.Camelize(f)))
				}
				field.Set(reflect.ValueOf(secretData))
			}
		}
	}

	str, err := jsonit.Marshal(specValue.Interface())
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

func updateStateField(kc kubernetes.Interface, namespace, providerName, filePath string, gvr schema.GroupVersionResource, obj *unstructured.Unstructured) error {
	gv := gvr.GroupVersion()
	stateJson, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	var state state
	err = json.Unmarshal(stateJson, &state)
	if err != nil {
		return err
	}

	data, err := meta.MarshalToJson(obj, gv)
	if err != nil {
		return err
	}

	typedObj, err := meta.UnmarshalFromJSON(data, gv)
	if err != nil {
		return err
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
	}.Froze()

	var raw []byte
	jsonByte, err := state.Resources[0].Instances[0].AttributesRaw.MarshalJSON()
	if err != nil {
		return err
	}

	raw = append(raw, []byte(`{"spec":`)...)
	raw = append(raw, jsonByte...)
	raw = append(raw, []byte(`}`)...)

	err = jsonit.Unmarshal(raw, &typedObj)
	if err != nil {
		return err
	}

	s := structs.New(typedObj)
	secretData := make(map[string]string, 0)
	processSensitiveFields(reflect.TypeOf(s.Field("Spec").Value()), reflect.ValueOf(s.Field("Spec").Value()), "", &secretData)

	if len(secretData) != 0 {
		var secretName string

		secretRef, _, err := unstructured.NestedFieldNoCopy(obj.Object, "spec", "secretRef")
		if err != nil {
			return err
		}

		if secretRef != nil {
			secretName = s.Field("Spec").Field("SecretRef").Field("Name").Value().(string)
		} else {
			secretName = obj.GetName() + "-" + obj.GetNamespace() + "-" + "sensitive"
		}

		var secret *corev1.Secret
		secret, err = kc.CoreV1().Secrets(namespace).Get(secretName, v1.GetOptions{})
		if err != nil {
			if errors.ReasonForError(err) == v1.StatusReasonNotFound {
				secret, err = kc.CoreV1().Secrets(namespace).Create(&corev1.Secret{
					ObjectMeta: v1.ObjectMeta{
						Name:      secretName,
						Namespace: namespace,
					},
					Type: corev1.SecretType("kfc.io/" + providerName),
				})
				if err != nil {
					return err
				}
			}
			return err
		}
		if secret.Data == nil {
			secret.Data = make(map[string][]byte, 0)
		}

		for key := range secretData {
			secret.Data["out."+key] = []byte(secretData[key])
		}

		_, err = kc.CoreV1().Secrets(namespace).Update(secret)
		if err != nil {
			return err
		}
	}

	err = setNestedFieldNoCopy(obj.Object, s.Field("Spec").Value(), "status", "output")
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

func createTFState(kc kubernetes.Interface, filePath, providerName string, gv schema.GroupVersion, u *unstructured.Unstructured) error {
	resourceName := providerName + "_" + flect.Underscore(u.GetKind())
	_, existErr := os.Stat(filePath)

	data, err := meta.MarshalToJson(u, gv)
	if err != nil {
		return err
	}

	typedObj, err := meta.UnmarshalFromJSON(data, gv)
	if err != nil {
		return err
	}

	typedStruct := structs.New(typedObj)
	stateValue := typedStruct.Field("Status").Field("State").Value()
	outputValue := reflect.ValueOf(typedStruct.Field("Status").Field("Output").Value())

	if os.IsNotExist(existErr) && stateValue.(*apis.State) != nil {
		stateData, err := json.Marshal(stateValue)
		if err != nil {
			return err
		}

		jsonit := jsoniter.Config{
			EscapeHTML:             true,
			SortMapKeys:            true,
			ValidateJsonRawMessage: true,
			TagKey:                 "tf",
		}.Froze()

		secretRef, _, err := unstructured.NestedFieldNoCopy(u.Object, "spec", "secretRef")
		if err != nil {
			return err
		}

		if secretRef != nil {
			secretName := typedStruct.Field("Spec").Field("SecretRef").Field("Name").Value()
			if secretName != nil {
				secret, err := kc.CoreV1().Secrets(u.GetNamespace()).Get(secretName.(string), v1.GetOptions{})
				if err != nil {
					return err
				}

				for key := range secret.Data {
					if !strings.Contains(key, "out.") {
						continue
					}

					value := secret.Data[key]

					var secretData interface{}
					tempMap := make(map[string]string, 0)
					buffer := new(bytes.Buffer)

					if err := json.Compact(buffer, value); err != nil {
						secretData = strings.ReplaceAll(string(value), "\n", "")
					} else {
						err = json.Unmarshal(buffer.Bytes(), &tempMap)
						if err != nil {
							return err
						}
						secretData = tempMap
					}

					key = strings.ReplaceAll(key, "out.", "")
					fieldsName := strings.Split(key, ".")

					field := outputValue.Elem()
					for _, fieldName := range fieldsName {
						if index, err := strconv.Atoi(fieldName); err == nil {
							field = field.Index(index)
							continue
						}
						if field.Kind() == reflect.Ptr {
							field = field.Elem()
						}
						field = field.FieldByName(flect.Capitalize(flect.Camelize(fieldName)))
					}

					field.Set(reflect.ValueOf(secretData))
				}
			}
		}

		outputData, err := jsonit.Marshal(outputValue.Interface())
		if err != nil {
			return err
		}

		var tfstate state
		err = json.Unmarshal(stateData, &tfstate)
		if err != nil {
			return err
		}

		tfstate.Resources = []resource{
			{
				Mode:           "managed",
				Type:           resourceName,
				Name:           u.GetName(),
				ProviderConfig: "provider." + providerName,
				Instances: []instance{
					{
						SchemaVersion: 0,
						AttributesRaw: outputData,
					},
				},
			},
		}
		tfstateData, err := json.Marshal(tfstate)
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(filePath, tfstateData, 0644)
		if err != nil {
			return fmt.Errorf("failed to write file hash : %s", err.Error())
		}
	}

	return nil
}

func updateTFStateFile(filePath string, gv schema.GroupVersion, u *unstructured.Unstructured) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	var tfstate *apis.State
	err = json.Unmarshal(data, &tfstate)
	if err != nil {
		return err
	}

	jsonData, err := meta.MarshalToJson(u, gv)
	if err != nil {
		return err
	}
	typedObj, err := meta.UnmarshalFromJSON(jsonData, gv)
	if err != nil {
		return err
	}

	typedStruct := structs.New(typedObj)
	stateValue := typedStruct.Field("Status").Field("State").Value()

	if stateValue.(*apis.State) == nil || stateValue.(*apis.State).Serial != tfstate.Serial {
		err = setNestedFieldNoCopy(u.Object, tfstate, "status", "state")
		if err != nil {
			return err
		}
		return nil
	}

	return nil
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

func processSensitiveFields(r reflect.Type, v reflect.Value, tfkey string, data *map[string]string) {
	d := *data
	n := r.NumField()
	for i := 0; i < n; i++ {
		field := r.Field(i)
		value := v.Field(i)
		tftag := strings.ReplaceAll(field.Tag.Get("tf"), ",omitempty", "")
		newtfkey := tftag
		if tfkey != "" {
			newtfkey = tfkey + "." + tftag
		}

		if field.Tag.Get("sensitive") == "true" && value.Kind() == reflect.String && value.Interface().(string) != "" {
			d[newtfkey] = value.String()
		} else if field.Tag.Get("sensitive") == "true" && value.Kind() == reflect.Map && value.Interface().(map[string]string) != nil && len(value.Interface().(map[string]string)) != 0 {
			secretJson, err := json.Marshal(value.Interface())
			if err != nil {
				log.Error(err)
			} else {
				d[newtfkey] = string(secretJson)
			}
		}

		if value.Kind() == reflect.Struct {
			processSensitiveFields(value.Type(), value, newtfkey, &d)
		}

		if value.Kind() == reflect.Slice {
			n := value.Len()
			for i := 0; i < n; i++ {
				if value.Index(i).Kind() == reflect.Struct {
					processSensitiveFields(value.Index(i).Type(), value.Index(i), newtfkey+"."+strconv.FormatInt(int64(i), 10), &d)
				}
			}
		}
	}
}
