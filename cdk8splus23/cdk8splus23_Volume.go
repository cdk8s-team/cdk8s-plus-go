// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-23 synthesizes Kubernetes manifests for Kubernetes 1.23.0
package cdk8splus23

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus23/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus23/v2/internal"
)

// Volume represents a named volume in a pod that may be accessed by any container in the pod.
//
// Docker also has a concept of volumes, though it is somewhat looser and less
// managed. In Docker, a volume is simply a directory on disk or in another
// Container. Lifetimes are not managed and until very recently there were only
// local-disk-backed volumes. Docker now provides volume drivers, but the
// functionality is very limited for now (e.g. as of Docker 1.7 only one volume
// driver is allowed per Container and there is no way to pass parameters to
// volumes).
//
// A Kubernetes volume, on the other hand, has an explicit lifetime - the same
// as the Pod that encloses it. Consequently, a volume outlives any Containers
// that run within the Pod, and data is preserved across Container restarts. Of
// course, when a Pod ceases to exist, the volume will cease to exist, too.
// Perhaps more importantly than this, Kubernetes supports many types of
// volumes, and a Pod can use any number of them simultaneously.
//
// At its core, a volume is just a directory, possibly with some data in it,
// which is accessible to the Containers in a Pod. How that directory comes to
// be, the medium that backs it, and the contents of it are determined by the
// particular volume type used.
//
// To use a volume, a Pod specifies what volumes to provide for the Pod (the
// .spec.volumes field) and where to mount those into Containers (the
// .spec.containers[*].volumeMounts field).
//
// A process in a container sees a filesystem view composed from their Docker
// image and volumes. The Docker image is at the root of the filesystem
// hierarchy, and any volumes are mounted at the specified paths within the
// image. Volumes can not mount onto other volumes
type Volume interface {
	constructs.Construct
	IStorage
	Name() *string
	// The tree node.
	Node() constructs.Node
	// Convert the piece of storage into a concrete volume.
	AsVolume() Volume
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Volume
type jsiiProxy_Volume struct {
	internal.Type__constructsConstruct
	jsiiProxy_IStorage
}

func (j *jsiiProxy_Volume) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Mounts an Amazon Web Services (AWS) EBS volume into your pod.
//
// Unlike emptyDir, which is erased when a pod is removed, the contents of an EBS volume are
// persisted and the volume is unmounted. This means that an EBS volume can be pre-populated with data,
// and that data can be shared between pods.
//
// There are some restrictions when using an awsElasticBlockStore volume:
//
// - the nodes on which pods are running must be AWS EC2 instances.
// - those instances need to be in the same region and availability zone as the EBS volume.
// - EBS only supports a single EC2 instance mounting a volume.
func Volume_FromAwsElasticBlockStore(scope constructs.Construct, id *string, volumeId *string, options *AwsElasticBlockStoreVolumeOptions) Volume {
	_init_.Initialize()

	var returns Volume

	_jsii_.StaticInvoke(
		"cdk8s-plus-23.Volume",
		"fromAwsElasticBlockStore",
		[]interface{}{scope, id, volumeId, options},
		&returns,
	)

	return returns
}

// Mounts a Microsoft Azure Data Disk into a pod.
func Volume_FromAzureDisk(scope constructs.Construct, id *string, diskName *string, diskUri *string, options *AzureDiskVolumeOptions) Volume {
	_init_.Initialize()

	var returns Volume

	_jsii_.StaticInvoke(
		"cdk8s-plus-23.Volume",
		"fromAzureDisk",
		[]interface{}{scope, id, diskName, diskUri, options},
		&returns,
	)

	return returns
}

// Populate the volume from a ConfigMap.
//
// The configMap resource provides a way to inject configuration data into
// Pods. The data stored in a ConfigMap object can be referenced in a volume
// of type configMap and then consumed by containerized applications running
// in a Pod.
//
// When referencing a configMap object, you can simply provide its name in the
// volume to reference it. You can also customize the path to use for a
// specific entry in the ConfigMap.
func Volume_FromConfigMap(scope constructs.Construct, id *string, configMap IConfigMap, options *ConfigMapVolumeOptions) Volume {
	_init_.Initialize()

	var returns Volume

	_jsii_.StaticInvoke(
		"cdk8s-plus-23.Volume",
		"fromConfigMap",
		[]interface{}{scope, id, configMap, options},
		&returns,
	)

	return returns
}

// An emptyDir volume is first created when a Pod is assigned to a Node, and exists as long as that Pod is running on that node.
//
// As the name says, it is
// initially empty. Containers in the Pod can all read and write the same
// files in the emptyDir volume, though that volume can be mounted at the same
// or different paths in each Container. When a Pod is removed from a node for
// any reason, the data in the emptyDir is deleted forever.
// See: http://kubernetes.io/docs/user-guide/volumes#emptydir
//
func Volume_FromEmptyDir(scope constructs.Construct, id *string, name *string, options *EmptyDirVolumeOptions) Volume {
	_init_.Initialize()

	var returns Volume

	_jsii_.StaticInvoke(
		"cdk8s-plus-23.Volume",
		"fromEmptyDir",
		[]interface{}{scope, id, name, options},
		&returns,
	)

	return returns
}

// Mounts a Google Compute Engine (GCE) persistent disk (PD) into your Pod.
//
// Unlike emptyDir, which is erased when a pod is removed, the contents of a PD are
// preserved and the volume is merely unmounted. This means that a PD can be pre-populated
// with data, and that data can be shared between pods.
//
// There are some restrictions when using a gcePersistentDisk:
//
// - the nodes on which Pods are running must be GCE VMs
// - those VMs need to be in the same GCE project and zone as the persistent disk.
func Volume_FromGcePersistentDisk(scope constructs.Construct, id *string, pdName *string, options *GCEPersistentDiskVolumeOptions) Volume {
	_init_.Initialize()

	var returns Volume

	_jsii_.StaticInvoke(
		"cdk8s-plus-23.Volume",
		"fromGcePersistentDisk",
		[]interface{}{scope, id, pdName, options},
		&returns,
	)

	return returns
}

// Used to mount a file or directory from the host node's filesystem into a Pod.
//
// This is not something that most Pods will need, but it offers a powerful
// escape hatch for some applications.
// See: https://kubernetes.io/docs/concepts/storage/volumes/#hostpath
//
func Volume_FromHostPath(scope constructs.Construct, id *string, name *string, options *HostPathVolumeOptions) Volume {
	_init_.Initialize()

	var returns Volume

	_jsii_.StaticInvoke(
		"cdk8s-plus-23.Volume",
		"fromHostPath",
		[]interface{}{scope, id, name, options},
		&returns,
	)

	return returns
}

// Used to mount a PersistentVolume into a Pod.
//
// PersistentVolumeClaims are a way for users to "claim" durable storage (such as a GCE PersistentDisk or an iSCSI volume)
// without knowing the details of the particular cloud environment.
// See: https://kubernetes.io/docs/concepts/storage/persistent-volumes/
//
func Volume_FromPersistentVolumeClaim(scope constructs.Construct, id *string, claim IPersistentVolumeClaim, options *PersistentVolumeClaimVolumeOptions) Volume {
	_init_.Initialize()

	var returns Volume

	_jsii_.StaticInvoke(
		"cdk8s-plus-23.Volume",
		"fromPersistentVolumeClaim",
		[]interface{}{scope, id, claim, options},
		&returns,
	)

	return returns
}

// Populate the volume from a Secret.
//
// A secret volume is used to pass sensitive information, such as passwords, to Pods.
// You can store secrets in the Kubernetes API and mount them as files for use by pods
// without coupling to Kubernetes directly.
//
// secret volumes are backed by tmpfs (a RAM-backed filesystem)
// so they are never written to non-volatile storage.
// See: https://kubernetes.io/docs/concepts/storage/volumes/#secret
//
func Volume_FromSecret(scope constructs.Construct, id *string, secr ISecret, options *SecretVolumeOptions) Volume {
	_init_.Initialize()

	var returns Volume

	_jsii_.StaticInvoke(
		"cdk8s-plus-23.Volume",
		"fromSecret",
		[]interface{}{scope, id, secr, options},
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
func Volume_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-23.Volume",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_Volume) AsVolume() Volume {
	var returns Volume

	_jsii_.Invoke(
		v,
		"asVolume",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_Volume) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

