package controllers

import (
	"bytes"
	"errors"
	"os"

	"k8s.io/apimachinery/pkg/runtime"

	"github.com/hashicorp/terraform/command"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func terraformInit(resPath string) error {
	err := os.Chdir(resPath)
	if err != nil {
		return err
	}

	codeUi := &CodeUi{
		OutputBuffer: new(bytes.Buffer),
	}
	initCommand := command.InitCommand{
		Meta: command.Meta{
			Ui: codeUi,
		},
	}

	x := initCommand.Run(nil)

	if x != 0 {
		return errors.New("failed to run terraform init command")
	}

	return nil
}

func terraformApply(resPath string) error {
	err := os.Chdir(resPath)
	if err != nil {
		return err
	}
	codeUi := &CodeUi{
		OutputBuffer: new(bytes.Buffer),
	}

	cmd := command.ApplyCommand{
		Meta: command.Meta{
			Ui: codeUi,
		},
	}

	args := []string{
		"-auto-approve",
	}
	x := cmd.Run(args)
	if x != 0 {
		return errors.New("failed to run terraform apply command")
	}

	return nil
}

func terraformDestroy(resPath string) error {
	err := os.Chdir(resPath)
	if err != nil {
		return err
	}
	codeUi := &CodeUi{
		OutputBuffer: new(bytes.Buffer),
	}

	cmd := command.ApplyCommand{
		Meta: command.Meta{
			Ui: codeUi,
		},
		Destroy: true,
	}

	args := []string{
		"-auto-approve",
	}
	x := cmd.Run(args)
	if x != 0 {
		return errors.New("failed to run terraform destroy command")
	}

	return nil
}

func updateStatusOut(u *unstructured.Unstructured, resPath string) error {
	err := os.Chdir(resPath)
	if err != nil {
		return err
	}
	codeUi := &CodeUi{
		OutputBuffer: new(bytes.Buffer),
	}
	showCmd := command.ShowCommand{
		Meta: command.Meta{
			Ui: codeUi,
		},
	}

	args := []string{
		"-json",
	}

	x := showCmd.Run(args)
	if x != 0 {
		return errors.New("failed to run terraform show command")
	}

	out := codeUi.OutputBuffer.Bytes()
	rawData := &runtime.RawExtension{
		Raw: out,
	}

	return setNestedFieldNoCopy(u.Object, rawData, "status", "output")
}
