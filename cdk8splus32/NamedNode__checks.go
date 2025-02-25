//go:build !no_runtime_type_checking

package cdk8splus32

import (
	"fmt"
)

func validateNewNamedNodeParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

