package controllers

import (
	"bytes"
	"fmt"

	"github.com/appscode/go/log"
)

// CodeUi is a UI that is for using the cli in code
type CodeUi struct {
	OutputBuffer *bytes.Buffer
}

func (u *CodeUi) Ask(query string) (string, error) {
	return "", nil
}

func (u *CodeUi) AskSecret(query string) (string, error) {
	return u.Ask(query)
}

func (u *CodeUi) Error(message string) {
	log.Error(message)
}

func (u *CodeUi) Info(message string) {
	log.Info(message)
}

func (u *CodeUi) Output(message string) {
	fmt.Fprint(u.OutputBuffer, message)
	fmt.Fprint(u.OutputBuffer, "\n")
}

func (u *CodeUi) Warn(message string) {
	log.Info(message)
}
