package cdk8splus25

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus25/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
type AzureDiskPersistentVolume interface {
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
	// Azure kind of this volume.
	AzureKind() AzureDiskPersistentVolumeKind
	// Caching mode of this volume.
	CachingMode() AzureDiskPersistentVolumeCachingMode
	// PVC this volume is bound to.
	//
	// Undefined means this volume is not yet
	// claimed by any PVC.
	Claim() IPersistentVolumeClaim
	// Disk name of this volume.
	DiskName() *string
	// Disk URI of this volume.
	DiskUri() *string
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

// The jsii proxy struct for AzureDiskPersistentVolume
type jsiiProxy_AzureDiskPersistentVolume struct {
	jsiiProxy_PersistentVolume
}

func (j *jsiiProxy_AzureDiskPersistentVolume) AccessModes() *[]PersistentVolumeAccessMode {
	var returns *[]PersistentVolumeAccessMode
	_jsii_.Get(
		j,
		"accessModes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureDiskPersistentVolume) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureDiskPersistentVolume) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureDiskPersistentVolume) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureDiskPersistentVolume) AzureKind() AzureDiskPersistentVolumeKind {
	var returns AzureDiskPersistentVolumeKind
	_jsii_.Get(
		j,
		"azureKind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureDiskPersistentVolume) CachingMode() AzureDiskPersistentVolumeCachingMode {
	var returns AzureDiskPersistentVolumeCachingMode
	_jsii_.Get(
		j,
		"cachingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureDiskPersistentVolume) Claim() IPersistentVolumeClaim {
	var returns IPersistentVolumeClaim
	_jsii_.Get(
		j,
		"claim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureDiskPersistentVolume) DiskName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureDiskPersistentVolume) DiskUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureDiskPersistentVolume) FsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fsType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureDiskPersistentVolume) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureDiskPersistentVolume) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureDiskPersistentVolume) Mode() PersistentVolumeMode {
	var returns PersistentVolumeMode
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureDiskPersistentVolume) MountOptions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"mountOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureDiskPersistentVolume) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureDiskPersistentVolume) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureDiskPersistentVolume) Permissions() ResourcePermissions {
	var returns ResourcePermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureDiskPersistentVolume) ReadOnly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureDiskPersistentVolume) ReclaimPolicy() PersistentVolumeReclaimPolicy {
	var returns PersistentVolumeReclaimPolicy
	_jsii_.Get(
		j,
		"reclaimPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureDiskPersistentVolume) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureDiskPersistentVolume) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureDiskPersistentVolume) Storage() cdk8s.Size {
	var returns cdk8s.Size
	_jsii_.Get(
		j,
		"storage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureDiskPersistentVolume) StorageClassName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClassName",
		&returns,
	)
	return returns
}


func NewAzureDiskPersistentVolume(scope constructs.Construct, id *string, props *AzureDiskPersistentVolumeProps) AzureDiskPersistentVolume {
	_init_.Initialize()

	if err := validateNewAzureDiskPersistentVolumeParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AzureDiskPersistentVolume{}

	_jsii_.Create(
		"cdk8s-plus-25.AzureDiskPersistentVolume",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewAzureDiskPersistentVolume_Override(a AzureDiskPersistentVolume, scope constructs.Construct, id *string, props *AzureDiskPersistentVolumeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-25.AzureDiskPersistentVolume",
		[]interface{}{scope, id, props},
		a,
	)
}

// Imports a pv from the cluster as a reference.
func AzureDiskPersistentVolume_FromPersistentVolumeName(scope constructs.Construct, id *string, volumeName *string) IPersistentVolume {
	_init_.Initialize()

	if err := validateAzureDiskPersistentVolume_FromPersistentVolumeNameParameters(scope, id, volumeName); err != nil {
		panic(err)
	}
	var returns IPersistentVolume

	_jsii_.StaticInvoke(
		"cdk8s-plus-25.AzureDiskPersistentVolume",
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
func AzureDiskPersistentVolume_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAzureDiskPersistentVolume_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-25.AzureDiskPersistentVolume",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureDiskPersistentVolume) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		a,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureDiskPersistentVolume) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureDiskPersistentVolume) AsVolume() Volume {
	var returns Volume

	_jsii_.Invoke(
		a,
		"asVolume",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureDiskPersistentVolume) Bind(claim IPersistentVolumeClaim) {
	if err := a.validateBindParameters(claim); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"bind",
		[]interface{}{claim},
	)
}

func (a *jsiiProxy_AzureDiskPersistentVolume) Reserve() PersistentVolumeClaim {
	var returns PersistentVolumeClaim

	_jsii_.Invoke(
		a,
		"reserve",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureDiskPersistentVolume) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

