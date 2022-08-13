// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus24/v2/internal"
)

// Describes a peer to allow traffic to/from.
type INetworkPolicyPeer interface {
	constructs.IConstruct
	// Return the configuration of this peer.
	ToNetworkPolicyPeerConfig() *NetworkPolicyPeerConfig
	// Convert the peer into a pod selector, if possible.
	ToPodSelector() IPodSelector
}

// The jsii proxy for INetworkPolicyPeer
type jsiiProxy_INetworkPolicyPeer struct {
	internal.Type__constructsIConstruct
}

func (i *jsiiProxy_INetworkPolicyPeer) ToNetworkPolicyPeerConfig() *NetworkPolicyPeerConfig {
	var returns *NetworkPolicyPeerConfig

	_jsii_.Invoke(
		i,
		"toNetworkPolicyPeerConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_INetworkPolicyPeer) ToPodSelector() IPodSelector {
	var returns IPodSelector

	_jsii_.Invoke(
		i,
		"toPodSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

