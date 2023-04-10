// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-26 synthesizes Kubernetes manifests for Kubernetes 1.26.0
package cdk8splus26

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus26/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// A Deployment provides declarative updates for Pods and ReplicaSets.
//
// You describe a desired state in a Deployment, and the Deployment Controller changes the actual
// state to the desired state at a controlled rate. You can define Deployments to create new ReplicaSets, or to remove
// existing Deployments and adopt all their resources with new Deployments.
//
// > Note: Do not manage ReplicaSets owned by a Deployment. Consider opening an issue in the main Kubernetes repository if your use case is not covered below.
//
// Use Case
//
// The following are typical use cases for Deployments:
//
// - Create a Deployment to rollout a ReplicaSet. The ReplicaSet creates Pods in the background.
//    Check the status of the rollout to see if it succeeds or not.
// - Declare the new state of the Pods by updating the PodTemplateSpec of the Deployment.
//    A new ReplicaSet is created and the Deployment manages moving the Pods from the old ReplicaSet to the new one at a controlled rate.
//    Each new ReplicaSet updates the revision of the Deployment.
// - Rollback to an earlier Deployment revision if the current state of the Deployment is not stable.
//    Each rollback updates the revision of the Deployment.
// - Scale up the Deployment to facilitate more load.
// - Pause the Deployment to apply multiple fixes to its PodTemplateSpec and then resume it to start a new rollout.
// - Use the status of the Deployment as an indicator that a rollout has stuck.
// - Clean up older ReplicaSets that you don't need anymore.
type Deployment interface {
	Workload
	IScalable
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
	DockerRegistryAuth() ISecret
	// If this is a target of an autoscaler.
	HasAutoscaler() *bool
	SetHasAutoscaler(val *bool)
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
	// Minimum duration for which a newly created pod should be ready without any of its container crashing, for it to be considered available.
	MinReady() cdk8s.Duration
	// The name of this API object.
	Name() *string
	// The tree node.
	Node() constructs.Node
	Permissions() ResourcePermissions
	// The metadata of pods in this workload.
	PodMetadata() cdk8s.ApiObjectMetadataDefinition
	// The maximum duration for a deployment to make progress before it is considered to be failed.
	ProgressDeadline() cdk8s.Duration
	// Number of desired pods.
	Replicas() *float64
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
	Strategy() DeploymentStrategy
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
	// Expose a deployment via an ingress.
	//
	// This will first expose the deployment with a service, and then expose the service via an ingress.
	ExposeViaIngress(path *string, options *ExposeDeploymentViaIngressOptions) Ingress
	// Expose a deployment via a service.
	//
	// This is equivalent to running `kubectl expose deployment <deployment-name>`.
	ExposeViaService(options *DeploymentExposeViaServiceOptions) Service
	// Called on all IScalable targets when they are associated with an autoscaler.
	// See: IScalable.markHasAutoscaler()
	//
	MarkHasAutoscaler()
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
	// Return the target spec properties of this Scalable.
	// See: IScalable.toScalingTarget()
	//
	ToScalingTarget() *ScalingTarget
	// Returns a string representation of this construct.
	ToString() *string
	// Return the subject configuration.
	// See: ISubect.toSubjectConfiguration()
	//
	ToSubjectConfiguration() *SubjectConfiguration
}

// The jsii proxy struct for Deployment
type jsiiProxy_Deployment struct {
	jsiiProxy_Workload
	jsiiProxy_IScalable
}

