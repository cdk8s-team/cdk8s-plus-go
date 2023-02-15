// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-25 synthesizes Kubernetes manifests for Kubernetes 1.25.0
package cdk8splus25

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus25/v2/jsii"
)

// Available topology domains.
type Topology interface {
	Key() *string
}

// The jsii proxy struct for Topology
type jsiiProxy_Topology struct {
	_ byte // padding
}

func (j *jsiiProxy_Topology) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}


// Custom key for the node label that the system uses to denote the topology domain.
func Topology_Custom(key *string) Topology {
	_init_.Initialize()

	if err := validateTopology_CustomParameters(key); err != nil {
		panic(err)
	}
	var returns Topology

	_jsii_.StaticInvoke(
		"cdk8s-plus-25.Topology",
		"custom",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func Topology_HOSTNAME() Topology {
	_init_.Initialize()
	var returns Topology
	_jsii_.StaticGet(
		"cdk8s-plus-25.Topology",
		"HOSTNAME",
		&returns,
	)
	return returns
}

func Topology_REGION() Topology {
	_init_.Initialize()
	var returns Topology
	_jsii_.StaticGet(
		"cdk8s-plus-25.Topology",
		"REGION",
		&returns,
	)
	return returns
}

func Topology_ZONE() Topology {
	_init_.Initialize()
	var returns Topology
	_jsii_.StaticGet(
		"cdk8s-plus-25.Topology",
		"ZONE",
		&returns,
	)
	return returns
}

