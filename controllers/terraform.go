package controllers

import (
	"bytes"
	"errors"

	"github.com/hashicorp/terraform/command"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func terraformInit(resPath string) error {
	// TODO: This initializes terraform in the current directory. Shouldn't it be moved to the resource directory?
	codeUi := &CodeUi{
		OutputBuffer: new(bytes.Buffer),
	}
	initCommand := command.InitCommand{
		Meta: command.Meta{
			Ui: codeUi,
		},
	}

	args := []string{
		resPath,
	}

	x := initCommand.Run(args)

	if x != 0 {
		return errors.New("failed to run terraform init command")
	}

	return nil
}

func terraformApply(resPath, stateFile string) error {
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
		"-state",
		stateFile,
		resPath,
	}
	x := cmd.Run(args)
	if x != 0 {
		return errors.New("failed to run terraform apply command")
	}

	return nil
}

func terraformDestroy(resPath, stateFile string) error {
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
		"-state",
		stateFile,
		resPath,
	}
	x := cmd.Run(args)
	if x != 0 {
		return errors.New("failed to run terraform destroy command")
	}

	return nil
}

func updateStatusOut(u *unstructured.Unstructured, resPath string) error {
	//TODO: Handle output
	//codeUi := &CodeUi{
	//	OutputBuffer: new(bytes.Buffer),
	//}
	//showCmd := command.ShowCommand{
	//	Meta: command.Meta{
	//		Ui: codeUi,
	//	},
	//}
	//
	//args := []string{
	//	"-json",
	//	resPath,
	//}
	//
	//x := showCmd.Run(args)
	//if x != 0 {
	//	return errors.New("failed to run terraform show command")
	//}
	//
	//out := codeUi.OutputBuffer.String()

	out := "updated"

	return unstructured.SetNestedField(u.Object, out, "status", "out")
}
