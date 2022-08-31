//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24

// Building without runtime type checking enabled, so all the below just return nil

func validateProbe_FromCommandParameters(command *[]*string, options *CommandProbeOptions) error {
	return nil
}

func validateProbe_FromHttpGetParameters(path *string, options *HttpGetProbeOptions) error {
	return nil
}

func validateProbe_FromTcpSocketParameters(options *TcpSocketProbeOptions) error {
	return nil
}

