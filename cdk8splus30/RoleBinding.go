package cdk8splus30

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// A RoleBinding grants permissions within a specific namespace to a user or set of users.
type RoleBinding interface {
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
	Role() IRole
	Subjects() *[]ISubject
	// Adds a subject to the role.
	AddSubjects(subjects ...ISubject)
	// Return the IApiResource this object represents.
	AsApiResource() IApiResource
	// Return the non resource url this object represents.
	AsNonApiResource() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for RoleBinding
type jsiiProxy_RoleBinding struct {
	jsiiProxy_Resource
}

func (j *jsiiProxy_RoleBinding) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleBinding) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleBinding) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleBinding) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleBinding) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleBinding) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleBinding) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleBinding) Permissions() ResourcePermissions {
	var returns ResourcePermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleBinding) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleBinding) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleBinding) Role() IRole {
	var returns IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleBinding) Subjects() *[]ISubject {
	var returns *[]ISubject
	_jsii_.Get(
		j,
		"subjects",
		&returns,
	)
	return returns
}


func NewRoleBinding(scope constructs.Construct, id *string, props *RoleBindingProps) RoleBinding {
	_init_.Initialize()

	if err := validateNewRoleBindingParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_RoleBinding{}

	_jsii_.Create(
		"cdk8s-plus-30.RoleBinding",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewRoleBinding_Override(r RoleBinding, scope constructs.Construct, id *string, props *RoleBindingProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-30.RoleBinding",
		[]interface{}{scope, id, props},
		r,
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
func RoleBinding_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRoleBinding_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-30.RoleBinding",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RoleBinding) AddSubjects(subjects ...ISubject) {
	args := []interface{}{}
	for _, a := range subjects {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"addSubjects",
		args,
	)
}

func (r *jsiiProxy_RoleBinding) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		r,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RoleBinding) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RoleBinding) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

