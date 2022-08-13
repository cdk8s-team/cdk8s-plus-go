// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus24/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus24/v2/internal"
)

// Represents a group of pods.
type Pods interface {
	constructs.Construct
	IPodSelector
	// The tree node.
	Node() constructs.Node
	// See: INetworkPolicyPeer.toNetworkPolicyPeerConfig()
	//
	ToNetworkPolicyPeerConfig() *NetworkPolicyPeerConfig
	// See: INetworkPolicyPeer.toPodSelector()
	//
	ToPodSelector() IPodSelector
	// Return the configuration of this selector.
	// See: IPodSelector.toPodSelectorConfig()
	//
	ToPodSelectorConfig() *PodSelectorConfig
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Pods
type jsiiProxy_Pods struct {
	internal.Type__constructsConstruct
	jsiiProxy_IPodSelector
}

func (j *jsiiProxy_Pods) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewPods(scope constructs.Construct, id *string, expressions *[]LabelExpression, labels *map[string]*string, namespaces INamespaceSelector) Pods {
	_init_.Initialize()

	j := jsiiProxy_Pods{}

	_jsii_.Create(
		"cdk8s-plus-24.Pods",
		[]interface{}{scope, id, expressions, labels, namespaces},
		&j,
	)

	return &j
}

func NewPods_Override(p Pods, scope constructs.Construct, id *string, expressions *[]LabelExpression, labels *map[string]*string, namespaces INamespaceSelector) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.Pods",
		[]interface{}{scope, id, expressions, labels, namespaces},
		p,
	)
}

// Select all pods.
func Pods_All(scope constructs.Construct, id *string, options *PodsAllOptions) Pods {
	_init_.Initialize()

	var returns Pods

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Pods",
		"all",
		[]interface{}{scope, id, options},
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
func Pods_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Pods",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Select pods in the cluster with various selectors.
func Pods_Select(scope constructs.Construct, id *string, options *PodsSelectOptions) Pods {
	_init_.Initialize()

	var returns Pods

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Pods",
		"select",
		[]interface{}{scope, id, options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pods) ToNetworkPolicyPeerConfig() *NetworkPolicyPeerConfig {
	var returns *NetworkPolicyPeerConfig

	_jsii_.Invoke(
		p,
		"toNetworkPolicyPeerConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pods) ToPodSelector() IPodSelector {
	var returns IPodSelector

	_jsii_.Invoke(
		p,
		"toPodSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pods) ToPodSelectorConfig() *PodSelectorConfig {
	var returns *PodSelectorConfig

	_jsii_.Invoke(
		p,
		"toPodSelectorConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pods) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

