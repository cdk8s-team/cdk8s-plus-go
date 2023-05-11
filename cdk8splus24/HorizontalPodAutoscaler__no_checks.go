//go:build no_runtime_type_checking

package cdk8splus24

// Building without runtime type checking enabled, so all the below just return nil

func validateHorizontalPodAutoscaler_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewHorizontalPodAutoscalerParameters(scope constructs.Construct, id *string, props *HorizontalPodAutoscalerProps) error {
	return nil
}

