//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeCronJob_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeCronJob_ManifestParameters(props *KubeCronJobProps) error {
	return nil
}

func validateKubeCronJob_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeCronJobParameters(scope constructs.Construct, id *string, props *KubeCronJobProps) error {
	return nil
}

