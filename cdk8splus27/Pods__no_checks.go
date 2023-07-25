//go:build no_runtime_type_checking

package cdk8splus27

// Building without runtime type checking enabled, so all the below just return nil

func validatePods_AllParameters(scope constructs.Construct, id *string, options *PodsAllOptions) error {
	return nil
}

func validatePods_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePods_SelectParameters(scope constructs.Construct, id *string, options *PodsSelectOptions) error {
	return nil
}

func validateNewPodsParameters(scope constructs.Construct, id *string) error {
	return nil
}

