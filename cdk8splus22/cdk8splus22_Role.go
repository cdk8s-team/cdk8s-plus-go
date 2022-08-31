// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-22 synthesizes Kubernetes manifests for Kubernetes 1.22.0
package cdk8splus22

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus22/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Role is a namespaced, logical grouping of PolicyRules that can be referenced as a unit by a RoleBinding.
type Role interface {
	Resource
	IRole
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
	// Rules associaated with this Role.
	//
	// Returns a copy, use `allow` to add rules.
	Rules() *[]*RolePolicyRule
	// Add permission to perform a list of HTTP verbs on a collection of resources.
	// See: https://kubernetes.io/docs/reference/access-authn-authz/authorization/#determine-the-request-verb
	//
	Allow(verbs *[]*string, resources ...IApiResource)
	// Add "create" permission for the resources.
	AllowCreate(resources ...IApiResource)
	// Add "delete" permission for the resources.
	AllowDelete(resources ...IApiResource)
	// Add "deletecollection" permission for the resources.
	AllowDeleteCollection(resources ...IApiResource)
	// Add "get" permission for the resources.
	AllowGet(resources ...IApiResource)
	// Add "list" permission for the resources.
	AllowList(resources ...IApiResource)
	// Add "patch" permission for the resources.
	AllowPatch(resources ...IApiResource)
	// Add "get", "list", and "watch" permissions for the resources.
	AllowRead(resources ...IApiResource)
	// Add "get", "list", "watch", "create", "update", "patch", "delete", and "deletecollection" permissions for the resources.
	AllowReadWrite(resources ...IApiResource)
	// Add "update" permission for the resources.
	AllowUpdate(resources ...IApiResource)
	// Add "watch" permission for the resources.
	AllowWatch(resources ...IApiResource)
	// Return the IApiResource this object represents.
	AsApiResource() IApiResource
	// Return the non resource url this object represents.
	AsNonApiResource() *string
	// Create a RoleBinding that binds the permissions in this Role to a list of subjects, that will only apply this role's namespace.
	Bind(subjects ...ISubject) RoleBinding
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Role
type jsiiProxy_Role struct {
	jsiiProxy_Resource
	jsiiProxy_IRole
}

func (j *jsiiProxy_Role) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) Permissions() ResourcePermissions {
	var returns ResourcePermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) Rules() *[]*RolePolicyRule {
	var returns *[]*RolePolicyRule
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}


func NewRole(scope constructs.Construct, id *string, props *RoleProps) Role {
	_init_.Initialize()

	if err := validateNewRoleParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Role{}

	_jsii_.Create(
		"cdk8s-plus-22.Role",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewRole_Override(r Role, scope constructs.Construct, id *string, props *RoleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-22.Role",
		[]interface{}{scope, id, props},
		r,
	)
}

// Imports a role from the cluster as a reference.
func Role_FromRoleName(scope constructs.Construct, id *string, name *string) IRole {
	_init_.Initialize()

	if err := validateRole_FromRoleNameParameters(scope, id, name); err != nil {
		panic(err)
	}
	var returns IRole

	_jsii_.StaticInvoke(
		"cdk8s-plus-22.Role",
		"fromRoleName",
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
func Role_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRole_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-22.Role",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Role) Allow(verbs *[]*string, resources ...IApiResource) {
	if err := r.validateAllowParameters(verbs); err != nil {
		panic(err)
	}
	args := []interface{}{verbs}
	for _, a := range resources {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"allow",
		args,
	)
}

func (r *jsiiProxy_Role) AllowCreate(resources ...IApiResource) {
	args := []interface{}{}
	for _, a := range resources {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"allowCreate",
		args,
	)
}

func (r *jsiiProxy_Role) AllowDelete(resources ...IApiResource) {
	args := []interface{}{}
	for _, a := range resources {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"allowDelete",
		args,
	)
}

func (r *jsiiProxy_Role) AllowDeleteCollection(resources ...IApiResource) {
	args := []interface{}{}
	for _, a := range resources {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"allowDeleteCollection",
		args,
	)
}

func (r *jsiiProxy_Role) AllowGet(resources ...IApiResource) {
	args := []interface{}{}
	for _, a := range resources {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"allowGet",
		args,
	)
}

func (r *jsiiProxy_Role) AllowList(resources ...IApiResource) {
	args := []interface{}{}
	for _, a := range resources {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"allowList",
		args,
	)
}

func (r *jsiiProxy_Role) AllowPatch(resources ...IApiResource) {
	args := []interface{}{}
	for _, a := range resources {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"allowPatch",
		args,
	)
}

func (r *jsiiProxy_Role) AllowRead(resources ...IApiResource) {
	args := []interface{}{}
	for _, a := range resources {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"allowRead",
		args,
	)
}

func (r *jsiiProxy_Role) AllowReadWrite(resources ...IApiResource) {
	args := []interface{}{}
	for _, a := range resources {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"allowReadWrite",
		args,
	)
}

func (r *jsiiProxy_Role) AllowUpdate(resources ...IApiResource) {
	args := []interface{}{}
	for _, a := range resources {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"allowUpdate",
		args,
	)
}

func (r *jsiiProxy_Role) AllowWatch(resources ...IApiResource) {
	args := []interface{}{}
	for _, a := range resources {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"allowWatch",
		args,
	)
}

func (r *jsiiProxy_Role) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		r,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Role) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Role) Bind(subjects ...ISubject) RoleBinding {
	args := []interface{}{}
	for _, a := range subjects {
		args = append(args, a)
	}

	var returns RoleBinding

	_jsii_.Invoke(
		r,
		"bind",
		args,
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Role) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