func (j *jsiiProxy_Deployment) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deployment) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deployment) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deployment) AutomountServiceAccountToken() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"automountServiceAccountToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deployment) Connections() PodConnections {
	var returns PodConnections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deployment) Containers() *[]Container {
	var returns *[]Container
	_jsii_.Get(
		j,
		"containers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deployment) Dns() PodDns {
	var returns PodDns
	_jsii_.Get(
		j,
		"dns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deployment) DockerRegistryAuth() ISecret {
	var returns ISecret
	_jsii_.Get(
		j,
		"dockerRegistryAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deployment) HasAutoscaler() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"hasAutoscaler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deployment) HostAliases() *[]*HostAlias {
	var returns *[]*HostAlias
	_jsii_.Get(
		j,
		"hostAliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deployment) HostNetwork() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"hostNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deployment) InitContainers() *[]Container {
	var returns *[]Container
	_jsii_.Get(
		j,
		"initContainers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deployment) Isolate() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isolate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deployment) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deployment) MatchExpressions() *[]*LabelSelectorRequirement {
	var returns *[]*LabelSelectorRequirement
	_jsii_.Get(
		j,
		"matchExpressions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deployment) MatchLabels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"matchLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deployment) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deployment) MinReady() cdk8s.Duration {
	var returns cdk8s.Duration
	_jsii_.Get(
		j,
		"minReady",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deployment) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deployment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deployment) Permissions() ResourcePermissions {
	var returns ResourcePermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deployment) PodMetadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"podMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deployment) ProgressDeadline() cdk8s.Duration {
	var returns cdk8s.Duration
	_jsii_.Get(
		j,
		"progressDeadline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deployment) Replicas() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deployment) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deployment) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deployment) RestartPolicy() RestartPolicy {
	var returns RestartPolicy
	_jsii_.Get(
		j,
		"restartPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deployment) Scheduling() WorkloadScheduling {
	var returns WorkloadScheduling
	_jsii_.Get(
		j,
		"scheduling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deployment) SecurityContext() PodSecurityContext {
	var returns PodSecurityContext
	_jsii_.Get(
		j,
		"securityContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deployment) ServiceAccount() IServiceAccount {
	var returns IServiceAccount
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deployment) Strategy() DeploymentStrategy {
	var returns DeploymentStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deployment) Volumes() *[]Volume {
	var returns *[]Volume
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}


func NewDeployment(scope constructs.Construct, id *string, props *DeploymentProps) Deployment {
	_init_.Initialize()

	if err := validateNewDeploymentParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Deployment{}

	_jsii_.Create(
		"cdk8s-plus-26.Deployment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDeployment_Override(d Deployment, scope constructs.Construct, id *string, props *DeploymentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-26.Deployment",
		[]interface{}{scope, id, props},
		d,
	)
}

func (j *jsiiProxy_Deployment)SetHasAutoscaler(val *bool) {
	if err := j.validateSetHasAutoscalerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hasAutoscaler",
		val,
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
func Deployment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDeployment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-26.Deployment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Deployment) AddContainer(cont *ContainerProps) Container {
	if err := d.validateAddContainerParameters(cont); err != nil {
		panic(err)
	}
	var returns Container

	_jsii_.Invoke(
		d,
		"addContainer",
		[]interface{}{cont},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Deployment) AddHostAlias(hostAlias *HostAlias) {
	if err := d.validateAddHostAliasParameters(hostAlias); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addHostAlias",
		[]interface{}{hostAlias},
	)
}

func (d *jsiiProxy_Deployment) AddInitContainer(cont *ContainerProps) Container {
	if err := d.validateAddInitContainerParameters(cont); err != nil {
		panic(err)
	}
	var returns Container

	_jsii_.Invoke(
		d,
		"addInitContainer",
		[]interface{}{cont},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Deployment) AddVolume(vol Volume) {
	if err := d.validateAddVolumeParameters(vol); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addVolume",
		[]interface{}{vol},
	)
}

func (d *jsiiProxy_Deployment) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		d,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Deployment) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Deployment) AttachContainer(cont Container) {
	if err := d.validateAttachContainerParameters(cont); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"attachContainer",
		[]interface{}{cont},
	)
}

func (d *jsiiProxy_Deployment) ExposeViaIngress(path *string, options *ExposeDeploymentViaIngressOptions) Ingress {
	if err := d.validateExposeViaIngressParameters(path, options); err != nil {
		panic(err)
	}
	var returns Ingress

	_jsii_.Invoke(
		d,
		"exposeViaIngress",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Deployment) ExposeViaService(options *DeploymentExposeViaServiceOptions) Service {
	if err := d.validateExposeViaServiceParameters(options); err != nil {
		panic(err)
	}
	var returns Service

	_jsii_.Invoke(
		d,
		"exposeViaService",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Deployment) MarkHasAutoscaler() {
	_jsii_.InvokeVoid(
		d,
		"markHasAutoscaler",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Deployment) Select(selectors ...LabelSelector) {
	args := []interface{}{}
	for _, a := range selectors {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		d,
		"select",
		args,
	)
}

func (d *jsiiProxy_Deployment) ToNetworkPolicyPeerConfig() *NetworkPolicyPeerConfig {
	var returns *NetworkPolicyPeerConfig

	_jsii_.Invoke(
		d,
		"toNetworkPolicyPeerConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Deployment) ToPodSelector() IPodSelector {
	var returns IPodSelector

	_jsii_.Invoke(
		d,
		"toPodSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Deployment) ToPodSelectorConfig() *PodSelectorConfig {
	var returns *PodSelectorConfig

	_jsii_.Invoke(
		d,
		"toPodSelectorConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Deployment) ToScalingTarget() *ScalingTarget {
	var returns *ScalingTarget

	_jsii_.Invoke(
		d,
		"toScalingTarget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Deployment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Deployment) ToSubjectConfiguration() *SubjectConfiguration {
	var returns *SubjectConfiguration

	_jsii_.Invoke(
		d,
		"toSubjectConfiguration",
		nil, // no parameters
		&returns,
	)

	return returns
}

