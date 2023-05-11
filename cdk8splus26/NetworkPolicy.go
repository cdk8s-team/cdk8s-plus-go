package cdk8splus26

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus26/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Control traffic flow at the IP address or port level (OSI layer 3 or 4), network policies are an application-centric construct which allow you to specify how a pod is allowed to communicate with various network peers.
//
// - Outgoing traffic is allowed if there are no network policies selecting
//    the pod (and cluster policy otherwise allows the traffic),
//    OR if the traffic matches at least one egress rule across all of the
//    network policies that select the pod.
//
// - Incoming traffic is allowed to a pod if there are no network policies
//    selecting the pod (and cluster policy otherwise allows the traffic),
//    OR if the traffic source is the pod's local node,
//    OR if the traffic matches at least one ingress rule across all of
//    the network policies that select the pod.
//
// Network policies do not conflict; they are additive.
// If any policy or policies apply to a given pod for a given
// direction, the connections allowed in that direction from
// that pod is the union of what the applicable policies allow.
// Thus, order of evaluation does not affect the policy result.
//
// For a connection from a source pod to a destination pod to be allowed,
// both the egress policy on the source pod and the ingress policy on the
// destination pod need to allow the connection.
// If either side does not allow the connection, it will not happen.
// See: https://kubernetes.io/docs/concepts/services-networking/network-policies/#networkpolicy-resource
//
type NetworkPolicy interface {
	Resource
	// The group portion of the API version (e.g. "authorization.k8s.io").
	ApiGroup() *string
	// The underlying cdk8s API object.
	// See: base.Resource.apiObject
	//
	ApiObject() cdk8s.ApiObject
	// The object's API version (e.g. "authorization.k8s.io/v1").
	ApiVersion() *string
	// The object kind (e.g. "Deployment").
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	// The name of this API object.
	Name() *string
	// The tree node.
	Node() constructs.Node
	Permissions() ResourcePermissions
	// The unique, namespace-global, name of an object inside the Kubernetes cluster.
	//
	// If this is omitted, the ApiResource should represent all objects of the given type.
	ResourceName() *string
	// The name of a resource type as it appears in the relevant API endpoint.
	ResourceType() *string
	// Allow outgoing traffic to the peer.
	//
	// If ports are not passed, traffic will be allowed on all ports.
	AddEgressRule(peer INetworkPolicyPeer, ports *[]NetworkPolicyPort)
	// Allow incoming traffic from the peer.
	//
	// If ports are not passed, traffic will be allowed on all ports.
	AddIngressRule(peer INetworkPolicyPeer, ports *[]NetworkPolicyPort)
	// Return the IApiResource this object represents.
	AsApiResource() IApiResource
	// Return the non resource url this object represents.
	AsNonApiResource() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for NetworkPolicy
type jsiiProxy_NetworkPolicy struct {
	jsiiProxy_Resource
}

func (j *jsiiProxy_NetworkPolicy) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicy) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicy) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicy) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicy) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicy) Permissions() ResourcePermissions {
	var returns ResourcePermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicy) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicy) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}


func NewNetworkPolicy(scope constructs.Construct, id *string, props *NetworkPolicyProps) NetworkPolicy {
	_init_.Initialize()

	if err := validateNewNetworkPolicyParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkPolicy{}

	_jsii_.Create(
		"cdk8s-plus-26.NetworkPolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewNetworkPolicy_Override(n NetworkPolicy, scope constructs.Construct, id *string, props *NetworkPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-26.NetworkPolicy",
		[]interface{}{scope, id, props},
		n,
	)
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
func NetworkPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkPolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-26.NetworkPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPolicy) AddEgressRule(peer INetworkPolicyPeer, ports *[]NetworkPolicyPort) {
	if err := n.validateAddEgressRuleParameters(peer); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addEgressRule",
		[]interface{}{peer, ports},
	)
}

func (n *jsiiProxy_NetworkPolicy) AddIngressRule(peer INetworkPolicyPeer, ports *[]NetworkPolicyPort) {
	if err := n.validateAddIngressRuleParameters(peer); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addIngressRule",
		[]interface{}{peer, ports},
	)
}

func (n *jsiiProxy_NetworkPolicy) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		n,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPolicy) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

