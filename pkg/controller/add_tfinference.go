package controller

import (
	"github.com/yzyDavid/tf-inference-operator/pkg/controller/tfinference"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, tfinference.Add)
}
