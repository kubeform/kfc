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
