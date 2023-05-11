package cdk8splus26

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus26/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus26/v2/internal"
)

// Represents a group of namespaces.
type Namespaces interface {
	constructs.Construct
	INamespaceSelector
	INetworkPolicyPeer
	// The tree node.
	Node() constructs.Node
	// Return the configuration of this selector.
	// See: INamespaceSelector.toNamespaceSelectorConfig()
	//
	ToNamespaceSelectorConfig() *NamespaceSelectorConfig
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

// The jsii proxy struct for Namespaces
type jsiiProxy_Namespaces struct {
	internal.Type__constructsConstruct
	jsiiProxy_INamespaceSelector
	jsiiProxy_INetworkPolicyPeer
}

func (j *jsiiProxy_Namespaces) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewNamespaces(scope constructs.Construct, id *string, expressions *[]LabelExpression, names *[]*string, labels *map[string]*string) Namespaces {
	_init_.Initialize()

	if err := validateNewNamespacesParameters(scope, id); err != nil {
		panic(err)
	}
	j := jsiiProxy_Namespaces{}

	_jsii_.Create(
		"cdk8s-plus-26.Namespaces",
		[]interface{}{scope, id, expressions, names, labels},
		&j,
	)

	return &j
}

func NewNamespaces_Override(n Namespaces, scope constructs.Construct, id *string, expressions *[]LabelExpression, names *[]*string, labels *map[string]*string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-26.Namespaces",
		[]interface{}{scope, id, expressions, names, labels},
		n,
	)
}

// Select all namespaces.
func Namespaces_All(scope constructs.Construct, id *string) Namespaces {
	_init_.Initialize()

	if err := validateNamespaces_AllParameters(scope, id); err != nil {
		panic(err)
	}
	var returns Namespaces

	_jsii_.StaticInvoke(
		"cdk8s-plus-26.Namespaces",
		"all",
		[]interface{}{scope, id},
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
func Namespaces_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNamespaces_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-26.Namespaces",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Select specific namespaces.
func Namespaces_Select(scope constructs.Construct, id *string, options *NamespacesSelectOptions) Namespaces {
	_init_.Initialize()

	if err := validateNamespaces_SelectParameters(scope, id, options); err != nil {
		panic(err)
	}
	var returns Namespaces

	_jsii_.StaticInvoke(
		"cdk8s-plus-26.Namespaces",
		"select",
		[]interface{}{scope, id, options},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Namespaces) ToNamespaceSelectorConfig() *NamespaceSelectorConfig {
	var returns *NamespaceSelectorConfig

	_jsii_.Invoke(
		n,
		"toNamespaceSelectorConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Namespaces) ToNetworkPolicyPeerConfig() *NetworkPolicyPeerConfig {
	var returns *NetworkPolicyPeerConfig

	_jsii_.Invoke(
		n,
		"toNetworkPolicyPeerConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Namespaces) ToPodSelector() IPodSelector {
	var returns IPodSelector

	_jsii_.Invoke(
		n,
		"toPodSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Namespaces) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

