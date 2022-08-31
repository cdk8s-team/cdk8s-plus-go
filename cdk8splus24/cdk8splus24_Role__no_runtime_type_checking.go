//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_Role) validateAllowParameters(verbs *[]*string) error {
	return nil
}

func validateRole_FromRoleNameParameters(scope constructs.Construct, id *string, name *string) error {
	return nil
}

func validateRole_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewRoleParameters(scope constructs.Construct, id *string, props *RoleProps) error {
	return nil
}

