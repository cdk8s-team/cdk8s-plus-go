//go:build !no_runtime_type_checking

package cdk8splus27

import (
	"fmt"
)

func validateNewResourcePermissionsParameters(instance Resource) error {
	if instance == nil {
		return fmt.Errorf("parameter instance is required, but nil was provided")
	}

	return nil
}

