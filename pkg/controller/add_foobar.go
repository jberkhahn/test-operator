package controller

import (
	"github.com/jberkhahn/test-operator/pkg/controller/foobar"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, foobar.Add)
}
