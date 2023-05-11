//go:build !no_runtime_type_checking

package cdk8splus25

import (
	"fmt"
)

func validateTopology_CustomParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

