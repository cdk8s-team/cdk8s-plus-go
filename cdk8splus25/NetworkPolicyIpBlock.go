package cdk8splus25

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus25/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus25/v2/internal"
)

// Describes a particular CIDR (Ex.
//
// "192.168.1.1/24","2001:db9::/64") that is
// allowed to the pods matched by a network policy selector.
// The except entry describes CIDRs that should not be included within this rule.
type NetworkPolicyIpBlock interface {
	constructs.Construct
	INetworkPolicyPeer
	// A string representing the IP Block Valid examples are "192.168.1.1/24" or "2001:db9::/64".
	Cidr() *string
	// A slice of CIDRs that should not be included within an IP Block Valid examples are "192.168.1.1/24" or "2001:db9::/64". Except values will be rejected if they are outside the CIDR range.
	Except() *[]*string
	// The tree node.
	Node() constructs.Node
	// Return the configuration of this peer.
	// See: INetworkPolicyPeer.toNetworkPolicyPeerConfig()
	//
	ToNetworkPolicyPeerConfig() *NetworkPolicyPeerConfig
	// Convert the peer into a pod selector, if possible.
	// See: INetworkPolicyPeer.toPodSelector()
	//
	ToPodSelector() IPodSelector
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for NetworkPolicyIpBlock
type jsiiProxy_NetworkPolicyIpBlock struct {
	internal.Type__constructsConstruct
	jsiiProxy_INetworkPolicyPeer
}

func (j *jsiiProxy_NetworkPolicyIpBlock) Cidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicyIpBlock) Except() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"except",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicyIpBlock) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Any IPv4 address.
func NetworkPolicyIpBlock_AnyIpv4(scope constructs.Construct, id *string) NetworkPolicyIpBlock {
	_init_.Initialize()

	if err := validateNetworkPolicyIpBlock_AnyIpv4Parameters(scope, id); err != nil {
		panic(err)
	}
	var returns NetworkPolicyIpBlock

	_jsii_.StaticInvoke(
		"cdk8s-plus-25.NetworkPolicyIpBlock",
		"anyIpv4",
		[]interface{}{scope, id},
		&returns,
	)

	return returns
}

// Any IPv6 address.
func NetworkPolicyIpBlock_AnyIpv6(scope constructs.Construct, id *string) NetworkPolicyIpBlock {
	_init_.Initialize()

	if err := validateNetworkPolicyIpBlock_AnyIpv6Parameters(scope, id); err != nil {
		panic(err)
	}
	var returns NetworkPolicyIpBlock

	_jsii_.StaticInvoke(
		"cdk8s-plus-25.NetworkPolicyIpBlock",
		"anyIpv6",
		[]interface{}{scope, id},
		&returns,
	)

	return returns
}

// Create an IPv4 peer from a CIDR.
func NetworkPolicyIpBlock_Ipv4(scope constructs.Construct, id *string, cidrIp *string, except *[]*string) NetworkPolicyIpBlock {
	_init_.Initialize()

	if err := validateNetworkPolicyIpBlock_Ipv4Parameters(scope, id, cidrIp); err != nil {
		panic(err)
	}
	var returns NetworkPolicyIpBlock

	_jsii_.StaticInvoke(
		"cdk8s-plus-25.NetworkPolicyIpBlock",
		"ipv4",
		[]interface{}{scope, id, cidrIp, except},
		&returns,
	)

	return returns
}

// Create an IPv6 peer from a CIDR.
func NetworkPolicyIpBlock_Ipv6(scope constructs.Construct, id *string, cidrIp *string, except *[]*string) NetworkPolicyIpBlock {
	_init_.Initialize()

	if err := validateNetworkPolicyIpBlock_Ipv6Parameters(scope, id, cidrIp); err != nil {
		panic(err)
	}
	var returns NetworkPolicyIpBlock

	_jsii_.StaticInvoke(
		"cdk8s-plus-25.NetworkPolicyIpBlock",
		"ipv6",
		[]interface{}{scope, id, cidrIp, except},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func NetworkPolicyIpBlock_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkPolicyIpBlock_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-25.NetworkPolicyIpBlock",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPolicyIpBlock) ToNetworkPolicyPeerConfig() *NetworkPolicyPeerConfig {
	var returns *NetworkPolicyPeerConfig

	_jsii_.Invoke(
		n,
		"toNetworkPolicyPeerConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPolicyIpBlock) ToPodSelector() IPodSelector {
	var returns IPodSelector

	_jsii_.Invoke(
		n,
		"toPodSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPolicyIpBlock) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

