// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-23 synthesizes Kubernetes manifests for Kubernetes 1.23.0
package cdk8splus23

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus23/v2/jsii"
)

// The backend for an ingress path.
type IngressBackend interface {
}

// The jsii proxy struct for IngressBackend
type jsiiProxy_IngressBackend struct {
	_ byte // padding
}

// A Resource backend is an ObjectRef to another Kubernetes resource within the same namespace as the Ingress object.
//
// A common usage for a Resource backend is to ingress data to an object
// storage backend with static assets.
func IngressBackend_FromResource(resource IResource) IngressBackend {
	_init_.Initialize()

	if err := validateIngressBackend_FromResourceParameters(resource); err != nil {
		panic(err)
	}
	var returns IngressBackend

	_jsii_.StaticInvoke(
		"cdk8s-plus-23.IngressBackend",
		"fromResource",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

// A Kubernetes `Service` to use as the backend for this path.
func IngressBackend_FromService(serv Service, options *ServiceIngressBackendOptions) IngressBackend {
	_init_.Initialize()

	if err := validateIngressBackend_FromServiceParameters(serv, options); err != nil {
		panic(err)
	}
	var returns IngressBackend

	_jsii_.StaticInvoke(
		"cdk8s-plus-23.IngressBackend",
		"fromService",
		[]interface{}{serv, options},
		&returns,
	)

	return returns
}

