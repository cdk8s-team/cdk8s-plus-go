//go:build !no_runtime_type_checking

package cdk8splus27

import (
	"fmt"
)

func validateNewLabeledNodeParameters(labelSelector *[]NodeLabelQuery) error {
	if labelSelector == nil {
		return fmt.Errorf("parameter labelSelector is required, but nil was provided")
	}

	return nil
}

