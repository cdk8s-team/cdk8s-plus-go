//go:build no_runtime_type_checking

package cdk8splus30

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_Ingress) validateAddDefaultBackendParameters(backend IngressBackend) error {
	return nil
}

func (i *jsiiProxy_Ingress) validateAddHostDefaultBackendParameters(host *string, backend IngressBackend) error {
	return nil
}

func (i *jsiiProxy_Ingress) validateAddHostRuleParameters(host *string, path *string, backend IngressBackend) error {
	return nil
}

func (i *jsiiProxy_Ingress) validateAddRuleParameters(path *string, backend IngressBackend) error {
	return nil
}

func (i *jsiiProxy_Ingress) validateAddRulesParameters(rules *[]*IngressRule) error {
	return nil
}

func (i *jsiiProxy_Ingress) validateAddTlsParameters(tls *[]*IngressTls) error {
	return nil
}

func validateIngress_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewIngressParameters(scope constructs.Construct, id *string, props *IngressProps) error {
	return nil
}

