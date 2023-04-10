// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus24/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// A workload is an application running on Kubernetes.
//
// Whether your workload is a single
// component or several that work together, on Kubernetes you run it inside a set of pods.
// In Kubernetes, a Pod represents a set of running containers on your cluster.
type Workload interface {
	AbstractPod
	// The group portion of the API version (e.g. "authorization.k8s.io").
	ApiGroup() *string
	// The underlying cdk8s API object.
	ApiObject() cdk8s.ApiObject
	// The object's API version (e.g. "authorization.k8s.io/v1").
	ApiVersion() *string
	AutomountServiceAccountToken() *bool
	Connections() PodConnections
	Containers() *[]Container
	Dns() PodDns
	DockerRegistryAuth() ISecret
	HostAliases() *[]*HostAlias
	HostNetwork() *bool
	InitContainers() *[]Container
	Isolate() *bool
	// The object kind (e.g. "Deployment").
	Kind() *string
	// The expression matchers this workload will use in order to select pods.
	//
	// Returns a a copy. Use `select()` to add expression matchers.
	MatchExpressions() *[]*LabelSelectorRequirement
	// The label matchers this workload will use in order to select pods.
	//
	// Returns a a copy. Use `select()` to add label matchers.
	MatchLabels() *map[string]*string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	// The name of this API object.
	Name() *string
	// The tree node.
	Node() constructs.Node
	Permissions() ResourcePermissions
	// The metadata of pods in this workload.
	PodMetadata() cdk8s.ApiObjectMetadataDefinition
	// The unique, namespace-global, name of an object inside the Kubernetes cluster.
	//
	// If this is omitted, the ApiResource should represent all objects of the given type.
	ResourceName() *string
	// The name of a resource type as it appears in the relevant API endpoint.
	ResourceType() *string
	RestartPolicy() RestartPolicy
	Scheduling() WorkloadScheduling
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
	// Configure selectors for this workload.
	Select(selectors ...LabelSelector)
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

// The jsii proxy struct for Workload
type jsiiProxy_Workload struct {
	jsiiProxy_AbstractPod
}

func (j *jsiiProxy_Workload) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) AutomountServiceAccountToken() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"automountServiceAccountToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) Connections() PodConnections {
	var returns PodConnections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) Containers() *[]Container {
	var returns *[]Container
	_jsii_.Get(
		j,
		"containers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) Dns() PodDns {
	var returns PodDns
	_jsii_.Get(
		j,
		"dns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) DockerRegistryAuth() ISecret {
	var returns ISecret
	_jsii_.Get(
		j,
		"dockerRegistryAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) HostAliases() *[]*HostAlias {
	var returns *[]*HostAlias
	_jsii_.Get(
		j,
		"hostAliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) HostNetwork() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"hostNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) InitContainers() *[]Container {
	var returns *[]Container
	_jsii_.Get(
		j,
		"initContainers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) Isolate() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isolate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) MatchExpressions() *[]*LabelSelectorRequirement {
	var returns *[]*LabelSelectorRequirement
	_jsii_.Get(
		j,
		"matchExpressions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) MatchLabels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"matchLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) Permissions() ResourcePermissions {
	var returns ResourcePermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) PodMetadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"podMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) RestartPolicy() RestartPolicy {
	var returns RestartPolicy
	_jsii_.Get(
		j,
		"restartPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) Scheduling() WorkloadScheduling {
	var returns WorkloadScheduling
	_jsii_.Get(
		j,
		"scheduling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) SecurityContext() PodSecurityContext {
	var returns PodSecurityContext
	_jsii_.Get(
		j,
		"securityContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) ServiceAccount() IServiceAccount {
	var returns IServiceAccount
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) Volumes() *[]Volume {
	var returns *[]Volume
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}


func NewWorkload_Override(w Workload, scope constructs.Construct, id *string, props *WorkloadProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.Workload",
		[]interface{}{scope, id, props},
		w,
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
func Workload_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWorkload_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Workload",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workload) AddContainer(cont *ContainerProps) Container {
	if err := w.validateAddContainerParameters(cont); err != nil {
		panic(err)
	}
	var returns Container

	_jsii_.Invoke(
		w,
		"addContainer",
		[]interface{}{cont},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workload) AddHostAlias(hostAlias *HostAlias) {
	if err := w.validateAddHostAliasParameters(hostAlias); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addHostAlias",
		[]interface{}{hostAlias},
	)
}

func (w *jsiiProxy_Workload) AddInitContainer(cont *ContainerProps) Container {
	if err := w.validateAddInitContainerParameters(cont); err != nil {
		panic(err)
	}
	var returns Container

	_jsii_.Invoke(
		w,
		"addInitContainer",
		[]interface{}{cont},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workload) AddVolume(vol Volume) {
	if err := w.validateAddVolumeParameters(vol); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addVolume",
		[]interface{}{vol},
	)
}

func (w *jsiiProxy_Workload) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		w,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workload) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workload) AttachContainer(cont Container) {
	if err := w.validateAttachContainerParameters(cont); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"attachContainer",
		[]interface{}{cont},
	)
}

func (w *jsiiProxy_Workload) Select(selectors ...LabelSelector) {
	args := []interface{}{}
	for _, a := range selectors {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		w,
		"select",
		args,
	)
}

func (w *jsiiProxy_Workload) ToNetworkPolicyPeerConfig() *NetworkPolicyPeerConfig {
	var returns *NetworkPolicyPeerConfig

	_jsii_.Invoke(
		w,
		"toNetworkPolicyPeerConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workload) ToPodSelector() IPodSelector {
	var returns IPodSelector

	_jsii_.Invoke(
		w,
		"toPodSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workload) ToPodSelectorConfig() *PodSelectorConfig {
	var returns *PodSelectorConfig

	_jsii_.Invoke(
		w,
		"toPodSelectorConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workload) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workload) ToSubjectConfiguration() *SubjectConfiguration {
	var returns *SubjectConfiguration

	_jsii_.Invoke(
		w,
		"toSubjectConfiguration",
		nil, // no parameters
		&returns,
	)

	return returns
}

