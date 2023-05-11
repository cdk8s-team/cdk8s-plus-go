//go:build !no_runtime_type_checking

package cdk8splus25

import (
	"fmt"
)

func validateNewTaintedNodeParameters(taintSelector *[]NodeTaintQuery) error {
	if taintSelector == nil {
		return fmt.Errorf("parameter taintSelector is required, but nil was provided")
	}

	return nil
}

