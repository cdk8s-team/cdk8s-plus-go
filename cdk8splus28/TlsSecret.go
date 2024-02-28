package cdk8splus28

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus28/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Create a secret for storing a TLS certificate and its associated key.
// See: https://kubernetes.io/docs/concepts/configuration/secret/#tls-secrets
//
type TlsSecret interface {
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
	// Returns EnvValue object from a secret's key.
	EnvValue(key *string, options *EnvValueFromSecretOptions) EnvValue
	// Gets a string data by key or undefined.
	GetStringData(key *string) *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for TlsSecret
type jsiiProxy_TlsSecret struct {
	jsiiProxy_Secret
}

func (j *jsiiProxy_TlsSecret) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TlsSecret) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TlsSecret) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TlsSecret) Immutable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"immutable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TlsSecret) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TlsSecret) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TlsSecret) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TlsSecret) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TlsSecret) Permissions() ResourcePermissions {
	var returns ResourcePermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TlsSecret) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TlsSecret) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}


func NewTlsSecret(scope constructs.Construct, id *string, props *TlsSecretProps) TlsSecret {
	_init_.Initialize()

	if err := validateNewTlsSecretParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_TlsSecret{}

	_jsii_.Create(
		"cdk8s-plus-28.TlsSecret",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewTlsSecret_Override(t TlsSecret, scope constructs.Construct, id *string, props *TlsSecretProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-28.TlsSecret",
		[]interface{}{scope, id, props},
		t,
	)
}

// Imports a secret from the cluster as a reference.
func TlsSecret_FromSecretName(scope constructs.Construct, id *string, name *string) ISecret {
	_init_.Initialize()

	if err := validateTlsSecret_FromSecretNameParameters(scope, id, name); err != nil {
		panic(err)
	}
	var returns ISecret

	_jsii_.StaticInvoke(
		"cdk8s-plus-28.TlsSecret",
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
func TlsSecret_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTlsSecret_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-28.TlsSecret",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TlsSecret) AddStringData(key *string, value *string) {
	if err := t.validateAddStringDataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addStringData",
		[]interface{}{key, value},
	)
}

func (t *jsiiProxy_TlsSecret) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		t,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TlsSecret) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TlsSecret) EnvValue(key *string, options *EnvValueFromSecretOptions) EnvValue {
	if err := t.validateEnvValueParameters(key, options); err != nil {
		panic(err)
	}
	var returns EnvValue

	_jsii_.Invoke(
		t,
		"envValue",
		[]interface{}{key, options},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TlsSecret) GetStringData(key *string) *string {
	if err := t.validateGetStringDataParameters(key); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringData",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TlsSecret) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

