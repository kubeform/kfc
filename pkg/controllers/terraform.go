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
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func terraformInit(resPath string) error {
	cmd := exec.Command("terraform", "init")
	cmd.Dir = resPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func terraformApply(resPath string) error {
	cmd := exec.Command("terraform", "apply", "-auto-approve")
	cmd.Dir = resPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func terraformDestroy(resPath string) error {
	cmd := exec.Command("terraform", "destroy", "-auto-approve")
	cmd.Dir = resPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func terraformOutput(resPath string) (string, error) {
	cmd := exec.Command("terraform", "output", "-json")
	cmd.Dir = resPath
	outputBuffer := new(bytes.Buffer)
	outputWriter := bufio.NewWriter(outputBuffer)
	cmd.Stdout = outputWriter

	errBuffer := new(bytes.Buffer)
	errWriter := bufio.NewWriter(errBuffer)
	cmd.Stderr = errWriter

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	if errBuffer.String() != "" {
		return "", fmt.Errorf("%v", errBuffer.String())
	}

	return outputBuffer.String(), nil
}
