package cdk8splus30

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// A PersistentVolumeClaim (PVC) is a request for storage by a user.
//
// It is similar to a Pod. Pods consume node resources and PVCs consume PV resources.
// Pods can request specific levels of resources (CPU and Memory).
// Claims can request specific size and access modes.
type PersistentVolumeClaim interface {
	Resource
	IPersistentVolumeClaim
	// Access modes requirement of this claim.
	AccessModes() *[]PersistentVolumeAccessMode
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
	// Storage requirement of this claim.
	Storage() cdk8s.Size
	// Storage class requirment of this claim.
	StorageClassName() *string
	// PV this claim is bound to.
	//
	// Undefined means the claim is not bound
	// to any specific volume.
	Volume() IPersistentVolume
	// Volume mode requirement of this claim.
	VolumeMode() PersistentVolumeMode
	// Return the IApiResource this object represents.
	AsApiResource() IApiResource
	// Return the non resource url this object represents.
	AsNonApiResource() *string
	// Bind a claim to a specific volume.
	//
	// Note that you must also bind the volume to the claim.
	// See: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#binding
	//
	Bind(vol IPersistentVolume)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for PersistentVolumeClaim
type jsiiProxy_PersistentVolumeClaim struct {
	jsiiProxy_Resource
	jsiiProxy_IPersistentVolumeClaim
}

func (j *jsiiProxy_PersistentVolumeClaim) AccessModes() *[]PersistentVolumeAccessMode {
	var returns *[]PersistentVolumeAccessMode
	_jsii_.Get(
		j,
		"accessModes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeClaim) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeClaim) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeClaim) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeClaim) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeClaim) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeClaim) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeClaim) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeClaim) Permissions() ResourcePermissions {
	var returns ResourcePermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeClaim) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeClaim) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeClaim) Storage() cdk8s.Size {
	var returns cdk8s.Size
	_jsii_.Get(
		j,
		"storage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeClaim) StorageClassName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClassName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeClaim) Volume() IPersistentVolume {
	var returns IPersistentVolume
	_jsii_.Get(
		j,
		"volume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeClaim) VolumeMode() PersistentVolumeMode {
	var returns PersistentVolumeMode
	_jsii_.Get(
		j,
		"volumeMode",
		&returns,
	)
	return returns
}


func NewPersistentVolumeClaim(scope constructs.Construct, id *string, props *PersistentVolumeClaimProps) PersistentVolumeClaim {
	_init_.Initialize()

	if err := validateNewPersistentVolumeClaimParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_PersistentVolumeClaim{}

	_jsii_.Create(
		"cdk8s-plus-30.PersistentVolumeClaim",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewPersistentVolumeClaim_Override(p PersistentVolumeClaim, scope constructs.Construct, id *string, props *PersistentVolumeClaimProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-30.PersistentVolumeClaim",
		[]interface{}{scope, id, props},
		p,
	)
}

// Imports a pvc from the cluster as a reference.
func PersistentVolumeClaim_FromClaimName(scope constructs.Construct, id *string, claimName *string) IPersistentVolumeClaim {
	_init_.Initialize()

	if err := validatePersistentVolumeClaim_FromClaimNameParameters(scope, id, claimName); err != nil {
		panic(err)
	}
	var returns IPersistentVolumeClaim

	_jsii_.StaticInvoke(
		"cdk8s-plus-30.PersistentVolumeClaim",
		"fromClaimName",
		[]interface{}{scope, id, claimName},
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
func PersistentVolumeClaim_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePersistentVolumeClaim_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-30.PersistentVolumeClaim",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeClaim) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		p,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeClaim) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeClaim) Bind(vol IPersistentVolume) {
	if err := p.validateBindParameters(vol); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"bind",
		[]interface{}{vol},
	)
}

func (p *jsiiProxy_PersistentVolumeClaim) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

