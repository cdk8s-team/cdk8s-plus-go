package cdk8splus29

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus29/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Ingress is a collection of rules that allow inbound connections to reach the endpoints defined by a backend.
//
// An Ingress can be configured to give services
// externally-reachable urls, load balance traffic, terminate SSL, offer name
// based virtual hosting etc.
type Ingress interface {
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
	// Defines the default backend for this ingress.
	//
	// A default backend capable of
	// servicing requests that don't match any rule.
	AddDefaultBackend(backend IngressBackend)
	// Specify a default backend for a specific host name.
	//
	// This backend will be used as a catch-all for requests
	// targeted to this host name (the `Host` header matches this value).
	AddHostDefaultBackend(host *string, backend IngressBackend)
	// Adds an ingress rule applied to requests to a specific host and a specific HTTP path (the `Host` header matches this value).
	AddHostRule(host *string, path *string, backend IngressBackend, pathType HttpIngressPathType)
	// Adds an ingress rule applied to requests sent to a specific HTTP path.
	AddRule(path *string, backend IngressBackend, pathType HttpIngressPathType)
	// Adds rules to this ingress.
	AddRules(rules ...*IngressRule)
	AddTls(tls *[]*IngressTls)
	// Return the IApiResource this object represents.
	AsApiResource() IApiResource
	// Return the non resource url this object represents.
	AsNonApiResource() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Ingress
type jsiiProxy_Ingress struct {
	jsiiProxy_Resource
}

func (j *jsiiProxy_Ingress) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ingress) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ingress) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ingress) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ingress) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ingress) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ingress) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ingress) Permissions() ResourcePermissions {
	var returns ResourcePermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ingress) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ingress) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}


func NewIngress(scope constructs.Construct, id *string, props *IngressProps) Ingress {
	_init_.Initialize()

	if err := validateNewIngressParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ingress{}

	_jsii_.Create(
		"cdk8s-plus-29.Ingress",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewIngress_Override(i Ingress, scope constructs.Construct, id *string, props *IngressProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-29.Ingress",
		[]interface{}{scope, id, props},
		i,
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
func Ingress_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIngress_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-29.Ingress",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Ingress) AddDefaultBackend(backend IngressBackend) {
	if err := i.validateAddDefaultBackendParameters(backend); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addDefaultBackend",
		[]interface{}{backend},
	)
}

func (i *jsiiProxy_Ingress) AddHostDefaultBackend(host *string, backend IngressBackend) {
	if err := i.validateAddHostDefaultBackendParameters(host, backend); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addHostDefaultBackend",
		[]interface{}{host, backend},
	)
}

func (i *jsiiProxy_Ingress) AddHostRule(host *string, path *string, backend IngressBackend, pathType HttpIngressPathType) {
	if err := i.validateAddHostRuleParameters(host, path, backend); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addHostRule",
		[]interface{}{host, path, backend, pathType},
	)
}

func (i *jsiiProxy_Ingress) AddRule(path *string, backend IngressBackend, pathType HttpIngressPathType) {
	if err := i.validateAddRuleParameters(path, backend); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addRule",
		[]interface{}{path, backend, pathType},
	)
}

func (i *jsiiProxy_Ingress) AddRules(rules ...*IngressRule) {
	if err := i.validateAddRulesParameters(&rules); err != nil {
		panic(err)
	}
	args := []interface{}{}
	for _, a := range rules {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		i,
		"addRules",
		args,
	)
}

func (i *jsiiProxy_Ingress) AddTls(tls *[]*IngressTls) {
	if err := i.validateAddTlsParameters(tls); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addTls",
		[]interface{}{tls},
	)
}

func (i *jsiiProxy_Ingress) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		i,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Ingress) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Ingress) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

