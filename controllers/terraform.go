package controllers

import (
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
