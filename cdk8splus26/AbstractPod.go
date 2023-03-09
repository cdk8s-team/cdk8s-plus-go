// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-26 synthesizes Kubernetes manifests for Kubernetes 1.26.0
package cdk8splus26

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus26/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type AbstractPod interface {
	Resource
	INetworkPolicyPeer
	IPodSelector
	ISubject
	// The group portion of the API version (e.g. "authorization.k8s.io").
	ApiGroup() *string
	// The underlying cdk8s API object.
	ApiObject() cdk8s.ApiObject
	// The object's API version (e.g. "authorization.k8s.io/v1").
	ApiVersion() *string
	AutomountServiceAccountToken() *bool
	Containers() *[]Container
	Dns() PodDns
	DockerRegistryAuth() ISecret
	HostAliases() *[]*HostAlias
	InitContainers() *[]Container
	Isolate() *bool
	// The object kind (e.g. "Deployment").
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	// The name of this API object.
	Name() *string
	// The tree node.
	Node() constructs.Node
	Permissions() ResourcePermissions
	PodMetadata() cdk8s.ApiObjectMetadataDefinition
	// The unique, namespace-global, name of an object inside the Kubernetes cluster.
	//
	// If this is omitted, the ApiResource should represent all objects of the given type.
	ResourceName() *string
	// The name of a resource type as it appears in the relevant API endpoint.
	ResourceType() *string
	RestartPolicy() RestartPolicy
	SecurityContext() PodSecurityContext
	ServiceAccount() IServiceAccount
	Volumes() *[]Volume
	AddContainer(cont *ContainerProps) Container
	AddHostAlias(hostAlias *HostAlias)
	AddInitContainer(cont *ContainerProps) Container
	AddVolume(vol Volume)
	// Return the IApiResource this object represents.
	AsApiResource() IApiResource
	// Return the non resource url this object represents.
	AsNonApiResource() *string
	AttachContainer(cont Container)
	// Return the configuration of this peer.
	// See: INetworkPolicyPeer.toNetworkPolicyPeerConfig()
	//
	ToNetworkPolicyPeerConfig() *NetworkPolicyPeerConfig
	// Convert the peer into a pod selector, if possible.
	// See: INetworkPolicyPeer.toPodSelector()
	//
	ToPodSelector() IPodSelector
	// Return the configuration of this selector.
	// See: IPodSelector.toPodSelectorConfig()
	//
	ToPodSelectorConfig() *PodSelectorConfig
	// Returns a string representation of this construct.
	ToString() *string
	// Return the subject configuration.
	// See: ISubect.toSubjectConfiguration()
	//
	ToSubjectConfiguration() *SubjectConfiguration
}

// The jsii proxy struct for AbstractPod
type jsiiProxy_AbstractPod struct {
	jsiiProxy_Resource
	jsiiProxy_INetworkPolicyPeer
	jsiiProxy_IPodSelector
	jsiiProxy_ISubject
}

func (j *jsiiProxy_AbstractPod) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AbstractPod) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AbstractPod) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AbstractPod) AutomountServiceAccountToken() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"automountServiceAccountToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AbstractPod) Containers() *[]Container {
	var returns *[]Container
	_jsii_.Get(
		j,
		"containers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AbstractPod) Dns() PodDns {
	var returns PodDns
	_jsii_.Get(
		j,
		"dns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AbstractPod) DockerRegistryAuth() ISecret {
	var returns ISecret
	_jsii_.Get(
		j,
		"dockerRegistryAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AbstractPod) HostAliases() *[]*HostAlias {
	var returns *[]*HostAlias
	_jsii_.Get(
		j,
		"hostAliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AbstractPod) InitContainers() *[]Container {
	var returns *[]Container
	_jsii_.Get(
		j,
		"initContainers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AbstractPod) Isolate() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isolate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AbstractPod) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AbstractPod) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AbstractPod) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AbstractPod) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AbstractPod) Permissions() ResourcePermissions {
	var returns ResourcePermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AbstractPod) PodMetadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"podMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AbstractPod) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AbstractPod) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AbstractPod) RestartPolicy() RestartPolicy {
	var returns RestartPolicy
	_jsii_.Get(
		j,
		"restartPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AbstractPod) SecurityContext() PodSecurityContext {
	var returns PodSecurityContext
	_jsii_.Get(
		j,
		"securityContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AbstractPod) ServiceAccount() IServiceAccount {
	var returns IServiceAccount
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AbstractPod) Volumes() *[]Volume {
	var returns *[]Volume
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}


func NewAbstractPod_Override(a AbstractPod, scope constructs.Construct, id *string, props *AbstractPodProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-26.AbstractPod",
		[]interface{}{scope, id, props},
		a,
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
func AbstractPod_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAbstractPod_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-26.AbstractPod",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AbstractPod) AddContainer(cont *ContainerProps) Container {
	if err := a.validateAddContainerParameters(cont); err != nil {
		panic(err)
	}
	var returns Container

	_jsii_.Invoke(
		a,
		"addContainer",
		[]interface{}{cont},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AbstractPod) AddHostAlias(hostAlias *HostAlias) {
	if err := a.validateAddHostAliasParameters(hostAlias); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addHostAlias",
		[]interface{}{hostAlias},
	)
}

func (a *jsiiProxy_AbstractPod) AddInitContainer(cont *ContainerProps) Container {
	if err := a.validateAddInitContainerParameters(cont); err != nil {
		panic(err)
	}
	var returns Container

	_jsii_.Invoke(
		a,
		"addInitContainer",
		[]interface{}{cont},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AbstractPod) AddVolume(vol Volume) {
	if err := a.validateAddVolumeParameters(vol); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addVolume",
		[]interface{}{vol},
	)
}

func (a *jsiiProxy_AbstractPod) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		a,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AbstractPod) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AbstractPod) AttachContainer(cont Container) {
	if err := a.validateAttachContainerParameters(cont); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"attachContainer",
		[]interface{}{cont},
	)
}

func (a *jsiiProxy_AbstractPod) ToNetworkPolicyPeerConfig() *NetworkPolicyPeerConfig {
	var returns *NetworkPolicyPeerConfig

	_jsii_.Invoke(
		a,
		"toNetworkPolicyPeerConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AbstractPod) ToPodSelector() IPodSelector {
	var returns IPodSelector

	_jsii_.Invoke(
		a,
		"toPodSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AbstractPod) ToPodSelectorConfig() *PodSelectorConfig {
	var returns *PodSelectorConfig

	_jsii_.Invoke(
		a,
		"toPodSelectorConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AbstractPod) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AbstractPod) ToSubjectConfiguration() *SubjectConfiguration {
	var returns *SubjectConfiguration

	_jsii_.Invoke(
		a,
		"toSubjectConfiguration",
		nil, // no parameters
		&returns,
	)

	return returns
}

