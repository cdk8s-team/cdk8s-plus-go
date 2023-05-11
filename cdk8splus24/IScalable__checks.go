//go:build !no_runtime_type_checking

package cdk8splus24

import (
	"fmt"
)

func (j *jsiiProxy_IScalable) validateSetHasAutoscalerParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

