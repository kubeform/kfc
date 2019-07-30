package controllers

import (
	"os"
	"os/exec"
)

func terraformInit(resPath string) error {
	cmd := exec.Command("terraform", "init")
	cmd.Dir = resPath
	cmd.Stdout = os.Stdout
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
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

//func updateStatusOut(u *unstructured.Unstructured, resPath string) error {
//	cmd := exec.Command("terraform", "show")
//	cmd.Dir = basePath
//	out, err := cmd.Output()
//	if err != nil {
//		return err
//	}
//	rawData := &runtime.RawExtension{
//		Raw: out,
//	}
//
//	return setNestedFieldNoCopy(u.Object, rawData, "status", "output")
//}
