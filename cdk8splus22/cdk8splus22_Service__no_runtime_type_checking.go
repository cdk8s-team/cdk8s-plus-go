//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-22 synthesizes Kubernetes manifests for Kubernetes 1.22.0
package cdk8splus22

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Service) validateBindParameters(port *float64, options *ServiceBindOptions) error {
	return nil
}

func (s *jsiiProxy_Service) validateExposeViaIngressParameters(path *string, options *ExposeServiceViaIngressOptions) error {
	return nil
}

func (s *jsiiProxy_Service) validateSelectParameters(selector IPodSelector) error {
	return nil
}

func (s *jsiiProxy_Service) validateSelectLabelParameters(key *string, value *string) error {
	return nil
}

func validateService_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewServiceParameters(scope constructs.Construct, id *string, props *ServiceProps) error {
	return nil
}

