// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-22 synthesizes Kubernetes manifests for Kubernetes 1.22.0
package cdk8splus22

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus22/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Create a secret for storing credentials for accessing a container image registry.
// See: https://kubernetes.io/docs/concepts/configuration/secret/#docker-config-secrets
//
type DockerConfigSecret interface {
	Secret
	// The group portion of the API version (e.g. "authorization.k8s.io").
	ApiGroup() *string
	// The underlying cdk8s API object.
	// See: base.Resource.apiObject
	//
	ApiObject() cdk8s.ApiObject
	// The object's API version (e.g. "authorization.k8s.io/v1").
	ApiVersion() *string
	// Whether or not the secret is immutable.
	Immutable() *bool
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
	// Adds a string data field to the secert.
	AddStringData(key *string, value *string)
	// Return the IApiResource this object represents.
	AsApiResource() IApiResource
	// Return the non resource url this object represents.
	AsNonApiResource() *string
	// Gets a string data by key or undefined.
	GetStringData(key *string) *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for DockerConfigSecret
type jsiiProxy_DockerConfigSecret struct {
	jsiiProxy_Secret
}

func (j *jsiiProxy_DockerConfigSecret) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerConfigSecret) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerConfigSecret) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerConfigSecret) Immutable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"immutable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerConfigSecret) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerConfigSecret) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerConfigSecret) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerConfigSecret) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerConfigSecret) Permissions() ResourcePermissions {
	var returns ResourcePermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerConfigSecret) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerConfigSecret) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}


func NewDockerConfigSecret(scope constructs.Construct, id *string, props *DockerConfigSecretProps) DockerConfigSecret {
	_init_.Initialize()

	if err := validateNewDockerConfigSecretParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DockerConfigSecret{}

	_jsii_.Create(
		"cdk8s-plus-22.DockerConfigSecret",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDockerConfigSecret_Override(d DockerConfigSecret, scope constructs.Construct, id *string, props *DockerConfigSecretProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-22.DockerConfigSecret",
		[]interface{}{scope, id, props},
		d,
	)
}

// Imports a secret from the cluster as a reference.
func DockerConfigSecret_FromSecretName(scope constructs.Construct, id *string, name *string) ISecret {
	_init_.Initialize()

	if err := validateDockerConfigSecret_FromSecretNameParameters(scope, id, name); err != nil {
		panic(err)
	}
	var returns ISecret

	_jsii_.StaticInvoke(
		"cdk8s-plus-22.DockerConfigSecret",
		"fromSecretName",
		[]interface{}{scope, id, name},
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
func DockerConfigSecret_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDockerConfigSecret_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-22.DockerConfigSecret",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DockerConfigSecret) AddStringData(key *string, value *string) {
	if err := d.validateAddStringDataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addStringData",
		[]interface{}{key, value},
	)
}

func (d *jsiiProxy_DockerConfigSecret) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		d,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DockerConfigSecret) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DockerConfigSecret) GetStringData(key *string) *string {
	if err := d.validateGetStringDataParameters(key); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringData",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DockerConfigSecret) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

