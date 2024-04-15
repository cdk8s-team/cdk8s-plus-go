//go:build !no_runtime_type_checking

package cdk8splus29

import (
	"fmt"
)

func validateNode_NamedParameters(nodeName *string) error {
	if nodeName == nil {
		return fmt.Errorf("parameter nodeName is required, but nil was provided")
	}

	return nil
}

