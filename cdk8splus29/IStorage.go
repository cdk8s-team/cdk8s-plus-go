package cdk8splus29

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus29/v2/internal"
)

// Represents a piece of storage in the cluster.
type IStorage interface {
	constructs.IConstruct
	// Convert the piece of storage into a concrete volume.
	AsVolume() Volume
}

// The jsii proxy for IStorage
type jsiiProxy_IStorage struct {
	internal.Type__constructsIConstruct
}

func (i *jsiiProxy_IStorage) AsVolume() Volume {
	var returns Volume

	_jsii_.Invoke(
		i,
		"asVolume",
		nil, // no parameters
		&returns,
	)

	return returns
}

