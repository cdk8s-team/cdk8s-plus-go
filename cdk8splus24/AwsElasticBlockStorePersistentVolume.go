// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus24/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod.
// See: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
//
type AwsElasticBlockStorePersistentVolume interface {
	PersistentVolume
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
	// PVC this volume is bound to.
	//
	// Undefined means this volume is not yet
	// claimed by any PVC.
	Claim() IPersistentVolumeClaim
	// File system type of this volume.
	FsType() *string
	// The object kind (e.g. "Deployment").
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	// Volume mode of this volume.
	Mode() PersistentVolumeMode
	// Mount options of this volume.
	MountOptions() *[]*string
	// The name of this API object.
	Name() *string
	// The tree node.
	Node() constructs.Node
	// Partition of this volume.
	Partition() *float64
	Permissions() ResourcePermissions
	// Whether or not it is mounted as a read-only volume.
	ReadOnly() *bool
	// Reclaim policy of this volume.
	ReclaimPolicy() PersistentVolumeReclaimPolicy
	// The unique, namespace-global, name of an object inside the Kubernetes cluster.
	//
	// If this is omitted, the ApiResource should represent all objects of the given type.
	ResourceName() *string
	// The name of a resource type as it appears in the relevant API endpoint.
	ResourceType() *string
	// Storage size of this volume.
	Storage() cdk8s.Size
	// Storage class this volume belongs to.
	StorageClassName() *string
	// Volume id of this volume.
	VolumeId() *string
	// Return the IApiResource this object represents.
	AsApiResource() IApiResource
	// Return the non resource url this object represents.
	AsNonApiResource() *string
	// Convert the piece of storage into a concrete volume.
	AsVolume() Volume
	// Bind a volume to a specific claim.
	//
	// Note that you must also bind the claim to the volume.
	// See: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#binding
	//
	Bind(claim IPersistentVolumeClaim)
	// Reserve a `PersistentVolume` by creating a `PersistentVolumeClaim` that is wired to claim this volume.
	//
	// Note that this method will throw in case the volume is already claimed.
	// See: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#reserving-a-persistentvolume
	//
	Reserve() PersistentVolumeClaim
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for AwsElasticBlockStorePersistentVolume
type jsiiProxy_AwsElasticBlockStorePersistentVolume struct {
	jsiiProxy_PersistentVolume
}

func (j *jsiiProxy_AwsElasticBlockStorePersistentVolume) AccessModes() *[]PersistentVolumeAccessMode {
	var returns *[]PersistentVolumeAccessMode
	_jsii_.Get(
		j,
		"accessModes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticBlockStorePersistentVolume) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticBlockStorePersistentVolume) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticBlockStorePersistentVolume) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticBlockStorePersistentVolume) Claim() IPersistentVolumeClaim {
	var returns IPersistentVolumeClaim
	_jsii_.Get(
		j,
		"claim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticBlockStorePersistentVolume) FsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fsType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticBlockStorePersistentVolume) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticBlockStorePersistentVolume) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticBlockStorePersistentVolume) Mode() PersistentVolumeMode {
	var returns PersistentVolumeMode
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticBlockStorePersistentVolume) MountOptions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"mountOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticBlockStorePersistentVolume) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticBlockStorePersistentVolume) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticBlockStorePersistentVolume) Partition() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"partition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticBlockStorePersistentVolume) Permissions() ResourcePermissions {
	var returns ResourcePermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticBlockStorePersistentVolume) ReadOnly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticBlockStorePersistentVolume) ReclaimPolicy() PersistentVolumeReclaimPolicy {
	var returns PersistentVolumeReclaimPolicy
	_jsii_.Get(
		j,
		"reclaimPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticBlockStorePersistentVolume) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticBlockStorePersistentVolume) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticBlockStorePersistentVolume) Storage() cdk8s.Size {
	var returns cdk8s.Size
	_jsii_.Get(
		j,
		"storage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticBlockStorePersistentVolume) StorageClassName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClassName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticBlockStorePersistentVolume) VolumeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeId",
		&returns,
	)
	return returns
}


func NewAwsElasticBlockStorePersistentVolume(scope constructs.Construct, id *string, props *AwsElasticBlockStorePersistentVolumeProps) AwsElasticBlockStorePersistentVolume {
	_init_.Initialize()

	if err := validateNewAwsElasticBlockStorePersistentVolumeParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsElasticBlockStorePersistentVolume{}

	_jsii_.Create(
		"cdk8s-plus-24.AwsElasticBlockStorePersistentVolume",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewAwsElasticBlockStorePersistentVolume_Override(a AwsElasticBlockStorePersistentVolume, scope constructs.Construct, id *string, props *AwsElasticBlockStorePersistentVolumeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.AwsElasticBlockStorePersistentVolume",
		[]interface{}{scope, id, props},
		a,
	)
}

// Imports a pv from the cluster as a reference.
func AwsElasticBlockStorePersistentVolume_FromPersistentVolumeName(scope constructs.Construct, id *string, volumeName *string) IPersistentVolume {
	_init_.Initialize()

	if err := validateAwsElasticBlockStorePersistentVolume_FromPersistentVolumeNameParameters(scope, id, volumeName); err != nil {
		panic(err)
	}
	var returns IPersistentVolume

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.AwsElasticBlockStorePersistentVolume",
		"fromPersistentVolumeName",
		[]interface{}{scope, id, volumeName},
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
func AwsElasticBlockStorePersistentVolume_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsElasticBlockStorePersistentVolume_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.AwsElasticBlockStorePersistentVolume",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsElasticBlockStorePersistentVolume) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		a,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsElasticBlockStorePersistentVolume) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsElasticBlockStorePersistentVolume) AsVolume() Volume {
	var returns Volume

	_jsii_.Invoke(
		a,
		"asVolume",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsElasticBlockStorePersistentVolume) Bind(claim IPersistentVolumeClaim) {
	if err := a.validateBindParameters(claim); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"bind",
		[]interface{}{claim},
	)
}

func (a *jsiiProxy_AwsElasticBlockStorePersistentVolume) Reserve() PersistentVolumeClaim {
	var returns PersistentVolumeClaim

	_jsii_.Invoke(
		a,
		"reserve",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsElasticBlockStorePersistentVolume) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

