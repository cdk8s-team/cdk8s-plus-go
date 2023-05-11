package cdk8splus26

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus26/v2/jsii"
)

// Controls network isolation rules for inter-pod communication.
type PodConnections interface {
	Instance() AbstractPod
	// Allow network traffic from the peer to this pod.
	//
	// By default, this will create an ingress network policy for this pod, and an egress
	// network policy for the peer. This is required if both sides are already isolated.
	// Use `options.isolation` to control this behavior.
	//
	// Example:
	//   // create only an egress policy that selects the 'web' pod to allow outgoing traffic
	//   // to the 'redis' pod. this requires the 'redis' pod to not be isolated for ingress.
	//   redis.connections.allowFrom(web, { isolation: Isolation.PEER })
	//
	//   // create only an ingress policy that selects the 'redis' peer to allow incoming traffic
	//   // from the 'web' pod. this requires the 'web' pod to not be isolated for egress.
	//   redis.connections.allowFrom(web, { isolation: Isolation.POD })
	//
	AllowFrom(peer INetworkPolicyPeer, options *PodConnectionsAllowFromOptions)
	// Allow network traffic from this pod to the peer.
	//
	// By default, this will create an egress network policy for this pod, and an ingress
	// network policy for the peer. This is required if both sides are already isolated.
	// Use `options.isolation` to control this behavior.
	//
	// Example:
	//   // create only an egress policy that selects the 'web' pod to allow outgoing traffic
	//   // to the 'redis' pod. this requires the 'redis' pod to not be isolated for ingress.
	//   web.connections.allowTo(redis, { isolation: Isolation.POD })
	//
	//   // create only an ingress policy that selects the 'redis' peer to allow incoming traffic
	//   // from the 'web' pod. this requires the 'web' pod to not be isolated for egress.
	//   web.connections.allowTo(redis, { isolation: Isolation.PEER })
	//
	AllowTo(peer INetworkPolicyPeer, options *PodConnectionsAllowToOptions)
	// Sets the default network policy for Pod/Workload to have all egress and ingress connections as disabled.
	Isolate()
}

// The jsii proxy struct for PodConnections
type jsiiProxy_PodConnections struct {
	_ byte // padding
}

func (j *jsiiProxy_PodConnections) Instance() AbstractPod {
	var returns AbstractPod
	_jsii_.Get(
		j,
		"instance",
		&returns,
	)
	return returns
}


func NewPodConnections(instance AbstractPod) PodConnections {
	_init_.Initialize()

	if err := validateNewPodConnectionsParameters(instance); err != nil {
		panic(err)
	}
	j := jsiiProxy_PodConnections{}

	_jsii_.Create(
		"cdk8s-plus-26.PodConnections",
		[]interface{}{instance},
		&j,
	)

	return &j
}

func NewPodConnections_Override(p PodConnections, instance AbstractPod) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-26.PodConnections",
		[]interface{}{instance},
		p,
	)
}

func (p *jsiiProxy_PodConnections) AllowFrom(peer INetworkPolicyPeer, options *PodConnectionsAllowFromOptions) {
	if err := p.validateAllowFromParameters(peer, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"allowFrom",
		[]interface{}{peer, options},
	)
}

func (p *jsiiProxy_PodConnections) AllowTo(peer INetworkPolicyPeer, options *PodConnectionsAllowToOptions) {
	if err := p.validateAllowToParameters(peer, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"allowTo",
		[]interface{}{peer, options},
	)
}

func (p *jsiiProxy_PodConnections) Isolate() {
	_jsii_.InvokeVoid(
		p,
		"isolate",
		nil, // no parameters
	)
}

