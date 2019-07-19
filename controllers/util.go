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

	"k8s.io/apimachinery/pkg/runtime"

	"k8s.io/klog"

	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/fatih/structs"
	"kmodules.xyz/client-go/meta"

	jsoniter "github.com/json-iterator/go"

	"github.com/gobuffalo/flect"
	corev1 "k8s.io/api/core/v1"
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

func crdToTFResource(gv schema.GroupVersion, kind, providerName string, obj *unstructured.Unstructured, mainFile string) error {
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
	value := spec.Value()
	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
	}.Froze()

	tfStr, err := jsonit.Marshal(value)
	if err != nil {
		klog.Error(err)
	}

	d1 := []byte(`{"resource":{ "` + providerName + "_" + flect.Underscore(kind) + `":{"` + obj.GetName() + `":`)

	d1 = append(d1, tfStr...)
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
