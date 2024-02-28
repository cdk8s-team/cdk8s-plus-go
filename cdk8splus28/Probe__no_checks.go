//go:build no_runtime_type_checking

package cdk8splus28

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

