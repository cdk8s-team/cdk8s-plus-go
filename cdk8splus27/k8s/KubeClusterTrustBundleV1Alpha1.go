package k8s

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus27/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus27/v2/k8s/internal"
)

// ClusterTrustBundle is a cluster-scoped container for X.509 trust anchors (root certificates).
//
// ClusterTrustBundle objects are considered to be readable by any authenticated user in the cluster, because they can be mounted by pods using the `clusterTrustBundle` projection.  All service accounts have read access to ClusterTrustBundles by default.  Users who only have namespace-level access to a cluster can read ClusterTrustBundles by impersonating a serviceaccount that they have access to.
//
// It can be optionally associated with a particular assigner, in which case it contains one valid set of trust anchors for that signer. Signers may have multiple associated ClusterTrustBundles; each is an independent set of trust anchors for that signer. Admission control is used to enforce that only users with permissions on the signer can create or modify the corresponding bundle.
type KubeClusterTrustBundleV1Alpha1 interface {
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

// The jsii proxy struct for KubeClusterTrustBundleV1Alpha1
type jsiiProxy_KubeClusterTrustBundleV1Alpha1 struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_KubeClusterTrustBundleV1Alpha1) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeClusterTrustBundleV1Alpha1) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeClusterTrustBundleV1Alpha1) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeClusterTrustBundleV1Alpha1) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeClusterTrustBundleV1Alpha1) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeClusterTrustBundleV1Alpha1) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeClusterTrustBundleV1Alpha1) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Defines a "io.k8s.api.certificates.v1alpha1.ClusterTrustBundle" API object.
func NewKubeClusterTrustBundleV1Alpha1(scope constructs.Construct, id *string, props *KubeClusterTrustBundleV1Alpha1Props) KubeClusterTrustBundleV1Alpha1 {
	_init_.Initialize()

	if err := validateNewKubeClusterTrustBundleV1Alpha1Parameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubeClusterTrustBundleV1Alpha1{}

	_jsii_.Create(
		"cdk8s-plus-27.k8s.KubeClusterTrustBundleV1Alpha1",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "io.k8s.api.certificates.v1alpha1.ClusterTrustBundle" API object.
func NewKubeClusterTrustBundleV1Alpha1_Override(k KubeClusterTrustBundleV1Alpha1, scope constructs.Construct, id *string, props *KubeClusterTrustBundleV1Alpha1Props) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-27.k8s.KubeClusterTrustBundleV1Alpha1",
		[]interface{}{scope, id, props},
		k,
	)
}

// Return whether the given object is an `ApiObject`.
//
// We do attribute detection since we can't reliably use 'instanceof'.
func KubeClusterTrustBundleV1Alpha1_IsApiObject(o interface{}) *bool {
	_init_.Initialize()

	if err := validateKubeClusterTrustBundleV1Alpha1_IsApiObjectParameters(o); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-27.k8s.KubeClusterTrustBundleV1Alpha1",
		"isApiObject",
		[]interface{}{o},
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
func KubeClusterTrustBundleV1Alpha1_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubeClusterTrustBundleV1Alpha1_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-27.k8s.KubeClusterTrustBundleV1Alpha1",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Renders a Kubernetes manifest for "io.k8s.api.certificates.v1alpha1.ClusterTrustBundle".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func KubeClusterTrustBundleV1Alpha1_Manifest(props *KubeClusterTrustBundleV1Alpha1Props) interface{} {
	_init_.Initialize()

	if err := validateKubeClusterTrustBundleV1Alpha1_ManifestParameters(props); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.StaticInvoke(
		"cdk8s-plus-27.k8s.KubeClusterTrustBundleV1Alpha1",
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
func KubeClusterTrustBundleV1Alpha1_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	if err := validateKubeClusterTrustBundleV1Alpha1_OfParameters(c); err != nil {
		panic(err)
	}
	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"cdk8s-plus-27.k8s.KubeClusterTrustBundleV1Alpha1",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func KubeClusterTrustBundleV1Alpha1_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"cdk8s-plus-27.k8s.KubeClusterTrustBundleV1Alpha1",
		"GVK",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KubeClusterTrustBundleV1Alpha1) AddDependency(dependencies ...constructs.IConstruct) {
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

func (k *jsiiProxy_KubeClusterTrustBundleV1Alpha1) AddJsonPatch(ops ...cdk8s.JsonPatch) {
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

func (k *jsiiProxy_KubeClusterTrustBundleV1Alpha1) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubeClusterTrustBundleV1Alpha1) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

