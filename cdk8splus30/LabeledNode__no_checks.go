//go:build no_runtime_type_checking

package cdk8splus30

// Building without runtime type checking enabled, so all the below just return nil

func validateNewLabeledNodeParameters(labelSelector *[]NodeLabelQuery) error {
	return nil
}

