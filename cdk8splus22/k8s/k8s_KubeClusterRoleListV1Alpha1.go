package k8s

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus22/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus22/v2/k8s/internal"
)

// ClusterRoleList is a collection of ClusterRoles.
//
// Deprecated in v1.17 in favor of rbac.authorization.k8s.io/v1 ClusterRoles, and will no longer be served in v1.22.
type KubeClusterRoleListV1Alpha1 interface {
	cdk8s.ApiObject
	// The group portion of the API version (e.g. `authorization.k8s.io`).
	ApiGroup() *string
	// The object's API version (e.g. `authorization.k8s.io/v1`).
	ApiVersion() *string
	// The chart in which this object is defined.
	Chart() cdk8s.Chart
	// The object kind.
	Kind() *string
	// Metadata associated with this API object.
	Metadata() cdk8s.ApiObjectMetadataDefinition
	// The name of the API object.
	//
	// If a name is specified in `metadata.name` this will be the name returned.
	// Otherwise, a name will be generated by calling
	// `Chart.of(this).generatedObjectName(this)`, which by default uses the
	// construct path to generate a DNS-compatible name for the resource.
	Name() *string
	// The tree node.
	Node() constructs.Node
	// Create a dependency between this ApiObject and other constructs.
	//
	// These can be other ApiObjects, Charts, or custom.
	AddDependency(dependencies ...constructs.IConstruct)
	// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
	//
	// Example:
	//     kubePod.addJsonPatch(JsonPatch.replace('/spec/enableServiceLinks', true));
	//
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	// Renders the object to Kubernetes JSON.
	ToJson() interface{}
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for KubeClusterRoleListV1Alpha1
type jsiiProxy_KubeClusterRoleListV1Alpha1 struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_KubeClusterRoleListV1Alpha1) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeClusterRoleListV1Alpha1) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeClusterRoleListV1Alpha1) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeClusterRoleListV1Alpha1) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeClusterRoleListV1Alpha1) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeClusterRoleListV1Alpha1) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeClusterRoleListV1Alpha1) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Defines a "io.k8s.api.rbac.v1alpha1.ClusterRoleList" API object.
func NewKubeClusterRoleListV1Alpha1(scope constructs.Construct, id *string, props *KubeClusterRoleListV1Alpha1Props) KubeClusterRoleListV1Alpha1 {
	_init_.Initialize()

	if err := validateNewKubeClusterRoleListV1Alpha1Parameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubeClusterRoleListV1Alpha1{}

	_jsii_.Create(
		"cdk8s-plus-22.k8s.KubeClusterRoleListV1Alpha1",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "io.k8s.api.rbac.v1alpha1.ClusterRoleList" API object.
func NewKubeClusterRoleListV1Alpha1_Override(k KubeClusterRoleListV1Alpha1, scope constructs.Construct, id *string, props *KubeClusterRoleListV1Alpha1Props) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-22.k8s.KubeClusterRoleListV1Alpha1",
		[]interface{}{scope, id, props},
		k,
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
func KubeClusterRoleListV1Alpha1_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubeClusterRoleListV1Alpha1_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-22.k8s.KubeClusterRoleListV1Alpha1",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Renders a Kubernetes manifest for "io.k8s.api.rbac.v1alpha1.ClusterRoleList".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func KubeClusterRoleListV1Alpha1_Manifest(props *KubeClusterRoleListV1Alpha1Props) interface{} {
	_init_.Initialize()

	if err := validateKubeClusterRoleListV1Alpha1_ManifestParameters(props); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.StaticInvoke(
		"cdk8s-plus-22.k8s.KubeClusterRoleListV1Alpha1",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
func KubeClusterRoleListV1Alpha1_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	if err := validateKubeClusterRoleListV1Alpha1_OfParameters(c); err != nil {
		panic(err)
	}
	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"cdk8s-plus-22.k8s.KubeClusterRoleListV1Alpha1",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func KubeClusterRoleListV1Alpha1_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"cdk8s-plus-22.k8s.KubeClusterRoleListV1Alpha1",
		"GVK",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KubeClusterRoleListV1Alpha1) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		k,
		"addDependency",
		args,
	)
}

func (k *jsiiProxy_KubeClusterRoleListV1Alpha1) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		k,
		"addJsonPatch",
		args,
	)
}

func (k *jsiiProxy_KubeClusterRoleListV1Alpha1) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubeClusterRoleListV1Alpha1) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

