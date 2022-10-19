// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-22 synthesizes Kubernetes manifests for Kubernetes 1.22.0
package cdk8splus22

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus22/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Pod is a collection of containers that can run on a host.
//
// This resource is
// created by clients and scheduled onto hosts.
type Pod interface {
	AbstractPod
	// The group portion of the API version (e.g. "authorization.k8s.io").
	ApiGroup() *string
	// The underlying cdk8s API object.
	// See: base.Resource.apiObject
	//
	ApiObject() cdk8s.ApiObject
	// The object's API version (e.g. "authorization.k8s.io/v1").
	ApiVersion() *string
	AutomountServiceAccountToken() *bool
	Connections() PodConnections
	Containers() *[]Container
	Dns() PodDns
	DockerRegistryAuth() DockerConfigSecret
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
	Scheduling() PodScheduling
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

// The jsii proxy struct for Pod
type jsiiProxy_Pod struct {
	jsiiProxy_AbstractPod
}

func (j *jsiiProxy_Pod) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pod) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pod) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pod) AutomountServiceAccountToken() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"automountServiceAccountToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pod) Connections() PodConnections {
	var returns PodConnections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pod) Containers() *[]Container {
	var returns *[]Container
	_jsii_.Get(
		j,
		"containers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pod) Dns() PodDns {
	var returns PodDns
	_jsii_.Get(
		j,
		"dns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pod) DockerRegistryAuth() DockerConfigSecret {
	var returns DockerConfigSecret
	_jsii_.Get(
		j,
		"dockerRegistryAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pod) HostAliases() *[]*HostAlias {
	var returns *[]*HostAlias
	_jsii_.Get(
		j,
		"hostAliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pod) InitContainers() *[]Container {
	var returns *[]Container
	_jsii_.Get(
		j,
		"initContainers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pod) Isolate() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isolate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pod) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pod) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pod) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pod) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pod) Permissions() ResourcePermissions {
	var returns ResourcePermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pod) PodMetadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"podMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pod) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pod) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pod) RestartPolicy() RestartPolicy {
	var returns RestartPolicy
	_jsii_.Get(
		j,
		"restartPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pod) Scheduling() PodScheduling {
	var returns PodScheduling
	_jsii_.Get(
		j,
		"scheduling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pod) SecurityContext() PodSecurityContext {
	var returns PodSecurityContext
	_jsii_.Get(
		j,
		"securityContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pod) ServiceAccount() IServiceAccount {
	var returns IServiceAccount
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pod) Volumes() *[]Volume {
	var returns *[]Volume
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}


func NewPod(scope constructs.Construct, id *string, props *PodProps) Pod {
	_init_.Initialize()

	if err := validateNewPodParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Pod{}

	_jsii_.Create(
		"cdk8s-plus-22.Pod",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewPod_Override(p Pod, scope constructs.Construct, id *string, props *PodProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-22.Pod",
		[]interface{}{scope, id, props},
		p,
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
func Pod_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePod_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-22.Pod",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Pod_ADDRESS_LABEL() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdk8s-plus-22.Pod",
		"ADDRESS_LABEL",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_Pod) AddContainer(cont *ContainerProps) Container {
	if err := p.validateAddContainerParameters(cont); err != nil {
		panic(err)
	}
	var returns Container

	_jsii_.Invoke(
		p,
		"addContainer",
		[]interface{}{cont},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pod) AddHostAlias(hostAlias *HostAlias) {
	if err := p.validateAddHostAliasParameters(hostAlias); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addHostAlias",
		[]interface{}{hostAlias},
	)
}

func (p *jsiiProxy_Pod) AddInitContainer(cont *ContainerProps) Container {
	if err := p.validateAddInitContainerParameters(cont); err != nil {
		panic(err)
	}
	var returns Container

	_jsii_.Invoke(
		p,
		"addInitContainer",
		[]interface{}{cont},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pod) AddVolume(vol Volume) {
	if err := p.validateAddVolumeParameters(vol); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addVolume",
		[]interface{}{vol},
	)
}

func (p *jsiiProxy_Pod) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		p,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pod) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pod) ToNetworkPolicyPeerConfig() *NetworkPolicyPeerConfig {
	var returns *NetworkPolicyPeerConfig

	_jsii_.Invoke(
		p,
		"toNetworkPolicyPeerConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pod) ToPodSelector() IPodSelector {
	var returns IPodSelector

	_jsii_.Invoke(
		p,
		"toPodSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pod) ToPodSelectorConfig() *PodSelectorConfig {
	var returns *PodSelectorConfig

	_jsii_.Invoke(
		p,
		"toPodSelectorConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pod) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pod) ToSubjectConfiguration() *SubjectConfiguration {
	var returns *SubjectConfiguration

	_jsii_.Invoke(
		p,
		"toSubjectConfiguration",
		nil, // no parameters
		&returns,
	)

	return returns
}

