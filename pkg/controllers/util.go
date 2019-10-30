/*
Copyright The Kubeform Authors.

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
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	base "kubeform.dev/kubeform/apis/base/v1alpha1"
	"kubeform.dev/kubeform/data"

	"ekyu.moe/base91"
	"github.com/appscode/go/log"
	"github.com/fatih/structs"
	"github.com/gobuffalo/flect"
	jsoniter "github.com/json-iterator/go"
	"gocloud.dev/secrets"
	_ "gocloud.dev/secrets/localsecrets"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	du "kmodules.xyz/client-go/dynamic"
	"kmodules.xyz/client-go/meta"
)

const KFCFinalizer = "kfc.io"

var (
	basePath = filepath.Join("/tmp", ".kfc")
)

func secretToTFProvider(secret *corev1.Secret, providerName, providerFile string) error {
	d1 := []byte(`{ "provider": { "` + providerName + `":`)
	tempData := make(map[string]string)
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
				var tempMap = make(map[string]string)
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

func crdToModule(kc kubernetes.Interface, gv schema.GroupVersion, obj *unstructured.Unstructured, source, mainFile, outputFile string) error {
	moduleName := obj.GetName()
	err := unstructured.SetNestedField(obj.Object, source, "spec", "source")
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
			secret, err := kc.CoreV1().Secrets(obj.GetNamespace()).Get(secretName.(string), v1.GetOptions{})
			if err != nil {
				return err
			}

			for key := range secret.Data {
				val := secret.Data[key]

				tempMap := make(map[string]string)
				buffer := new(bytes.Buffer)
				var secretData interface{}

				if err := json.Compact(buffer, val); err != nil {
					secretData = strings.ReplaceAll(string(val), "\n", "")
				} else {
					err = json.Unmarshal(buffer.Bytes(), &tempMap)
					if err != nil {
						return err
					}
					secretData = tempMap
				}

				specValue.Elem().FieldByName(flect.Capitalize(flect.Camelize(key))).Set(reflect.ValueOf(secretData))
			}
		}
	}

	str, err := jsonit.Marshal(specValue.Interface())
	if err != nil {
		return err
	}

	moduleData :=
		[]byte(`{"module":{ "` + moduleName + `":`)

	moduleData = append(moduleData, str...)
	moduleData = append(moduleData, []byte("} }")...)
	prettyData, err := prettyJSON(moduleData)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(mainFile, prettyData, 0644)
	if err != nil {
		return err
	}

	outputData := []byte(``)
	output := reflect.TypeOf(typedStruct.Field("Status").Field("Output").Value()).Elem()

	for i := 0; i < output.NumField(); i++ {
		field := output.Field(i).Tag.Get("tf")
		outputData = append(outputData, []byte(`output "`+field+`" { 
value = module.`+moduleName+`.`+field+` 
}
`)...)
	}

	err = ioutil.WriteFile(outputFile, outputData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func updateStateField(c *Controller, namespace, providerName, filePath string, gvr schema.GroupVersionResource, obj *unstructured.Unstructured) error {
	kc := c.kubeclientset
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
	secretData := make(map[string]string)
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
				_, err = kc.CoreV1().Secrets(namespace).Create(&corev1.Secret{
					ObjectMeta: v1.ObjectMeta{
						Name:      secretName,
						Namespace: namespace,
					},
					Type: corev1.SecretType("kubeform.com/" + providerName),
				})
				if err != nil {
					return err
				}
			}
			return err
		}
		if secret.Data == nil {
			secret.Data = make(map[string][]byte)
		}

		for key := range secretData {
			secret.Data["out."+key] = []byte(secretData[key])
		}

		_, err = kc.CoreV1().Secrets(namespace).Update(secret)
		if err != nil {
			return err
		}
	}

	output := s.Field("Spec").Value()
	specByte, err := json.Marshal(output)
	if err != nil {
		return err
	}

	var specMap map[string]interface{}
	err = json.Unmarshal(specByte, &specMap)
	if err != nil {
		return err
	}

	_, err = du.UpdateStatus(c.dynamicclient, gvr, obj, func(in *unstructured.Unstructured) *unstructured.Unstructured {
		err := unstructured.SetNestedField(in.Object, specMap, "status", "output")
		if err != nil {
			log.Error("failed to update status output")
		}

		return in
	})
	if err != nil {
		return err
	}

	return nil
}

func updateOutputField(c *Controller, respath, namespace, providerName string, gvr schema.GroupVersionResource, obj *unstructured.Unstructured) error {
	kc := c.kubeclientset
	value, err := terraformOutput(respath)
	if err != nil {
		return err
	}

	secretData := make(map[string][]byte)

	outputs := make(map[string]output)

	err = json.Unmarshal([]byte(value), &outputs)
	if err != nil {
		return err
	}

	for name, output := range outputs {
		if !output.Sensitive {
			val, err := output.ValueRaw.MarshalJSON()
			if err != nil {
				return err
			}

			_, err = du.UpdateStatus(c.dynamicclient, gvr, obj, func(in *unstructured.Unstructured) *unstructured.Unstructured {
				err := setNestedFieldNoCopy(in.Object, string(val), "status", "output", flect.Camelize(name))
				if err != nil {
					log.Error("failed to update status output")
				}

				return in
			})
			if err != nil {
				log.Error(err)
			}
		} else {
			secretData[name] = output.ValueRaw
		}
	}

	if len(secretData) != 0 {
		var secretName interface{}

		secretRef, _, err := unstructured.NestedFieldNoCopy(obj.Object, "spec", "secretRef")
		if err != nil {
			return err
		}

		if secretRef != nil {
			secretName, _, _ = unstructured.NestedFieldNoCopy(obj.Object, "spec", "secretRef", "name")
		} else {
			secretName = obj.GetName() + "-" + obj.GetNamespace() + "-" + "sensitive"
		}

		var secret *corev1.Secret
		secret, err = kc.CoreV1().Secrets(namespace).Get(secretName.(string), v1.GetOptions{})
		if err != nil {
			if errors.ReasonForError(err) == v1.StatusReasonNotFound {
				_, err = kc.CoreV1().Secrets(namespace).Create(&corev1.Secret{
					ObjectMeta: v1.ObjectMeta{
						Name:      secretName.(string),
						Namespace: namespace,
					},
					Type: corev1.SecretType("kubeform.com/" + providerName),
				})
				if err != nil {
					return err
				}
			}
			return err
		}

		for key := range secretData {
			secret.Data[key] = secretData[key]
		}

		_, err = kc.CoreV1().Secrets(namespace).Update(secret)
		if err != nil {
			return err
		}
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

func addFinalizer(dynamicclient dynamic.Interface, gvr schema.GroupVersionResource, u *unstructured.Unstructured, finalizer string) error {

	finalizers := u.GetFinalizers()
	for _, v := range finalizers {
		if v == finalizer {
			return nil
		}
	}

	finalizers = append(finalizers, finalizer)

	_, err := du.TryUpdate(dynamicclient, gvr, v1.ObjectMeta{Name: u.GetName(), Namespace: u.GetNamespace()}, func(in *unstructured.Unstructured) *unstructured.Unstructured {
		err := unstructured.SetNestedStringSlice(in.Object, finalizers, "metadata", "finalizers")
		if err != nil {
			log.Error(err)
		}

		return in
	})

	return err
}

func removeFinalizer(dynamicclient dynamic.Interface, gvr schema.GroupVersionResource, u *unstructured.Unstructured, finalizer string) error {
	finalizers := u.GetFinalizers()
	for i, v := range finalizers {
		if v == finalizer {
			finalizers = append(finalizers[:i], finalizers[i+1:]...)
			break
		}
	}

	_, err := du.TryUpdate(dynamicclient, gvr, v1.ObjectMeta{Name: u.GetName(), Namespace: u.GetNamespace()}, func(in *unstructured.Unstructured) *unstructured.Unstructured {
		err := unstructured.SetNestedStringSlice(in.Object, finalizers, "metadata", "finalizers")
		if err != nil {
			log.Error(err)
		}

		return in
	})

	return err
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

func createTFState(kc kubernetes.Interface, filePath, providerName string, isModule bool, gv schema.GroupVersion, u *unstructured.Unstructured) error {

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
	if isModule {
		if os.IsNotExist(existErr) && stateValue.(string) != "" {
			decodedData, err := decodeState(stateValue.(string))
			if err != nil {
				return err
			}

			err = ioutil.WriteFile(filePath, decodedData, 0644)
			if err != nil {
				return fmt.Errorf("failed to write file hash : %s", err.Error())
			}
		}

		return nil
	}
	outputValue := reflect.ValueOf(typedStruct.Field("Status").Field("Output").Value())

	if os.IsNotExist(existErr) && stateValue.(*base.State) != nil {
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
					tempMap := make(map[string]string)
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

func updateTFStateFile(c *Controller, filePath string, isModule bool, gvr schema.GroupVersionResource, u *unstructured.Unstructured) error {
	gv := gvr.GroupVersion()
	data, err := ioutil.ReadFile(filePath)
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

	if isModule {
		if stateValue.(string) == "" || !reflect.DeepEqual([]byte(stateValue.(string)), data) {
			processedData, err := encodeState(data)
			if err != nil {
				return err
			}

			_, err = du.UpdateStatus(c.dynamicclient, gvr, u, func(in *unstructured.Unstructured) *unstructured.Unstructured {
				err := unstructured.SetNestedField(in.Object, processedData, "status", "state")
				if err != nil {
					log.Error("failed to update status state")
				}

				return in
			})
			if err != nil {
				return err
			}
		}
		return nil
	}
	var tfstate *base.State
	err = json.Unmarshal(data, &tfstate)
	if err != nil {
		return err
	}

	if stateValue.(*base.State) == nil || stateValue.(*base.State).Serial != tfstate.Serial {
		tfstateByte, err := json.Marshal(tfstate)
		if err != nil {
			return err
		}
		var tfstateMap map[string]interface{}
		err = json.Unmarshal(tfstateByte, &tfstateMap)
		if err != nil {
			return err
		}

		_, err = du.UpdateStatus(c.dynamicclient, gvr, u, func(in *unstructured.Unstructured) *unstructured.Unstructured {
			err := unstructured.SetNestedField(in.Object, tfstateMap, "status", "state")
			if err != nil {
				log.Error("failed to update status state")
			}

			return in
		})
		if err != nil {
			return err
		}
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

func isModule(group string) bool {
	s := strings.Split(group, ".")

	return s[0] == "modules"
}

func encodeState(data []byte) (string, error) {
	// zip
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	if _, err := zw.Write(data); err != nil {
		log.Fatal(err)
	}

	if err := zw.Close(); err != nil {
		log.Fatal(err)
	}

	// encrypt
	savedKeyKeeper, err := secrets.OpenKeeper(context.Background(), "base64key://"+SecretKey)
	if err != nil {
		return "", err
	}
	defer savedKeyKeeper.Close()

	cipherText, err := savedKeyKeeper.Encrypt(context.Background(), buf.Bytes())
	if err != nil {
		return "", err
	}

	// base91

	return base91.EncodeToString(cipherText), nil
}

func decodeState(data string) ([]byte, error) {
	cipherText := base91.DecodeString(data)

	savedKeyKeeper, err := secrets.OpenKeeper(context.Background(), "base64key://"+SecretKey)
	if err != nil {
		return nil, err
	}
	defer savedKeyKeeper.Close()

	plainText, err := savedKeyKeeper.Decrypt(context.Background(), cipherText)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(plainText)

	zr, err := gzip.NewReader(buf)
	if err != nil {
		return nil, err
	}

	result, err := ioutil.ReadAll(zr)
	if err != nil {
		return nil, err
	}

	if err := zr.Close(); err != nil {
		return nil, err
	}

	return result, nil
}

func getModuleProviderAndSource(name string) (string, string) {
	for _, moduleConfig := range data.ModuleConfig {
		if moduleConfig.Name == name {
			return moduleConfig.Provider, moduleConfig.Source
		}
	}

	return "", ""
}
