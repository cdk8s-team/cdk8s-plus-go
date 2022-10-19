// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus24/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// A CronJob is responsible for creating a Job and scheduling it based on provided cron schedule.
//
// This helps running Jobs in a recurring manner.
type CronJob interface {
	Workload
	// The group portion of the API version (e.g. "authorization.k8s.io").
	ApiGroup() *string
	// The underlying cdk8s API object.
	// See: base.Resource.apiObject
	//
	ApiObject() cdk8s.ApiObject
	// The object's API version (e.g. "authorization.k8s.io/v1").
	ApiVersion() *string
	AutomountServiceAccountToken() *bool
	// The policy used by this cron job to determine the concurrency mode in which to schedule jobs.
	ConcurrencyPolicy() *string
	Connections() PodConnections
	Containers() *[]Container
	Dns() PodDns
	DockerRegistryAuth() DockerConfigSecret
	// The number of failed jobs retained by this cron job.
	FailedJobsRetained() *float64
	HostAliases() *[]*HostAlias
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
	// Represents the resource type.
	ResourceType() *string
	RestartPolicy() RestartPolicy
	// The schedule this cron job is scheduled to run in.
	Schedule() cdk8s.Cron
	Scheduling() WorkloadScheduling
	SecurityContext() PodSecurityContext
	ServiceAccount() IServiceAccount
	// The time by which the running cron job needs to schedule the next job execution.
	//
	// The job is considered as failed if it misses this deadline.
	StartingDeadline() cdk8s.Duration
	// The number of successful jobs retained by this cron job.
	SuccessfulJobsRetained() *float64
	// Whether or not the cron job is currently suspended or not.
	Suspend() *bool
	// The timezone which this cron job would follow to schedule jobs.
	TimeZone() *string
	Volumes() *[]Volume
	AddContainer(cont *ContainerProps) Container
	AddHostAlias(hostAlias *HostAlias)
	AddInitContainer(cont *ContainerProps) Container
	AddVolume(vol Volume)
	// Return the IApiResource this object represents.
	AsApiResource() IApiResource
	// Return the non resource url this object represents.
	AsNonApiResource() *string
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

// The jsii proxy struct for CronJob
type jsiiProxy_CronJob struct {
	jsiiProxy_Workload
}

func (j *jsiiProxy_CronJob) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJob) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJob) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJob) AutomountServiceAccountToken() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"automountServiceAccountToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJob) ConcurrencyPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"concurrencyPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJob) Connections() PodConnections {
	var returns PodConnections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJob) Containers() *[]Container {
	var returns *[]Container
	_jsii_.Get(
		j,
		"containers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJob) Dns() PodDns {
	var returns PodDns
	_jsii_.Get(
		j,
		"dns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJob) DockerRegistryAuth() DockerConfigSecret {
	var returns DockerConfigSecret
	_jsii_.Get(
		j,
		"dockerRegistryAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJob) FailedJobsRetained() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failedJobsRetained",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJob) HostAliases() *[]*HostAlias {
	var returns *[]*HostAlias
	_jsii_.Get(
		j,
		"hostAliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJob) InitContainers() *[]Container {
	var returns *[]Container
	_jsii_.Get(
		j,
		"initContainers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJob) Isolate() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isolate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJob) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJob) MatchExpressions() *[]*LabelSelectorRequirement {
	var returns *[]*LabelSelectorRequirement
	_jsii_.Get(
		j,
		"matchExpressions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJob) MatchLabels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"matchLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJob) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJob) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJob) Permissions() ResourcePermissions {
	var returns ResourcePermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJob) PodMetadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"podMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJob) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJob) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJob) RestartPolicy() RestartPolicy {
	var returns RestartPolicy
	_jsii_.Get(
		j,
		"restartPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJob) Schedule() cdk8s.Cron {
	var returns cdk8s.Cron
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJob) Scheduling() WorkloadScheduling {
	var returns WorkloadScheduling
	_jsii_.Get(
		j,
		"scheduling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJob) SecurityContext() PodSecurityContext {
	var returns PodSecurityContext
	_jsii_.Get(
		j,
		"securityContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJob) ServiceAccount() IServiceAccount {
	var returns IServiceAccount
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJob) StartingDeadline() cdk8s.Duration {
	var returns cdk8s.Duration
	_jsii_.Get(
		j,
		"startingDeadline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJob) SuccessfulJobsRetained() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successfulJobsRetained",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJob) Suspend() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"suspend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJob) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJob) Volumes() *[]Volume {
	var returns *[]Volume
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}


func NewCronJob(scope constructs.Construct, id *string, props *CronJobProps) CronJob {
	_init_.Initialize()

	if err := validateNewCronJobParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CronJob{}

	_jsii_.Create(
		"cdk8s-plus-24.CronJob",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCronJob_Override(c CronJob, scope constructs.Construct, id *string, props *CronJobProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.CronJob",
		[]interface{}{scope, id, props},
		c,
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
func CronJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCronJob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.CronJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJob) AddContainer(cont *ContainerProps) Container {
	if err := c.validateAddContainerParameters(cont); err != nil {
		panic(err)
	}
	var returns Container

	_jsii_.Invoke(
		c,
		"addContainer",
		[]interface{}{cont},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJob) AddHostAlias(hostAlias *HostAlias) {
	if err := c.validateAddHostAliasParameters(hostAlias); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addHostAlias",
		[]interface{}{hostAlias},
	)
}

func (c *jsiiProxy_CronJob) AddInitContainer(cont *ContainerProps) Container {
	if err := c.validateAddInitContainerParameters(cont); err != nil {
		panic(err)
	}
	var returns Container

	_jsii_.Invoke(
		c,
		"addInitContainer",
		[]interface{}{cont},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJob) AddVolume(vol Volume) {
	if err := c.validateAddVolumeParameters(vol); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addVolume",
		[]interface{}{vol},
	)
}

func (c *jsiiProxy_CronJob) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		c,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJob) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJob) Select(selectors ...LabelSelector) {
	args := []interface{}{}
	for _, a := range selectors {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"select",
		args,
	)
}

func (c *jsiiProxy_CronJob) ToNetworkPolicyPeerConfig() *NetworkPolicyPeerConfig {
	var returns *NetworkPolicyPeerConfig

	_jsii_.Invoke(
		c,
		"toNetworkPolicyPeerConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJob) ToPodSelector() IPodSelector {
	var returns IPodSelector

	_jsii_.Invoke(
		c,
		"toPodSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJob) ToPodSelectorConfig() *PodSelectorConfig {
	var returns *PodSelectorConfig

	_jsii_.Invoke(
		c,
		"toPodSelectorConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJob) ToSubjectConfiguration() *SubjectConfiguration {
	var returns *SubjectConfiguration

	_jsii_.Invoke(
		c,
		"toSubjectConfiguration",
		nil, // no parameters
		&returns,
	)

	return returns
}

