package cdk8splus31

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus31/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// ClusterRole is a cluster level, logical grouping of PolicyRules that can be referenced as a unit by a RoleBinding or ClusterRoleBinding.
type ClusterRole interface {
	Resource
	IClusterRole
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
	Rules() *[]*ClusterRolePolicyRule
	// Aggregate rules from roles matching this label selector.
	Aggregate(key *string, value *string)
	// Add permission to perform a list of HTTP verbs on a collection of resources.
	// See: https://kubernetes.io/docs/reference/access-authn-authz/authorization/#determine-the-request-verb
	//
	Allow(verbs *[]*string, endpoints ...IApiEndpoint)
	// Add "create" permission for the resources.
	AllowCreate(endpoints ...IApiEndpoint)
	// Add "delete" permission for the resources.
	AllowDelete(endpoints ...IApiEndpoint)
	// Add "deletecollection" permission for the resources.
	AllowDeleteCollection(endpoints ...IApiEndpoint)
	// Add "get" permission for the resources.
	AllowGet(endpoints ...IApiEndpoint)
	// Add "list" permission for the resources.
	AllowList(endpoints ...IApiEndpoint)
	// Add "patch" permission for the resources.
	AllowPatch(endpoints ...IApiEndpoint)
	// Add "get", "list", and "watch" permissions for the resources.
	AllowRead(endpoints ...IApiEndpoint)
	// Add "get", "list", "watch", "create", "update", "patch", "delete", and "deletecollection" permissions for the resources.
	AllowReadWrite(endpoints ...IApiEndpoint)
	// Add "update" permission for the resources.
	AllowUpdate(endpoints ...IApiEndpoint)
	// Add "watch" permission for the resources.
	AllowWatch(endpoints ...IApiEndpoint)
	// Return the IApiResource this object represents.
	AsApiResource() IApiResource
	// Return the non resource url this object represents.
	AsNonApiResource() *string
	// Create a ClusterRoleBinding that binds the permissions in this ClusterRole to a list of subjects, without namespace restrictions.
	Bind(subjects ...ISubject) ClusterRoleBinding
	// Create a RoleBinding that binds the permissions in this ClusterRole to a list of subjects, that will only apply to the given namespace.
	BindInNamespace(namespace *string, subjects ...ISubject) RoleBinding
	// Combines the rules of the argument ClusterRole into this ClusterRole using aggregation labels.
	Combine(rol ClusterRole)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for ClusterRole
type jsiiProxy_ClusterRole struct {
	jsiiProxy_Resource
	jsiiProxy_IClusterRole
	jsiiProxy_IRole
}

func (j *jsiiProxy_ClusterRole) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRole) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRole) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRole) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRole) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRole) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRole) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRole) Permissions() ResourcePermissions {
	var returns ResourcePermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRole) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRole) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRole) Rules() *[]*ClusterRolePolicyRule {
	var returns *[]*ClusterRolePolicyRule
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}


func NewClusterRole(scope constructs.Construct, id *string, props *ClusterRoleProps) ClusterRole {
	_init_.Initialize()

	if err := validateNewClusterRoleParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ClusterRole{}

	_jsii_.Create(
		"cdk8s-plus-31.ClusterRole",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewClusterRole_Override(c ClusterRole, scope constructs.Construct, id *string, props *ClusterRoleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-31.ClusterRole",
		[]interface{}{scope, id, props},
		c,
	)
}

// Imports a role from the cluster as a reference.
func ClusterRole_FromClusterRoleName(scope constructs.Construct, id *string, name *string) IClusterRole {
	_init_.Initialize()

	if err := validateClusterRole_FromClusterRoleNameParameters(scope, id, name); err != nil {
		panic(err)
	}
	var returns IClusterRole

	_jsii_.StaticInvoke(
		"cdk8s-plus-31.ClusterRole",
		"fromClusterRoleName",
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
func ClusterRole_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateClusterRole_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-31.ClusterRole",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterRole) Aggregate(key *string, value *string) {
	if err := c.validateAggregateParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"aggregate",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_ClusterRole) Allow(verbs *[]*string, endpoints ...IApiEndpoint) {
	if err := c.validateAllowParameters(verbs); err != nil {
		panic(err)
	}
	args := []interface{}{verbs}
	for _, a := range endpoints {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"allow",
		args,
	)
}

func (c *jsiiProxy_ClusterRole) AllowCreate(endpoints ...IApiEndpoint) {
	args := []interface{}{}
	for _, a := range endpoints {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"allowCreate",
		args,
	)
}

func (c *jsiiProxy_ClusterRole) AllowDelete(endpoints ...IApiEndpoint) {
	args := []interface{}{}
	for _, a := range endpoints {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"allowDelete",
		args,
	)
}

func (c *jsiiProxy_ClusterRole) AllowDeleteCollection(endpoints ...IApiEndpoint) {
	args := []interface{}{}
	for _, a := range endpoints {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"allowDeleteCollection",
		args,
	)
}

func (c *jsiiProxy_ClusterRole) AllowGet(endpoints ...IApiEndpoint) {
	args := []interface{}{}
	for _, a := range endpoints {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"allowGet",
		args,
	)
}

func (c *jsiiProxy_ClusterRole) AllowList(endpoints ...IApiEndpoint) {
	args := []interface{}{}
	for _, a := range endpoints {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"allowList",
		args,
	)
}

func (c *jsiiProxy_ClusterRole) AllowPatch(endpoints ...IApiEndpoint) {
	args := []interface{}{}
	for _, a := range endpoints {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"allowPatch",
		args,
	)
}

func (c *jsiiProxy_ClusterRole) AllowRead(endpoints ...IApiEndpoint) {
	args := []interface{}{}
	for _, a := range endpoints {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"allowRead",
		args,
	)
}

func (c *jsiiProxy_ClusterRole) AllowReadWrite(endpoints ...IApiEndpoint) {
	args := []interface{}{}
	for _, a := range endpoints {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"allowReadWrite",
		args,
	)
}

func (c *jsiiProxy_ClusterRole) AllowUpdate(endpoints ...IApiEndpoint) {
	args := []interface{}{}
	for _, a := range endpoints {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"allowUpdate",
		args,
	)
}

func (c *jsiiProxy_ClusterRole) AllowWatch(endpoints ...IApiEndpoint) {
	args := []interface{}{}
	for _, a := range endpoints {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"allowWatch",
		args,
	)
}

func (c *jsiiProxy_ClusterRole) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		c,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterRole) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterRole) Bind(subjects ...ISubject) ClusterRoleBinding {
	args := []interface{}{}
	for _, a := range subjects {
		args = append(args, a)
	}

	var returns ClusterRoleBinding

	_jsii_.Invoke(
		c,
		"bind",
		args,
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterRole) BindInNamespace(namespace *string, subjects ...ISubject) RoleBinding {
	if err := c.validateBindInNamespaceParameters(namespace); err != nil {
		panic(err)
	}
	args := []interface{}{namespace}
	for _, a := range subjects {
		args = append(args, a)
	}

	var returns RoleBinding

	_jsii_.Invoke(
		c,
		"bindInNamespace",
		args,
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterRole) Combine(rol ClusterRole) {
	if err := c.validateCombineParameters(rol); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"combine",
		[]interface{}{rol},
	)
}

func (c *jsiiProxy_ClusterRole) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

