// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus24/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// In Kubernetes, namespaces provides a mechanism for isolating groups of resources within a single cluster.
//
// Names of resources need to be unique within a namespace, but not across namespaces.
// Namespace-based scoping is applicable only for namespaced objects (e.g. Deployments, Services, etc) and
// not for cluster-wide objects (e.g. StorageClass, Nodes, PersistentVolumes, etc).
type Namespace interface {
	Resource
	INamespaceSelector
	INetworkPolicyPeer
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
	// Return the IApiResource this object represents.
	AsApiResource() IApiResource
	// Return the non resource url this object represents.
	AsNonApiResource() *string
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

// The jsii proxy struct for Namespace
type jsiiProxy_Namespace struct {
	jsiiProxy_Resource
	jsiiProxy_INamespaceSelector
	jsiiProxy_INetworkPolicyPeer
}

func (j *jsiiProxy_Namespace) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Namespace) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Namespace) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Namespace) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Namespace) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Namespace) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Namespace) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Namespace) Permissions() ResourcePermissions {
	var returns ResourcePermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Namespace) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Namespace) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}


func NewNamespace(scope constructs.Construct, id *string, props *NamespaceProps) Namespace {
	_init_.Initialize()

	if err := validateNewNamespaceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Namespace{}

	_jsii_.Create(
		"cdk8s-plus-24.Namespace",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewNamespace_Override(n Namespace, scope constructs.Construct, id *string, props *NamespaceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.Namespace",
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
func Namespace_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNamespace_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Namespace",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Namespace_NAME_LABEL() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdk8s-plus-24.Namespace",
		"NAME_LABEL",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_Namespace) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		n,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Namespace) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Namespace) ToNamespaceSelectorConfig() *NamespaceSelectorConfig {
	var returns *NamespaceSelectorConfig

	_jsii_.Invoke(
		n,
		"toNamespaceSelectorConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Namespace) ToNetworkPolicyPeerConfig() *NetworkPolicyPeerConfig {
	var returns *NetworkPolicyPeerConfig

	_jsii_.Invoke(
		n,
		"toNetworkPolicyPeerConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Namespace) ToPodSelector() IPodSelector {
	var returns IPodSelector

	_jsii_.Invoke(
		n,
		"toPodSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Namespace) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

