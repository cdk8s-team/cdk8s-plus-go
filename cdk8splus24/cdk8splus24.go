// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus24/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus24/v2/internal"
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
	DockerRegistryAuth() DockerConfigSecret
	HostAliases() *[]*HostAlias
	InitContainers() *[]Container
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

func (j *jsiiProxy_AbstractPod) DockerRegistryAuth() DockerConfigSecret {
	var returns DockerConfigSecret
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
		"cdk8s-plus-24.AbstractPod",
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

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.AbstractPod",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AbstractPod) AddContainer(cont *ContainerProps) Container {
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
	_jsii_.InvokeVoid(
		a,
		"addHostAlias",
		[]interface{}{hostAlias},
	)
}

func (a *jsiiProxy_AbstractPod) AddInitContainer(cont *ContainerProps) Container {
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

// Properties for `AbstractPod`.
type AbstractPodProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// Indicates whether a service account token should be automatically mounted.
	// See: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/#use-the-default-service-account-to-access-the-api-server
	//
	AutomountServiceAccountToken *bool `field:"optional" json:"automountServiceAccountToken" yaml:"automountServiceAccountToken"`
	// List of containers belonging to the pod.
	//
	// Containers cannot currently be
	// added or removed. There must be at least one container in a Pod.
	//
	// You can add additionnal containers using `podSpec.addContainer()`
	Containers *[]*ContainerProps `field:"optional" json:"containers" yaml:"containers"`
	// DNS settings for the pod.
	// See: https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/
	//
	Dns *PodDnsProps `field:"optional" json:"dns" yaml:"dns"`
	// A secret containing docker credentials for authenticating to a registry.
	DockerRegistryAuth DockerConfigSecret `field:"optional" json:"dockerRegistryAuth" yaml:"dockerRegistryAuth"`
	// HostAlias holds the mapping between IP and hostnames that will be injected as an entry in the pod's hosts file.
	HostAliases *[]*HostAlias `field:"optional" json:"hostAliases" yaml:"hostAliases"`
	// List of initialization containers belonging to the pod.
	//
	// Init containers are executed in order prior to containers being started.
	// If any init container fails, the pod is considered to have failed and is handled according to its restartPolicy.
	// The name for an init container or normal container must be unique among all containers.
	// Init containers may not have Lifecycle actions, Readiness probes, Liveness probes, or Startup probes.
	// The resourceRequirements of an init container are taken into account during scheduling by finding the highest request/limit
	// for each resource type, and then using the max of of that value or the sum of the normal containers.
	// Limits are applied to init containers in a similar fashion.
	//
	// Init containers cannot currently be added ,removed or updated.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/
	//
	InitContainers *[]*ContainerProps `field:"optional" json:"initContainers" yaml:"initContainers"`
	// Restart policy for all containers within the pod.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy
	//
	RestartPolicy RestartPolicy `field:"optional" json:"restartPolicy" yaml:"restartPolicy"`
	// SecurityContext holds pod-level security attributes and common container settings.
	SecurityContext *PodSecurityContextProps `field:"optional" json:"securityContext" yaml:"securityContext"`
	// A service account provides an identity for processes that run in a Pod.
	//
	// When you (a human) access the cluster (for example, using kubectl), you are
	// authenticated by the apiserver as a particular User Account (currently this
	// is usually admin, unless your cluster administrator has customized your
	// cluster). Processes in containers inside pods can also contact the
	// apiserver. When they do, they are authenticated as a particular Service
	// Account (for example, default).
	// See: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
	//
	ServiceAccount IServiceAccount `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// List of volumes that can be mounted by containers belonging to the pod.
	//
	// You can also add volumes later using `podSpec.addVolume()`
	// See: https://kubernetes.io/docs/concepts/storage/volumes
	//
	Volumes *[]Volume `field:"optional" json:"volumes" yaml:"volumes"`
}

// Options to add a deployment to a service.
type AddDeploymentOptions struct {
	// The name of this port within the service.
	//
	// This must be a DNS_LABEL. All
	// ports within a ServiceSpec must have unique names. This maps to the 'Name'
	// field in EndpointPort objects. Optional if only one ServicePort is defined
	// on this service.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The port on each node on which this service is exposed when type=NodePort or LoadBalancer.
	//
	// Usually assigned by the system. If specified, it will be
	// allocated to the service if unused or else creation of the service will
	// fail. Default is to auto-allocate a port if the ServiceType of this Service
	// requires one.
	// See: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport
	//
	NodePort *float64 `field:"optional" json:"nodePort" yaml:"nodePort"`
	// The IP protocol for this port.
	//
	// Supports "TCP", "UDP", and "SCTP". Default is TCP.
	Protocol Protocol `field:"optional" json:"protocol" yaml:"protocol"`
	// The port number the service will redirect to.
	TargetPort *float64 `field:"optional" json:"targetPort" yaml:"targetPort"`
	// The port number the service will bind to.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

// Options for `configmap.addDirectory()`.
type AddDirectoryOptions struct {
	// Glob patterns to exclude when adding files.
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// A prefix to add to all keys in the config map.
	KeyPrefix *string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
}

// Represents information about an API resource type.
type ApiResource interface {
	IApiEndpoint
	IApiResource
	// The group portion of the API version (e.g. `authorization.k8s.io`).
	ApiGroup() *string
	// The name of the resource type as it appears in the relevant API endpoint.
	//
	// Example:
	//   - "pods" or "pods/log"
	//
	// See: https://kubernetes.io/docs/reference/access-authn-authz/rbac/#referring-to-resources
	//
	ResourceType() *string
	// Return the IApiResource this object represents.
	AsApiResource() IApiResource
	// Return the non resource url this object represents.
	AsNonApiResource() *string
}

// The jsii proxy struct for ApiResource
type jsiiProxy_ApiResource struct {
	jsiiProxy_IApiEndpoint
	jsiiProxy_IApiResource
}

func (j *jsiiProxy_ApiResource) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiResource) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}


// API resource information for a custom resource type.
func ApiResource_Custom(options *ApiResourceOptions) ApiResource {
	_init_.Initialize()

	var returns ApiResource

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.ApiResource",
		"custom",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func ApiResource_API_SERVICES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"API_SERVICES",
		&returns,
	)
	return returns
}

func ApiResource_BINDINGS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"BINDINGS",
		&returns,
	)
	return returns
}

func ApiResource_CERTIFICATE_SIGNING_REQUESTS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"CERTIFICATE_SIGNING_REQUESTS",
		&returns,
	)
	return returns
}

func ApiResource_CLUSTER_ROLE_BINDINGS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"CLUSTER_ROLE_BINDINGS",
		&returns,
	)
	return returns
}

func ApiResource_CLUSTER_ROLES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"CLUSTER_ROLES",
		&returns,
	)
	return returns
}

func ApiResource_COMPONENT_STATUSES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"COMPONENT_STATUSES",
		&returns,
	)
	return returns
}

func ApiResource_CONFIG_MAPS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"CONFIG_MAPS",
		&returns,
	)
	return returns
}

func ApiResource_CONTROLLER_REVISIONS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"CONTROLLER_REVISIONS",
		&returns,
	)
	return returns
}

func ApiResource_CRON_JOBS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"CRON_JOBS",
		&returns,
	)
	return returns
}

func ApiResource_CSI_DRIVERS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"CSI_DRIVERS",
		&returns,
	)
	return returns
}

func ApiResource_CSI_NODES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"CSI_NODES",
		&returns,
	)
	return returns
}

func ApiResource_CSI_STORAGE_CAPACITIES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"CSI_STORAGE_CAPACITIES",
		&returns,
	)
	return returns
}

func ApiResource_CUSTOM_RESOURCE_DEFINITIONS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"CUSTOM_RESOURCE_DEFINITIONS",
		&returns,
	)
	return returns
}

func ApiResource_DAEMON_SETS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"DAEMON_SETS",
		&returns,
	)
	return returns
}

func ApiResource_DEPLOYMENTS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"DEPLOYMENTS",
		&returns,
	)
	return returns
}

func ApiResource_ENDPOINT_SLICES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"ENDPOINT_SLICES",
		&returns,
	)
	return returns
}

func ApiResource_ENDPOINTS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"ENDPOINTS",
		&returns,
	)
	return returns
}

func ApiResource_EVENTS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"EVENTS",
		&returns,
	)
	return returns
}

func ApiResource_FLOW_SCHEMAS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"FLOW_SCHEMAS",
		&returns,
	)
	return returns
}

func ApiResource_HORIZONTAL_POD_AUTOSCALERS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"HORIZONTAL_POD_AUTOSCALERS",
		&returns,
	)
	return returns
}

func ApiResource_INGRESS_CLASSES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"INGRESS_CLASSES",
		&returns,
	)
	return returns
}

func ApiResource_INGRESSES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"INGRESSES",
		&returns,
	)
	return returns
}

func ApiResource_JOBS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"JOBS",
		&returns,
	)
	return returns
}

func ApiResource_LEASES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"LEASES",
		&returns,
	)
	return returns
}

func ApiResource_LIMIT_RANGES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"LIMIT_RANGES",
		&returns,
	)
	return returns
}

func ApiResource_LOCAL_SUBJECT_ACCESS_REVIEWS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"LOCAL_SUBJECT_ACCESS_REVIEWS",
		&returns,
	)
	return returns
}

func ApiResource_MUTATING_WEBHOOK_CONFIGURATIONS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"MUTATING_WEBHOOK_CONFIGURATIONS",
		&returns,
	)
	return returns
}

func ApiResource_NAMESPACES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"NAMESPACES",
		&returns,
	)
	return returns
}

func ApiResource_NETWORK_POLICIES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"NETWORK_POLICIES",
		&returns,
	)
	return returns
}

func ApiResource_NODES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"NODES",
		&returns,
	)
	return returns
}

func ApiResource_PERSISTENT_VOLUME_CLAIMS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"PERSISTENT_VOLUME_CLAIMS",
		&returns,
	)
	return returns
}

func ApiResource_PERSISTENT_VOLUMES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"PERSISTENT_VOLUMES",
		&returns,
	)
	return returns
}

func ApiResource_POD_DISRUPTION_BUDGETS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"POD_DISRUPTION_BUDGETS",
		&returns,
	)
	return returns
}

func ApiResource_POD_SECURITY_POLICIES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"POD_SECURITY_POLICIES",
		&returns,
	)
	return returns
}

func ApiResource_POD_TEMPLATES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"POD_TEMPLATES",
		&returns,
	)
	return returns
}

func ApiResource_PODS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"PODS",
		&returns,
	)
	return returns
}

func ApiResource_PRIORITY_CLASSES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"PRIORITY_CLASSES",
		&returns,
	)
	return returns
}

func ApiResource_PRIORITY_LEVEL_CONFIGURATIONS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"PRIORITY_LEVEL_CONFIGURATIONS",
		&returns,
	)
	return returns
}

func ApiResource_REPLICA_SETS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"REPLICA_SETS",
		&returns,
	)
	return returns
}

func ApiResource_REPLICATION_CONTROLLERS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"REPLICATION_CONTROLLERS",
		&returns,
	)
	return returns
}

func ApiResource_RESOURCE_QUOTAS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"RESOURCE_QUOTAS",
		&returns,
	)
	return returns
}

func ApiResource_ROLE_BINDINGS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"ROLE_BINDINGS",
		&returns,
	)
	return returns
}

func ApiResource_ROLES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"ROLES",
		&returns,
	)
	return returns
}

func ApiResource_RUNTIME_CLASSES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"RUNTIME_CLASSES",
		&returns,
	)
	return returns
}

func ApiResource_SECRETS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"SECRETS",
		&returns,
	)
	return returns
}

func ApiResource_SELF_SUBJECT_ACCESS_REVIEWS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"SELF_SUBJECT_ACCESS_REVIEWS",
		&returns,
	)
	return returns
}

func ApiResource_SELF_SUBJECT_RULES_REVIEWS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"SELF_SUBJECT_RULES_REVIEWS",
		&returns,
	)
	return returns
}

func ApiResource_SERVICE_ACCOUNTS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"SERVICE_ACCOUNTS",
		&returns,
	)
	return returns
}

func ApiResource_SERVICES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"SERVICES",
		&returns,
	)
	return returns
}

func ApiResource_STATEFUL_SETS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"STATEFUL_SETS",
		&returns,
	)
	return returns
}

func ApiResource_STORAGE_CLASSES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"STORAGE_CLASSES",
		&returns,
	)
	return returns
}

func ApiResource_SUBJECT_ACCESS_REVIEWS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"SUBJECT_ACCESS_REVIEWS",
		&returns,
	)
	return returns
}

func ApiResource_TOKEN_REVIEWS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"TOKEN_REVIEWS",
		&returns,
	)
	return returns
}

func ApiResource_VALIDATING_WEBHOOK_CONFIGURATIONS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"VALIDATING_WEBHOOK_CONFIGURATIONS",
		&returns,
	)
	return returns
}

func ApiResource_VOLUME_ATTACHMENTS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-24.ApiResource",
		"VOLUME_ATTACHMENTS",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ApiResource) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		a,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiResource) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options for `ApiResource`.
type ApiResourceOptions struct {
	// The group portion of the API version (e.g. `authorization.k8s.io`).
	ApiGroup *string `field:"required" json:"apiGroup" yaml:"apiGroup"`
	// The name of the resource type as it appears in the relevant API endpoint.
	//
	// Example:
	//   - "pods" or "pods/log"
	//
	// See: https://kubernetes.io/docs/reference/access-authn-authz/rbac/#referring-to-resources
	//
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
}

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

// Properties for `AwsElasticBlockStorePersistentVolume`.
type AwsElasticBlockStorePersistentVolumeProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// Contains all ways the volume can be mounted.
	// See: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes
	//
	AccessModes *[]PersistentVolumeAccessMode `field:"optional" json:"accessModes" yaml:"accessModes"`
	// Part of a bi-directional binding between PersistentVolume and PersistentVolumeClaim.
	//
	// Expected to be non-nil when bound.
	// See: https://kubernetes.io/docs/concepts/storage/persistent-volumes#binding
	//
	Claim IPersistentVolumeClaim `field:"optional" json:"claim" yaml:"claim"`
	// A list of mount options, e.g. ["ro", "soft"]. Not validated - mount will simply fail if one is invalid.
	// See: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#mount-options
	//
	MountOptions *[]*string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
	// When a user is done with their volume, they can delete the PVC objects from the API that allows reclamation of the resource.
	//
	// The reclaim policy tells the cluster what to do with
	// the volume after it has been released of its claim.
	// See: https://kubernetes.io/docs/concepts/storage/persistent-volumes#reclaiming
	//
	ReclaimPolicy PersistentVolumeReclaimPolicy `field:"optional" json:"reclaimPolicy" yaml:"reclaimPolicy"`
	// What is the storage capacity of this volume.
	// See: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources
	//
	Storage cdk8s.Size `field:"optional" json:"storage" yaml:"storage"`
	// Name of StorageClass to which this persistent volume belongs.
	StorageClassName *string `field:"optional" json:"storageClassName" yaml:"storageClassName"`
	// Defines what type of volume is required by the claim.
	VolumeMode PersistentVolumeMode `field:"optional" json:"volumeMode" yaml:"volumeMode"`
	// Unique ID of the persistent disk resource in AWS (Amazon EBS volume).
	//
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	// See: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	//
	VolumeId *string `field:"required" json:"volumeId" yaml:"volumeId"`
	// Filesystem type of the volume that you want to mount.
	//
	// Tip: Ensure that the filesystem type is supported by the host operating system.
	// See: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	//
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// The partition in the volume that you want to mount.
	//
	// If omitted, the default is to mount by volume name.
	// Examples: For volume /dev/sda1, you specify the partition as "1".
	// Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty).
	Partition *float64 `field:"optional" json:"partition" yaml:"partition"`
	// Specify "true" to force and set the ReadOnly property in VolumeMounts to "true".
	// See: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	//
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
}

// Options of `Volume.fromAwsElasticBlockStore`.
type AwsElasticBlockStoreVolumeOptions struct {
	// Filesystem type of the volume that you want to mount.
	//
	// Tip: Ensure that the filesystem type is supported by the host operating system.
	// See: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	//
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// The volume name.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The partition in the volume that you want to mount.
	//
	// If omitted, the default is to mount by volume name.
	// Examples: For volume /dev/sda1, you specify the partition as "1".
	// Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty).
	Partition *float64 `field:"optional" json:"partition" yaml:"partition"`
	// Specify "true" to force and set the ReadOnly property in VolumeMounts to "true".
	// See: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	//
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
}

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

	j := jsiiProxy_AzureDiskPersistentVolume{}

	_jsii_.Create(
		"cdk8s-plus-24.AzureDiskPersistentVolume",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewAzureDiskPersistentVolume_Override(a AzureDiskPersistentVolume, scope constructs.Construct, id *string, props *AzureDiskPersistentVolumeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.AzureDiskPersistentVolume",
		[]interface{}{scope, id, props},
		a,
	)
}

// Imports a pv from the cluster as a reference.
func AzureDiskPersistentVolume_FromPersistentVolumeName(scope constructs.Construct, id *string, volumeName *string) IPersistentVolume {
	_init_.Initialize()

	var returns IPersistentVolume

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.AzureDiskPersistentVolume",
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

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.AzureDiskPersistentVolume",
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

// Azure disk caching modes.
type AzureDiskPersistentVolumeCachingMode string

const (
	// None.
	AzureDiskPersistentVolumeCachingMode_NONE AzureDiskPersistentVolumeCachingMode = "NONE"
	// ReadOnly.
	AzureDiskPersistentVolumeCachingMode_READ_ONLY AzureDiskPersistentVolumeCachingMode = "READ_ONLY"
	// ReadWrite.
	AzureDiskPersistentVolumeCachingMode_READ_WRITE AzureDiskPersistentVolumeCachingMode = "READ_WRITE"
)

// Azure Disk kinds.
type AzureDiskPersistentVolumeKind string

const (
	// Multiple blob disks per storage account.
	AzureDiskPersistentVolumeKind_SHARED AzureDiskPersistentVolumeKind = "SHARED"
	// Single blob disk per storage account.
	AzureDiskPersistentVolumeKind_DEDICATED AzureDiskPersistentVolumeKind = "DEDICATED"
	// Azure managed data disk.
	AzureDiskPersistentVolumeKind_MANAGED AzureDiskPersistentVolumeKind = "MANAGED"
)

// Properties for `AzureDiskPersistentVolume`.
type AzureDiskPersistentVolumeProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// Contains all ways the volume can be mounted.
	// See: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes
	//
	AccessModes *[]PersistentVolumeAccessMode `field:"optional" json:"accessModes" yaml:"accessModes"`
	// Part of a bi-directional binding between PersistentVolume and PersistentVolumeClaim.
	//
	// Expected to be non-nil when bound.
	// See: https://kubernetes.io/docs/concepts/storage/persistent-volumes#binding
	//
	Claim IPersistentVolumeClaim `field:"optional" json:"claim" yaml:"claim"`
	// A list of mount options, e.g. ["ro", "soft"]. Not validated - mount will simply fail if one is invalid.
	// See: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#mount-options
	//
	MountOptions *[]*string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
	// When a user is done with their volume, they can delete the PVC objects from the API that allows reclamation of the resource.
	//
	// The reclaim policy tells the cluster what to do with
	// the volume after it has been released of its claim.
	// See: https://kubernetes.io/docs/concepts/storage/persistent-volumes#reclaiming
	//
	ReclaimPolicy PersistentVolumeReclaimPolicy `field:"optional" json:"reclaimPolicy" yaml:"reclaimPolicy"`
	// What is the storage capacity of this volume.
	// See: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources
	//
	Storage cdk8s.Size `field:"optional" json:"storage" yaml:"storage"`
	// Name of StorageClass to which this persistent volume belongs.
	StorageClassName *string `field:"optional" json:"storageClassName" yaml:"storageClassName"`
	// Defines what type of volume is required by the claim.
	VolumeMode PersistentVolumeMode `field:"optional" json:"volumeMode" yaml:"volumeMode"`
	// The Name of the data disk in the blob storage.
	DiskName *string `field:"required" json:"diskName" yaml:"diskName"`
	// The URI the data disk in the blob storage.
	DiskUri *string `field:"required" json:"diskUri" yaml:"diskUri"`
	// Host Caching mode.
	CachingMode AzureDiskPersistentVolumeCachingMode `field:"optional" json:"cachingMode" yaml:"cachingMode"`
	// Filesystem type to mount.
	//
	// Must be a filesystem type supported by the host operating system.
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// Kind of disk.
	Kind AzureDiskPersistentVolumeKind `field:"optional" json:"kind" yaml:"kind"`
	// Force the ReadOnly setting in VolumeMounts.
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
}

// Options of `Volume.fromAzureDisk`.
type AzureDiskVolumeOptions struct {
	// Host Caching mode.
	CachingMode AzureDiskPersistentVolumeCachingMode `field:"optional" json:"cachingMode" yaml:"cachingMode"`
	// Filesystem type to mount.
	//
	// Must be a filesystem type supported by the host operating system.
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// Kind of disk.
	Kind AzureDiskPersistentVolumeKind `field:"optional" json:"kind" yaml:"kind"`
	// The volume name.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Force the ReadOnly setting in VolumeMounts.
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
}

// Create a secret for basic authentication.
// See: https://kubernetes.io/docs/concepts/configuration/secret/#basic-authentication-secret
//
type BasicAuthSecret interface {
	Secret
	// The group portion of the API version (e.g. "authorization.k8s.io").
	ApiGroup() *string
	// The underlying cdk8s API object.
	// See: base.Resource.apiObject
	//
	ApiObject() cdk8s.ApiObject
	// The object's API version (e.g. "authorization.k8s.io/v1").
	ApiVersion() *string
	// Whether or not the secret is immutable.
	Immutable() *bool
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
	// Adds a string data field to the secert.
	AddStringData(key *string, value *string)
	// Return the IApiResource this object represents.
	AsApiResource() IApiResource
	// Return the non resource url this object represents.
	AsNonApiResource() *string
	// Gets a string data by key or undefined.
	GetStringData(key *string) *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for BasicAuthSecret
type jsiiProxy_BasicAuthSecret struct {
	jsiiProxy_Secret
}

func (j *jsiiProxy_BasicAuthSecret) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BasicAuthSecret) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BasicAuthSecret) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BasicAuthSecret) Immutable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"immutable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BasicAuthSecret) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BasicAuthSecret) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BasicAuthSecret) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BasicAuthSecret) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BasicAuthSecret) Permissions() ResourcePermissions {
	var returns ResourcePermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BasicAuthSecret) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BasicAuthSecret) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}


func NewBasicAuthSecret(scope constructs.Construct, id *string, props *BasicAuthSecretProps) BasicAuthSecret {
	_init_.Initialize()

	j := jsiiProxy_BasicAuthSecret{}

	_jsii_.Create(
		"cdk8s-plus-24.BasicAuthSecret",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewBasicAuthSecret_Override(b BasicAuthSecret, scope constructs.Construct, id *string, props *BasicAuthSecretProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.BasicAuthSecret",
		[]interface{}{scope, id, props},
		b,
	)
}

// Imports a secret from the cluster as a reference.
func BasicAuthSecret_FromSecretName(scope constructs.Construct, id *string, name *string) ISecret {
	_init_.Initialize()

	var returns ISecret

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.BasicAuthSecret",
		"fromSecretName",
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
func BasicAuthSecret_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.BasicAuthSecret",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BasicAuthSecret) AddStringData(key *string, value *string) {
	_jsii_.InvokeVoid(
		b,
		"addStringData",
		[]interface{}{key, value},
	)
}

func (b *jsiiProxy_BasicAuthSecret) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		b,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BasicAuthSecret) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BasicAuthSecret) GetStringData(key *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringData",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BasicAuthSecret) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options for `BasicAuthSecret`.
type BasicAuthSecretProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// If set to true, ensures that data stored in the Secret cannot be updated (only object metadata can be modified).
	//
	// If not set to true, the field can be modified at any time.
	Immutable *bool `field:"optional" json:"immutable" yaml:"immutable"`
	// The password or token for authentication.
	Password *string `field:"required" json:"password" yaml:"password"`
	// The user name for authentication.
	Username *string `field:"required" json:"username" yaml:"username"`
}

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

	j := jsiiProxy_ClusterRole{}

	_jsii_.Create(
		"cdk8s-plus-24.ClusterRole",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewClusterRole_Override(c ClusterRole, scope constructs.Construct, id *string, props *ClusterRoleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.ClusterRole",
		[]interface{}{scope, id, props},
		c,
	)
}

// Imports a role from the cluster as a reference.
func ClusterRole_FromClusterRoleName(scope constructs.Construct, id *string, name *string) IClusterRole {
	_init_.Initialize()

	var returns IClusterRole

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.ClusterRole",
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

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.ClusterRole",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterRole) Aggregate(key *string, value *string) {
	_jsii_.InvokeVoid(
		c,
		"aggregate",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_ClusterRole) Allow(verbs *[]*string, endpoints ...IApiEndpoint) {
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

// A ClusterRoleBinding grants permissions cluster-wide to a user or set of users.
type ClusterRoleBinding interface {
	Resource
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
	Role() IClusterRole
	Subjects() *[]ISubject
	// Adds a subject to the role.
	AddSubjects(subjects ...ISubject)
	// Return the IApiResource this object represents.
	AsApiResource() IApiResource
	// Return the non resource url this object represents.
	AsNonApiResource() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for ClusterRoleBinding
type jsiiProxy_ClusterRoleBinding struct {
	jsiiProxy_Resource
}

func (j *jsiiProxy_ClusterRoleBinding) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRoleBinding) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRoleBinding) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRoleBinding) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRoleBinding) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRoleBinding) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRoleBinding) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRoleBinding) Permissions() ResourcePermissions {
	var returns ResourcePermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRoleBinding) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRoleBinding) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRoleBinding) Role() IClusterRole {
	var returns IClusterRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRoleBinding) Subjects() *[]ISubject {
	var returns *[]ISubject
	_jsii_.Get(
		j,
		"subjects",
		&returns,
	)
	return returns
}


func NewClusterRoleBinding(scope constructs.Construct, id *string, props *ClusterRoleBindingProps) ClusterRoleBinding {
	_init_.Initialize()

	j := jsiiProxy_ClusterRoleBinding{}

	_jsii_.Create(
		"cdk8s-plus-24.ClusterRoleBinding",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewClusterRoleBinding_Override(c ClusterRoleBinding, scope constructs.Construct, id *string, props *ClusterRoleBindingProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.ClusterRoleBinding",
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
func ClusterRoleBinding_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.ClusterRoleBinding",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterRoleBinding) AddSubjects(subjects ...ISubject) {
	args := []interface{}{}
	for _, a := range subjects {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addSubjects",
		args,
	)
}

func (c *jsiiProxy_ClusterRoleBinding) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		c,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterRoleBinding) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterRoleBinding) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for `ClusterRoleBinding`.
type ClusterRoleBindingProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// The role to bind to.
	Role IClusterRole `field:"required" json:"role" yaml:"role"`
}

// Policy rule of a `ClusterRole.
type ClusterRolePolicyRule struct {
	// Endpoints this rule applies to.
	//
	// Can be either api resources
	// or non api resources.
	Endpoints *[]IApiEndpoint `field:"required" json:"endpoints" yaml:"endpoints"`
	// Verbs to allow.
	//
	// (e.g ['get', 'watch'])
	Verbs *[]*string `field:"required" json:"verbs" yaml:"verbs"`
}

// Properties for `ClusterRole`.
type ClusterRoleProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// Specify labels that should be used to locate ClusterRoles, whose rules will be automatically filled into this ClusterRole's rules.
	AggregationLabels *map[string]*string `field:"optional" json:"aggregationLabels" yaml:"aggregationLabels"`
	// A list of rules the role should allow.
	Rules *[]*ClusterRolePolicyRule `field:"optional" json:"rules" yaml:"rules"`
}

// Options for `Probe.fromCommand()`.
type CommandProbeOptions struct {
	// Minimum consecutive failures for the probe to be considered failed after having succeeded.
	//
	// Defaults to 3. Minimum value is 1.
	FailureThreshold *float64 `field:"optional" json:"failureThreshold" yaml:"failureThreshold"`
	// Number of seconds after the container has started before liveness probes are initiated.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	//
	InitialDelaySeconds cdk8s.Duration `field:"optional" json:"initialDelaySeconds" yaml:"initialDelaySeconds"`
	// How often (in seconds) to perform the probe.
	//
	// Default to 10 seconds. Minimum value is 1.
	PeriodSeconds cdk8s.Duration `field:"optional" json:"periodSeconds" yaml:"periodSeconds"`
	// Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1.
	//
	// Must be 1 for liveness and startup. Minimum value is 1.
	SuccessThreshold *float64 `field:"optional" json:"successThreshold" yaml:"successThreshold"`
	// Number of seconds after which the probe times out.
	//
	// Defaults to 1 second. Minimum value is 1.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	//
	TimeoutSeconds cdk8s.Duration `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
}

// Common properties for `Secret`.
type CommonSecretProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// If set to true, ensures that data stored in the Secret cannot be updated (only object metadata can be modified).
	//
	// If not set to true, the field can be modified at any time.
	Immutable *bool `field:"optional" json:"immutable" yaml:"immutable"`
}

// ConfigMap holds configuration data for pods to consume.
type ConfigMap interface {
	Resource
	IConfigMap
	// The group portion of the API version (e.g. "authorization.k8s.io").
	ApiGroup() *string
	// The underlying cdk8s API object.
	// See: base.Resource.apiObject
	//
	ApiObject() cdk8s.ApiObject
	// The object's API version (e.g. "authorization.k8s.io/v1").
	ApiVersion() *string
	// The binary data associated with this config map.
	//
	// Returns a copy. To add data records, use `addBinaryData()` or `addData()`.
	BinaryData() *map[string]*string
	// The data associated with this config map.
	//
	// Returns an copy. To add data records, use `addData()` or `addBinaryData()`.
	Data() *map[string]*string
	// Whether or not this config map is immutable.
	Immutable() *bool
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
	// Adds a binary data entry to the config map.
	//
	// BinaryData can contain byte
	// sequences that are not in the UTF-8 range.
	AddBinaryData(key *string, value *string)
	// Adds a data entry to the config map.
	AddData(key *string, value *string)
	// Adds a directory to the ConfigMap.
	AddDirectory(localDir *string, options *AddDirectoryOptions)
	// Adds a file to the ConfigMap.
	AddFile(localFile *string, key *string)
	// Return the IApiResource this object represents.
	AsApiResource() IApiResource
	// Return the non resource url this object represents.
	AsNonApiResource() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for ConfigMap
type jsiiProxy_ConfigMap struct {
	jsiiProxy_Resource
	jsiiProxy_IConfigMap
}

func (j *jsiiProxy_ConfigMap) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigMap) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigMap) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigMap) BinaryData() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"binaryData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigMap) Data() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"data",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigMap) Immutable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"immutable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigMap) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigMap) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigMap) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigMap) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigMap) Permissions() ResourcePermissions {
	var returns ResourcePermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigMap) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigMap) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}


func NewConfigMap(scope constructs.Construct, id *string, props *ConfigMapProps) ConfigMap {
	_init_.Initialize()

	j := jsiiProxy_ConfigMap{}

	_jsii_.Create(
		"cdk8s-plus-24.ConfigMap",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewConfigMap_Override(c ConfigMap, scope constructs.Construct, id *string, props *ConfigMapProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.ConfigMap",
		[]interface{}{scope, id, props},
		c,
	)
}

// Represents a ConfigMap created elsewhere.
func ConfigMap_FromConfigMapName(scope constructs.Construct, id *string, name *string) IConfigMap {
	_init_.Initialize()

	var returns IConfigMap

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.ConfigMap",
		"fromConfigMapName",
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
func ConfigMap_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.ConfigMap",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigMap) AddBinaryData(key *string, value *string) {
	_jsii_.InvokeVoid(
		c,
		"addBinaryData",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_ConfigMap) AddData(key *string, value *string) {
	_jsii_.InvokeVoid(
		c,
		"addData",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_ConfigMap) AddDirectory(localDir *string, options *AddDirectoryOptions) {
	_jsii_.InvokeVoid(
		c,
		"addDirectory",
		[]interface{}{localDir, options},
	)
}

func (c *jsiiProxy_ConfigMap) AddFile(localFile *string, key *string) {
	_jsii_.InvokeVoid(
		c,
		"addFile",
		[]interface{}{localFile, key},
	)
}

func (c *jsiiProxy_ConfigMap) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		c,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigMap) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigMap) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for initialization of `ConfigMap`.
type ConfigMapProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// BinaryData contains the binary data.
	//
	// Each key must consist of alphanumeric characters, '-', '_' or '.'.
	// BinaryData can contain byte sequences that are not in the UTF-8 range. The
	// keys stored in BinaryData must not overlap with the ones in the Data field,
	// this is enforced during validation process.
	//
	// You can also add binary data using `configMap.addBinaryData()`.
	BinaryData *map[string]*string `field:"optional" json:"binaryData" yaml:"binaryData"`
	// Data contains the configuration data.
	//
	// Each key must consist of alphanumeric characters, '-', '_' or '.'. Values
	// with non-UTF-8 byte sequences must use the BinaryData field. The keys
	// stored in Data must not overlap with the keys in the BinaryData field, this
	// is enforced during validation process.
	//
	// You can also add data using `configMap.addData()`.
	Data *map[string]*string `field:"optional" json:"data" yaml:"data"`
	// If set to true, ensures that data stored in the ConfigMap cannot be updated (only object metadata can be modified).
	//
	// If not set to true, the field can be modified at any time.
	Immutable *bool `field:"optional" json:"immutable" yaml:"immutable"`
}

// Options for the ConfigMap-based volume.
type ConfigMapVolumeOptions struct {
	// Mode bits to use on created files by default.
	//
	// Must be a value between 0 and
	// 0777. Defaults to 0644. Directories within the path are not affected by
	// this setting. This might be in conflict with other options that affect the
	// file mode, like fsGroup, and the result can be other mode bits set.
	DefaultMode *float64 `field:"optional" json:"defaultMode" yaml:"defaultMode"`
	// If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value.
	//
	// If specified, the listed keys will be projected
	// into the specified paths, and unlisted keys will not be present. If a key
	// is specified which is not present in the ConfigMap, the volume setup will
	// error unless it is marked optional. Paths must be relative and may not
	// contain the '..' path or start with '..'.
	Items *map[string]*PathMapping `field:"optional" json:"items" yaml:"items"`
	// The volume name.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specify whether the ConfigMap or its keys must be defined.
	Optional *bool `field:"optional" json:"optional" yaml:"optional"`
}

// A single application container that you want to run within a pod.
type Container interface {
	// Arguments to the entrypoint.
	//
	// Returns: a copy of the arguments array, cannot be modified.
	Args() *[]*string
	// Entrypoint array (the command to execute when the container starts).
	//
	// Returns: a copy of the entrypoint array, cannot be modified.
	Command() *[]*string
	// The environment of the container.
	Env() Env
	// The container image.
	Image() *string
	// Image pull policy for this container.
	ImagePullPolicy() ImagePullPolicy
	// Volume mounts configured for this container.
	Mounts() *[]*VolumeMount
	// The name of the container.
	Name() *string
	// The port this container exposes.
	Port() *float64
	// Compute resources (CPU and memory requests and limits) required by the container.
	// See: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	//
	Resources() *ContainerResources
	// The security context of the container.
	SecurityContext() ContainerSecurityContext
	// The working directory inside the container.
	WorkingDir() *string
	// Mount a volume to a specific path so that it is accessible by the container.
	//
	// Every pod that is configured to use this container will autmoatically have access to the volume.
	Mount(path *string, storage IStorage, options *MountOptions)
}

// The jsii proxy struct for Container
type jsiiProxy_Container struct {
	_ byte // padding
}

func (j *jsiiProxy_Container) Args() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"args",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Container) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Container) Env() Env {
	var returns Env
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Container) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Container) ImagePullPolicy() ImagePullPolicy {
	var returns ImagePullPolicy
	_jsii_.Get(
		j,
		"imagePullPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Container) Mounts() *[]*VolumeMount {
	var returns *[]*VolumeMount
	_jsii_.Get(
		j,
		"mounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Container) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Container) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Container) Resources() *ContainerResources {
	var returns *ContainerResources
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Container) SecurityContext() ContainerSecurityContext {
	var returns ContainerSecurityContext
	_jsii_.Get(
		j,
		"securityContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Container) WorkingDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDir",
		&returns,
	)
	return returns
}


func NewContainer(props *ContainerProps) Container {
	_init_.Initialize()

	j := jsiiProxy_Container{}

	_jsii_.Create(
		"cdk8s-plus-24.Container",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewContainer_Override(c Container, props *ContainerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.Container",
		[]interface{}{props},
		c,
	)
}

func (c *jsiiProxy_Container) Mount(path *string, storage IStorage, options *MountOptions) {
	_jsii_.InvokeVoid(
		c,
		"mount",
		[]interface{}{path, storage, options},
	)
}

// Container lifecycle properties.
type ContainerLifecycle struct {
	// This hook is executed immediately after a container is created.
	//
	// However,
	// there is no guarantee that the hook will execute before the container ENTRYPOINT.
	PostStart Handler `field:"optional" json:"postStart" yaml:"postStart"`
	// This hook is called immediately before a container is terminated due to an API request or management event such as a liveness/startup probe failure, preemption, resource contention and others.
	//
	// A call to the PreStop hook fails if the container is already in a terminated or completed state
	// and the hook must complete before the TERM signal to stop the container can be sent.
	// The Pod's termination grace period countdown begins before the PreStop hook is executed,
	// so regardless of the outcome of the handler, the container will eventually terminate
	// within the Pod's termination grace period. No parameters are passed to the handler.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#pod-termination
	//
	PreStop Handler `field:"optional" json:"preStop" yaml:"preStop"`
}

// Properties for creating a container.
type ContainerProps struct {
	// Docker image name.
	Image *string `field:"required" json:"image" yaml:"image"`
	// Arguments to the entrypoint. The docker image's CMD is used if `command` is not provided.
	//
	// Variable references $(VAR_NAME) are expanded using the container's
	// environment. If a variable cannot be resolved, the reference in the input
	// string will be unchanged. The $(VAR_NAME) syntax can be escaped with a
	// double $$, ie: $$(VAR_NAME). Escaped references will never be expanded,
	// regardless of whether the variable exists or not.
	//
	// Cannot be updated.
	// See: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	//
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// Entrypoint array.
	//
	// Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment.
	// If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME).
	// Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated.
	// More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// List of sources to populate environment variables in the container.
	//
	// When a key exists in multiple sources, the value associated with
	// the last source will take precedence. Values defined by the `envVariables` property
	// with a duplicate key will take precedence.
	EnvFrom *[]EnvFrom `field:"optional" json:"envFrom" yaml:"envFrom"`
	// Environment variables to set in the container.
	EnvVariables *map[string]EnvValue `field:"optional" json:"envVariables" yaml:"envVariables"`
	// Image pull policy for this container.
	ImagePullPolicy ImagePullPolicy `field:"optional" json:"imagePullPolicy" yaml:"imagePullPolicy"`
	// Describes actions that the management system should take in response to container lifecycle events.
	Lifecycle *ContainerLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Periodic probe of container liveness.
	//
	// Container will be restarted if the probe fails.
	Liveness Probe `field:"optional" json:"liveness" yaml:"liveness"`
	// Name of the container specified as a DNS_LABEL.
	//
	// Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Number of port to expose on the pod's IP address.
	//
	// This must be a valid port number, 0 < x < 65536.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Determines when the container is ready to serve traffic.
	Readiness Probe `field:"optional" json:"readiness" yaml:"readiness"`
	// Compute resources (CPU and memory requests and limits) required by the container.
	// See: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	//
	Resources *ContainerResources `field:"optional" json:"resources" yaml:"resources"`
	// SecurityContext defines the security options the container should be run with.
	//
	// If set, the fields override equivalent fields of the pod's security context.
	// See: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
	//
	SecurityContext *ContainerSecurityContextProps `field:"optional" json:"securityContext" yaml:"securityContext"`
	// StartupProbe indicates that the Pod has successfully initialized.
	//
	// If specified, no other probes are executed until this completes successfully.
	Startup Probe `field:"optional" json:"startup" yaml:"startup"`
	// Pod volumes to mount into the container's filesystem.
	//
	// Cannot be updated.
	VolumeMounts *[]*VolumeMount `field:"optional" json:"volumeMounts" yaml:"volumeMounts"`
	// Container's working directory.
	//
	// If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.
	WorkingDir *string `field:"optional" json:"workingDir" yaml:"workingDir"`
}

// CPU and memory compute resources.
type ContainerResources struct {
	Cpu *CpuResources `field:"optional" json:"cpu" yaml:"cpu"`
	Memory *MemoryResources `field:"optional" json:"memory" yaml:"memory"`
}

// Container security attributes and settings.
type ContainerSecurityContext interface {
	EnsureNonRoot() *bool
	Group() *float64
	Privileged() *bool
	ReadOnlyRootFilesystem() *bool
	User() *float64
}

// The jsii proxy struct for ContainerSecurityContext
type jsiiProxy_ContainerSecurityContext struct {
	_ byte // padding
}

func (j *jsiiProxy_ContainerSecurityContext) EnsureNonRoot() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"ensureNonRoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerSecurityContext) Group() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerSecurityContext) Privileged() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"privileged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerSecurityContext) ReadOnlyRootFilesystem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readOnlyRootFilesystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerSecurityContext) User() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}


func NewContainerSecurityContext(props *ContainerSecurityContextProps) ContainerSecurityContext {
	_init_.Initialize()

	j := jsiiProxy_ContainerSecurityContext{}

	_jsii_.Create(
		"cdk8s-plus-24.ContainerSecurityContext",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewContainerSecurityContext_Override(c ContainerSecurityContext, props *ContainerSecurityContextProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.ContainerSecurityContext",
		[]interface{}{props},
		c,
	)
}

// Properties for `ContainerSecurityContext`.
type ContainerSecurityContextProps struct {
	// Indicates that the container must run as a non-root user.
	//
	// If true, the Kubelet will validate the image at runtime to ensure that it does
	// not run as UID 0 (root) and fail to start the container if it does.
	EnsureNonRoot *bool `field:"optional" json:"ensureNonRoot" yaml:"ensureNonRoot"`
	// The GID to run the entrypoint of the container process.
	Group *float64 `field:"optional" json:"group" yaml:"group"`
	// Run container in privileged mode.
	//
	// Processes in privileged containers are essentially equivalent to root on the host.
	Privileged *bool `field:"optional" json:"privileged" yaml:"privileged"`
	// Whether this container has a read-only root filesystem.
	ReadOnlyRootFilesystem *bool `field:"optional" json:"readOnlyRootFilesystem" yaml:"readOnlyRootFilesystem"`
	// The UID to run the entrypoint of the container process.
	User *float64 `field:"optional" json:"user" yaml:"user"`
}

// Represents the amount of CPU.
//
// The amount can be passed as millis or units.
type Cpu interface {
	Amount() *string
	SetAmount(val *string)
}

// The jsii proxy struct for Cpu
type jsiiProxy_Cpu struct {
	_ byte // padding
}

func (j *jsiiProxy_Cpu) Amount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amount",
		&returns,
	)
	return returns
}


func (j *jsiiProxy_Cpu) SetAmount(val *string) {
	_jsii_.Set(
		j,
		"amount",
		val,
	)
}

func Cpu_Millis(amount *float64) Cpu {
	_init_.Initialize()

	var returns Cpu

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Cpu",
		"millis",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

func Cpu_Units(amount *float64) Cpu {
	_init_.Initialize()

	var returns Cpu

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Cpu",
		"units",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

// CPU request and limit.
type CpuResources struct {
	Limit Cpu `field:"optional" json:"limit" yaml:"limit"`
	Request Cpu `field:"optional" json:"request" yaml:"request"`
}

// A DaemonSet ensures that all (or some) Nodes run a copy of a Pod.
//
// As nodes are added to the cluster, Pods are added to them.
// As nodes are removed from the cluster, those Pods are garbage collected.
// Deleting a DaemonSet will clean up the Pods it created.
//
// Some typical uses of a DaemonSet are:
//
// - running a cluster storage daemon on every node
// - running a logs collection daemon on every node
// - running a node monitoring daemon on every node
//
// In a simple case, one DaemonSet, covering all nodes, would be used for each type of daemon.
// A more complex setup might use multiple DaemonSets for a single type of daemon,
// but with different flags and/or different memory and cpu requests for different hardware types.
type DaemonSet interface {
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
	Connections() PodConnections
	Containers() *[]Container
	Dns() PodDns
	DockerRegistryAuth() DockerConfigSecret
	HostAliases() *[]*HostAlias
	InitContainers() *[]Container
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
	MinReadySeconds() *float64
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

// The jsii proxy struct for DaemonSet
type jsiiProxy_DaemonSet struct {
	jsiiProxy_Workload
}

func (j *jsiiProxy_DaemonSet) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSet) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSet) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSet) AutomountServiceAccountToken() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"automountServiceAccountToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSet) Connections() PodConnections {
	var returns PodConnections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSet) Containers() *[]Container {
	var returns *[]Container
	_jsii_.Get(
		j,
		"containers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSet) Dns() PodDns {
	var returns PodDns
	_jsii_.Get(
		j,
		"dns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSet) DockerRegistryAuth() DockerConfigSecret {
	var returns DockerConfigSecret
	_jsii_.Get(
		j,
		"dockerRegistryAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSet) HostAliases() *[]*HostAlias {
	var returns *[]*HostAlias
	_jsii_.Get(
		j,
		"hostAliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSet) InitContainers() *[]Container {
	var returns *[]Container
	_jsii_.Get(
		j,
		"initContainers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSet) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSet) MatchExpressions() *[]*LabelSelectorRequirement {
	var returns *[]*LabelSelectorRequirement
	_jsii_.Get(
		j,
		"matchExpressions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSet) MatchLabels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"matchLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSet) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSet) MinReadySeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minReadySeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSet) Permissions() ResourcePermissions {
	var returns ResourcePermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSet) PodMetadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"podMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSet) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSet) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSet) RestartPolicy() RestartPolicy {
	var returns RestartPolicy
	_jsii_.Get(
		j,
		"restartPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSet) Scheduling() WorkloadScheduling {
	var returns WorkloadScheduling
	_jsii_.Get(
		j,
		"scheduling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSet) SecurityContext() PodSecurityContext {
	var returns PodSecurityContext
	_jsii_.Get(
		j,
		"securityContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSet) ServiceAccount() IServiceAccount {
	var returns IServiceAccount
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSet) Volumes() *[]Volume {
	var returns *[]Volume
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}


func NewDaemonSet(scope constructs.Construct, id *string, props *DaemonSetProps) DaemonSet {
	_init_.Initialize()

	j := jsiiProxy_DaemonSet{}

	_jsii_.Create(
		"cdk8s-plus-24.DaemonSet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDaemonSet_Override(d DaemonSet, scope constructs.Construct, id *string, props *DaemonSetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.DaemonSet",
		[]interface{}{scope, id, props},
		d,
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
func DaemonSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.DaemonSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonSet) AddContainer(cont *ContainerProps) Container {
	var returns Container

	_jsii_.Invoke(
		d,
		"addContainer",
		[]interface{}{cont},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonSet) AddHostAlias(hostAlias *HostAlias) {
	_jsii_.InvokeVoid(
		d,
		"addHostAlias",
		[]interface{}{hostAlias},
	)
}

func (d *jsiiProxy_DaemonSet) AddInitContainer(cont *ContainerProps) Container {
	var returns Container

	_jsii_.Invoke(
		d,
		"addInitContainer",
		[]interface{}{cont},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonSet) AddVolume(vol Volume) {
	_jsii_.InvokeVoid(
		d,
		"addVolume",
		[]interface{}{vol},
	)
}

func (d *jsiiProxy_DaemonSet) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		d,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonSet) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonSet) Select(selectors ...LabelSelector) {
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

func (d *jsiiProxy_DaemonSet) ToNetworkPolicyPeerConfig() *NetworkPolicyPeerConfig {
	var returns *NetworkPolicyPeerConfig

	_jsii_.Invoke(
		d,
		"toNetworkPolicyPeerConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonSet) ToPodSelector() IPodSelector {
	var returns IPodSelector

	_jsii_.Invoke(
		d,
		"toPodSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonSet) ToPodSelectorConfig() *PodSelectorConfig {
	var returns *PodSelectorConfig

	_jsii_.Invoke(
		d,
		"toPodSelectorConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonSet) ToSubjectConfiguration() *SubjectConfiguration {
	var returns *SubjectConfiguration

	_jsii_.Invoke(
		d,
		"toSubjectConfiguration",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for `DaemonSet`.
type DaemonSetProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// Indicates whether a service account token should be automatically mounted.
	// See: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/#use-the-default-service-account-to-access-the-api-server
	//
	AutomountServiceAccountToken *bool `field:"optional" json:"automountServiceAccountToken" yaml:"automountServiceAccountToken"`
	// List of containers belonging to the pod.
	//
	// Containers cannot currently be
	// added or removed. There must be at least one container in a Pod.
	//
	// You can add additionnal containers using `podSpec.addContainer()`
	Containers *[]*ContainerProps `field:"optional" json:"containers" yaml:"containers"`
	// DNS settings for the pod.
	// See: https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/
	//
	Dns *PodDnsProps `field:"optional" json:"dns" yaml:"dns"`
	// A secret containing docker credentials for authenticating to a registry.
	DockerRegistryAuth DockerConfigSecret `field:"optional" json:"dockerRegistryAuth" yaml:"dockerRegistryAuth"`
	// HostAlias holds the mapping between IP and hostnames that will be injected as an entry in the pod's hosts file.
	HostAliases *[]*HostAlias `field:"optional" json:"hostAliases" yaml:"hostAliases"`
	// List of initialization containers belonging to the pod.
	//
	// Init containers are executed in order prior to containers being started.
	// If any init container fails, the pod is considered to have failed and is handled according to its restartPolicy.
	// The name for an init container or normal container must be unique among all containers.
	// Init containers may not have Lifecycle actions, Readiness probes, Liveness probes, or Startup probes.
	// The resourceRequirements of an init container are taken into account during scheduling by finding the highest request/limit
	// for each resource type, and then using the max of of that value or the sum of the normal containers.
	// Limits are applied to init containers in a similar fashion.
	//
	// Init containers cannot currently be added ,removed or updated.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/
	//
	InitContainers *[]*ContainerProps `field:"optional" json:"initContainers" yaml:"initContainers"`
	// Restart policy for all containers within the pod.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy
	//
	RestartPolicy RestartPolicy `field:"optional" json:"restartPolicy" yaml:"restartPolicy"`
	// SecurityContext holds pod-level security attributes and common container settings.
	SecurityContext *PodSecurityContextProps `field:"optional" json:"securityContext" yaml:"securityContext"`
	// A service account provides an identity for processes that run in a Pod.
	//
	// When you (a human) access the cluster (for example, using kubectl), you are
	// authenticated by the apiserver as a particular User Account (currently this
	// is usually admin, unless your cluster administrator has customized your
	// cluster). Processes in containers inside pods can also contact the
	// apiserver. When they do, they are authenticated as a particular Service
	// Account (for example, default).
	// See: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
	//
	ServiceAccount IServiceAccount `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// List of volumes that can be mounted by containers belonging to the pod.
	//
	// You can also add volumes later using `podSpec.addVolume()`
	// See: https://kubernetes.io/docs/concepts/storage/volumes
	//
	Volumes *[]Volume `field:"optional" json:"volumes" yaml:"volumes"`
	// The pod metadata of this workload.
	PodMetadata *cdk8s.ApiObjectMetadata `field:"optional" json:"podMetadata" yaml:"podMetadata"`
	// Automatically allocates a pod label selector for this workload and add it to the pod metadata.
	//
	// This ensures this workload manages pods created by
	// its pod template.
	Select *bool `field:"optional" json:"select" yaml:"select"`
	// Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available.
	MinReadySeconds *float64 `field:"optional" json:"minReadySeconds" yaml:"minReadySeconds"`
}

// A Deployment provides declarative updates for Pods and ReplicaSets.
//
// You describe a desired state in a Deployment, and the Deployment Controller changes the actual
// state to the desired state at a controlled rate. You can define Deployments to create new ReplicaSets, or to remove
// existing Deployments and adopt all their resources with new Deployments.
//
// > Note: Do not manage ReplicaSets owned by a Deployment. Consider opening an issue in the main Kubernetes repository if your use case is not covered below.
//
// Use Case
// ---------
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
	// Expose a deployment via an ingress.
	//
	// This will first expose the deployment with a service, and then expose the service via an ingress.
	ExposeViaIngress(path *string, options *ExposeDeploymentViaIngressOptions) Ingress
	// Expose a deployment via a service.
	//
	// This is equivalent to running `kubectl expose deployment <deployment-name>`.
	ExposeViaService(options *DeploymentExposeViaServiceOptions) Service
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

// The jsii proxy struct for Deployment
type jsiiProxy_Deployment struct {
	jsiiProxy_Workload
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

func (j *jsiiProxy_Deployment) DockerRegistryAuth() DockerConfigSecret {
	var returns DockerConfigSecret
	_jsii_.Get(
		j,
		"dockerRegistryAuth",
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

func (j *jsiiProxy_Deployment) InitContainers() *[]Container {
	var returns *[]Container
	_jsii_.Get(
		j,
		"initContainers",
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

	j := jsiiProxy_Deployment{}

	_jsii_.Create(
		"cdk8s-plus-24.Deployment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDeployment_Override(d Deployment, scope constructs.Construct, id *string, props *DeploymentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.Deployment",
		[]interface{}{scope, id, props},
		d,
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

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Deployment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Deployment) AddContainer(cont *ContainerProps) Container {
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
	_jsii_.InvokeVoid(
		d,
		"addHostAlias",
		[]interface{}{hostAlias},
	)
}

func (d *jsiiProxy_Deployment) AddInitContainer(cont *ContainerProps) Container {
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

func (d *jsiiProxy_Deployment) ExposeViaIngress(path *string, options *ExposeDeploymentViaIngressOptions) Ingress {
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
	var returns Service

	_jsii_.Invoke(
		d,
		"exposeViaService",
		[]interface{}{options},
		&returns,
	)

	return returns
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

// Options for `Deployment.exposeViaService`.
type DeploymentExposeViaServiceOptions struct {
	// The name of the service to expose.
	//
	// If you'd like to expose the deployment multiple times,
	// you must explicitly set a name starting from the second expose call.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ports that the service should bind to.
	Ports *[]*ServicePort `field:"optional" json:"ports" yaml:"ports"`
	// The type of the exposed service.
	ServiceType ServiceType `field:"optional" json:"serviceType" yaml:"serviceType"`
}

// Properties for `Deployment`.
type DeploymentProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// Indicates whether a service account token should be automatically mounted.
	// See: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/#use-the-default-service-account-to-access-the-api-server
	//
	AutomountServiceAccountToken *bool `field:"optional" json:"automountServiceAccountToken" yaml:"automountServiceAccountToken"`
	// List of containers belonging to the pod.
	//
	// Containers cannot currently be
	// added or removed. There must be at least one container in a Pod.
	//
	// You can add additionnal containers using `podSpec.addContainer()`
	Containers *[]*ContainerProps `field:"optional" json:"containers" yaml:"containers"`
	// DNS settings for the pod.
	// See: https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/
	//
	Dns *PodDnsProps `field:"optional" json:"dns" yaml:"dns"`
	// A secret containing docker credentials for authenticating to a registry.
	DockerRegistryAuth DockerConfigSecret `field:"optional" json:"dockerRegistryAuth" yaml:"dockerRegistryAuth"`
	// HostAlias holds the mapping between IP and hostnames that will be injected as an entry in the pod's hosts file.
	HostAliases *[]*HostAlias `field:"optional" json:"hostAliases" yaml:"hostAliases"`
	// List of initialization containers belonging to the pod.
	//
	// Init containers are executed in order prior to containers being started.
	// If any init container fails, the pod is considered to have failed and is handled according to its restartPolicy.
	// The name for an init container or normal container must be unique among all containers.
	// Init containers may not have Lifecycle actions, Readiness probes, Liveness probes, or Startup probes.
	// The resourceRequirements of an init container are taken into account during scheduling by finding the highest request/limit
	// for each resource type, and then using the max of of that value or the sum of the normal containers.
	// Limits are applied to init containers in a similar fashion.
	//
	// Init containers cannot currently be added ,removed or updated.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/
	//
	InitContainers *[]*ContainerProps `field:"optional" json:"initContainers" yaml:"initContainers"`
	// Restart policy for all containers within the pod.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy
	//
	RestartPolicy RestartPolicy `field:"optional" json:"restartPolicy" yaml:"restartPolicy"`
	// SecurityContext holds pod-level security attributes and common container settings.
	SecurityContext *PodSecurityContextProps `field:"optional" json:"securityContext" yaml:"securityContext"`
	// A service account provides an identity for processes that run in a Pod.
	//
	// When you (a human) access the cluster (for example, using kubectl), you are
	// authenticated by the apiserver as a particular User Account (currently this
	// is usually admin, unless your cluster administrator has customized your
	// cluster). Processes in containers inside pods can also contact the
	// apiserver. When they do, they are authenticated as a particular Service
	// Account (for example, default).
	// See: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
	//
	ServiceAccount IServiceAccount `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// List of volumes that can be mounted by containers belonging to the pod.
	//
	// You can also add volumes later using `podSpec.addVolume()`
	// See: https://kubernetes.io/docs/concepts/storage/volumes
	//
	Volumes *[]Volume `field:"optional" json:"volumes" yaml:"volumes"`
	// The pod metadata of this workload.
	PodMetadata *cdk8s.ApiObjectMetadata `field:"optional" json:"podMetadata" yaml:"podMetadata"`
	// Automatically allocates a pod label selector for this workload and add it to the pod metadata.
	//
	// This ensures this workload manages pods created by
	// its pod template.
	Select *bool `field:"optional" json:"select" yaml:"select"`
	// Minimum duration for which a newly created pod should be ready without any of its container crashing, for it to be considered available.
	//
	// Zero means the pod will be considered available as soon as it is ready.
	// See: https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#min-ready-seconds
	//
	MinReady cdk8s.Duration `field:"optional" json:"minReady" yaml:"minReady"`
	// The maximum duration for a deployment to make progress before it is considered to be failed.
	//
	// The deployment controller will continue
	// to process failed deployments and a condition with a ProgressDeadlineExceeded
	// reason will be surfaced in the deployment status.
	//
	// Note that progress will not be estimated during the time a deployment is paused.
	// See: https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#progress-deadline-seconds
	//
	ProgressDeadline cdk8s.Duration `field:"optional" json:"progressDeadline" yaml:"progressDeadline"`
	// Number of desired pods.
	Replicas *float64 `field:"optional" json:"replicas" yaml:"replicas"`
	// Specifies the strategy used to replace old Pods by new ones.
	Strategy DeploymentStrategy `field:"optional" json:"strategy" yaml:"strategy"`
}

// Deployment strategies.
type DeploymentStrategy interface {
}

// The jsii proxy struct for DeploymentStrategy
type jsiiProxy_DeploymentStrategy struct {
	_ byte // padding
}

// All existing Pods are killed before new ones are created.
// See: https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#recreate-deployment
//
func DeploymentStrategy_Recreate() DeploymentStrategy {
	_init_.Initialize()

	var returns DeploymentStrategy

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.DeploymentStrategy",
		"recreate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func DeploymentStrategy_RollingUpdate(options *DeploymentStrategyRollingUpdateOptions) DeploymentStrategy {
	_init_.Initialize()

	var returns DeploymentStrategy

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.DeploymentStrategy",
		"rollingUpdate",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Options for `DeploymentStrategy.rollingUpdate`.
type DeploymentStrategyRollingUpdateOptions struct {
	// The maximum number of pods that can be scheduled above the desired number of pods.
	//
	// Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%).
	// Absolute number is calculated from percentage by rounding up.
	// This can not be 0 if `maxUnavailable` is 0.
	//
	// Example: when this is set to 30%, the new ReplicaSet can be scaled up immediately when the rolling update
	// starts, such that the total number of old and new pods do not exceed 130% of desired pods.
	// Once old pods have been killed, new ReplicaSet can be scaled up further, ensuring that
	// total number of pods running at any time during the update is at most 130% of desired pods.
	MaxSurge PercentOrAbsolute `field:"optional" json:"maxSurge" yaml:"maxSurge"`
	// The maximum number of pods that can be unavailable during the update.
	//
	// Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%).
	// Absolute number is calculated from percentage by rounding down.
	// This can not be 0 if `maxSurge` is 0.
	//
	// Example: when this is set to 30%, the old ReplicaSet can be scaled down to 70% of desired
	// pods immediately when the rolling update starts. Once new pods are ready, old ReplicaSet can
	// be scaled down further, followed by scaling up the new ReplicaSet, ensuring that the total
	// number of pods available at all times during the update is at least 70% of desired pods.
	MaxUnavailable PercentOrAbsolute `field:"optional" json:"maxUnavailable" yaml:"maxUnavailable"`
}

// Custom DNS option.
type DnsOption struct {
	// Option name.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Option value.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

// Pod DNS policies.
type DnsPolicy string

const (
	// Any DNS query that does not match the configured cluster domain suffix, such as "www.kubernetes.io", is forwarded to the upstream nameserver inherited from the node. Cluster administrators may have extra stub-domain and upstream DNS servers configured.
	DnsPolicy_CLUSTER_FIRST DnsPolicy = "CLUSTER_FIRST"
	// For Pods running with hostNetwork, you should explicitly set its DNS policy "ClusterFirstWithHostNet".
	DnsPolicy_CLUSTER_FIRST_WITH_HOST_NET DnsPolicy = "CLUSTER_FIRST_WITH_HOST_NET"
	// The Pod inherits the name resolution configuration from the node that the pods run on.
	DnsPolicy_DEFAULT DnsPolicy = "DEFAULT"
	// It allows a Pod to ignore DNS settings from the Kubernetes environment.
	//
	// All DNS settings are supposed to be provided using the dnsConfig
	// field in the Pod Spec.
	DnsPolicy_NONE DnsPolicy = "NONE"
)

// Create a secret for storing credentials for accessing a container image registry.
// See: https://kubernetes.io/docs/concepts/configuration/secret/#docker-config-secrets
//
type DockerConfigSecret interface {
	Secret
	// The group portion of the API version (e.g. "authorization.k8s.io").
	ApiGroup() *string
	// The underlying cdk8s API object.
	// See: base.Resource.apiObject
	//
	ApiObject() cdk8s.ApiObject
	// The object's API version (e.g. "authorization.k8s.io/v1").
	ApiVersion() *string
	// Whether or not the secret is immutable.
	Immutable() *bool
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
	// Adds a string data field to the secert.
	AddStringData(key *string, value *string)
	// Return the IApiResource this object represents.
	AsApiResource() IApiResource
	// Return the non resource url this object represents.
	AsNonApiResource() *string
	// Gets a string data by key or undefined.
	GetStringData(key *string) *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for DockerConfigSecret
type jsiiProxy_DockerConfigSecret struct {
	jsiiProxy_Secret
}

func (j *jsiiProxy_DockerConfigSecret) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerConfigSecret) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerConfigSecret) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerConfigSecret) Immutable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"immutable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerConfigSecret) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerConfigSecret) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerConfigSecret) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerConfigSecret) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerConfigSecret) Permissions() ResourcePermissions {
	var returns ResourcePermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerConfigSecret) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerConfigSecret) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}


func NewDockerConfigSecret(scope constructs.Construct, id *string, props *DockerConfigSecretProps) DockerConfigSecret {
	_init_.Initialize()

	j := jsiiProxy_DockerConfigSecret{}

	_jsii_.Create(
		"cdk8s-plus-24.DockerConfigSecret",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDockerConfigSecret_Override(d DockerConfigSecret, scope constructs.Construct, id *string, props *DockerConfigSecretProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.DockerConfigSecret",
		[]interface{}{scope, id, props},
		d,
	)
}

// Imports a secret from the cluster as a reference.
func DockerConfigSecret_FromSecretName(scope constructs.Construct, id *string, name *string) ISecret {
	_init_.Initialize()

	var returns ISecret

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.DockerConfigSecret",
		"fromSecretName",
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
func DockerConfigSecret_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.DockerConfigSecret",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DockerConfigSecret) AddStringData(key *string, value *string) {
	_jsii_.InvokeVoid(
		d,
		"addStringData",
		[]interface{}{key, value},
	)
}

func (d *jsiiProxy_DockerConfigSecret) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		d,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DockerConfigSecret) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DockerConfigSecret) GetStringData(key *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringData",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DockerConfigSecret) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options for `DockerConfigSecret`.
type DockerConfigSecretProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// If set to true, ensures that data stored in the Secret cannot be updated (only object metadata can be modified).
	//
	// If not set to true, the field can be modified at any time.
	Immutable *bool `field:"optional" json:"immutable" yaml:"immutable"`
	// JSON content to provide for the `~/.docker/config.json` file. This will be stringified and inserted as stringData.
	// See: https://docs.docker.com/engine/reference/commandline/cli/#sample-configuration-file
	//
	Data *map[string]interface{} `field:"required" json:"data" yaml:"data"`
}

// The medium on which to store the volume.
type EmptyDirMedium string

const (
	// The default volume of the backing node.
	EmptyDirMedium_DEFAULT EmptyDirMedium = "DEFAULT"
	// Mount a tmpfs (RAM-backed filesystem) for you instead.
	//
	// While tmpfs is very
	// fast, be aware that unlike disks, tmpfs is cleared on node reboot and any
	// files you write will count against your Container's memory limit.
	EmptyDirMedium_MEMORY EmptyDirMedium = "MEMORY"
)

// Options for volumes populated with an empty directory.
type EmptyDirVolumeOptions struct {
	// By default, emptyDir volumes are stored on whatever medium is backing the node - that might be disk or SSD or network storage, depending on your environment.
	//
	// However, you can set the emptyDir.medium field to
	// `EmptyDirMedium.MEMORY` to tell Kubernetes to mount a tmpfs (RAM-backed
	// filesystem) for you instead. While tmpfs is very fast, be aware that unlike
	// disks, tmpfs is cleared on node reboot and any files you write will count
	// against your Container's memory limit.
	Medium EmptyDirMedium `field:"optional" json:"medium" yaml:"medium"`
	// Total amount of local storage required for this EmptyDir volume.
	//
	// The size
	// limit is also applicable for memory medium. The maximum usage on memory
	// medium EmptyDir would be the minimum value between the SizeLimit specified
	// here and the sum of memory limits of all containers in a pod.
	SizeLimit cdk8s.Size `field:"optional" json:"sizeLimit" yaml:"sizeLimit"`
}

// Container environment variables.
type Env interface {
	// The list of sources used to populate the container environment, in addition to the `variables`.
	//
	// Returns a copy. To add a source use `container.env.copyFrom()`.
	Sources() *[]EnvFrom
	// The environment variables for this container.
	//
	// Returns a copy. To add environment variables use `container.env.addVariable()`.
	Variables() *map[string]EnvValue
	// Add a single variable by name and value.
	//
	// The variable value can come from various dynamic sources such a secrets of config maps.
	// Use `EnvValue.fromXXX` to select sources.
	AddVariable(name *string, value EnvValue)
	// Add a collection of variables by copying from another source.
	//
	// Use `Env.fromXXX` functions to select sources.
	CopyFrom(from EnvFrom)
}

// The jsii proxy struct for Env
type jsiiProxy_Env struct {
	_ byte // padding
}

func (j *jsiiProxy_Env) Sources() *[]EnvFrom {
	var returns *[]EnvFrom
	_jsii_.Get(
		j,
		"sources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Env) Variables() *map[string]EnvValue {
	var returns *map[string]EnvValue
	_jsii_.Get(
		j,
		"variables",
		&returns,
	)
	return returns
}


func NewEnv(sources *[]EnvFrom, variables *map[string]EnvValue) Env {
	_init_.Initialize()

	j := jsiiProxy_Env{}

	_jsii_.Create(
		"cdk8s-plus-24.Env",
		[]interface{}{sources, variables},
		&j,
	)

	return &j
}

func NewEnv_Override(e Env, sources *[]EnvFrom, variables *map[string]EnvValue) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.Env",
		[]interface{}{sources, variables},
		e,
	)
}

// Selects a ConfigMap to populate the environment variables with.
//
// The contents of the target ConfigMap's Data field will represent
// the key-value pairs as environment variables.
func Env_FromConfigMap(configMap IConfigMap, prefix *string) EnvFrom {
	_init_.Initialize()

	var returns EnvFrom

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Env",
		"fromConfigMap",
		[]interface{}{configMap, prefix},
		&returns,
	)

	return returns
}

// Selects a Secret to populate the environment variables with.
//
// The contents of the target Secret's Data field will represent
// the key-value pairs as environment variables.
func Env_FromSecret(secr ISecret) EnvFrom {
	_init_.Initialize()

	var returns EnvFrom

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Env",
		"fromSecret",
		[]interface{}{secr},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Env) AddVariable(name *string, value EnvValue) {
	_jsii_.InvokeVoid(
		e,
		"addVariable",
		[]interface{}{name, value},
	)
}

func (e *jsiiProxy_Env) CopyFrom(from EnvFrom) {
	_jsii_.InvokeVoid(
		e,
		"copyFrom",
		[]interface{}{from},
	)
}

type EnvFieldPaths string

const (
	// The name of the pod.
	EnvFieldPaths_POD_NAME EnvFieldPaths = "POD_NAME"
	// The namespace of the pod.
	EnvFieldPaths_POD_NAMESPACE EnvFieldPaths = "POD_NAMESPACE"
	// The uid of the pod.
	EnvFieldPaths_POD_UID EnvFieldPaths = "POD_UID"
	// The labels of the pod.
	EnvFieldPaths_POD_LABEL EnvFieldPaths = "POD_LABEL"
	// The annotations of the pod.
	EnvFieldPaths_POD_ANNOTATION EnvFieldPaths = "POD_ANNOTATION"
	// The ipAddress of the pod.
	EnvFieldPaths_POD_IP EnvFieldPaths = "POD_IP"
	// The service account name of the pod.
	EnvFieldPaths_SERVICE_ACCOUNT_NAME EnvFieldPaths = "SERVICE_ACCOUNT_NAME"
	// The name of the node.
	EnvFieldPaths_NODE_NAME EnvFieldPaths = "NODE_NAME"
	// The ipAddress of the node.
	EnvFieldPaths_NODE_IP EnvFieldPaths = "NODE_IP"
	// The ipAddresess of the pod.
	EnvFieldPaths_POD_IPS EnvFieldPaths = "POD_IPS"
)

// A collection of env variables defined in other resources.
type EnvFrom interface {
}

// The jsii proxy struct for EnvFrom
type jsiiProxy_EnvFrom struct {
	_ byte // padding
}

func NewEnvFrom(configMap IConfigMap, prefix *string, sec ISecret) EnvFrom {
	_init_.Initialize()

	j := jsiiProxy_EnvFrom{}

	_jsii_.Create(
		"cdk8s-plus-24.EnvFrom",
		[]interface{}{configMap, prefix, sec},
		&j,
	)

	return &j
}

func NewEnvFrom_Override(e EnvFrom, configMap IConfigMap, prefix *string, sec ISecret) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.EnvFrom",
		[]interface{}{configMap, prefix, sec},
		e,
	)
}

// Utility class for creating reading env values from various sources.
type EnvValue interface {
	Value() interface{}
	ValueFrom() interface{}
}

// The jsii proxy struct for EnvValue
type jsiiProxy_EnvValue struct {
	_ byte // padding
}

func (j *jsiiProxy_EnvValue) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvValue) ValueFrom() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"valueFrom",
		&returns,
	)
	return returns
}


// Create a value by reading a specific key inside a config map.
func EnvValue_FromConfigMap(configMap IConfigMap, key *string, options *EnvValueFromConfigMapOptions) EnvValue {
	_init_.Initialize()

	var returns EnvValue

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.EnvValue",
		"fromConfigMap",
		[]interface{}{configMap, key, options},
		&returns,
	)

	return returns
}

// Create a value from a field reference.
func EnvValue_FromFieldRef(fieldPath EnvFieldPaths, options *EnvValueFromFieldRefOptions) EnvValue {
	_init_.Initialize()

	var returns EnvValue

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.EnvValue",
		"fromFieldRef",
		[]interface{}{fieldPath, options},
		&returns,
	)

	return returns
}

// Create a value from a key in the current process environment.
func EnvValue_FromProcess(key *string, options *EnvValueFromProcessOptions) EnvValue {
	_init_.Initialize()

	var returns EnvValue

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.EnvValue",
		"fromProcess",
		[]interface{}{key, options},
		&returns,
	)

	return returns
}

// Create a value from a resource.
func EnvValue_FromResource(resource ResourceFieldPaths, options *EnvValueFromResourceOptions) EnvValue {
	_init_.Initialize()

	var returns EnvValue

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.EnvValue",
		"fromResource",
		[]interface{}{resource, options},
		&returns,
	)

	return returns
}

// Defines an environment value from a secret JSON value.
func EnvValue_FromSecretValue(secretValue *SecretValue, options *EnvValueFromSecretOptions) EnvValue {
	_init_.Initialize()

	var returns EnvValue

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.EnvValue",
		"fromSecretValue",
		[]interface{}{secretValue, options},
		&returns,
	)

	return returns
}

// Create a value from the given argument.
func EnvValue_FromValue(value *string) EnvValue {
	_init_.Initialize()

	var returns EnvValue

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.EnvValue",
		"fromValue",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Options to specify an envionment variable value from a ConfigMap key.
type EnvValueFromConfigMapOptions struct {
	// Specify whether the ConfigMap or its key must be defined.
	Optional *bool `field:"optional" json:"optional" yaml:"optional"`
}

// Options to specify an environment variable value from a field reference.
type EnvValueFromFieldRefOptions struct {
	// Version of the schema the FieldPath is written in terms of.
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// The key to select the pod label or annotation.
	Key *string `field:"optional" json:"key" yaml:"key"`
}

// Options to specify an environment variable value from the process environment.
type EnvValueFromProcessOptions struct {
	// Specify whether the key must exist in the environment.
	//
	// If this is set to true, and the key does not exist, an error will thrown.
	Required *bool `field:"optional" json:"required" yaml:"required"`
}

// Options to specify an environment variable value from a resource.
type EnvValueFromResourceOptions struct {
	// The container to select the value from.
	Container Container `field:"optional" json:"container" yaml:"container"`
	// The output format of the exposed resource.
	Divisor *string `field:"optional" json:"divisor" yaml:"divisor"`
}

// Options to specify an environment variable value from a Secret.
type EnvValueFromSecretOptions struct {
	// Specify whether the Secret or its key must be defined.
	Optional *bool `field:"optional" json:"optional" yaml:"optional"`
}

// Options for exposing a deployment via an ingress.
type ExposeDeploymentViaIngressOptions struct {
	// The name of the service to expose.
	//
	// If you'd like to expose the deployment multiple times,
	// you must explicitly set a name starting from the second expose call.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ports that the service should bind to.
	Ports *[]*ServicePort `field:"optional" json:"ports" yaml:"ports"`
	// The type of the exposed service.
	ServiceType ServiceType `field:"optional" json:"serviceType" yaml:"serviceType"`
	// The ingress to add rules to.
	Ingress Ingress `field:"optional" json:"ingress" yaml:"ingress"`
	// The type of the path.
	PathType HttpIngressPathType `field:"optional" json:"pathType" yaml:"pathType"`
}

// Options for exposing a service using an ingress.
type ExposeServiceViaIngressOptions struct {
	// The ingress to add rules to.
	Ingress Ingress `field:"optional" json:"ingress" yaml:"ingress"`
	// The type of the path.
	PathType HttpIngressPathType `field:"optional" json:"pathType" yaml:"pathType"`
}

type FsGroupChangePolicy string

const (
	// Only change permissions and ownership if permission and ownership of root directory does not match with expected permissions of the volume.
	//
	// This could help shorten the time it takes to change ownership and permission of a volume.
	FsGroupChangePolicy_ON_ROOT_MISMATCH FsGroupChangePolicy = "ON_ROOT_MISMATCH"
	// Always change permission and ownership of the volume when volume is mounted.
	FsGroupChangePolicy_ALWAYS FsGroupChangePolicy = "ALWAYS"
)

// GCEPersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod.
//
// Provisioned by an admin.
// See: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
//
type GCEPersistentDiskPersistentVolume interface {
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
	// PD resource in GCE of this volume.
	PdName() *string
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

// The jsii proxy struct for GCEPersistentDiskPersistentVolume
type jsiiProxy_GCEPersistentDiskPersistentVolume struct {
	jsiiProxy_PersistentVolume
}

func (j *jsiiProxy_GCEPersistentDiskPersistentVolume) AccessModes() *[]PersistentVolumeAccessMode {
	var returns *[]PersistentVolumeAccessMode
	_jsii_.Get(
		j,
		"accessModes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GCEPersistentDiskPersistentVolume) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GCEPersistentDiskPersistentVolume) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GCEPersistentDiskPersistentVolume) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GCEPersistentDiskPersistentVolume) Claim() IPersistentVolumeClaim {
	var returns IPersistentVolumeClaim
	_jsii_.Get(
		j,
		"claim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GCEPersistentDiskPersistentVolume) FsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fsType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GCEPersistentDiskPersistentVolume) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GCEPersistentDiskPersistentVolume) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GCEPersistentDiskPersistentVolume) Mode() PersistentVolumeMode {
	var returns PersistentVolumeMode
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GCEPersistentDiskPersistentVolume) MountOptions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"mountOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GCEPersistentDiskPersistentVolume) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GCEPersistentDiskPersistentVolume) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GCEPersistentDiskPersistentVolume) Partition() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"partition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GCEPersistentDiskPersistentVolume) PdName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pdName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GCEPersistentDiskPersistentVolume) Permissions() ResourcePermissions {
	var returns ResourcePermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GCEPersistentDiskPersistentVolume) ReadOnly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GCEPersistentDiskPersistentVolume) ReclaimPolicy() PersistentVolumeReclaimPolicy {
	var returns PersistentVolumeReclaimPolicy
	_jsii_.Get(
		j,
		"reclaimPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GCEPersistentDiskPersistentVolume) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GCEPersistentDiskPersistentVolume) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GCEPersistentDiskPersistentVolume) Storage() cdk8s.Size {
	var returns cdk8s.Size
	_jsii_.Get(
		j,
		"storage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GCEPersistentDiskPersistentVolume) StorageClassName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClassName",
		&returns,
	)
	return returns
}


func NewGCEPersistentDiskPersistentVolume(scope constructs.Construct, id *string, props *GCEPersistentDiskPersistentVolumeProps) GCEPersistentDiskPersistentVolume {
	_init_.Initialize()

	j := jsiiProxy_GCEPersistentDiskPersistentVolume{}

	_jsii_.Create(
		"cdk8s-plus-24.GCEPersistentDiskPersistentVolume",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewGCEPersistentDiskPersistentVolume_Override(g GCEPersistentDiskPersistentVolume, scope constructs.Construct, id *string, props *GCEPersistentDiskPersistentVolumeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.GCEPersistentDiskPersistentVolume",
		[]interface{}{scope, id, props},
		g,
	)
}

// Imports a pv from the cluster as a reference.
func GCEPersistentDiskPersistentVolume_FromPersistentVolumeName(scope constructs.Construct, id *string, volumeName *string) IPersistentVolume {
	_init_.Initialize()

	var returns IPersistentVolume

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.GCEPersistentDiskPersistentVolume",
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
func GCEPersistentDiskPersistentVolume_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.GCEPersistentDiskPersistentVolume",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GCEPersistentDiskPersistentVolume) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		g,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GCEPersistentDiskPersistentVolume) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GCEPersistentDiskPersistentVolume) AsVolume() Volume {
	var returns Volume

	_jsii_.Invoke(
		g,
		"asVolume",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GCEPersistentDiskPersistentVolume) Bind(claim IPersistentVolumeClaim) {
	_jsii_.InvokeVoid(
		g,
		"bind",
		[]interface{}{claim},
	)
}

func (g *jsiiProxy_GCEPersistentDiskPersistentVolume) Reserve() PersistentVolumeClaim {
	var returns PersistentVolumeClaim

	_jsii_.Invoke(
		g,
		"reserve",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GCEPersistentDiskPersistentVolume) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for `GCEPersistentDiskPersistentVolume`.
type GCEPersistentDiskPersistentVolumeProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// Contains all ways the volume can be mounted.
	// See: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes
	//
	AccessModes *[]PersistentVolumeAccessMode `field:"optional" json:"accessModes" yaml:"accessModes"`
	// Part of a bi-directional binding between PersistentVolume and PersistentVolumeClaim.
	//
	// Expected to be non-nil when bound.
	// See: https://kubernetes.io/docs/concepts/storage/persistent-volumes#binding
	//
	Claim IPersistentVolumeClaim `field:"optional" json:"claim" yaml:"claim"`
	// A list of mount options, e.g. ["ro", "soft"]. Not validated - mount will simply fail if one is invalid.
	// See: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#mount-options
	//
	MountOptions *[]*string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
	// When a user is done with their volume, they can delete the PVC objects from the API that allows reclamation of the resource.
	//
	// The reclaim policy tells the cluster what to do with
	// the volume after it has been released of its claim.
	// See: https://kubernetes.io/docs/concepts/storage/persistent-volumes#reclaiming
	//
	ReclaimPolicy PersistentVolumeReclaimPolicy `field:"optional" json:"reclaimPolicy" yaml:"reclaimPolicy"`
	// What is the storage capacity of this volume.
	// See: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources
	//
	Storage cdk8s.Size `field:"optional" json:"storage" yaml:"storage"`
	// Name of StorageClass to which this persistent volume belongs.
	StorageClassName *string `field:"optional" json:"storageClassName" yaml:"storageClassName"`
	// Defines what type of volume is required by the claim.
	VolumeMode PersistentVolumeMode `field:"optional" json:"volumeMode" yaml:"volumeMode"`
	// Unique name of the PD resource in GCE.
	//
	// Used to identify the disk in GCE.
	// See: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	//
	PdName *string `field:"required" json:"pdName" yaml:"pdName"`
	// Filesystem type of the volume that you want to mount.
	//
	// Tip: Ensure that the filesystem type is supported by the host operating system.
	// See: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	//
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// The partition in the volume that you want to mount.
	//
	// If omitted, the default is to mount by volume name.
	// Examples: For volume /dev/sda1, you specify the partition as "1".
	// Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty).
	Partition *float64 `field:"optional" json:"partition" yaml:"partition"`
	// Specify "true" to force and set the ReadOnly property in VolumeMounts to "true".
	// See: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	//
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
}

// Options of `Volume.fromGcePersistentDisk`.
type GCEPersistentDiskVolumeOptions struct {
	// Filesystem type of the volume that you want to mount.
	//
	// Tip: Ensure that the filesystem type is supported by the host operating system.
	// See: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	//
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// The volume name.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The partition in the volume that you want to mount.
	//
	// If omitted, the default is to mount by volume name.
	// Examples: For volume /dev/sda1, you specify the partition as "1".
	// Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty).
	Partition *float64 `field:"optional" json:"partition" yaml:"partition"`
	// Specify "true" to force and set the ReadOnly property in VolumeMounts to "true".
	// See: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	//
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
}

// Represents a group.
type Group interface {
	constructs.Construct
	ISubject
	ApiGroup() *string
	Kind() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
	// Return the subject configuration.
	// See: ISubect.toSubjectConfiguration()
	//
	ToSubjectConfiguration() *SubjectConfiguration
}

// The jsii proxy struct for Group
type jsiiProxy_Group struct {
	internal.Type__constructsConstruct
	jsiiProxy_ISubject
}

func (j *jsiiProxy_Group) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Reference a group by name.
func Group_FromName(scope constructs.Construct, id *string, name *string) Group {
	_init_.Initialize()

	var returns Group

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Group",
		"fromName",
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
func Group_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Group",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) ToSubjectConfiguration() *SubjectConfiguration {
	var returns *SubjectConfiguration

	_jsii_.Invoke(
		g,
		"toSubjectConfiguration",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Defines a specific action that should be taken.
type Handler interface {
}

// The jsii proxy struct for Handler
type jsiiProxy_Handler struct {
	_ byte // padding
}

// Defines a handler based on a command which is executed within the container.
func Handler_FromCommand(command *[]*string) Handler {
	_init_.Initialize()

	var returns Handler

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Handler",
		"fromCommand",
		[]interface{}{command},
		&returns,
	)

	return returns
}

// Defines a handler based on an HTTP GET request to the IP address of the container.
func Handler_FromHttpGet(path *string, options *HandlerFromHttpGetOptions) Handler {
	_init_.Initialize()

	var returns Handler

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Handler",
		"fromHttpGet",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Defines a handler based opening a connection to a TCP socket on the container.
func Handler_FromTcpSocket(options *HandlerFromTcpSocketOptions) Handler {
	_init_.Initialize()

	var returns Handler

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Handler",
		"fromTcpSocket",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Options for `Handler.fromHttpGet`.
type HandlerFromHttpGetOptions struct {
	// The TCP port to use when sending the GET request.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

// Options for `Handler.fromTcpSocket`.
type HandlerFromTcpSocketOptions struct {
	// The host name to connect to on the container.
	Host *string `field:"optional" json:"host" yaml:"host"`
	// The TCP port to connect to on the container.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

// HostAlias holds the mapping between IP and hostnames that will be injected as an entry in the pod's /etc/hosts file.
type HostAlias struct {
	// Hostnames for the chosen IP address.
	Hostnames *[]*string `field:"required" json:"hostnames" yaml:"hostnames"`
	// IP address of the host file entry.
	Ip *string `field:"required" json:"ip" yaml:"ip"`
}

// Options for a HostPathVolume-based volume.
type HostPathVolumeOptions struct {
	// The path of the directory on the host.
	Path *string `field:"required" json:"path" yaml:"path"`
	// The expected type of the path found on the host.
	Type HostPathVolumeType `field:"optional" json:"type" yaml:"type"`
}

// Host path types.
type HostPathVolumeType string

const (
	// Empty string (default) is for backward compatibility, which means that no checks will be performed before mounting the hostPath volume.
	HostPathVolumeType_DEFAULT HostPathVolumeType = "DEFAULT"
	// If nothing exists at the given path, an empty directory will be created there as needed with permission set to 0755, having the same group and ownership with Kubelet.
	HostPathVolumeType_DIRECTORY_OR_CREATE HostPathVolumeType = "DIRECTORY_OR_CREATE"
	// A directory must exist at the given path.
	HostPathVolumeType_DIRECTORY HostPathVolumeType = "DIRECTORY"
	// If nothing exists at the given path, an empty file will be created there as needed with permission set to 0644, having the same group and ownership with Kubelet.
	HostPathVolumeType_FILE_OR_CREATE HostPathVolumeType = "FILE_OR_CREATE"
	// A file must exist at the given path.
	HostPathVolumeType_FILE HostPathVolumeType = "FILE"
	// A UNIX socket must exist at the given path.
	HostPathVolumeType_SOCKET HostPathVolumeType = "SOCKET"
	// A character device must exist at the given path.
	HostPathVolumeType_CHAR_DEVICE HostPathVolumeType = "CHAR_DEVICE"
	// A block device must exist at the given path.
	HostPathVolumeType_BLOCK_DEVICE HostPathVolumeType = "BLOCK_DEVICE"
)

// Options for `Probe.fromHttpGet()`.
type HttpGetProbeOptions struct {
	// Minimum consecutive failures for the probe to be considered failed after having succeeded.
	//
	// Defaults to 3. Minimum value is 1.
	FailureThreshold *float64 `field:"optional" json:"failureThreshold" yaml:"failureThreshold"`
	// Number of seconds after the container has started before liveness probes are initiated.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	//
	InitialDelaySeconds cdk8s.Duration `field:"optional" json:"initialDelaySeconds" yaml:"initialDelaySeconds"`
	// How often (in seconds) to perform the probe.
	//
	// Default to 10 seconds. Minimum value is 1.
	PeriodSeconds cdk8s.Duration `field:"optional" json:"periodSeconds" yaml:"periodSeconds"`
	// Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1.
	//
	// Must be 1 for liveness and startup. Minimum value is 1.
	SuccessThreshold *float64 `field:"optional" json:"successThreshold" yaml:"successThreshold"`
	// Number of seconds after which the probe times out.
	//
	// Defaults to 1 second. Minimum value is 1.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	//
	TimeoutSeconds cdk8s.Duration `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
	// The TCP port to use when sending the GET request.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

// Specify how the path is matched against request paths.
// See: https://kubernetes.io/docs/concepts/services-networking/ingress/#path-types
//
type HttpIngressPathType string

const (
	// Matches the URL path exactly.
	HttpIngressPathType_PREFIX HttpIngressPathType = "PREFIX"
	// Matches based on a URL path prefix split by '/'.
	HttpIngressPathType_EXACT HttpIngressPathType = "EXACT"
	// Matching is specified by the underlying IngressClass.
	HttpIngressPathType_IMPLEMENTATION_SPECIFIC HttpIngressPathType = "IMPLEMENTATION_SPECIFIC"
)

// An API Endpoint can either be a resource descriptor (e.g /pods) or a non resource url (e.g /healthz). It must be one or the other, and not both.
type IApiEndpoint interface {
	// Return the IApiResource this object represents.
	AsApiResource() IApiResource
	// Return the non resource url this object represents.
	AsNonApiResource() *string
}

// The jsii proxy for IApiEndpoint
type jsiiProxy_IApiEndpoint struct {
	_ byte // padding
}

func (i *jsiiProxy_IApiEndpoint) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		i,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApiEndpoint) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Represents a resource or collection of resources.
type IApiResource interface {
	// The group portion of the API version (e.g. `authorization.k8s.io`).
	ApiGroup() *string
	// The unique, namespace-global, name of an object inside the Kubernetes cluster.
	//
	// If this is omitted, the ApiResource should represent all objects of the given type.
	ResourceName() *string
	// The name of a resource type as it appears in the relevant API endpoint.
	//
	// Example:
	//   - "pods" or "pods/log"
	//
	// See: https://kubernetes.io/docs/reference/access-authn-authz/rbac/#referring-to-resources
	//
	ResourceType() *string
}

// The jsii proxy for IApiResource
type jsiiProxy_IApiResource struct {
	_ byte // padding
}

func (j *jsiiProxy_IApiResource) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApiResource) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApiResource) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

// Represents a cluster-level role.
type IClusterRole interface {
	IResource
}

// The jsii proxy for IClusterRole
type jsiiProxy_IClusterRole struct {
	jsiiProxy_IResource
}

// Represents a config map.
type IConfigMap interface {
	IResource
}

// The jsii proxy for IConfigMap
type jsiiProxy_IConfigMap struct {
	jsiiProxy_IResource
}

// Represents an object that can select namespaces.
type INamespaceSelector interface {
	constructs.IConstruct
	// Return the configuration of this selector.
	ToNamespaceSelectorConfig() *NamespaceSelectorConfig
}

// The jsii proxy for INamespaceSelector
type jsiiProxy_INamespaceSelector struct {
	internal.Type__constructsIConstruct
}

func (i *jsiiProxy_INamespaceSelector) ToNamespaceSelectorConfig() *NamespaceSelectorConfig {
	var returns *NamespaceSelectorConfig

	_jsii_.Invoke(
		i,
		"toNamespaceSelectorConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Describes a peer to allow traffic to/from.
type INetworkPolicyPeer interface {
	constructs.IConstruct
	// Return the configuration of this peer.
	ToNetworkPolicyPeerConfig() *NetworkPolicyPeerConfig
	// Convert the peer into a pod selector, if possible.
	ToPodSelector() IPodSelector
}

// The jsii proxy for INetworkPolicyPeer
type jsiiProxy_INetworkPolicyPeer struct {
	internal.Type__constructsIConstruct
}

func (i *jsiiProxy_INetworkPolicyPeer) ToNetworkPolicyPeerConfig() *NetworkPolicyPeerConfig {
	var returns *NetworkPolicyPeerConfig

	_jsii_.Invoke(
		i,
		"toNetworkPolicyPeerConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_INetworkPolicyPeer) ToPodSelector() IPodSelector {
	var returns IPodSelector

	_jsii_.Invoke(
		i,
		"toPodSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Contract of a `PersistentVolumeClaim`.
type IPersistentVolume interface {
	IResource
}

// The jsii proxy for IPersistentVolume
type jsiiProxy_IPersistentVolume struct {
	jsiiProxy_IResource
}

// Contract of a `PersistentVolumeClaim`.
type IPersistentVolumeClaim interface {
	IResource
}

// The jsii proxy for IPersistentVolumeClaim
type jsiiProxy_IPersistentVolumeClaim struct {
	jsiiProxy_IResource
}

// Represents an object that can select pods.
type IPodSelector interface {
	constructs.IConstruct
	// Return the configuration of this selector.
	ToPodSelectorConfig() *PodSelectorConfig
}

// The jsii proxy for IPodSelector
type jsiiProxy_IPodSelector struct {
	internal.Type__constructsIConstruct
}

func (i *jsiiProxy_IPodSelector) ToPodSelectorConfig() *PodSelectorConfig {
	var returns *PodSelectorConfig

	_jsii_.Invoke(
		i,
		"toPodSelectorConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Represents a resource.
type IResource interface {
	constructs.IConstruct
	// The group portion of the API version (e.g. "authorization.k8s.io").
	ApiGroup() *string
	// The object's API version (e.g. "authorization.k8s.io/v1").
	ApiVersion() *string
	// The object kind (e.g. "Deployment").
	Kind() *string
	// The Kubernetes name of this resource.
	Name() *string
}

// The jsii proxy for IResource
type jsiiProxy_IResource struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IResource) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResource) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResource) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

// A reference to any Role or ClusterRole.
type IRole interface {
	IResource
}

// The jsii proxy for IRole
type jsiiProxy_IRole struct {
	jsiiProxy_IResource
}

type ISecret interface {
	IResource
}

// The jsii proxy for ISecret
type jsiiProxy_ISecret struct {
	jsiiProxy_IResource
}

type IServiceAccount interface {
	IResource
}

// The jsii proxy for IServiceAccount
type jsiiProxy_IServiceAccount struct {
	jsiiProxy_IResource
}

// Represents a piece of storage in the cluster.
type IStorage interface {
	constructs.IConstruct
	// Convert the piece of storage into a concrete volume.
	AsVolume() Volume
}

// The jsii proxy for IStorage
type jsiiProxy_IStorage struct {
	internal.Type__constructsIConstruct
}

func (i *jsiiProxy_IStorage) AsVolume() Volume {
	var returns Volume

	_jsii_.Invoke(
		i,
		"asVolume",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Represents an object that can be used as a role binding subject.
type ISubject interface {
	constructs.IConstruct
	// Return the subject configuration.
	ToSubjectConfiguration() *SubjectConfiguration
}

// The jsii proxy for ISubject
type jsiiProxy_ISubject struct {
	internal.Type__constructsIConstruct
}

func (i *jsiiProxy_ISubject) ToSubjectConfiguration() *SubjectConfiguration {
	var returns *SubjectConfiguration

	_jsii_.Invoke(
		i,
		"toSubjectConfiguration",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ImagePullPolicy string

const (
	// Every time the kubelet launches a container, the kubelet queries the container image registry to resolve the name to an image digest.
	//
	// If the kubelet has a container image with that exact
	// digest cached locally, the kubelet uses its cached image; otherwise, the kubelet downloads
	// (pulls) the image with the resolved digest, and uses that image to launch the container.
	//
	// Default is Always if ImagePullPolicy is omitted and either the image tag is :latest or
	// the image tag is omitted.
	ImagePullPolicy_ALWAYS ImagePullPolicy = "ALWAYS"
	// The image is pulled only if it is not already present locally.
	//
	// Default is IfNotPresent if ImagePullPolicy is omitted and the image tag is present but
	// not :latest.
	ImagePullPolicy_IF_NOT_PRESENT ImagePullPolicy = "IF_NOT_PRESENT"
	// The image is assumed to exist locally.
	//
	// No attempt is made to pull the image.
	ImagePullPolicy_NEVER ImagePullPolicy = "NEVER"
)

// Ingress is a collection of rules that allow inbound connections to reach the endpoints defined by a backend.
//
// An Ingress can be configured to give services
// externally-reachable urls, load balance traffic, terminate SSL, offer name
// based virtual hosting etc.
type Ingress interface {
	Resource
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
	// Defines the default backend for this ingress.
	//
	// A default backend capable of
	// servicing requests that don't match any rule.
	AddDefaultBackend(backend IngressBackend)
	// Specify a default backend for a specific host name.
	//
	// This backend will be used as a catch-all for requests
	// targeted to this host name (the `Host` header matches this value).
	AddHostDefaultBackend(host *string, backend IngressBackend)
	// Adds an ingress rule applied to requests to a specific host and a specific HTTP path (the `Host` header matches this value).
	AddHostRule(host *string, path *string, backend IngressBackend, pathType HttpIngressPathType)
	// Adds an ingress rule applied to requests sent to a specific HTTP path.
	AddRule(path *string, backend IngressBackend, pathType HttpIngressPathType)
	// Adds rules to this ingress.
	AddRules(rules ...*IngressRule)
	AddTls(tls *[]*IngressTls)
	// Return the IApiResource this object represents.
	AsApiResource() IApiResource
	// Return the non resource url this object represents.
	AsNonApiResource() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Ingress
type jsiiProxy_Ingress struct {
	jsiiProxy_Resource
}

func (j *jsiiProxy_Ingress) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ingress) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ingress) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ingress) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ingress) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ingress) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ingress) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ingress) Permissions() ResourcePermissions {
	var returns ResourcePermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ingress) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ingress) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}


func NewIngress(scope constructs.Construct, id *string, props *IngressProps) Ingress {
	_init_.Initialize()

	j := jsiiProxy_Ingress{}

	_jsii_.Create(
		"cdk8s-plus-24.Ingress",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewIngress_Override(i Ingress, scope constructs.Construct, id *string, props *IngressProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.Ingress",
		[]interface{}{scope, id, props},
		i,
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
func Ingress_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Ingress",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Ingress) AddDefaultBackend(backend IngressBackend) {
	_jsii_.InvokeVoid(
		i,
		"addDefaultBackend",
		[]interface{}{backend},
	)
}

func (i *jsiiProxy_Ingress) AddHostDefaultBackend(host *string, backend IngressBackend) {
	_jsii_.InvokeVoid(
		i,
		"addHostDefaultBackend",
		[]interface{}{host, backend},
	)
}

func (i *jsiiProxy_Ingress) AddHostRule(host *string, path *string, backend IngressBackend, pathType HttpIngressPathType) {
	_jsii_.InvokeVoid(
		i,
		"addHostRule",
		[]interface{}{host, path, backend, pathType},
	)
}

func (i *jsiiProxy_Ingress) AddRule(path *string, backend IngressBackend, pathType HttpIngressPathType) {
	_jsii_.InvokeVoid(
		i,
		"addRule",
		[]interface{}{path, backend, pathType},
	)
}

func (i *jsiiProxy_Ingress) AddRules(rules ...*IngressRule) {
	args := []interface{}{}
	for _, a := range rules {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		i,
		"addRules",
		args,
	)
}

func (i *jsiiProxy_Ingress) AddTls(tls *[]*IngressTls) {
	_jsii_.InvokeVoid(
		i,
		"addTls",
		[]interface{}{tls},
	)
}

func (i *jsiiProxy_Ingress) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		i,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Ingress) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Ingress) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The backend for an ingress path.
type IngressBackend interface {
}

// The jsii proxy struct for IngressBackend
type jsiiProxy_IngressBackend struct {
	_ byte // padding
}

// A Resource backend is an ObjectRef to another Kubernetes resource within the same namespace as the Ingress object.
//
// A common usage for a Resource backend is to ingress data to an object
// storage backend with static assets.
func IngressBackend_FromResource(resource IResource) IngressBackend {
	_init_.Initialize()

	var returns IngressBackend

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.IngressBackend",
		"fromResource",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

// A Kubernetes `Service` to use as the backend for this path.
func IngressBackend_FromService(serv Service, options *ServiceIngressBackendOptions) IngressBackend {
	_init_.Initialize()

	var returns IngressBackend

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.IngressBackend",
		"fromService",
		[]interface{}{serv, options},
		&returns,
	)

	return returns
}

// Properties for `Ingress`.
type IngressProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// The default backend services requests that do not match any rule.
	//
	// Using this option or the `addDefaultBackend()` method is equivalent to
	// adding a rule with both `path` and `host` undefined.
	DefaultBackend IngressBackend `field:"optional" json:"defaultBackend" yaml:"defaultBackend"`
	// Routing rules for this ingress.
	//
	// Each rule must define an `IngressBackend` that will receive the requests
	// that match this rule. If both `host` and `path` are not specifiec, this
	// backend will be used as the default backend of the ingress.
	//
	// You can also add rules later using `addRule()`, `addHostRule()`,
	// `addDefaultBackend()` and `addHostDefaultBackend()`.
	Rules *[]*IngressRule `field:"optional" json:"rules" yaml:"rules"`
	// TLS settings for this ingress.
	//
	// Using this option tells the ingress controller to expose a TLS endpoint.
	// Currently the Ingress only supports a single TLS port, 443. If multiple
	// members of this list specify different hosts, they will be multiplexed on
	// the same port according to the hostname specified through the SNI TLS
	// extension, if the ingress controller fulfilling the ingress supports SNI.
	Tls *[]*IngressTls `field:"optional" json:"tls" yaml:"tls"`
}

// Represents the rules mapping the paths under a specified host to the related backend services.
//
// Incoming requests are first evaluated for a host match,
// then routed to the backend associated with the matching path.
type IngressRule struct {
	// Backend defines the referenced service endpoint to which the traffic will be forwarded to.
	Backend IngressBackend `field:"required" json:"backend" yaml:"backend"`
	// Host is the fully qualified domain name of a network host, as defined by RFC 3986.
	//
	// Note the following deviations from the "host" part of the URI as
	// defined in the RFC: 1. IPs are not allowed. Currently an IngressRuleValue
	// can only apply to the IP in the Spec of the parent Ingress. 2. The `:`
	// delimiter is not respected because ports are not allowed. Currently the
	// port of an Ingress is implicitly :80 for http and :443 for https. Both
	// these may change in the future. Incoming requests are matched against the
	// host before the IngressRuleValue.
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Path is an extended POSIX regex as defined by IEEE Std 1003.1, (i.e this follows the egrep/unix syntax, not the perl syntax) matched against the path of an incoming request. Currently it can contain characters disallowed from the conventional "path" part of a URL as defined by RFC 3986. Paths must begin with a '/'.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Specify how the path is matched against request paths.
	//
	// By default, path
	// types will be matched by prefix.
	// See: https://kubernetes.io/docs/concepts/services-networking/ingress/#path-types
	//
	PathType HttpIngressPathType `field:"optional" json:"pathType" yaml:"pathType"`
}

// Represents the TLS configuration mapping that is passed to the ingress controller for SSL termination.
type IngressTls struct {
	// Hosts are a list of hosts included in the TLS certificate.
	//
	// The values in
	// this list must match the name/s used in the TLS Secret.
	Hosts *[]*string `field:"optional" json:"hosts" yaml:"hosts"`
	// Secret is the secret that contains the certificate and key used to terminate SSL traffic on 443.
	//
	// If the SNI host in a listener conflicts with
	// the "Host" header field used by an IngressRule, the SNI host is used for
	// termination and value of the Host header is used for routing.
	Secret ISecret `field:"optional" json:"secret" yaml:"secret"`
}

// A Job creates one or more Pods and ensures that a specified number of them successfully terminate.
//
// As pods successfully complete,
// the Job tracks the successful completions. When a specified number of successful completions is reached, the task (ie, Job) is complete.
// Deleting a Job will clean up the Pods it created. A simple case is to create one Job object in order to reliably run one Pod to completion.
// The Job object will start a new Pod if the first Pod fails or is deleted (for example due to a node hardware failure or a node reboot).
// You can also use a Job to run multiple Pods in parallel.
type Job interface {
	Workload
	// Duration before job is terminated.
	//
	// If undefined, there is no deadline.
	ActiveDeadline() cdk8s.Duration
	// The group portion of the API version (e.g. "authorization.k8s.io").
	ApiGroup() *string
	// The underlying cdk8s API object.
	// See: base.Resource.apiObject
	//
	ApiObject() cdk8s.ApiObject
	// The object's API version (e.g. "authorization.k8s.io/v1").
	ApiVersion() *string
	AutomountServiceAccountToken() *bool
	// Number of retries before marking failed.
	BackoffLimit() *float64
	Connections() PodConnections
	Containers() *[]Container
	Dns() PodDns
	DockerRegistryAuth() DockerConfigSecret
	HostAliases() *[]*HostAlias
	InitContainers() *[]Container
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
	// TTL before the job is deleted after it is finished.
	TtlAfterFinished() cdk8s.Duration
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

// The jsii proxy struct for Job
type jsiiProxy_Job struct {
	jsiiProxy_Workload
}

func (j *jsiiProxy_Job) ActiveDeadline() cdk8s.Duration {
	var returns cdk8s.Duration
	_jsii_.Get(
		j,
		"activeDeadline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) AutomountServiceAccountToken() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"automountServiceAccountToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) BackoffLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backoffLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Connections() PodConnections {
	var returns PodConnections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Containers() *[]Container {
	var returns *[]Container
	_jsii_.Get(
		j,
		"containers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Dns() PodDns {
	var returns PodDns
	_jsii_.Get(
		j,
		"dns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) DockerRegistryAuth() DockerConfigSecret {
	var returns DockerConfigSecret
	_jsii_.Get(
		j,
		"dockerRegistryAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) HostAliases() *[]*HostAlias {
	var returns *[]*HostAlias
	_jsii_.Get(
		j,
		"hostAliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) InitContainers() *[]Container {
	var returns *[]Container
	_jsii_.Get(
		j,
		"initContainers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) MatchExpressions() *[]*LabelSelectorRequirement {
	var returns *[]*LabelSelectorRequirement
	_jsii_.Get(
		j,
		"matchExpressions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) MatchLabels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"matchLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Permissions() ResourcePermissions {
	var returns ResourcePermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) PodMetadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"podMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) RestartPolicy() RestartPolicy {
	var returns RestartPolicy
	_jsii_.Get(
		j,
		"restartPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Scheduling() WorkloadScheduling {
	var returns WorkloadScheduling
	_jsii_.Get(
		j,
		"scheduling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) SecurityContext() PodSecurityContext {
	var returns PodSecurityContext
	_jsii_.Get(
		j,
		"securityContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) ServiceAccount() IServiceAccount {
	var returns IServiceAccount
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) TtlAfterFinished() cdk8s.Duration {
	var returns cdk8s.Duration
	_jsii_.Get(
		j,
		"ttlAfterFinished",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Volumes() *[]Volume {
	var returns *[]Volume
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}


func NewJob(scope constructs.Construct, id *string, props *JobProps) Job {
	_init_.Initialize()

	j := jsiiProxy_Job{}

	_jsii_.Create(
		"cdk8s-plus-24.Job",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewJob_Override(j Job, scope constructs.Construct, id *string, props *JobProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.Job",
		[]interface{}{scope, id, props},
		j,
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
func Job_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Job",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) AddContainer(cont *ContainerProps) Container {
	var returns Container

	_jsii_.Invoke(
		j,
		"addContainer",
		[]interface{}{cont},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) AddHostAlias(hostAlias *HostAlias) {
	_jsii_.InvokeVoid(
		j,
		"addHostAlias",
		[]interface{}{hostAlias},
	)
}

func (j *jsiiProxy_Job) AddInitContainer(cont *ContainerProps) Container {
	var returns Container

	_jsii_.Invoke(
		j,
		"addInitContainer",
		[]interface{}{cont},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) AddVolume(vol Volume) {
	_jsii_.InvokeVoid(
		j,
		"addVolume",
		[]interface{}{vol},
	)
}

func (j *jsiiProxy_Job) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		j,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) Select(selectors ...LabelSelector) {
	args := []interface{}{}
	for _, a := range selectors {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		j,
		"select",
		args,
	)
}

func (j *jsiiProxy_Job) ToNetworkPolicyPeerConfig() *NetworkPolicyPeerConfig {
	var returns *NetworkPolicyPeerConfig

	_jsii_.Invoke(
		j,
		"toNetworkPolicyPeerConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) ToPodSelector() IPodSelector {
	var returns IPodSelector

	_jsii_.Invoke(
		j,
		"toPodSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) ToPodSelectorConfig() *PodSelectorConfig {
	var returns *PodSelectorConfig

	_jsii_.Invoke(
		j,
		"toPodSelectorConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) ToSubjectConfiguration() *SubjectConfiguration {
	var returns *SubjectConfiguration

	_jsii_.Invoke(
		j,
		"toSubjectConfiguration",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for `Job`.
type JobProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// Indicates whether a service account token should be automatically mounted.
	// See: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/#use-the-default-service-account-to-access-the-api-server
	//
	AutomountServiceAccountToken *bool `field:"optional" json:"automountServiceAccountToken" yaml:"automountServiceAccountToken"`
	// List of containers belonging to the pod.
	//
	// Containers cannot currently be
	// added or removed. There must be at least one container in a Pod.
	//
	// You can add additionnal containers using `podSpec.addContainer()`
	Containers *[]*ContainerProps `field:"optional" json:"containers" yaml:"containers"`
	// DNS settings for the pod.
	// See: https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/
	//
	Dns *PodDnsProps `field:"optional" json:"dns" yaml:"dns"`
	// A secret containing docker credentials for authenticating to a registry.
	DockerRegistryAuth DockerConfigSecret `field:"optional" json:"dockerRegistryAuth" yaml:"dockerRegistryAuth"`
	// HostAlias holds the mapping between IP and hostnames that will be injected as an entry in the pod's hosts file.
	HostAliases *[]*HostAlias `field:"optional" json:"hostAliases" yaml:"hostAliases"`
	// List of initialization containers belonging to the pod.
	//
	// Init containers are executed in order prior to containers being started.
	// If any init container fails, the pod is considered to have failed and is handled according to its restartPolicy.
	// The name for an init container or normal container must be unique among all containers.
	// Init containers may not have Lifecycle actions, Readiness probes, Liveness probes, or Startup probes.
	// The resourceRequirements of an init container are taken into account during scheduling by finding the highest request/limit
	// for each resource type, and then using the max of of that value or the sum of the normal containers.
	// Limits are applied to init containers in a similar fashion.
	//
	// Init containers cannot currently be added ,removed or updated.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/
	//
	InitContainers *[]*ContainerProps `field:"optional" json:"initContainers" yaml:"initContainers"`
	// Restart policy for all containers within the pod.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy
	//
	RestartPolicy RestartPolicy `field:"optional" json:"restartPolicy" yaml:"restartPolicy"`
	// SecurityContext holds pod-level security attributes and common container settings.
	SecurityContext *PodSecurityContextProps `field:"optional" json:"securityContext" yaml:"securityContext"`
	// A service account provides an identity for processes that run in a Pod.
	//
	// When you (a human) access the cluster (for example, using kubectl), you are
	// authenticated by the apiserver as a particular User Account (currently this
	// is usually admin, unless your cluster administrator has customized your
	// cluster). Processes in containers inside pods can also contact the
	// apiserver. When they do, they are authenticated as a particular Service
	// Account (for example, default).
	// See: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
	//
	ServiceAccount IServiceAccount `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// List of volumes that can be mounted by containers belonging to the pod.
	//
	// You can also add volumes later using `podSpec.addVolume()`
	// See: https://kubernetes.io/docs/concepts/storage/volumes
	//
	Volumes *[]Volume `field:"optional" json:"volumes" yaml:"volumes"`
	// The pod metadata of this workload.
	PodMetadata *cdk8s.ApiObjectMetadata `field:"optional" json:"podMetadata" yaml:"podMetadata"`
	// Automatically allocates a pod label selector for this workload and add it to the pod metadata.
	//
	// This ensures this workload manages pods created by
	// its pod template.
	Select *bool `field:"optional" json:"select" yaml:"select"`
	// Specifies the duration the job may be active before the system tries to terminate it.
	ActiveDeadline cdk8s.Duration `field:"optional" json:"activeDeadline" yaml:"activeDeadline"`
	// Specifies the number of retries before marking this job failed.
	BackoffLimit *float64 `field:"optional" json:"backoffLimit" yaml:"backoffLimit"`
	// Limits the lifetime of a Job that has finished execution (either Complete or Failed).
	//
	// If this field is set, after the Job finishes, it is eligible to
	// be automatically deleted. When the Job is being deleted, its lifecycle
	// guarantees (e.g. finalizers) will be honored. If this field is set to zero,
	// the Job becomes eligible to be deleted immediately after it finishes. This
	// field is alpha-level and is only honored by servers that enable the
	// `TTLAfterFinished` feature.
	TtlAfterFinished cdk8s.Duration `field:"optional" json:"ttlAfterFinished" yaml:"ttlAfterFinished"`
}

// Represents a query that can be performed against resources with labels.
type LabelExpression interface {
	Key() *string
	Operator() *string
	Values() *[]*string
}

// The jsii proxy struct for LabelExpression
type jsiiProxy_LabelExpression struct {
	_ byte // padding
}

func (j *jsiiProxy_LabelExpression) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabelExpression) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabelExpression) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}


// Requires label `key` to not exist.
func LabelExpression_DoesNotExist(key *string) LabelExpression {
	_init_.Initialize()

	var returns LabelExpression

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.LabelExpression",
		"doesNotExist",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Requires label `key` to exist.
func LabelExpression_Exists(key *string) LabelExpression {
	_init_.Initialize()

	var returns LabelExpression

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.LabelExpression",
		"exists",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Requires value of label `key` to be one of `values`.
func LabelExpression_In(key *string, values *[]*string) LabelExpression {
	_init_.Initialize()

	var returns LabelExpression

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.LabelExpression",
		"in",
		[]interface{}{key, values},
		&returns,
	)

	return returns
}

// Requires value of label `key` to be none of `values`.
func LabelExpression_NotIn(key *string, values *[]*string) LabelExpression {
	_init_.Initialize()

	var returns LabelExpression

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.LabelExpression",
		"notIn",
		[]interface{}{key, values},
		&returns,
	)

	return returns
}

// Match a resource by labels.
type LabelSelector interface {
	IsEmpty() *bool
}

// The jsii proxy struct for LabelSelector
type jsiiProxy_LabelSelector struct {
	_ byte // padding
}

func LabelSelector_Of(options *LabelSelectorOptions) LabelSelector {
	_init_.Initialize()

	var returns LabelSelector

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.LabelSelector",
		"of",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LabelSelector) IsEmpty() *bool {
	var returns *bool

	_jsii_.Invoke(
		l,
		"isEmpty",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options for `LabelSelector.of`.
type LabelSelectorOptions struct {
	// Expression based label matchers.
	Expressions *[]LabelExpression `field:"optional" json:"expressions" yaml:"expressions"`
	// Strict label matchers.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
}

// A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.
type LabelSelectorRequirement struct {
	// The label key that the selector applies to.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Represents a key's relationship to a set of values.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// An array of string values.
	//
	// If the operator is In or NotIn, the values array
	// must be non-empty. If the operator is Exists or DoesNotExist,
	// the values array must be empty. This array is replaced during a strategic merge patch.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

// A node that is matched by label selectors.
type LabeledNode interface {
	LabelSelector() *[]NodeLabelQuery
}

// The jsii proxy struct for LabeledNode
type jsiiProxy_LabeledNode struct {
	_ byte // padding
}

func (j *jsiiProxy_LabeledNode) LabelSelector() *[]NodeLabelQuery {
	var returns *[]NodeLabelQuery
	_jsii_.Get(
		j,
		"labelSelector",
		&returns,
	)
	return returns
}


func NewLabeledNode(labelSelector *[]NodeLabelQuery) LabeledNode {
	_init_.Initialize()

	j := jsiiProxy_LabeledNode{}

	_jsii_.Create(
		"cdk8s-plus-24.LabeledNode",
		[]interface{}{labelSelector},
		&j,
	)

	return &j
}

func NewLabeledNode_Override(l LabeledNode, labelSelector *[]NodeLabelQuery) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.LabeledNode",
		[]interface{}{labelSelector},
		l,
	)
}

// Memory request and limit.
type MemoryResources struct {
	Limit cdk8s.Size `field:"optional" json:"limit" yaml:"limit"`
	Request cdk8s.Size `field:"optional" json:"request" yaml:"request"`
}

// Options for mounts.
type MountOptions struct {
	// Determines how mounts are propagated from the host to container and the other way around.
	//
	// When not set, MountPropagationNone is used.
	//
	// Mount propagation allows for sharing volumes mounted by a Container to
	// other Containers in the same Pod, or even to other Pods on the same node.
	Propagation MountPropagation `field:"optional" json:"propagation" yaml:"propagation"`
	// Mounted read-only if true, read-write otherwise (false or unspecified).
	//
	// Defaults to false.
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Path within the volume from which the container's volume should be mounted.).
	SubPath *string `field:"optional" json:"subPath" yaml:"subPath"`
	// Expanded path within the volume from which the container's volume should be mounted.
	//
	// Behaves similarly to SubPath but environment variable references
	// $(VAR_NAME) are expanded using the container's environment. Defaults to ""
	// (volume's root).
	//
	// `subPathExpr` and `subPath` are mutually exclusive.
	SubPathExpr *string `field:"optional" json:"subPathExpr" yaml:"subPathExpr"`
}

type MountPropagation string

const (
	// This volume mount will not receive any subsequent mounts that are mounted to this volume or any of its subdirectories by the host.
	//
	// In similar
	// fashion, no mounts created by the Container will be visible on the host.
	//
	// This is the default mode.
	//
	// This mode is equal to `private` mount propagation as described in the Linux
	// kernel documentation.
	MountPropagation_NONE MountPropagation = "NONE"
	// This volume mount will receive all subsequent mounts that are mounted to this volume or any of its subdirectories.
	//
	// In other words, if the host mounts anything inside the volume mount, the
	// Container will see it mounted there.
	//
	// Similarly, if any Pod with Bidirectional mount propagation to the same
	// volume mounts anything there, the Container with HostToContainer mount
	// propagation will see it.
	//
	// This mode is equal to `rslave` mount propagation as described in the Linux
	// kernel documentation.
	MountPropagation_HOST_TO_CONTAINER MountPropagation = "HOST_TO_CONTAINER"
	// This volume mount behaves the same the HostToContainer mount.
	//
	// In addition,
	// all volume mounts created by the Container will be propagated back to the
	// host and to all Containers of all Pods that use the same volume
	//
	// A typical use case for this mode is a Pod with a FlexVolume or CSI driver
	// or a Pod that needs to mount something on the host using a hostPath volume.
	//
	// This mode is equal to `rshared` mount propagation as described in the Linux
	// kernel documentation
	//
	// Caution: Bidirectional mount propagation can be dangerous. It can damage
	// the host operating system and therefore it is allowed only in privileged
	// Containers. Familiarity with Linux kernel behavior is strongly recommended.
	// In addition, any volume mounts created by Containers in Pods must be
	// destroyed (unmounted) by the Containers on termination.
	MountPropagation_BIDIRECTIONAL MountPropagation = "BIDIRECTIONAL"
)

// A node that is matched by its name.
type NamedNode interface {
	Name() *string
}

// The jsii proxy struct for NamedNode
type jsiiProxy_NamedNode struct {
	_ byte // padding
}

func (j *jsiiProxy_NamedNode) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


func NewNamedNode(name *string) NamedNode {
	_init_.Initialize()

	j := jsiiProxy_NamedNode{}

	_jsii_.Create(
		"cdk8s-plus-24.NamedNode",
		[]interface{}{name},
		&j,
	)

	return &j
}

func NewNamedNode_Override(n NamedNode, name *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.NamedNode",
		[]interface{}{name},
		n,
	)
}

// In Kubernetes, namespaces provides a mechanism for isolating groups of resources within a single cluster.
//
// Names of resources need to be unique within a namespace, but not across namespaces.
// Namespace-based scoping is applicable only for namespaced objects (e.g. Deployments, Services, etc) and
// not for cluster-wide objects (e.g. StorageClass, Nodes, PersistentVolumes, etc).
type Namespace interface {
	Resource
	INamespaceSelector
	INetworkPolicyPeer
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
	// Return the IApiResource this object represents.
	AsApiResource() IApiResource
	// Return the non resource url this object represents.
	AsNonApiResource() *string
	// Return the configuration of this selector.
	// See: INamespaceSelector.toNamespaceSelectorConfig()
	//
	ToNamespaceSelectorConfig() *NamespaceSelectorConfig
	// Return the configuration of this peer.
	// See: INetworkPolicyPeer.toNetworkPolicyPeerConfig()
	//
	ToNetworkPolicyPeerConfig() *NetworkPolicyPeerConfig
	// Convert the peer into a pod selector, if possible.
	// See: INetworkPolicyPeer.toPodSelector()
	//
	ToPodSelector() IPodSelector
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Namespace
type jsiiProxy_Namespace struct {
	jsiiProxy_Resource
	jsiiProxy_INamespaceSelector
	jsiiProxy_INetworkPolicyPeer
}

func (j *jsiiProxy_Namespace) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Namespace) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Namespace) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Namespace) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Namespace) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Namespace) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Namespace) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Namespace) Permissions() ResourcePermissions {
	var returns ResourcePermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Namespace) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Namespace) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}


func NewNamespace(scope constructs.Construct, id *string, props *NamespaceProps) Namespace {
	_init_.Initialize()

	j := jsiiProxy_Namespace{}

	_jsii_.Create(
		"cdk8s-plus-24.Namespace",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewNamespace_Override(n Namespace, scope constructs.Construct, id *string, props *NamespaceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.Namespace",
		[]interface{}{scope, id, props},
		n,
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
func Namespace_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Namespace",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Namespace_NAME_LABEL() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdk8s-plus-24.Namespace",
		"NAME_LABEL",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_Namespace) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		n,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Namespace) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Namespace) ToNamespaceSelectorConfig() *NamespaceSelectorConfig {
	var returns *NamespaceSelectorConfig

	_jsii_.Invoke(
		n,
		"toNamespaceSelectorConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Namespace) ToNetworkPolicyPeerConfig() *NetworkPolicyPeerConfig {
	var returns *NetworkPolicyPeerConfig

	_jsii_.Invoke(
		n,
		"toNetworkPolicyPeerConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Namespace) ToPodSelector() IPodSelector {
	var returns IPodSelector

	_jsii_.Invoke(
		n,
		"toPodSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Namespace) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for `Namespace`.
type NamespaceProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
}

// Configuration for selecting namespaces.
type NamespaceSelectorConfig struct {
	// A selector to select namespaces by labels.
	LabelSelector LabelSelector `field:"optional" json:"labelSelector" yaml:"labelSelector"`
	// A list of names to select namespaces by names.
	Names *[]*string `field:"optional" json:"names" yaml:"names"`
}

// Represents a group of namespaces.
type Namespaces interface {
	constructs.Construct
	INamespaceSelector
	INetworkPolicyPeer
	// The tree node.
	Node() constructs.Node
	// Return the configuration of this selector.
	// See: INamespaceSelector.toNamespaceSelectorConfig()
	//
	ToNamespaceSelectorConfig() *NamespaceSelectorConfig
	// Return the configuration of this peer.
	// See: INetworkPolicyPeer.toNetworkPolicyPeerConfig()
	//
	ToNetworkPolicyPeerConfig() *NetworkPolicyPeerConfig
	// Convert the peer into a pod selector, if possible.
	// See: INetworkPolicyPeer.toPodSelector()
	//
	ToPodSelector() IPodSelector
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Namespaces
type jsiiProxy_Namespaces struct {
	internal.Type__constructsConstruct
	jsiiProxy_INamespaceSelector
	jsiiProxy_INetworkPolicyPeer
}

func (j *jsiiProxy_Namespaces) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewNamespaces(scope constructs.Construct, id *string, expressions *[]LabelExpression, names *[]*string, labels *map[string]*string) Namespaces {
	_init_.Initialize()

	j := jsiiProxy_Namespaces{}

	_jsii_.Create(
		"cdk8s-plus-24.Namespaces",
		[]interface{}{scope, id, expressions, names, labels},
		&j,
	)

	return &j
}

func NewNamespaces_Override(n Namespaces, scope constructs.Construct, id *string, expressions *[]LabelExpression, names *[]*string, labels *map[string]*string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.Namespaces",
		[]interface{}{scope, id, expressions, names, labels},
		n,
	)
}

// Select all namespaces.
func Namespaces_All(scope constructs.Construct, id *string) Namespaces {
	_init_.Initialize()

	var returns Namespaces

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Namespaces",
		"all",
		[]interface{}{scope, id},
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
func Namespaces_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Namespaces",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Select specific namespaces.
func Namespaces_Select(scope constructs.Construct, id *string, options *NamespacesSelectOptions) Namespaces {
	_init_.Initialize()

	var returns Namespaces

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Namespaces",
		"select",
		[]interface{}{scope, id, options},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Namespaces) ToNamespaceSelectorConfig() *NamespaceSelectorConfig {
	var returns *NamespaceSelectorConfig

	_jsii_.Invoke(
		n,
		"toNamespaceSelectorConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Namespaces) ToNetworkPolicyPeerConfig() *NetworkPolicyPeerConfig {
	var returns *NetworkPolicyPeerConfig

	_jsii_.Invoke(
		n,
		"toNetworkPolicyPeerConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Namespaces) ToPodSelector() IPodSelector {
	var returns IPodSelector

	_jsii_.Invoke(
		n,
		"toPodSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Namespaces) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options for `Namespaces.select`.
type NamespacesSelectOptions struct {
	// Namespaces must satisfy these selectors.
	//
	// The selectors query labels, just like the `labels` property, but they
	// provide a more advanced matching mechanism.
	Expressions *[]LabelExpression `field:"optional" json:"expressions" yaml:"expressions"`
	// Labels the namespaces must have.
	//
	// This is equivalent to using an 'Is' selector.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Namespaces names must be one of these.
	Names *[]*string `field:"optional" json:"names" yaml:"names"`
}

// Control traffic flow at the IP address or port level (OSI layer 3 or 4), network policies are an application-centric construct which allow you to specify how a pod is allowed to communicate with various network peers.
//
// - Outgoing traffic is allowed if there are no network policies selecting
//    the pod (and cluster policy otherwise allows the traffic),
//    OR if the traffic matches at least one egress rule across all of the
//    network policies that select the pod.
//
// - Incoming traffic is allowed to a pod if there are no network policies
//    selecting the pod (and cluster policy otherwise allows the traffic),
//    OR if the traffic source is the pod's local node,
//    OR if the traffic matches at least one ingress rule across all of
//    the network policies that select the pod.
//
// Network policies do not conflict; they are additive.
// If any policy or policies apply to a given pod for a given
// direction, the connections allowed in that direction from
// that pod is the union of what the applicable policies allow.
// Thus, order of evaluation does not affect the policy result.
//
// For a connection from a source pod to a destination pod to be allowed,
// both the egress policy on the source pod and the ingress policy on the
// destination pod need to allow the connection.
// If either side does not allow the connection, it will not happen.
// See: https://kubernetes.io/docs/concepts/services-networking/network-policies/#networkpolicy-resource
//
type NetworkPolicy interface {
	Resource
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
	// Allow outgoing traffic to the peer.
	//
	// If ports are not passed, traffic will be allowed on all ports.
	AddEgressRule(peer INetworkPolicyPeer, ports *[]NetworkPolicyPort)
	// Allow incoming traffic from the peer.
	//
	// If ports are not passed, traffic will be allowed on all ports.
	AddIngressRule(peer INetworkPolicyPeer, ports *[]NetworkPolicyPort)
	// Return the IApiResource this object represents.
	AsApiResource() IApiResource
	// Return the non resource url this object represents.
	AsNonApiResource() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for NetworkPolicy
type jsiiProxy_NetworkPolicy struct {
	jsiiProxy_Resource
}

func (j *jsiiProxy_NetworkPolicy) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicy) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicy) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicy) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicy) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicy) Permissions() ResourcePermissions {
	var returns ResourcePermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicy) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicy) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}


func NewNetworkPolicy(scope constructs.Construct, id *string, props *NetworkPolicyProps) NetworkPolicy {
	_init_.Initialize()

	j := jsiiProxy_NetworkPolicy{}

	_jsii_.Create(
		"cdk8s-plus-24.NetworkPolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewNetworkPolicy_Override(n NetworkPolicy, scope constructs.Construct, id *string, props *NetworkPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.NetworkPolicy",
		[]interface{}{scope, id, props},
		n,
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
func NetworkPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.NetworkPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPolicy) AddEgressRule(peer INetworkPolicyPeer, ports *[]NetworkPolicyPort) {
	_jsii_.InvokeVoid(
		n,
		"addEgressRule",
		[]interface{}{peer, ports},
	)
}

func (n *jsiiProxy_NetworkPolicy) AddIngressRule(peer INetworkPolicyPeer, ports *[]NetworkPolicyPort) {
	_jsii_.InvokeVoid(
		n,
		"addIngressRule",
		[]interface{}{peer, ports},
	)
}

func (n *jsiiProxy_NetworkPolicy) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		n,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPolicy) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options for `NetworkPolicy.addEgressRule`.
type NetworkPolicyAddEgressRuleOptions struct {
	// Ports the rule should allow outgoing traffic to.
	Ports *[]NetworkPolicyPort `field:"optional" json:"ports" yaml:"ports"`
}

// Describes a particular CIDR (Ex.
//
// "192.168.1.1/24","2001:db9::/64") that is
// allowed to the pods matched by a network policy selector.
// The except entry describes CIDRs that should not be included within this rule.
type NetworkPolicyIpBlock interface {
	constructs.Construct
	INetworkPolicyPeer
	// A string representing the IP Block Valid examples are "192.168.1.1/24" or "2001:db9::/64".
	Cidr() *string
	// A slice of CIDRs that should not be included within an IP Block Valid examples are "192.168.1.1/24" or "2001:db9::/64". Except values will be rejected if they are outside the CIDR range.
	Except() *[]*string
	// The tree node.
	Node() constructs.Node
	// Return the configuration of this peer.
	// See: INetworkPolicyPeer.toNetworkPolicyPeerConfig()
	//
	ToNetworkPolicyPeerConfig() *NetworkPolicyPeerConfig
	// Convert the peer into a pod selector, if possible.
	// See: INetworkPolicyPeer.toPodSelector()
	//
	ToPodSelector() IPodSelector
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for NetworkPolicyIpBlock
type jsiiProxy_NetworkPolicyIpBlock struct {
	internal.Type__constructsConstruct
	jsiiProxy_INetworkPolicyPeer
}

func (j *jsiiProxy_NetworkPolicyIpBlock) Cidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicyIpBlock) Except() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"except",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicyIpBlock) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Any IPv4 address.
func NetworkPolicyIpBlock_AnyIpv4(scope constructs.Construct, id *string) NetworkPolicyIpBlock {
	_init_.Initialize()

	var returns NetworkPolicyIpBlock

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.NetworkPolicyIpBlock",
		"anyIpv4",
		[]interface{}{scope, id},
		&returns,
	)

	return returns
}

// Any IPv6 address.
func NetworkPolicyIpBlock_AnyIpv6(scope constructs.Construct, id *string) NetworkPolicyIpBlock {
	_init_.Initialize()

	var returns NetworkPolicyIpBlock

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.NetworkPolicyIpBlock",
		"anyIpv6",
		[]interface{}{scope, id},
		&returns,
	)

	return returns
}

// Create an IPv4 peer from a CIDR.
func NetworkPolicyIpBlock_Ipv4(scope constructs.Construct, id *string, cidrIp *string, except *[]*string) NetworkPolicyIpBlock {
	_init_.Initialize()

	var returns NetworkPolicyIpBlock

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.NetworkPolicyIpBlock",
		"ipv4",
		[]interface{}{scope, id, cidrIp, except},
		&returns,
	)

	return returns
}

// Create an IPv6 peer from a CIDR.
func NetworkPolicyIpBlock_Ipv6(scope constructs.Construct, id *string, cidrIp *string, except *[]*string) NetworkPolicyIpBlock {
	_init_.Initialize()

	var returns NetworkPolicyIpBlock

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.NetworkPolicyIpBlock",
		"ipv6",
		[]interface{}{scope, id, cidrIp, except},
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
func NetworkPolicyIpBlock_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.NetworkPolicyIpBlock",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPolicyIpBlock) ToNetworkPolicyPeerConfig() *NetworkPolicyPeerConfig {
	var returns *NetworkPolicyPeerConfig

	_jsii_.Invoke(
		n,
		"toNetworkPolicyPeerConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPolicyIpBlock) ToPodSelector() IPodSelector {
	var returns IPodSelector

	_jsii_.Invoke(
		n,
		"toPodSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPolicyIpBlock) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Configuration for network peers.
//
// A peer can either by an ip block, or a selection of pods, not both.
type NetworkPolicyPeerConfig struct {
	// The ip block this peer represents.
	IpBlock NetworkPolicyIpBlock `field:"optional" json:"ipBlock" yaml:"ipBlock"`
	// The pod selector this peer represents.
	PodSelector *PodSelectorConfig `field:"optional" json:"podSelector" yaml:"podSelector"`
}

// Describes a port to allow traffic on.
type NetworkPolicyPort interface {
}

// The jsii proxy struct for NetworkPolicyPort
type jsiiProxy_NetworkPolicyPort struct {
	_ byte // padding
}

// Any TCP traffic.
func NetworkPolicyPort_AllTcp() NetworkPolicyPort {
	_init_.Initialize()

	var returns NetworkPolicyPort

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.NetworkPolicyPort",
		"allTcp",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Any UDP traffic.
func NetworkPolicyPort_AllUdp() NetworkPolicyPort {
	_init_.Initialize()

	var returns NetworkPolicyPort

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.NetworkPolicyPort",
		"allUdp",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Custom port configuration.
func NetworkPolicyPort_Of(props *NetworkPolicyPortProps) NetworkPolicyPort {
	_init_.Initialize()

	var returns NetworkPolicyPort

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.NetworkPolicyPort",
		"of",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Distinct TCP ports.
func NetworkPolicyPort_Tcp(port *float64) NetworkPolicyPort {
	_init_.Initialize()

	var returns NetworkPolicyPort

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.NetworkPolicyPort",
		"tcp",
		[]interface{}{port},
		&returns,
	)

	return returns
}

// A TCP port range.
func NetworkPolicyPort_TcpRange(startPort *float64, endPort *float64) NetworkPolicyPort {
	_init_.Initialize()

	var returns NetworkPolicyPort

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.NetworkPolicyPort",
		"tcpRange",
		[]interface{}{startPort, endPort},
		&returns,
	)

	return returns
}

// Distinct UDP ports.
func NetworkPolicyPort_Udp(port *float64) NetworkPolicyPort {
	_init_.Initialize()

	var returns NetworkPolicyPort

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.NetworkPolicyPort",
		"udp",
		[]interface{}{port},
		&returns,
	)

	return returns
}

// A UDP port range.
func NetworkPolicyPort_UdpRange(startPort *float64, endPort *float64) NetworkPolicyPort {
	_init_.Initialize()

	var returns NetworkPolicyPort

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.NetworkPolicyPort",
		"udpRange",
		[]interface{}{startPort, endPort},
		&returns,
	)

	return returns
}

// Properties for `NetworkPolicyPort`.
type NetworkPolicyPortProps struct {
	// End port (relative to `port`).
	//
	// Only applies if `port` is defined.
	// Use this to specify a port range, rather that a specific one.
	EndPort *float64 `field:"optional" json:"endPort" yaml:"endPort"`
	// Specific port number.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Protocol.
	Protocol NetworkProtocol `field:"optional" json:"protocol" yaml:"protocol"`
}

// Properties for `NetworkPolicy`.
type NetworkPolicyProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// Egress traffic configuration.
	Egress *NetworkPolicyTraffic `field:"optional" json:"egress" yaml:"egress"`
	// Ingress traffic configuration.
	Ingress *NetworkPolicyTraffic `field:"optional" json:"ingress" yaml:"ingress"`
	// Which pods does this policy object applies to.
	//
	// This can either be a single pod / workload, or a grouping of pods selected
	// via the `Pods.select` function. Rules is applied to any pods selected by this property.
	// Multiple network policies can select the same set of pods.
	// In this case, the rules for each are combined additively.
	//
	// Note that.
	Selector IPodSelector `field:"optional" json:"selector" yaml:"selector"`
}

// Describes a rule allowing traffic from / to pods matched by a network policy selector.
type NetworkPolicyRule struct {
	// Peer this rule interacts with.
	Peer INetworkPolicyPeer `field:"required" json:"peer" yaml:"peer"`
	// The ports of the rule.
	Ports *[]NetworkPolicyPort `field:"optional" json:"ports" yaml:"ports"`
}

// Describes how the network policy should configure egress / ingress traffic.
type NetworkPolicyTraffic struct {
	// Specifies the default behavior of the policy when no rules are defined.
	Default NetworkPolicyTrafficDefault `field:"optional" json:"default" yaml:"default"`
	// List of rules to be applied to the selected pods.
	//
	// If empty, the behavior of the policy is dictated by the `default` property.
	Rules *[]*NetworkPolicyRule `field:"optional" json:"rules" yaml:"rules"`
}

// Default behaviors of network traffic in policies.
type NetworkPolicyTrafficDefault string

const (
	// The policy denies all traffic.
	//
	// Since rules are additive, additional rules or policies can allow
	// specific traffic.
	NetworkPolicyTrafficDefault_DENY NetworkPolicyTrafficDefault = "DENY"
	// The policy allows all traffic (either ingress or egress).
	//
	// Since rules are additive, no additional rule or policies can
	// subsequently deny the traffic.
	NetworkPolicyTrafficDefault_ALLOW NetworkPolicyTrafficDefault = "ALLOW"
)

// Network protocols.
type NetworkProtocol string

const (
	// TCP.
	NetworkProtocol_TCP NetworkProtocol = "TCP"
	// UDP.
	NetworkProtocol_UDP NetworkProtocol = "UDP"
	// SCTP.
	NetworkProtocol_SCTP NetworkProtocol = "SCTP"
)

// Represents a node in the cluster.
type Node interface {
}

// The jsii proxy struct for Node
type jsiiProxy_Node struct {
	_ byte // padding
}

func NewNode() Node {
	_init_.Initialize()

	j := jsiiProxy_Node{}

	_jsii_.Create(
		"cdk8s-plus-24.Node",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewNode_Override(n Node) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.Node",
		nil, // no parameters
		n,
	)
}

// Match a node by its labels.
func Node_Labeled(labelSelector ...NodeLabelQuery) LabeledNode {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range labelSelector {
		args = append(args, a)
	}

	var returns LabeledNode

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Node",
		"labeled",
		args,
		&returns,
	)

	return returns
}

// Match a node by its name.
func Node_Named(nodeName *string) NamedNode {
	_init_.Initialize()

	var returns NamedNode

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Node",
		"named",
		[]interface{}{nodeName},
		&returns,
	)

	return returns
}

// Match a node by its taints.
func Node_Tainted(taintSelector ...NodeTaintQuery) TaintedNode {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range taintSelector {
		args = append(args, a)
	}

	var returns TaintedNode

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Node",
		"tainted",
		args,
		&returns,
	)

	return returns
}

// Represents a query that can be performed against nodes with labels.
type NodeLabelQuery interface {
}

// The jsii proxy struct for NodeLabelQuery
type jsiiProxy_NodeLabelQuery struct {
	_ byte // padding
}

// Requires label `key` to not exist.
func NodeLabelQuery_DoesNotExist(key *string) NodeLabelQuery {
	_init_.Initialize()

	var returns NodeLabelQuery

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.NodeLabelQuery",
		"doesNotExist",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Requires label `key` to exist.
func NodeLabelQuery_Exists(key *string) NodeLabelQuery {
	_init_.Initialize()

	var returns NodeLabelQuery

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.NodeLabelQuery",
		"exists",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Requires value of label `key` to greater than all elements in `values`.
func NodeLabelQuery_Gt(key *string, values *[]*string) NodeLabelQuery {
	_init_.Initialize()

	var returns NodeLabelQuery

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.NodeLabelQuery",
		"gt",
		[]interface{}{key, values},
		&returns,
	)

	return returns
}

// Requires value of label `key` to be one of `values`.
func NodeLabelQuery_In(key *string, values *[]*string) NodeLabelQuery {
	_init_.Initialize()

	var returns NodeLabelQuery

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.NodeLabelQuery",
		"in",
		[]interface{}{key, values},
		&returns,
	)

	return returns
}

// Requires value of label `key` to equal `value`.
func NodeLabelQuery_Is(key *string, value *string) NodeLabelQuery {
	_init_.Initialize()

	var returns NodeLabelQuery

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.NodeLabelQuery",
		"is",
		[]interface{}{key, value},
		&returns,
	)

	return returns
}

// Requires value of label `key` to less than all elements in `values`.
func NodeLabelQuery_Lt(key *string, values *[]*string) NodeLabelQuery {
	_init_.Initialize()

	var returns NodeLabelQuery

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.NodeLabelQuery",
		"lt",
		[]interface{}{key, values},
		&returns,
	)

	return returns
}

// Requires value of label `key` to be none of `values`.
func NodeLabelQuery_NotIn(key *string, values *[]*string) NodeLabelQuery {
	_init_.Initialize()

	var returns NodeLabelQuery

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.NodeLabelQuery",
		"notIn",
		[]interface{}{key, values},
		&returns,
	)

	return returns
}

// Taint queries that can be perfomed against nodes.
type NodeTaintQuery interface {
}

// The jsii proxy struct for NodeTaintQuery
type jsiiProxy_NodeTaintQuery struct {
	_ byte // padding
}

// Matches any taint.
func NodeTaintQuery_Any() NodeTaintQuery {
	_init_.Initialize()

	var returns NodeTaintQuery

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.NodeTaintQuery",
		"any",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Matches a tain with any value of a specific key.
func NodeTaintQuery_Exists(key *string, options *NodeTaintQueryOptions) NodeTaintQuery {
	_init_.Initialize()

	var returns NodeTaintQuery

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.NodeTaintQuery",
		"exists",
		[]interface{}{key, options},
		&returns,
	)

	return returns
}

// Matches a taint with a specific key and value.
func NodeTaintQuery_Is(key *string, value *string, options *NodeTaintQueryOptions) NodeTaintQuery {
	_init_.Initialize()

	var returns NodeTaintQuery

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.NodeTaintQuery",
		"is",
		[]interface{}{key, value, options},
		&returns,
	)

	return returns
}

// Options for `NodeTaintQuery`.
type NodeTaintQueryOptions struct {
	// The taint effect to match.
	Effect TaintEffect `field:"optional" json:"effect" yaml:"effect"`
	// How much time should a pod that tolerates the `NO_EXECUTE` effect be bound to the node.
	//
	// Only applies for the `NO_EXECUTE` effect.
	EvictAfter cdk8s.Duration `field:"optional" json:"evictAfter" yaml:"evictAfter"`
}

// Factory for creating non api resources.
type NonApiResource interface {
	IApiEndpoint
	// Return the IApiResource this object represents.
	AsApiResource() IApiResource
	// Return the non resource url this object represents.
	AsNonApiResource() *string
}

// The jsii proxy struct for NonApiResource
type jsiiProxy_NonApiResource struct {
	jsiiProxy_IApiEndpoint
}

func NonApiResource_Of(url *string) NonApiResource {
	_init_.Initialize()

	var returns NonApiResource

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.NonApiResource",
		"of",
		[]interface{}{url},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NonApiResource) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		n,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NonApiResource) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Maps a string key to a path within a volume.
type PathMapping struct {
	// The relative path of the file to map the key to.
	//
	// May not be an absolute
	// path. May not contain the path element '..'. May not start with the string
	// '..'.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Optional: mode bits to use on this file, must be a value between 0 and 0777.
	//
	// If not specified, the volume defaultMode will be used. This might be
	// in conflict with other options that affect the file mode, like fsGroup, and
	// the result can be other mode bits set.
	Mode *float64 `field:"optional" json:"mode" yaml:"mode"`
}

// Union like class repsenting either a ration in percents or an absolute number.
type PercentOrAbsolute interface {
	Value() interface{}
	IsZero() *bool
}

// The jsii proxy struct for PercentOrAbsolute
type jsiiProxy_PercentOrAbsolute struct {
	_ byte // padding
}

func (j *jsiiProxy_PercentOrAbsolute) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Absolute number.
func PercentOrAbsolute_Absolute(num *float64) PercentOrAbsolute {
	_init_.Initialize()

	var returns PercentOrAbsolute

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.PercentOrAbsolute",
		"absolute",
		[]interface{}{num},
		&returns,
	)

	return returns
}

// Percent ratio.
func PercentOrAbsolute_Percent(percent *float64) PercentOrAbsolute {
	_init_.Initialize()

	var returns PercentOrAbsolute

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.PercentOrAbsolute",
		"percent",
		[]interface{}{percent},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PercentOrAbsolute) IsZero() *bool {
	var returns *bool

	_jsii_.Invoke(
		p,
		"isZero",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A PersistentVolume (PV) is a piece of storage in the cluster that has been provisioned by an administrator or dynamically provisioned using Storage Classes.
//
// It is a resource in the cluster just like a node is a cluster resource.
// PVs are volume plugins like Volumes, but have a lifecycle independent of any
// individual Pod that uses the PV. This API object captures the details of the
// implementation of the storage, be that NFS, iSCSI, or a
// cloud-provider-specific storage system.
type PersistentVolume interface {
	Resource
	IPersistentVolume
	IStorage
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

// The jsii proxy struct for PersistentVolume
type jsiiProxy_PersistentVolume struct {
	jsiiProxy_Resource
	jsiiProxy_IPersistentVolume
	jsiiProxy_IStorage
}

func (j *jsiiProxy_PersistentVolume) AccessModes() *[]PersistentVolumeAccessMode {
	var returns *[]PersistentVolumeAccessMode
	_jsii_.Get(
		j,
		"accessModes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolume) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolume) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolume) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolume) Claim() IPersistentVolumeClaim {
	var returns IPersistentVolumeClaim
	_jsii_.Get(
		j,
		"claim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolume) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolume) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolume) Mode() PersistentVolumeMode {
	var returns PersistentVolumeMode
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolume) MountOptions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"mountOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolume) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolume) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolume) Permissions() ResourcePermissions {
	var returns ResourcePermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolume) ReclaimPolicy() PersistentVolumeReclaimPolicy {
	var returns PersistentVolumeReclaimPolicy
	_jsii_.Get(
		j,
		"reclaimPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolume) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolume) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolume) Storage() cdk8s.Size {
	var returns cdk8s.Size
	_jsii_.Get(
		j,
		"storage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolume) StorageClassName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClassName",
		&returns,
	)
	return returns
}


func NewPersistentVolume(scope constructs.Construct, id *string, props *PersistentVolumeProps) PersistentVolume {
	_init_.Initialize()

	j := jsiiProxy_PersistentVolume{}

	_jsii_.Create(
		"cdk8s-plus-24.PersistentVolume",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewPersistentVolume_Override(p PersistentVolume, scope constructs.Construct, id *string, props *PersistentVolumeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.PersistentVolume",
		[]interface{}{scope, id, props},
		p,
	)
}

// Imports a pv from the cluster as a reference.
func PersistentVolume_FromPersistentVolumeName(scope constructs.Construct, id *string, volumeName *string) IPersistentVolume {
	_init_.Initialize()

	var returns IPersistentVolume

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.PersistentVolume",
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
func PersistentVolume_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.PersistentVolume",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolume) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		p,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolume) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolume) AsVolume() Volume {
	var returns Volume

	_jsii_.Invoke(
		p,
		"asVolume",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolume) Bind(claim IPersistentVolumeClaim) {
	_jsii_.InvokeVoid(
		p,
		"bind",
		[]interface{}{claim},
	)
}

func (p *jsiiProxy_PersistentVolume) Reserve() PersistentVolumeClaim {
	var returns PersistentVolumeClaim

	_jsii_.Invoke(
		p,
		"reserve",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolume) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Access Modes.
type PersistentVolumeAccessMode string

const (
	// The volume can be mounted as read-write by a single node.
	//
	// ReadWriteOnce access mode still can allow multiple pods to access
	// the volume when the pods are running on the same node.
	PersistentVolumeAccessMode_READ_WRITE_ONCE PersistentVolumeAccessMode = "READ_WRITE_ONCE"
	// The volume can be mounted as read-only by many nodes.
	PersistentVolumeAccessMode_READ_ONLY_MANY PersistentVolumeAccessMode = "READ_ONLY_MANY"
	// The volume can be mounted as read-write by many nodes.
	PersistentVolumeAccessMode_READ_WRITE_MANY PersistentVolumeAccessMode = "READ_WRITE_MANY"
	// The volume can be mounted as read-write by a single Pod.
	//
	// Use ReadWriteOncePod access mode if you want to ensure that
	// only one pod across whole cluster can read that PVC or write to it.
	// This is only supported for CSI volumes and Kubernetes version 1.22+.
	PersistentVolumeAccessMode_READ_WRITE_ONCE_POD PersistentVolumeAccessMode = "READ_WRITE_ONCE_POD"
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

	j := jsiiProxy_PersistentVolumeClaim{}

	_jsii_.Create(
		"cdk8s-plus-24.PersistentVolumeClaim",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewPersistentVolumeClaim_Override(p PersistentVolumeClaim, scope constructs.Construct, id *string, props *PersistentVolumeClaimProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.PersistentVolumeClaim",
		[]interface{}{scope, id, props},
		p,
	)
}

// Imports a pvc from the cluster as a reference.
func PersistentVolumeClaim_FromClaimName(scope constructs.Construct, id *string, claimName *string) IPersistentVolumeClaim {
	_init_.Initialize()

	var returns IPersistentVolumeClaim

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.PersistentVolumeClaim",
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

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.PersistentVolumeClaim",
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

// Properties for `PersistentVolumeClaim`.
type PersistentVolumeClaimProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// Contains the access modes the volume should support.
	// See: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1
	//
	AccessModes *[]PersistentVolumeAccessMode `field:"optional" json:"accessModes" yaml:"accessModes"`
	// Minimum storage size the volume should have.
	// See: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources
	//
	Storage cdk8s.Size `field:"optional" json:"storage" yaml:"storage"`
	// Name of the StorageClass required by the claim. When this property is not set, the behavior is as follows:.
	//
	// - If the admission plugin is turned on, the storage class marked as default will be used.
	// - If the admission plugin is turned off, the pvc can only be bound to volumes without a storage class.
	// See: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1
	//
	StorageClassName *string `field:"optional" json:"storageClassName" yaml:"storageClassName"`
	// The PersistentVolume backing this claim.
	//
	// The control plane still checks that storage class, access modes,
	// and requested storage size on the volume are valid.
	//
	// Note that in order to guarantee a proper binding, the volume should
	// also define a `claimRef` referring to this claim. Otherwise, the volume may be
	// claimed be other pvc's before it gets a chance to bind to this one.
	//
	// If the volume is managed (i.e not imported), you can use `pv.claim()` to easily
	// create a bi-directional bounded claim.
	// See: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#binding.
	//
	Volume IPersistentVolume `field:"optional" json:"volume" yaml:"volume"`
	// Defines what type of volume is required by the claim.
	VolumeMode PersistentVolumeMode `field:"optional" json:"volumeMode" yaml:"volumeMode"`
}

// Options for a PersistentVolumeClaim-based volume.
type PersistentVolumeClaimVolumeOptions struct {
	// The volume name.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Will force the ReadOnly setting in VolumeMounts.
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
}

// Volume Modes.
type PersistentVolumeMode string

const (
	// Volume is ounted into Pods into a directory.
	//
	// If the volume is backed by a block device and the device is empty,
	// Kubernetes creates a filesystem on the device before mounting it
	// for the first time.
	PersistentVolumeMode_FILE_SYSTEM PersistentVolumeMode = "FILE_SYSTEM"
	// Use a volume as a raw block device.
	//
	// Such volume is presented into a Pod as a block device,
	// without any filesystem on it. This mode is useful to provide a Pod the fastest possible way
	// to access a volume, without any filesystem layer between the Pod
	// and the volume. On the other hand, the application running in
	// the Pod must know how to handle a raw block device.
	PersistentVolumeMode_BLOCK PersistentVolumeMode = "BLOCK"
)

// Properties for `PersistentVolume`.
type PersistentVolumeProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// Contains all ways the volume can be mounted.
	// See: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes
	//
	AccessModes *[]PersistentVolumeAccessMode `field:"optional" json:"accessModes" yaml:"accessModes"`
	// Part of a bi-directional binding between PersistentVolume and PersistentVolumeClaim.
	//
	// Expected to be non-nil when bound.
	// See: https://kubernetes.io/docs/concepts/storage/persistent-volumes#binding
	//
	Claim IPersistentVolumeClaim `field:"optional" json:"claim" yaml:"claim"`
	// A list of mount options, e.g. ["ro", "soft"]. Not validated - mount will simply fail if one is invalid.
	// See: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#mount-options
	//
	MountOptions *[]*string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
	// When a user is done with their volume, they can delete the PVC objects from the API that allows reclamation of the resource.
	//
	// The reclaim policy tells the cluster what to do with
	// the volume after it has been released of its claim.
	// See: https://kubernetes.io/docs/concepts/storage/persistent-volumes#reclaiming
	//
	ReclaimPolicy PersistentVolumeReclaimPolicy `field:"optional" json:"reclaimPolicy" yaml:"reclaimPolicy"`
	// What is the storage capacity of this volume.
	// See: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources
	//
	Storage cdk8s.Size `field:"optional" json:"storage" yaml:"storage"`
	// Name of StorageClass to which this persistent volume belongs.
	StorageClassName *string `field:"optional" json:"storageClassName" yaml:"storageClassName"`
	// Defines what type of volume is required by the claim.
	VolumeMode PersistentVolumeMode `field:"optional" json:"volumeMode" yaml:"volumeMode"`
}

// Reclaim Policies.
type PersistentVolumeReclaimPolicy string

const (
	// The Retain reclaim policy allows for manual reclamation of the resource.
	//
	// When the PersistentVolumeClaim is deleted, the PersistentVolume still exists and the
	// volume is considered "released". But it is not yet available for another claim
	// because the previous claimant's data remains on the volume.
	// An administrator can manually reclaim the volume with the following steps:
	//
	// 1. Delete the PersistentVolume. The associated storage asset in external
	//     infrastructure (such as an AWS EBS, GCE PD, Azure Disk, or Cinder volume) still exists after the PV is deleted.
	// 2. Manually clean up the data on the associated storage asset accordingly.
	// 3. Manually delete the associated storage asset.
	//
	// If you want to reuse the same storage asset, create a new PersistentVolume
	// with the same storage asset definition.
	PersistentVolumeReclaimPolicy_RETAIN PersistentVolumeReclaimPolicy = "RETAIN"
	// For volume plugins that support the Delete reclaim policy, deletion removes both the PersistentVolume object from Kubernetes, as well as the associated storage asset in the external infrastructure, such as an AWS EBS, GCE PD, Azure Disk, or Cinder volume.
	//
	// Volumes that were dynamically provisioned inherit the reclaim policy of their StorageClass, which defaults to Delete.
	// The administrator should configure the StorageClass according to users' expectations; otherwise,
	// the PV must be edited or patched after it is created.
	PersistentVolumeReclaimPolicy_DELETE PersistentVolumeReclaimPolicy = "DELETE"
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

	j := jsiiProxy_Pod{}

	_jsii_.Create(
		"cdk8s-plus-24.Pod",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewPod_Override(p Pod, scope constructs.Construct, id *string, props *PodProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.Pod",
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

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Pod",
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
		"cdk8s-plus-24.Pod",
		"ADDRESS_LABEL",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_Pod) AddContainer(cont *ContainerProps) Container {
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
	_jsii_.InvokeVoid(
		p,
		"addHostAlias",
		[]interface{}{hostAlias},
	)
}

func (p *jsiiProxy_Pod) AddInitContainer(cont *ContainerProps) Container {
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

// Controls network isolation rules for inter-pod communication.
type PodConnections interface {
	Instance() AbstractPod
	// Allow network traffic from the peer to this pod.
	//
	// By default, this will create an ingress network policy for this pod, and an egress
	// network policy for the peer. This is required if both sides are already isolated.
	// Use `options.isolation` to control this behavior.
	//
	// Example:
	//   // create only an egress policy that selects the 'web' pod to allow outgoing traffic
	//   // to the 'redis' pod. this requires the 'redis' pod to not be isolated for ingress.
	//   redis.connections.allowFrom(web, { isolation: Isolation.PEER })
	//
	//   // create only an ingress policy that selects the 'redis' peer to allow incoming traffic
	//   // from the 'web' pod. this requires the 'web' pod to not be isolated for egress.
	//   redis.connections.allowFrom(web, { isolation: Isolation.POD })
	//
	AllowFrom(peer INetworkPolicyPeer, options *PodConnectionsAllowFromOptions)
	// Allow network traffic from this pod to the peer.
	//
	// By default, this will create an egress network policy for this pod, and an ingress
	// network policy for the peer. This is required if both sides are already isolated.
	// Use `options.isolation` to control this behavior.
	//
	// Example:
	//   // create only an egress policy that selects the 'web' pod to allow outgoing traffic
	//   // to the 'redis' pod. this requires the 'redis' pod to not be isolated for ingress.
	//   web.connections.allowTo(redis, { isolation: Isolation.POD })
	//
	//   // create only an ingress policy that selects the 'redis' peer to allow incoming traffic
	//   // from the 'web' pod. this requires the 'web' pod to not be isolated for egress.
	//   web.connections.allowTo(redis, { isolation: Isolation.PEER })
	//
	AllowTo(peer INetworkPolicyPeer, options *PodConnectionsAllowToOptions)
}

// The jsii proxy struct for PodConnections
type jsiiProxy_PodConnections struct {
	_ byte // padding
}

func (j *jsiiProxy_PodConnections) Instance() AbstractPod {
	var returns AbstractPod
	_jsii_.Get(
		j,
		"instance",
		&returns,
	)
	return returns
}


func NewPodConnections(instance AbstractPod) PodConnections {
	_init_.Initialize()

	j := jsiiProxy_PodConnections{}

	_jsii_.Create(
		"cdk8s-plus-24.PodConnections",
		[]interface{}{instance},
		&j,
	)

	return &j
}

func NewPodConnections_Override(p PodConnections, instance AbstractPod) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.PodConnections",
		[]interface{}{instance},
		p,
	)
}

func (p *jsiiProxy_PodConnections) AllowFrom(peer INetworkPolicyPeer, options *PodConnectionsAllowFromOptions) {
	_jsii_.InvokeVoid(
		p,
		"allowFrom",
		[]interface{}{peer, options},
	)
}

func (p *jsiiProxy_PodConnections) AllowTo(peer INetworkPolicyPeer, options *PodConnectionsAllowToOptions) {
	_jsii_.InvokeVoid(
		p,
		"allowTo",
		[]interface{}{peer, options},
	)
}

// Options for `PodConnections.allowFrom`.
type PodConnectionsAllowFromOptions struct {
	// Which isolation should be applied to establish the connection.
	Isolation PodConnectionsIsolation `field:"optional" json:"isolation" yaml:"isolation"`
	// Ports to allow incoming traffic to.
	Ports *[]NetworkPolicyPort `field:"optional" json:"ports" yaml:"ports"`
}

// Options for `PodConnections.allowTo`.
type PodConnectionsAllowToOptions struct {
	// Which isolation should be applied to establish the connection.
	Isolation PodConnectionsIsolation `field:"optional" json:"isolation" yaml:"isolation"`
	// Ports to allow outgoing traffic to.
	Ports *[]NetworkPolicyPort `field:"optional" json:"ports" yaml:"ports"`
}

// Isolation determines which policies are created when allowing connections from a a pod / workload to peers.
type PodConnectionsIsolation string

const (
	// Only creates network policies that select the pod.
	PodConnectionsIsolation_POD PodConnectionsIsolation = "POD"
	// Only creates network policies that select the peer.
	PodConnectionsIsolation_PEER PodConnectionsIsolation = "PEER"
)

// Holds dns settings of the pod.
type PodDns interface {
	// The configured hostname of the pod.
	//
	// Undefined means its set to a system-defined value.
	Hostname() *string
	// Whether or not the pods hostname is set to its FQDN.
	HostnameAsFQDN() *bool
	// Nameservers defined for this pod.
	Nameservers() *[]*string
	// Custom dns options defined for this pod.
	Options() *[]*DnsOption
	// The DNS policy of this pod.
	Policy() DnsPolicy
	// Search domains defined for this pod.
	Searches() *[]*string
	// The configured subdomain of the pod.
	Subdomain() *string
	// Add a nameserver.
	AddNameserver(nameservers ...*string)
	// Add a custom option.
	AddOption(options ...*DnsOption)
	// Add a search domain.
	AddSearch(searches ...*string)
}

// The jsii proxy struct for PodDns
type jsiiProxy_PodDns struct {
	_ byte // padding
}

func (j *jsiiProxy_PodDns) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodDns) HostnameAsFQDN() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"hostnameAsFQDN",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodDns) Nameservers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nameservers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodDns) Options() *[]*DnsOption {
	var returns *[]*DnsOption
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodDns) Policy() DnsPolicy {
	var returns DnsPolicy
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodDns) Searches() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"searches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodDns) Subdomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subdomain",
		&returns,
	)
	return returns
}


func NewPodDns(props *PodDnsProps) PodDns {
	_init_.Initialize()

	j := jsiiProxy_PodDns{}

	_jsii_.Create(
		"cdk8s-plus-24.PodDns",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewPodDns_Override(p PodDns, props *PodDnsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.PodDns",
		[]interface{}{props},
		p,
	)
}

func (p *jsiiProxy_PodDns) AddNameserver(nameservers ...*string) {
	args := []interface{}{}
	for _, a := range nameservers {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		p,
		"addNameserver",
		args,
	)
}

func (p *jsiiProxy_PodDns) AddOption(options ...*DnsOption) {
	args := []interface{}{}
	for _, a := range options {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		p,
		"addOption",
		args,
	)
}

func (p *jsiiProxy_PodDns) AddSearch(searches ...*string) {
	args := []interface{}{}
	for _, a := range searches {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		p,
		"addSearch",
		args,
	)
}

// Properties for `PodDns`.
type PodDnsProps struct {
	// Specifies the hostname of the Pod.
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// If true the pod's hostname will be configured as the pod's FQDN, rather than the leaf name (the default).
	//
	// In Linux containers, this means setting the FQDN in the hostname field of the kernel (the nodename field of struct utsname).
	// In Windows containers, this means setting the registry value of hostname for the registry
	// key HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Services\Tcpip\Parameters to FQDN.
	// If a pod does not have FQDN, this has no effect.
	HostnameAsFQDN *bool `field:"optional" json:"hostnameAsFQDN" yaml:"hostnameAsFQDN"`
	// A list of IP addresses that will be used as DNS servers for the Pod.
	//
	// There can be at most 3 IP addresses specified.
	// When the policy is set to "NONE", the list must contain at least one IP address,
	// otherwise this property is optional.
	// The servers listed will be combined to the base nameservers generated from
	// the specified DNS policy with duplicate addresses removed.
	Nameservers *[]*string `field:"optional" json:"nameservers" yaml:"nameservers"`
	// List of objects where each object may have a name property (required) and a value property (optional).
	//
	// The contents in this property
	// will be merged to the options generated from the specified DNS policy.
	// Duplicate entries are removed.
	Options *[]*DnsOption `field:"optional" json:"options" yaml:"options"`
	// Set DNS policy for the pod.
	//
	// If policy is set to `None`, other configuration must be supplied.
	Policy DnsPolicy `field:"optional" json:"policy" yaml:"policy"`
	// A list of DNS search domains for hostname lookup in the Pod.
	//
	// When specified, the provided list will be merged into the base
	// search domain names generated from the chosen DNS policy.
	// Duplicate domain names are removed.
	//
	// Kubernetes allows for at most 6 search domains.
	Searches *[]*string `field:"optional" json:"searches" yaml:"searches"`
	// If specified, the fully qualified Pod hostname will be "<hostname>.<subdomain>.<pod namespace>.svc.<cluster domain>".
	Subdomain *string `field:"optional" json:"subdomain" yaml:"subdomain"`
}

// Controls how pods are created during initial scale up, when replacing pods on nodes, or when scaling down.
//
// The default policy is `OrderedReady`, where pods are created in increasing order
// (pod-0, then pod-1, etc) and the controller will wait until each pod is ready before
// continuing. When scaling down, the pods are removed in the opposite order.
//
// The alternative policy is `Parallel` which will create pods in parallel to match the
// desired scale without waiting, and on scale down will delete all pods at once.
type PodManagementPolicy string

const (
	PodManagementPolicy_ORDERED_READY PodManagementPolicy = "ORDERED_READY"
	PodManagementPolicy_PARALLEL PodManagementPolicy = "PARALLEL"
)

// Properties for `Pod`.
type PodProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// Indicates whether a service account token should be automatically mounted.
	// See: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/#use-the-default-service-account-to-access-the-api-server
	//
	AutomountServiceAccountToken *bool `field:"optional" json:"automountServiceAccountToken" yaml:"automountServiceAccountToken"`
	// List of containers belonging to the pod.
	//
	// Containers cannot currently be
	// added or removed. There must be at least one container in a Pod.
	//
	// You can add additionnal containers using `podSpec.addContainer()`
	Containers *[]*ContainerProps `field:"optional" json:"containers" yaml:"containers"`
	// DNS settings for the pod.
	// See: https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/
	//
	Dns *PodDnsProps `field:"optional" json:"dns" yaml:"dns"`
	// A secret containing docker credentials for authenticating to a registry.
	DockerRegistryAuth DockerConfigSecret `field:"optional" json:"dockerRegistryAuth" yaml:"dockerRegistryAuth"`
	// HostAlias holds the mapping between IP and hostnames that will be injected as an entry in the pod's hosts file.
	HostAliases *[]*HostAlias `field:"optional" json:"hostAliases" yaml:"hostAliases"`
	// List of initialization containers belonging to the pod.
	//
	// Init containers are executed in order prior to containers being started.
	// If any init container fails, the pod is considered to have failed and is handled according to its restartPolicy.
	// The name for an init container or normal container must be unique among all containers.
	// Init containers may not have Lifecycle actions, Readiness probes, Liveness probes, or Startup probes.
	// The resourceRequirements of an init container are taken into account during scheduling by finding the highest request/limit
	// for each resource type, and then using the max of of that value or the sum of the normal containers.
	// Limits are applied to init containers in a similar fashion.
	//
	// Init containers cannot currently be added ,removed or updated.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/
	//
	InitContainers *[]*ContainerProps `field:"optional" json:"initContainers" yaml:"initContainers"`
	// Restart policy for all containers within the pod.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy
	//
	RestartPolicy RestartPolicy `field:"optional" json:"restartPolicy" yaml:"restartPolicy"`
	// SecurityContext holds pod-level security attributes and common container settings.
	SecurityContext *PodSecurityContextProps `field:"optional" json:"securityContext" yaml:"securityContext"`
	// A service account provides an identity for processes that run in a Pod.
	//
	// When you (a human) access the cluster (for example, using kubectl), you are
	// authenticated by the apiserver as a particular User Account (currently this
	// is usually admin, unless your cluster administrator has customized your
	// cluster). Processes in containers inside pods can also contact the
	// apiserver. When they do, they are authenticated as a particular Service
	// Account (for example, default).
	// See: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
	//
	ServiceAccount IServiceAccount `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// List of volumes that can be mounted by containers belonging to the pod.
	//
	// You can also add volumes later using `podSpec.addVolume()`
	// See: https://kubernetes.io/docs/concepts/storage/volumes
	//
	Volumes *[]Volume `field:"optional" json:"volumes" yaml:"volumes"`
}

// Controls the pod scheduling strategy.
type PodScheduling interface {
	Instance() AbstractPod
	// Assign this pod a specific node by name.
	//
	// The scheduler ignores the Pod, and the kubelet on the named node
	// tries to place the Pod on that node. Overrules any affinity rules of the pod.
	//
	// Some limitations of static assignment are:
	//
	// - If the named node does not exist, the Pod will not run, and in some
	//    cases may be automatically deleted.
	// - If the named node does not have the resources to accommodate the Pod,
	//    the Pod will fail and its reason will indicate why, for example OutOfmemory or OutOfcpu.
	// - Node names in cloud environments are not always predictable or stable.
	//
	// Will throw is the pod is already assigned to named node.
	//
	// Under the hood, this method utilizes the `nodeName` property.
	Assign(node NamedNode)
	// Attract this pod to a node matched by selectors. You can select a node by using `Node.labeled()`.
	//
	// Attracting to multiple nodes (i.e invoking this method multiple times) acts as
	// an OR condition, meaning the pod will be assigned to either one of the nodes.
	//
	// Under the hood, this method utilizes the `nodeAffinity` property.
	// See: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#node-affinity
	//
	Attract(node LabeledNode, options *PodSchedulingAttractOptions)
	// Co-locate this pod with a scheduling selection.
	//
	// A selection can be one of:
	//
	// - An instance of a `Pod`.
	// - An instance of a `Workload` (e.g `Deployment`, `StatefulSet`).
	// - An un-managed pod that can be selected via `Pods.select()`.
	//
	// Co-locating with multiple selections ((i.e invoking this method multiple times)) acts as
	// an AND condition. meaning the pod will be assigned to a node that satisfies all
	// selections (i.e runs at least one pod that satisifies each selection).
	//
	// Under the hood, this method utilizes the `podAffinity` property.
	// See: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#inter-pod-affinity-and-anti-affinity
	//
	Colocate(selector IPodSelector, options *PodSchedulingColocateOptions)
	// Seperate this pod from a scheduling selection.
	//
	// A selection can be one of:
	//
	// - An instance of a `Pod`.
	// - An instance of a `Workload` (e.g `Deployment`, `StatefulSet`).
	// - An un-managed pod that can be selected via `Pods.select()`.
	//
	// Seperating from multiple selections acts as an AND condition. meaning the pod
	// will not be assigned to a node that satisfies all selections (i.e runs at least one pod that satisifies each selection).
	//
	// Under the hood, this method utilizes the `podAntiAffinity` property.
	// See: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#inter-pod-affinity-and-anti-affinity
	//
	Separate(selector IPodSelector, options *PodSchedulingSeparateOptions)
	// Allow this pod to tolerate taints matching these tolerations.
	//
	// You can put multiple taints on the same node and multiple tolerations on the same pod.
	// The way Kubernetes processes multiple taints and tolerations is like a filter: start with
	// all of a node's taints, then ignore the ones for which the pod has a matching toleration;
	// the remaining un-ignored taints have the indicated effects on the pod. In particular:
	//
	// - if there is at least one un-ignored taint with effect NoSchedule then Kubernetes will
	//    not schedule the pod onto that node
	// - if there is no un-ignored taint with effect NoSchedule but there is at least one un-ignored
	//    taint with effect PreferNoSchedule then Kubernetes will try to not schedule the pod onto the node
	// - if there is at least one un-ignored taint with effect NoExecute then the pod will be evicted from
	//    the node (if it is already running on the node), and will not be scheduled onto the node (if it is
	//    not yet running on the node).
	//
	// Under the hood, this method utilizes the `tolerations` property.
	// See: https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/
	//
	Tolerate(node TaintedNode)
}

// The jsii proxy struct for PodScheduling
type jsiiProxy_PodScheduling struct {
	_ byte // padding
}

func (j *jsiiProxy_PodScheduling) Instance() AbstractPod {
	var returns AbstractPod
	_jsii_.Get(
		j,
		"instance",
		&returns,
	)
	return returns
}


func NewPodScheduling(instance AbstractPod) PodScheduling {
	_init_.Initialize()

	j := jsiiProxy_PodScheduling{}

	_jsii_.Create(
		"cdk8s-plus-24.PodScheduling",
		[]interface{}{instance},
		&j,
	)

	return &j
}

func NewPodScheduling_Override(p PodScheduling, instance AbstractPod) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.PodScheduling",
		[]interface{}{instance},
		p,
	)
}

func (p *jsiiProxy_PodScheduling) Assign(node NamedNode) {
	_jsii_.InvokeVoid(
		p,
		"assign",
		[]interface{}{node},
	)
}

func (p *jsiiProxy_PodScheduling) Attract(node LabeledNode, options *PodSchedulingAttractOptions) {
	_jsii_.InvokeVoid(
		p,
		"attract",
		[]interface{}{node, options},
	)
}

func (p *jsiiProxy_PodScheduling) Colocate(selector IPodSelector, options *PodSchedulingColocateOptions) {
	_jsii_.InvokeVoid(
		p,
		"colocate",
		[]interface{}{selector, options},
	)
}

func (p *jsiiProxy_PodScheduling) Separate(selector IPodSelector, options *PodSchedulingSeparateOptions) {
	_jsii_.InvokeVoid(
		p,
		"separate",
		[]interface{}{selector, options},
	)
}

func (p *jsiiProxy_PodScheduling) Tolerate(node TaintedNode) {
	_jsii_.InvokeVoid(
		p,
		"tolerate",
		[]interface{}{node},
	)
}

// Options for `PodScheduling.attract`.
type PodSchedulingAttractOptions struct {
	// Indicates the attraction is optional (soft), with this weight score.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

// Options for `PodScheduling.colocate`.
type PodSchedulingColocateOptions struct {
	// Which topology to coloate on.
	Topology Topology `field:"optional" json:"topology" yaml:"topology"`
	// Indicates the co-location is optional (soft), with this weight score.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

// Options for `PodScheduling.separate`.
type PodSchedulingSeparateOptions struct {
	// Which topology to separate on.
	Topology Topology `field:"optional" json:"topology" yaml:"topology"`
	// Indicates the separation is optional (soft), with this weight score.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

// Holds pod-level security attributes and common container settings.
type PodSecurityContext interface {
	EnsureNonRoot() *bool
	FsGroup() *float64
	FsGroupChangePolicy() FsGroupChangePolicy
	Group() *float64
	Sysctls() *[]*Sysctl
	User() *float64
}

// The jsii proxy struct for PodSecurityContext
type jsiiProxy_PodSecurityContext struct {
	_ byte // padding
}

func (j *jsiiProxy_PodSecurityContext) EnsureNonRoot() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"ensureNonRoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityContext) FsGroup() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fsGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityContext) FsGroupChangePolicy() FsGroupChangePolicy {
	var returns FsGroupChangePolicy
	_jsii_.Get(
		j,
		"fsGroupChangePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityContext) Group() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityContext) Sysctls() *[]*Sysctl {
	var returns *[]*Sysctl
	_jsii_.Get(
		j,
		"sysctls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityContext) User() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}


func NewPodSecurityContext(props *PodSecurityContextProps) PodSecurityContext {
	_init_.Initialize()

	j := jsiiProxy_PodSecurityContext{}

	_jsii_.Create(
		"cdk8s-plus-24.PodSecurityContext",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewPodSecurityContext_Override(p PodSecurityContext, props *PodSecurityContextProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.PodSecurityContext",
		[]interface{}{props},
		p,
	)
}

// Properties for `PodSecurityContext`.
type PodSecurityContextProps struct {
	// Indicates that the container must run as a non-root user.
	//
	// If true, the Kubelet will validate the image at runtime to ensure that it does
	// not run as UID 0 (root) and fail to start the container if it does.
	EnsureNonRoot *bool `field:"optional" json:"ensureNonRoot" yaml:"ensureNonRoot"`
	// Modify the ownership and permissions of pod volumes to this GID.
	FsGroup *float64 `field:"optional" json:"fsGroup" yaml:"fsGroup"`
	// Defines behavior of changing ownership and permission of the volume before being exposed inside Pod.
	//
	// This field will only apply to volume types which support fsGroup based ownership(and permissions).
	// It will have no effect on ephemeral volume types such as: secret, configmaps and emptydir.
	FsGroupChangePolicy FsGroupChangePolicy `field:"optional" json:"fsGroupChangePolicy" yaml:"fsGroupChangePolicy"`
	// The GID to run the entrypoint of the container process.
	Group *float64 `field:"optional" json:"group" yaml:"group"`
	// Sysctls hold a list of namespaced sysctls used for the pod.
	//
	// Pods with unsupported sysctls (by the container runtime) might fail to launch.
	Sysctls *[]*Sysctl `field:"optional" json:"sysctls" yaml:"sysctls"`
	// The UID to run the entrypoint of the container process.
	User *float64 `field:"optional" json:"user" yaml:"user"`
}

// Configuration for selecting pods, optionally in particular namespaces.
type PodSelectorConfig struct {
	// A selector to select pods by labels.
	LabelSelector LabelSelector `field:"required" json:"labelSelector" yaml:"labelSelector"`
	// Configuration for selecting which namepsaces are the pods allowed to be in.
	Namespaces *NamespaceSelectorConfig `field:"optional" json:"namespaces" yaml:"namespaces"`
}

// Represents a group of pods.
type Pods interface {
	constructs.Construct
	IPodSelector
	// The tree node.
	Node() constructs.Node
	// See: INetworkPolicyPeer.toNetworkPolicyPeerConfig()
	//
	ToNetworkPolicyPeerConfig() *NetworkPolicyPeerConfig
	// See: INetworkPolicyPeer.toPodSelector()
	//
	ToPodSelector() IPodSelector
	// Return the configuration of this selector.
	// See: IPodSelector.toPodSelectorConfig()
	//
	ToPodSelectorConfig() *PodSelectorConfig
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Pods
type jsiiProxy_Pods struct {
	internal.Type__constructsConstruct
	jsiiProxy_IPodSelector
}

func (j *jsiiProxy_Pods) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewPods(scope constructs.Construct, id *string, expressions *[]LabelExpression, labels *map[string]*string, namespaces INamespaceSelector) Pods {
	_init_.Initialize()

	j := jsiiProxy_Pods{}

	_jsii_.Create(
		"cdk8s-plus-24.Pods",
		[]interface{}{scope, id, expressions, labels, namespaces},
		&j,
	)

	return &j
}

func NewPods_Override(p Pods, scope constructs.Construct, id *string, expressions *[]LabelExpression, labels *map[string]*string, namespaces INamespaceSelector) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.Pods",
		[]interface{}{scope, id, expressions, labels, namespaces},
		p,
	)
}

// Select all pods.
func Pods_All(scope constructs.Construct, id *string, options *PodsAllOptions) Pods {
	_init_.Initialize()

	var returns Pods

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Pods",
		"all",
		[]interface{}{scope, id, options},
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
func Pods_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Pods",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Select pods in the cluster with various selectors.
func Pods_Select(scope constructs.Construct, id *string, options *PodsSelectOptions) Pods {
	_init_.Initialize()

	var returns Pods

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Pods",
		"select",
		[]interface{}{scope, id, options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pods) ToNetworkPolicyPeerConfig() *NetworkPolicyPeerConfig {
	var returns *NetworkPolicyPeerConfig

	_jsii_.Invoke(
		p,
		"toNetworkPolicyPeerConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pods) ToPodSelector() IPodSelector {
	var returns IPodSelector

	_jsii_.Invoke(
		p,
		"toPodSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pods) ToPodSelectorConfig() *PodSelectorConfig {
	var returns *PodSelectorConfig

	_jsii_.Invoke(
		p,
		"toPodSelectorConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pods) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options for `Pods.all`.
type PodsAllOptions struct {
	// Namespaces the pods are allowed to be in.
	//
	// Use `Namespaces.all()` to allow all namespaces.
	Namespaces Namespaces `field:"optional" json:"namespaces" yaml:"namespaces"`
}

// Options for `Pods.select`.
type PodsSelectOptions struct {
	// Expressions the pods must satisify.
	Expressions *[]LabelExpression `field:"optional" json:"expressions" yaml:"expressions"`
	// Labels the pods must have.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Namespaces the pods are allowed to be in.
	//
	// Use `Namespaces.all()` to allow all namespaces.
	Namespaces Namespaces `field:"optional" json:"namespaces" yaml:"namespaces"`
}

// Probe describes a health check to be performed against a container to determine whether it is alive or ready to receive traffic.
type Probe interface {
}

// The jsii proxy struct for Probe
type jsiiProxy_Probe struct {
	_ byte // padding
}

// Defines a probe based on a command which is executed within the container.
func Probe_FromCommand(command *[]*string, options *CommandProbeOptions) Probe {
	_init_.Initialize()

	var returns Probe

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Probe",
		"fromCommand",
		[]interface{}{command, options},
		&returns,
	)

	return returns
}

// Defines a probe based on an HTTP GET request to the IP address of the container.
func Probe_FromHttpGet(path *string, options *HttpGetProbeOptions) Probe {
	_init_.Initialize()

	var returns Probe

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Probe",
		"fromHttpGet",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Defines a probe based opening a connection to a TCP socket on the container.
func Probe_FromTcpSocket(options *TcpSocketProbeOptions) Probe {
	_init_.Initialize()

	var returns Probe

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Probe",
		"fromTcpSocket",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Probe options.
type ProbeOptions struct {
	// Minimum consecutive failures for the probe to be considered failed after having succeeded.
	//
	// Defaults to 3. Minimum value is 1.
	FailureThreshold *float64 `field:"optional" json:"failureThreshold" yaml:"failureThreshold"`
	// Number of seconds after the container has started before liveness probes are initiated.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	//
	InitialDelaySeconds cdk8s.Duration `field:"optional" json:"initialDelaySeconds" yaml:"initialDelaySeconds"`
	// How often (in seconds) to perform the probe.
	//
	// Default to 10 seconds. Minimum value is 1.
	PeriodSeconds cdk8s.Duration `field:"optional" json:"periodSeconds" yaml:"periodSeconds"`
	// Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1.
	//
	// Must be 1 for liveness and startup. Minimum value is 1.
	SuccessThreshold *float64 `field:"optional" json:"successThreshold" yaml:"successThreshold"`
	// Number of seconds after which the probe times out.
	//
	// Defaults to 1 second. Minimum value is 1.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	//
	TimeoutSeconds cdk8s.Duration `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
}

type Protocol string

const (
	Protocol_TCP Protocol = "TCP"
	Protocol_UDP Protocol = "UDP"
	Protocol_SCTP Protocol = "SCTP"
)

// Base class for all Kubernetes objects in stdk8s.
//
// Represents a single
// resource.
type Resource interface {
	constructs.Construct
	IApiEndpoint
	IApiResource
	IResource
	// The group portion of the API version (e.g. "authorization.k8s.io").
	ApiGroup() *string
	// The underlying cdk8s API object.
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
	// Return the IApiResource this object represents.
	AsApiResource() IApiResource
	// Return the non resource url this object represents.
	AsNonApiResource() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Resource
type jsiiProxy_Resource struct {
	internal.Type__constructsConstruct
	jsiiProxy_IApiEndpoint
	jsiiProxy_IApiResource
	jsiiProxy_IResource
}

func (j *jsiiProxy_Resource) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) Permissions() ResourcePermissions {
	var returns ResourcePermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}


func NewResource_Override(r Resource, scope constructs.Construct, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.Resource",
		[]interface{}{scope, id},
		r,
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
func Resource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Resource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resource) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		r,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resource) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ResourceFieldPaths string

const (
	// CPU limit of the container.
	ResourceFieldPaths_CPU_LIMIT ResourceFieldPaths = "CPU_LIMIT"
	// Memory limit of the container.
	ResourceFieldPaths_MEMORY_LIMIT ResourceFieldPaths = "MEMORY_LIMIT"
	// CPU request of the container.
	ResourceFieldPaths_CPU_REQUEST ResourceFieldPaths = "CPU_REQUEST"
	// Memory request of the container.
	ResourceFieldPaths_MEMORY_REQUEST ResourceFieldPaths = "MEMORY_REQUEST"
	// Ephemeral storage limit of the container.
	ResourceFieldPaths_STORAGE_LIMIT ResourceFieldPaths = "STORAGE_LIMIT"
	// Ephemeral storage request of the container.
	ResourceFieldPaths_STORAGE_REQUEST ResourceFieldPaths = "STORAGE_REQUEST"
)

// Controls permissions for operations on resources.
type ResourcePermissions interface {
	Instance() Resource
	// Grants the list of subjects permissions to read this resource.
	GrantRead(subjects ...ISubject) RoleBinding
	// Grants the list of subjects permissions to read and write this resource.
	GrantReadWrite(subjects ...ISubject) RoleBinding
}

// The jsii proxy struct for ResourcePermissions
type jsiiProxy_ResourcePermissions struct {
	_ byte // padding
}

func (j *jsiiProxy_ResourcePermissions) Instance() Resource {
	var returns Resource
	_jsii_.Get(
		j,
		"instance",
		&returns,
	)
	return returns
}


func NewResourcePermissions(instance Resource) ResourcePermissions {
	_init_.Initialize()

	j := jsiiProxy_ResourcePermissions{}

	_jsii_.Create(
		"cdk8s-plus-24.ResourcePermissions",
		[]interface{}{instance},
		&j,
	)

	return &j
}

func NewResourcePermissions_Override(r ResourcePermissions, instance Resource) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.ResourcePermissions",
		[]interface{}{instance},
		r,
	)
}

func (r *jsiiProxy_ResourcePermissions) GrantRead(subjects ...ISubject) RoleBinding {
	args := []interface{}{}
	for _, a := range subjects {
		args = append(args, a)
	}

	var returns RoleBinding

	_jsii_.Invoke(
		r,
		"grantRead",
		args,
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePermissions) GrantReadWrite(subjects ...ISubject) RoleBinding {
	args := []interface{}{}
	for _, a := range subjects {
		args = append(args, a)
	}

	var returns RoleBinding

	_jsii_.Invoke(
		r,
		"grantReadWrite",
		args,
		&returns,
	)

	return returns
}

// Initialization properties for resources.
type ResourceProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
}

// Restart policy for all containers within the pod.
type RestartPolicy string

const (
	// Always restart the pod after it exits.
	RestartPolicy_ALWAYS RestartPolicy = "ALWAYS"
	// Only restart if the pod exits with a non-zero exit code.
	RestartPolicy_ON_FAILURE RestartPolicy = "ON_FAILURE"
	// Never restart the pod.
	RestartPolicy_NEVER RestartPolicy = "NEVER"
)

// Role is a namespaced, logical grouping of PolicyRules that can be referenced as a unit by a RoleBinding.
type Role interface {
	Resource
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
	Rules() *[]*RolePolicyRule
	// Add permission to perform a list of HTTP verbs on a collection of resources.
	// See: https://kubernetes.io/docs/reference/access-authn-authz/authorization/#determine-the-request-verb
	//
	Allow(verbs *[]*string, resources ...IApiResource)
	// Add "create" permission for the resources.
	AllowCreate(resources ...IApiResource)
	// Add "delete" permission for the resources.
	AllowDelete(resources ...IApiResource)
	// Add "deletecollection" permission for the resources.
	AllowDeleteCollection(resources ...IApiResource)
	// Add "get" permission for the resources.
	AllowGet(resources ...IApiResource)
	// Add "list" permission for the resources.
	AllowList(resources ...IApiResource)
	// Add "patch" permission for the resources.
	AllowPatch(resources ...IApiResource)
	// Add "get", "list", and "watch" permissions for the resources.
	AllowRead(resources ...IApiResource)
	// Add "get", "list", "watch", "create", "update", "patch", "delete", and "deletecollection" permissions for the resources.
	AllowReadWrite(resources ...IApiResource)
	// Add "update" permission for the resources.
	AllowUpdate(resources ...IApiResource)
	// Add "watch" permission for the resources.
	AllowWatch(resources ...IApiResource)
	// Return the IApiResource this object represents.
	AsApiResource() IApiResource
	// Return the non resource url this object represents.
	AsNonApiResource() *string
	// Create a RoleBinding that binds the permissions in this Role to a list of subjects, that will only apply this role's namespace.
	Bind(subjects ...ISubject) RoleBinding
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Role
type jsiiProxy_Role struct {
	jsiiProxy_Resource
	jsiiProxy_IRole
}

func (j *jsiiProxy_Role) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) Permissions() ResourcePermissions {
	var returns ResourcePermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) Rules() *[]*RolePolicyRule {
	var returns *[]*RolePolicyRule
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}


func NewRole(scope constructs.Construct, id *string, props *RoleProps) Role {
	_init_.Initialize()

	j := jsiiProxy_Role{}

	_jsii_.Create(
		"cdk8s-plus-24.Role",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewRole_Override(r Role, scope constructs.Construct, id *string, props *RoleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.Role",
		[]interface{}{scope, id, props},
		r,
	)
}

// Imports a role from the cluster as a reference.
func Role_FromRoleName(scope constructs.Construct, id *string, name *string) IRole {
	_init_.Initialize()

	var returns IRole

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Role",
		"fromRoleName",
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
func Role_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Role",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Role) Allow(verbs *[]*string, resources ...IApiResource) {
	args := []interface{}{verbs}
	for _, a := range resources {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"allow",
		args,
	)
}

func (r *jsiiProxy_Role) AllowCreate(resources ...IApiResource) {
	args := []interface{}{}
	for _, a := range resources {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"allowCreate",
		args,
	)
}

func (r *jsiiProxy_Role) AllowDelete(resources ...IApiResource) {
	args := []interface{}{}
	for _, a := range resources {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"allowDelete",
		args,
	)
}

func (r *jsiiProxy_Role) AllowDeleteCollection(resources ...IApiResource) {
	args := []interface{}{}
	for _, a := range resources {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"allowDeleteCollection",
		args,
	)
}

func (r *jsiiProxy_Role) AllowGet(resources ...IApiResource) {
	args := []interface{}{}
	for _, a := range resources {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"allowGet",
		args,
	)
}

func (r *jsiiProxy_Role) AllowList(resources ...IApiResource) {
	args := []interface{}{}
	for _, a := range resources {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"allowList",
		args,
	)
}

func (r *jsiiProxy_Role) AllowPatch(resources ...IApiResource) {
	args := []interface{}{}
	for _, a := range resources {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"allowPatch",
		args,
	)
}

func (r *jsiiProxy_Role) AllowRead(resources ...IApiResource) {
	args := []interface{}{}
	for _, a := range resources {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"allowRead",
		args,
	)
}

func (r *jsiiProxy_Role) AllowReadWrite(resources ...IApiResource) {
	args := []interface{}{}
	for _, a := range resources {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"allowReadWrite",
		args,
	)
}

func (r *jsiiProxy_Role) AllowUpdate(resources ...IApiResource) {
	args := []interface{}{}
	for _, a := range resources {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"allowUpdate",
		args,
	)
}

func (r *jsiiProxy_Role) AllowWatch(resources ...IApiResource) {
	args := []interface{}{}
	for _, a := range resources {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"allowWatch",
		args,
	)
}

func (r *jsiiProxy_Role) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		r,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Role) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Role) Bind(subjects ...ISubject) RoleBinding {
	args := []interface{}{}
	for _, a := range subjects {
		args = append(args, a)
	}

	var returns RoleBinding

	_jsii_.Invoke(
		r,
		"bind",
		args,
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Role) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A RoleBinding grants permissions within a specific namespace to a user or set of users.
type RoleBinding interface {
	Resource
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
	Role() IRole
	Subjects() *[]ISubject
	// Adds a subject to the role.
	AddSubjects(subjects ...ISubject)
	// Return the IApiResource this object represents.
	AsApiResource() IApiResource
	// Return the non resource url this object represents.
	AsNonApiResource() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for RoleBinding
type jsiiProxy_RoleBinding struct {
	jsiiProxy_Resource
}

func (j *jsiiProxy_RoleBinding) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleBinding) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleBinding) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleBinding) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleBinding) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleBinding) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleBinding) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleBinding) Permissions() ResourcePermissions {
	var returns ResourcePermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleBinding) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleBinding) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleBinding) Role() IRole {
	var returns IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RoleBinding) Subjects() *[]ISubject {
	var returns *[]ISubject
	_jsii_.Get(
		j,
		"subjects",
		&returns,
	)
	return returns
}


func NewRoleBinding(scope constructs.Construct, id *string, props *RoleBindingProps) RoleBinding {
	_init_.Initialize()

	j := jsiiProxy_RoleBinding{}

	_jsii_.Create(
		"cdk8s-plus-24.RoleBinding",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewRoleBinding_Override(r RoleBinding, scope constructs.Construct, id *string, props *RoleBindingProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.RoleBinding",
		[]interface{}{scope, id, props},
		r,
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
func RoleBinding_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.RoleBinding",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RoleBinding) AddSubjects(subjects ...ISubject) {
	args := []interface{}{}
	for _, a := range subjects {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"addSubjects",
		args,
	)
}

func (r *jsiiProxy_RoleBinding) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		r,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RoleBinding) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RoleBinding) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for `RoleBinding`.
type RoleBindingProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// The role to bind to.
	//
	// A RoleBinding can reference a Role or a ClusterRole.
	Role IRole `field:"required" json:"role" yaml:"role"`
}

// Policy rule of a `Role.
type RolePolicyRule struct {
	// Resources this rule applies to.
	Resources *[]IApiResource `field:"required" json:"resources" yaml:"resources"`
	// Verbs to allow.
	//
	// (e.g ['get', 'watch'])
	Verbs *[]*string `field:"required" json:"verbs" yaml:"verbs"`
}

// Properties for `Role`.
type RoleProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// A list of rules the role should allow.
	Rules *[]*RolePolicyRule `field:"optional" json:"rules" yaml:"rules"`
}

// Kubernetes Secrets let you store and manage sensitive information, such as passwords, OAuth tokens, and ssh keys.
//
// Storing confidential information in a
// Secret is safer and more flexible than putting it verbatim in a Pod
// definition or in a container image.
// See: https://kubernetes.io/docs/concepts/configuration/secret
//
type Secret interface {
	Resource
	ISecret
	// The group portion of the API version (e.g. "authorization.k8s.io").
	ApiGroup() *string
	// The underlying cdk8s API object.
	// See: base.Resource.apiObject
	//
	ApiObject() cdk8s.ApiObject
	// The object's API version (e.g. "authorization.k8s.io/v1").
	ApiVersion() *string
	// Whether or not the secret is immutable.
	Immutable() *bool
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
	// Adds a string data field to the secert.
	AddStringData(key *string, value *string)
	// Return the IApiResource this object represents.
	AsApiResource() IApiResource
	// Return the non resource url this object represents.
	AsNonApiResource() *string
	// Gets a string data by key or undefined.
	GetStringData(key *string) *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Secret
type jsiiProxy_Secret struct {
	jsiiProxy_Resource
	jsiiProxy_ISecret
}

func (j *jsiiProxy_Secret) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Secret) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Secret) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Secret) Immutable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"immutable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Secret) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Secret) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Secret) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Secret) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Secret) Permissions() ResourcePermissions {
	var returns ResourcePermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Secret) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Secret) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}


func NewSecret(scope constructs.Construct, id *string, props *SecretProps) Secret {
	_init_.Initialize()

	j := jsiiProxy_Secret{}

	_jsii_.Create(
		"cdk8s-plus-24.Secret",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewSecret_Override(s Secret, scope constructs.Construct, id *string, props *SecretProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.Secret",
		[]interface{}{scope, id, props},
		s,
	)
}

// Imports a secret from the cluster as a reference.
func Secret_FromSecretName(scope constructs.Construct, id *string, name *string) ISecret {
	_init_.Initialize()

	var returns ISecret

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Secret",
		"fromSecretName",
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
func Secret_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Secret",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Secret) AddStringData(key *string, value *string) {
	_jsii_.InvokeVoid(
		s,
		"addStringData",
		[]interface{}{key, value},
	)
}

func (s *jsiiProxy_Secret) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		s,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Secret) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Secret) GetStringData(key *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringData",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Secret) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options for `Secret`.
type SecretProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// If set to true, ensures that data stored in the Secret cannot be updated (only object metadata can be modified).
	//
	// If not set to true, the field can be modified at any time.
	Immutable *bool `field:"optional" json:"immutable" yaml:"immutable"`
	// stringData allows specifying non-binary secret data in string form.
	//
	// It is
	// provided as a write-only convenience method. All keys and values are merged
	// into the data field on write, overwriting any existing values. It is never
	// output when reading from the API.
	StringData *map[string]*string `field:"optional" json:"stringData" yaml:"stringData"`
	// Optional type associated with the secret.
	//
	// Used to facilitate programmatic
	// handling of secret data by various controllers.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

// Represents a specific value in JSON secret.
type SecretValue struct {
	// The JSON key.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The secret.
	Secret ISecret `field:"required" json:"secret" yaml:"secret"`
}

// Options for the Secret-based volume.
type SecretVolumeOptions struct {
	// Mode bits to use on created files by default.
	//
	// Must be a value between 0 and
	// 0777. Defaults to 0644. Directories within the path are not affected by
	// this setting. This might be in conflict with other options that affect the
	// file mode, like fsGroup, and the result can be other mode bits set.
	DefaultMode *float64 `field:"optional" json:"defaultMode" yaml:"defaultMode"`
	// If unspecified, each key-value pair in the Data field of the referenced secret will be projected into the volume as a file whose name is the key and content is the value.
	//
	// If specified, the listed keys will be projected
	// into the specified paths, and unlisted keys will not be present. If a key
	// is specified which is not present in the secret, the volume setup will
	// error unless it is marked optional. Paths must be relative and may not
	// contain the '..' path or start with '..'.
	Items *map[string]*PathMapping `field:"optional" json:"items" yaml:"items"`
	// The volume name.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specify whether the secret or its keys must be defined.
	Optional *bool `field:"optional" json:"optional" yaml:"optional"`
}

// An abstract way to expose an application running on a set of Pods as a network service.
//
// With Kubernetes you don't need to modify your application to use an unfamiliar service discovery mechanism.
// Kubernetes gives Pods their own IP addresses and a single DNS name for a set of Pods, and can load-balance across them.
//
// For example, consider a stateless image-processing backend which is running with 3 replicas. Those replicas are fungible???frontends do not care which backend they use.
// While the actual Pods that compose the backend set may change, the frontend clients should not need to be aware of that,
// nor should they need to keep track of the set of backends themselves.
// The Service abstraction enables this decoupling.
//
// If you're able to use Kubernetes APIs for service discovery in your application, you can query the API server for Endpoints,
// that get updated whenever the set of Pods in a Service changes. For non-native applications, Kubernetes offers ways to place a network port
// or load balancer in between your application and the backend Pods.
type Service interface {
	Resource
	// The group portion of the API version (e.g. "authorization.k8s.io").
	ApiGroup() *string
	// The underlying cdk8s API object.
	// See: base.Resource.apiObject
	//
	ApiObject() cdk8s.ApiObject
	// The object's API version (e.g. "authorization.k8s.io/v1").
	ApiVersion() *string
	// The IP address of the service and is usually assigned randomly by the master.
	ClusterIP() *string
	// The externalName to be used for EXTERNAL_NAME types.
	ExternalName() *string
	// The object kind (e.g. "Deployment").
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	// The name of this API object.
	Name() *string
	// The tree node.
	Node() constructs.Node
	Permissions() ResourcePermissions
	// Ports for this service.
	//
	// Use `bind()` to bind additional service ports.
	Ports() *[]*ServicePort
	// The unique, namespace-global, name of an object inside the Kubernetes cluster.
	//
	// If this is omitted, the ApiResource should represent all objects of the given type.
	ResourceName() *string
	// The name of a resource type as it appears in the relevant API endpoint.
	ResourceType() *string
	// Determines how the Service is exposed.
	Type() ServiceType
	// Return the IApiResource this object represents.
	AsApiResource() IApiResource
	// Return the non resource url this object represents.
	AsNonApiResource() *string
	// Configure a port the service will bind to.
	//
	// This method can be called multiple times.
	Bind(port *float64, options *ServiceBindOptions)
	// Expose a service via an ingress using the specified path.
	//
	// Returns: The `Ingress` resource that was used.
	ExposeViaIngress(path *string, options *ExposeServiceViaIngressOptions) Ingress
	// Require this service to select pods matching the selector.
	//
	// Note that invoking this method multiple times acts as an AND operator
	// on the resulting labels.
	Select(selector IPodSelector)
	// Require this service to select pods with this label.
	//
	// Note that invoking this method multiple times acts as an AND operator
	// on the resulting labels.
	SelectLabel(key *string, value *string)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Service
type jsiiProxy_Service struct {
	jsiiProxy_Resource
}

func (j *jsiiProxy_Service) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) ClusterIP() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIP",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) ExternalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Permissions() ResourcePermissions {
	var returns ResourcePermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Ports() *[]*ServicePort {
	var returns *[]*ServicePort
	_jsii_.Get(
		j,
		"ports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Type() ServiceType {
	var returns ServiceType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


func NewService(scope constructs.Construct, id *string, props *ServiceProps) Service {
	_init_.Initialize()

	j := jsiiProxy_Service{}

	_jsii_.Create(
		"cdk8s-plus-24.Service",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewService_Override(s Service, scope constructs.Construct, id *string, props *ServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.Service",
		[]interface{}{scope, id, props},
		s,
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
func Service_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Service",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Service) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		s,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Service) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Service) Bind(port *float64, options *ServiceBindOptions) {
	_jsii_.InvokeVoid(
		s,
		"bind",
		[]interface{}{port, options},
	)
}

func (s *jsiiProxy_Service) ExposeViaIngress(path *string, options *ExposeServiceViaIngressOptions) Ingress {
	var returns Ingress

	_jsii_.Invoke(
		s,
		"exposeViaIngress",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Service) Select(selector IPodSelector) {
	_jsii_.InvokeVoid(
		s,
		"select",
		[]interface{}{selector},
	)
}

func (s *jsiiProxy_Service) SelectLabel(key *string, value *string) {
	_jsii_.InvokeVoid(
		s,
		"selectLabel",
		[]interface{}{key, value},
	)
}

func (s *jsiiProxy_Service) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A service account provides an identity for processes that run in a Pod.
//
// When you (a human) access the cluster (for example, using kubectl), you are
// authenticated by the apiserver as a particular User Account (currently this
// is usually admin, unless your cluster administrator has customized your
// cluster). Processes in containers inside pods can also contact the apiserver.
// When they do, they are authenticated as a particular Service Account (for
// example, default).
// See: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account
//
type ServiceAccount interface {
	Resource
	IServiceAccount
	ISubject
	// The group portion of the API version (e.g. "authorization.k8s.io").
	ApiGroup() *string
	// The underlying cdk8s API object.
	// See: base.Resource.apiObject
	//
	ApiObject() cdk8s.ApiObject
	// The object's API version (e.g. "authorization.k8s.io/v1").
	ApiVersion() *string
	// Whether or not a token is automatically mounted for this service account.
	AutomountToken() *bool
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
	// List of secrets allowed to be used by pods running using this service account.
	//
	// Returns a copy. To add a secret, use `addSecret()`.
	Secrets() *[]ISecret
	// Allow a secret to be accessed by pods using this service account.
	AddSecret(secr ISecret)
	// Return the IApiResource this object represents.
	AsApiResource() IApiResource
	// Return the non resource url this object represents.
	AsNonApiResource() *string
	// Returns a string representation of this construct.
	ToString() *string
	// Return the subject configuration.
	// See: ISubect.toSubjectConfiguration()
	//
	ToSubjectConfiguration() *SubjectConfiguration
}

// The jsii proxy struct for ServiceAccount
type jsiiProxy_ServiceAccount struct {
	jsiiProxy_Resource
	jsiiProxy_IServiceAccount
	jsiiProxy_ISubject
}

func (j *jsiiProxy_ServiceAccount) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceAccount) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceAccount) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceAccount) AutomountToken() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"automountToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceAccount) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceAccount) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceAccount) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceAccount) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceAccount) Permissions() ResourcePermissions {
	var returns ResourcePermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceAccount) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceAccount) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceAccount) Secrets() *[]ISecret {
	var returns *[]ISecret
	_jsii_.Get(
		j,
		"secrets",
		&returns,
	)
	return returns
}


func NewServiceAccount(scope constructs.Construct, id *string, props *ServiceAccountProps) ServiceAccount {
	_init_.Initialize()

	j := jsiiProxy_ServiceAccount{}

	_jsii_.Create(
		"cdk8s-plus-24.ServiceAccount",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewServiceAccount_Override(s ServiceAccount, scope constructs.Construct, id *string, props *ServiceAccountProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.ServiceAccount",
		[]interface{}{scope, id, props},
		s,
	)
}

// Imports a service account from the cluster as a reference.
func ServiceAccount_FromServiceAccountName(scope constructs.Construct, id *string, name *string) IServiceAccount {
	_init_.Initialize()

	var returns IServiceAccount

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.ServiceAccount",
		"fromServiceAccountName",
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
func ServiceAccount_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.ServiceAccount",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceAccount) AddSecret(secr ISecret) {
	_jsii_.InvokeVoid(
		s,
		"addSecret",
		[]interface{}{secr},
	)
}

func (s *jsiiProxy_ServiceAccount) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		s,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceAccount) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceAccount) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceAccount) ToSubjectConfiguration() *SubjectConfiguration {
	var returns *SubjectConfiguration

	_jsii_.Invoke(
		s,
		"toSubjectConfiguration",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for initialization of `ServiceAccount`.
type ServiceAccountProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// Indicates whether pods running as this service account should have an API token automatically mounted.
	//
	// Can be overridden at the pod level.
	// See: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/#use-the-default-service-account-to-access-the-api-server
	//
	AutomountToken *bool `field:"optional" json:"automountToken" yaml:"automountToken"`
	// List of secrets allowed to be used by pods running using this ServiceAccount.
	// See: https://kubernetes.io/docs/concepts/configuration/secret
	//
	Secrets *[]ISecret `field:"optional" json:"secrets" yaml:"secrets"`
}

// Create a secret for a service account token.
// See: https://kubernetes.io/docs/concepts/configuration/secret/#service-account-token-secrets
//
type ServiceAccountTokenSecret interface {
	Secret
	// The group portion of the API version (e.g. "authorization.k8s.io").
	ApiGroup() *string
	// The underlying cdk8s API object.
	// See: base.Resource.apiObject
	//
	ApiObject() cdk8s.ApiObject
	// The object's API version (e.g. "authorization.k8s.io/v1").
	ApiVersion() *string
	// Whether or not the secret is immutable.
	Immutable() *bool
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
	// Adds a string data field to the secert.
	AddStringData(key *string, value *string)
	// Return the IApiResource this object represents.
	AsApiResource() IApiResource
	// Return the non resource url this object represents.
	AsNonApiResource() *string
	// Gets a string data by key or undefined.
	GetStringData(key *string) *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for ServiceAccountTokenSecret
type jsiiProxy_ServiceAccountTokenSecret struct {
	jsiiProxy_Secret
}

func (j *jsiiProxy_ServiceAccountTokenSecret) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceAccountTokenSecret) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceAccountTokenSecret) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceAccountTokenSecret) Immutable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"immutable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceAccountTokenSecret) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceAccountTokenSecret) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceAccountTokenSecret) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceAccountTokenSecret) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceAccountTokenSecret) Permissions() ResourcePermissions {
	var returns ResourcePermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceAccountTokenSecret) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceAccountTokenSecret) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}


func NewServiceAccountTokenSecret(scope constructs.Construct, id *string, props *ServiceAccountTokenSecretProps) ServiceAccountTokenSecret {
	_init_.Initialize()

	j := jsiiProxy_ServiceAccountTokenSecret{}

	_jsii_.Create(
		"cdk8s-plus-24.ServiceAccountTokenSecret",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewServiceAccountTokenSecret_Override(s ServiceAccountTokenSecret, scope constructs.Construct, id *string, props *ServiceAccountTokenSecretProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.ServiceAccountTokenSecret",
		[]interface{}{scope, id, props},
		s,
	)
}

// Imports a secret from the cluster as a reference.
func ServiceAccountTokenSecret_FromSecretName(scope constructs.Construct, id *string, name *string) ISecret {
	_init_.Initialize()

	var returns ISecret

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.ServiceAccountTokenSecret",
		"fromSecretName",
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
func ServiceAccountTokenSecret_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.ServiceAccountTokenSecret",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceAccountTokenSecret) AddStringData(key *string, value *string) {
	_jsii_.InvokeVoid(
		s,
		"addStringData",
		[]interface{}{key, value},
	)
}

func (s *jsiiProxy_ServiceAccountTokenSecret) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		s,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceAccountTokenSecret) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceAccountTokenSecret) GetStringData(key *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringData",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceAccountTokenSecret) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options for `ServiceAccountTokenSecret`.
type ServiceAccountTokenSecretProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// If set to true, ensures that data stored in the Secret cannot be updated (only object metadata can be modified).
	//
	// If not set to true, the field can be modified at any time.
	Immutable *bool `field:"optional" json:"immutable" yaml:"immutable"`
	// The service account to store a secret for.
	ServiceAccount IServiceAccount `field:"required" json:"serviceAccount" yaml:"serviceAccount"`
}

// Options for `Service.bind`.
type ServiceBindOptions struct {
	// The name of this port within the service.
	//
	// This must be a DNS_LABEL. All
	// ports within a ServiceSpec must have unique names. This maps to the 'Name'
	// field in EndpointPort objects. Optional if only one ServicePort is defined
	// on this service.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The port on each node on which this service is exposed when type=NodePort or LoadBalancer.
	//
	// Usually assigned by the system. If specified, it will be
	// allocated to the service if unused or else creation of the service will
	// fail. Default is to auto-allocate a port if the ServiceType of this Service
	// requires one.
	// See: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport
	//
	NodePort *float64 `field:"optional" json:"nodePort" yaml:"nodePort"`
	// The IP protocol for this port.
	//
	// Supports "TCP", "UDP", and "SCTP". Default is TCP.
	Protocol Protocol `field:"optional" json:"protocol" yaml:"protocol"`
	// The port number the service will redirect to.
	TargetPort *float64 `field:"optional" json:"targetPort" yaml:"targetPort"`
}

// Options for setting up backends for ingress rules.
type ServiceIngressBackendOptions struct {
	// The port to use to access the service.
	//
	// - This option will fail if the service does not expose any ports.
	// - If the service exposes multiple ports, this option must be specified.
	// - If the service exposes a single port, this option is optional and if
	//    specified, it must be the same port exposed by the service.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

// Definition of a service port.
type ServicePort struct {
	// The name of this port within the service.
	//
	// This must be a DNS_LABEL. All
	// ports within a ServiceSpec must have unique names. This maps to the 'Name'
	// field in EndpointPort objects. Optional if only one ServicePort is defined
	// on this service.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The port on each node on which this service is exposed when type=NodePort or LoadBalancer.
	//
	// Usually assigned by the system. If specified, it will be
	// allocated to the service if unused or else creation of the service will
	// fail. Default is to auto-allocate a port if the ServiceType of this Service
	// requires one.
	// See: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport
	//
	NodePort *float64 `field:"optional" json:"nodePort" yaml:"nodePort"`
	// The IP protocol for this port.
	//
	// Supports "TCP", "UDP", and "SCTP". Default is TCP.
	Protocol Protocol `field:"optional" json:"protocol" yaml:"protocol"`
	// The port number the service will redirect to.
	TargetPort *float64 `field:"optional" json:"targetPort" yaml:"targetPort"`
	// The port number the service will bind to.
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

// Properties for `Service`.
type ServiceProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// The IP address of the service and is usually assigned randomly by the master.
	//
	// If an address is specified manually and is not in use by others, it
	// will be allocated to the service; otherwise, creation of the service will
	// fail. This field can not be changed through updates. Valid values are
	// "None", empty string (""), or a valid IP address. "None" can be specified
	// for headless services when proxying is not required. Only applies to types
	// ClusterIP, NodePort, and LoadBalancer. Ignored if type is ExternalName.
	// See: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	//
	ClusterIP *string `field:"optional" json:"clusterIP" yaml:"clusterIP"`
	// A list of IP addresses for which nodes in the cluster will also accept traffic for this service.
	//
	// These IPs are not managed by Kubernetes. The user
	// is responsible for ensuring that traffic arrives at a node with this IP. A
	// common example is external load-balancers that are not part of the
	// Kubernetes system.
	ExternalIPs *[]*string `field:"optional" json:"externalIPs" yaml:"externalIPs"`
	// The externalName to be used when ServiceType.EXTERNAL_NAME is set.
	ExternalName *string `field:"optional" json:"externalName" yaml:"externalName"`
	// A list of CIDR IP addresses, if specified and supported by the platform, will restrict traffic through the cloud-provider load-balancer to the specified client IPs.
	//
	// More info: https://kubernetes.io/docs/tasks/access-application-cluster/configure-cloud-provider-firewall/
	LoadBalancerSourceRanges *[]*string `field:"optional" json:"loadBalancerSourceRanges" yaml:"loadBalancerSourceRanges"`
	// The ports this service binds to.
	//
	// If the selector of the service is a managed pod / workload,
	// its ports will are automatically extracted and used as the default value.
	// Otherwise, no ports are bound.
	Ports *[]*ServicePort `field:"optional" json:"ports" yaml:"ports"`
	// Which pods should the service select and route to.
	//
	// You can pass one of the following:
	//
	// - An instance of `Pod` or any workload resource (e.g `Deployment`, `StatefulSet`, ...)
	// - Pods selected by the `Pods.select` function. Note that in this case only labels can be specified.
	//
	// Example:
	//   // select the pods of a specific deployment
	//   const backend = new kplus.Deployment(this, 'Backend', ...);
	//   new kplus.Service(this, 'Service', { selector: backend });
	//
	//   // select all pods labeled with the `tier=backend` label
	//   const backend = kplus.Pod.labeled({ tier: 'backend' });
	//   new kplus.Service(this, 'Service', { selector: backend });
	//
	Selector IPodSelector `field:"optional" json:"selector" yaml:"selector"`
	// Determines how the Service is exposed.
	//
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types
	Type ServiceType `field:"optional" json:"type" yaml:"type"`
}

// For some parts of your application (for example, frontends) you may want to expose a Service onto an external IP address, that's outside of your cluster.
//
// Kubernetes ServiceTypes allow you to specify what kind of Service you want.
// The default is ClusterIP.
type ServiceType string

const (
	// Exposes the Service on a cluster-internal IP.
	//
	// Choosing this value makes the Service only reachable from within the cluster.
	// This is the default ServiceType.
	ServiceType_CLUSTER_IP ServiceType = "CLUSTER_IP"
	// Exposes the Service on each Node's IP at a static port (the NodePort).
	//
	// A ClusterIP Service, to which the NodePort Service routes, is automatically created.
	// You'll be able to contact the NodePort Service, from outside the cluster,
	// by requesting <NodeIP>:<NodePort>.
	ServiceType_NODE_PORT ServiceType = "NODE_PORT"
	// Exposes the Service externally using a cloud provider's load balancer.
	//
	// NodePort and ClusterIP Services, to which the external load balancer routes,
	// are automatically created.
	ServiceType_LOAD_BALANCER ServiceType = "LOAD_BALANCER"
	// Maps the Service to the contents of the externalName field (e.g. foo.bar.example.com), by returning a CNAME record with its value. No proxying of any kind is set up.
	//
	// > Note: You need either kube-dns version 1.7 or CoreDNS version 0.0.8 or higher to use the ExternalName type.
	ServiceType_EXTERNAL_NAME ServiceType = "EXTERNAL_NAME"
)

// Create a secret for ssh authentication.
// See: https://kubernetes.io/docs/concepts/configuration/secret/#ssh-authentication-secrets
//
type SshAuthSecret interface {
	Secret
	// The group portion of the API version (e.g. "authorization.k8s.io").
	ApiGroup() *string
	// The underlying cdk8s API object.
	// See: base.Resource.apiObject
	//
	ApiObject() cdk8s.ApiObject
	// The object's API version (e.g. "authorization.k8s.io/v1").
	ApiVersion() *string
	// Whether or not the secret is immutable.
	Immutable() *bool
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
	// Adds a string data field to the secert.
	AddStringData(key *string, value *string)
	// Return the IApiResource this object represents.
	AsApiResource() IApiResource
	// Return the non resource url this object represents.
	AsNonApiResource() *string
	// Gets a string data by key or undefined.
	GetStringData(key *string) *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for SshAuthSecret
type jsiiProxy_SshAuthSecret struct {
	jsiiProxy_Secret
}

func (j *jsiiProxy_SshAuthSecret) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshAuthSecret) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshAuthSecret) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshAuthSecret) Immutable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"immutable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshAuthSecret) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshAuthSecret) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshAuthSecret) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshAuthSecret) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshAuthSecret) Permissions() ResourcePermissions {
	var returns ResourcePermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshAuthSecret) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshAuthSecret) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}


func NewSshAuthSecret(scope constructs.Construct, id *string, props *SshAuthSecretProps) SshAuthSecret {
	_init_.Initialize()

	j := jsiiProxy_SshAuthSecret{}

	_jsii_.Create(
		"cdk8s-plus-24.SshAuthSecret",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewSshAuthSecret_Override(s SshAuthSecret, scope constructs.Construct, id *string, props *SshAuthSecretProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.SshAuthSecret",
		[]interface{}{scope, id, props},
		s,
	)
}

// Imports a secret from the cluster as a reference.
func SshAuthSecret_FromSecretName(scope constructs.Construct, id *string, name *string) ISecret {
	_init_.Initialize()

	var returns ISecret

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.SshAuthSecret",
		"fromSecretName",
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
func SshAuthSecret_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.SshAuthSecret",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SshAuthSecret) AddStringData(key *string, value *string) {
	_jsii_.InvokeVoid(
		s,
		"addStringData",
		[]interface{}{key, value},
	)
}

func (s *jsiiProxy_SshAuthSecret) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		s,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SshAuthSecret) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SshAuthSecret) GetStringData(key *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringData",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SshAuthSecret) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options for `SshAuthSecret`.
type SshAuthSecretProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// If set to true, ensures that data stored in the Secret cannot be updated (only object metadata can be modified).
	//
	// If not set to true, the field can be modified at any time.
	Immutable *bool `field:"optional" json:"immutable" yaml:"immutable"`
	// The SSH private key to use.
	SshPrivateKey *string `field:"required" json:"sshPrivateKey" yaml:"sshPrivateKey"`
}

// StatefulSet is the workload API object used to manage stateful applications.
//
// Manages the deployment and scaling of a set of Pods, and provides guarantees
// about the ordering and uniqueness of these Pods.
//
// Like a Deployment, a StatefulSet manages Pods that are based on an identical
// container spec. Unlike a Deployment, a StatefulSet maintains a sticky identity
// for each of their Pods. These pods are created from the same spec, but are not
// interchangeable: each has a persistent identifier that it maintains across any
// rescheduling.
//
// If you want to use storage volumes to provide persistence for your workload, you
// can use a StatefulSet as part of the solution. Although individual Pods in a StatefulSet
// are susceptible to failure, the persistent Pod identifiers make it easier to match existing
// volumes to the new Pods that replace any that have failed.
//
// Using StatefulSets
// ------------------
// StatefulSets are valuable for applications that require one or more of the following.
//
// - Stable, unique network identifiers.
// - Stable, persistent storage.
// - Ordered, graceful deployment and scaling.
// - Ordered, automated rolling updates.
type StatefulSet interface {
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
	Connections() PodConnections
	Containers() *[]Container
	Dns() PodDns
	DockerRegistryAuth() DockerConfigSecret
	HostAliases() *[]*HostAlias
	InitContainers() *[]Container
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
	// Management policy to use for the set.
	PodManagementPolicy() PodManagementPolicy
	// The metadata of pods in this workload.
	PodMetadata() cdk8s.ApiObjectMetadataDefinition
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
	// The update startegy of this stateful set.
	Strategy() StatefulSetUpdateStrategy
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

// The jsii proxy struct for StatefulSet
type jsiiProxy_StatefulSet struct {
	jsiiProxy_Workload
}

func (j *jsiiProxy_StatefulSet) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSet) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSet) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSet) AutomountServiceAccountToken() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"automountServiceAccountToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSet) Connections() PodConnections {
	var returns PodConnections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSet) Containers() *[]Container {
	var returns *[]Container
	_jsii_.Get(
		j,
		"containers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSet) Dns() PodDns {
	var returns PodDns
	_jsii_.Get(
		j,
		"dns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSet) DockerRegistryAuth() DockerConfigSecret {
	var returns DockerConfigSecret
	_jsii_.Get(
		j,
		"dockerRegistryAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSet) HostAliases() *[]*HostAlias {
	var returns *[]*HostAlias
	_jsii_.Get(
		j,
		"hostAliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSet) InitContainers() *[]Container {
	var returns *[]Container
	_jsii_.Get(
		j,
		"initContainers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSet) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSet) MatchExpressions() *[]*LabelSelectorRequirement {
	var returns *[]*LabelSelectorRequirement
	_jsii_.Get(
		j,
		"matchExpressions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSet) MatchLabels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"matchLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSet) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSet) MinReady() cdk8s.Duration {
	var returns cdk8s.Duration
	_jsii_.Get(
		j,
		"minReady",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSet) Permissions() ResourcePermissions {
	var returns ResourcePermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSet) PodManagementPolicy() PodManagementPolicy {
	var returns PodManagementPolicy
	_jsii_.Get(
		j,
		"podManagementPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSet) PodMetadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"podMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSet) Replicas() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSet) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSet) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSet) RestartPolicy() RestartPolicy {
	var returns RestartPolicy
	_jsii_.Get(
		j,
		"restartPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSet) Scheduling() WorkloadScheduling {
	var returns WorkloadScheduling
	_jsii_.Get(
		j,
		"scheduling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSet) SecurityContext() PodSecurityContext {
	var returns PodSecurityContext
	_jsii_.Get(
		j,
		"securityContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSet) ServiceAccount() IServiceAccount {
	var returns IServiceAccount
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSet) Strategy() StatefulSetUpdateStrategy {
	var returns StatefulSetUpdateStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSet) Volumes() *[]Volume {
	var returns *[]Volume
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}


func NewStatefulSet(scope constructs.Construct, id *string, props *StatefulSetProps) StatefulSet {
	_init_.Initialize()

	j := jsiiProxy_StatefulSet{}

	_jsii_.Create(
		"cdk8s-plus-24.StatefulSet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewStatefulSet_Override(s StatefulSet, scope constructs.Construct, id *string, props *StatefulSetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.StatefulSet",
		[]interface{}{scope, id, props},
		s,
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
func StatefulSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.StatefulSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSet) AddContainer(cont *ContainerProps) Container {
	var returns Container

	_jsii_.Invoke(
		s,
		"addContainer",
		[]interface{}{cont},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSet) AddHostAlias(hostAlias *HostAlias) {
	_jsii_.InvokeVoid(
		s,
		"addHostAlias",
		[]interface{}{hostAlias},
	)
}

func (s *jsiiProxy_StatefulSet) AddInitContainer(cont *ContainerProps) Container {
	var returns Container

	_jsii_.Invoke(
		s,
		"addInitContainer",
		[]interface{}{cont},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSet) AddVolume(vol Volume) {
	_jsii_.InvokeVoid(
		s,
		"addVolume",
		[]interface{}{vol},
	)
}

func (s *jsiiProxy_StatefulSet) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		s,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSet) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSet) Select(selectors ...LabelSelector) {
	args := []interface{}{}
	for _, a := range selectors {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		s,
		"select",
		args,
	)
}

func (s *jsiiProxy_StatefulSet) ToNetworkPolicyPeerConfig() *NetworkPolicyPeerConfig {
	var returns *NetworkPolicyPeerConfig

	_jsii_.Invoke(
		s,
		"toNetworkPolicyPeerConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSet) ToPodSelector() IPodSelector {
	var returns IPodSelector

	_jsii_.Invoke(
		s,
		"toPodSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSet) ToPodSelectorConfig() *PodSelectorConfig {
	var returns *PodSelectorConfig

	_jsii_.Invoke(
		s,
		"toPodSelectorConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSet) ToSubjectConfiguration() *SubjectConfiguration {
	var returns *SubjectConfiguration

	_jsii_.Invoke(
		s,
		"toSubjectConfiguration",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for initialization of `StatefulSet`.
type StatefulSetProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// Indicates whether a service account token should be automatically mounted.
	// See: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/#use-the-default-service-account-to-access-the-api-server
	//
	AutomountServiceAccountToken *bool `field:"optional" json:"automountServiceAccountToken" yaml:"automountServiceAccountToken"`
	// List of containers belonging to the pod.
	//
	// Containers cannot currently be
	// added or removed. There must be at least one container in a Pod.
	//
	// You can add additionnal containers using `podSpec.addContainer()`
	Containers *[]*ContainerProps `field:"optional" json:"containers" yaml:"containers"`
	// DNS settings for the pod.
	// See: https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/
	//
	Dns *PodDnsProps `field:"optional" json:"dns" yaml:"dns"`
	// A secret containing docker credentials for authenticating to a registry.
	DockerRegistryAuth DockerConfigSecret `field:"optional" json:"dockerRegistryAuth" yaml:"dockerRegistryAuth"`
	// HostAlias holds the mapping between IP and hostnames that will be injected as an entry in the pod's hosts file.
	HostAliases *[]*HostAlias `field:"optional" json:"hostAliases" yaml:"hostAliases"`
	// List of initialization containers belonging to the pod.
	//
	// Init containers are executed in order prior to containers being started.
	// If any init container fails, the pod is considered to have failed and is handled according to its restartPolicy.
	// The name for an init container or normal container must be unique among all containers.
	// Init containers may not have Lifecycle actions, Readiness probes, Liveness probes, or Startup probes.
	// The resourceRequirements of an init container are taken into account during scheduling by finding the highest request/limit
	// for each resource type, and then using the max of of that value or the sum of the normal containers.
	// Limits are applied to init containers in a similar fashion.
	//
	// Init containers cannot currently be added ,removed or updated.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/
	//
	InitContainers *[]*ContainerProps `field:"optional" json:"initContainers" yaml:"initContainers"`
	// Restart policy for all containers within the pod.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy
	//
	RestartPolicy RestartPolicy `field:"optional" json:"restartPolicy" yaml:"restartPolicy"`
	// SecurityContext holds pod-level security attributes and common container settings.
	SecurityContext *PodSecurityContextProps `field:"optional" json:"securityContext" yaml:"securityContext"`
	// A service account provides an identity for processes that run in a Pod.
	//
	// When you (a human) access the cluster (for example, using kubectl), you are
	// authenticated by the apiserver as a particular User Account (currently this
	// is usually admin, unless your cluster administrator has customized your
	// cluster). Processes in containers inside pods can also contact the
	// apiserver. When they do, they are authenticated as a particular Service
	// Account (for example, default).
	// See: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
	//
	ServiceAccount IServiceAccount `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// List of volumes that can be mounted by containers belonging to the pod.
	//
	// You can also add volumes later using `podSpec.addVolume()`
	// See: https://kubernetes.io/docs/concepts/storage/volumes
	//
	Volumes *[]Volume `field:"optional" json:"volumes" yaml:"volumes"`
	// The pod metadata of this workload.
	PodMetadata *cdk8s.ApiObjectMetadata `field:"optional" json:"podMetadata" yaml:"podMetadata"`
	// Automatically allocates a pod label selector for this workload and add it to the pod metadata.
	//
	// This ensures this workload manages pods created by
	// its pod template.
	Select *bool `field:"optional" json:"select" yaml:"select"`
	// Service to associate with the statefulset.
	Service Service `field:"required" json:"service" yaml:"service"`
	// Minimum duration for which a newly created pod should be ready without any of its container crashing, for it to be considered available.
	//
	// Zero means the pod will be considered available as soon as it is ready.
	//
	// This is an alpha field and requires enabling StatefulSetMinReadySeconds feature gate.
	// See: https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#min-ready-seconds
	//
	MinReady cdk8s.Duration `field:"optional" json:"minReady" yaml:"minReady"`
	// Pod management policy to use for this statefulset.
	PodManagementPolicy PodManagementPolicy `field:"optional" json:"podManagementPolicy" yaml:"podManagementPolicy"`
	// Number of desired pods.
	Replicas *float64 `field:"optional" json:"replicas" yaml:"replicas"`
	// Indicates the StatefulSetUpdateStrategy that will be employed to update Pods in the StatefulSet when a revision is made to Template.
	Strategy StatefulSetUpdateStrategy `field:"optional" json:"strategy" yaml:"strategy"`
}

// StatefulSet update strategies.
type StatefulSetUpdateStrategy interface {
}

// The jsii proxy struct for StatefulSetUpdateStrategy
type jsiiProxy_StatefulSetUpdateStrategy struct {
	_ byte // padding
}

// The controller will not automatically update the Pods in a StatefulSet.
//
// Users must manually delete Pods to cause the controller to create new Pods
// that reflect modifications.
func StatefulSetUpdateStrategy_OnDelete() StatefulSetUpdateStrategy {
	_init_.Initialize()

	var returns StatefulSetUpdateStrategy

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.StatefulSetUpdateStrategy",
		"onDelete",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The controller will delete and recreate each Pod in the StatefulSet.
//
// It will proceed in the same order as Pod termination (from the largest ordinal to the smallest),
// updating each Pod one at a time. The Kubernetes control plane waits until an updated
// Pod is Running and Ready prior to updating its predecessor.
func StatefulSetUpdateStrategy_RollingUpdate(options *StatefulSetUpdateStrategyRollingUpdateOptions) StatefulSetUpdateStrategy {
	_init_.Initialize()

	var returns StatefulSetUpdateStrategy

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.StatefulSetUpdateStrategy",
		"rollingUpdate",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Options for `StatefulSetUpdateStrategy.rollingUpdate`.
type StatefulSetUpdateStrategyRollingUpdateOptions struct {
	// If specified, all Pods with an ordinal that is greater than or equal to the partition will be updated when the StatefulSet's .spec.template is updated. All Pods with an ordinal that is less than the partition will not be updated, and, even if they are deleted, they will be recreated at the previous version.
	//
	// If the partition is greater than replicas, updates to the pod template will not be propagated to Pods.
	// In most cases you will not need to use a partition, but they are useful if you want to stage an
	// update, roll out a canary, or perform a phased roll out.
	// See: https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/#partitions
	//
	Partition *float64 `field:"optional" json:"partition" yaml:"partition"`
}

// Subject contains a reference to the object or user identities a role binding applies to.
//
// This can either hold a direct API object reference, or a value
// for non-objects such as user and group names.
type SubjectConfiguration struct {
	// Kind of object being referenced.
	//
	// Values defined by this API group are
	// "User", "Group", and "ServiceAccount". If the Authorizer does not
	// recognized the kind value, the Authorizer should report an error.
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// Name of the object being referenced.
	Name *string `field:"required" json:"name" yaml:"name"`
	// APIGroup holds the API group of the referenced subject.
	//
	// Defaults to "" for
	// ServiceAccount subjects. Defaults to "rbac.authorization.k8s.io" for User
	// and Group subjects.
	ApiGroup *string `field:"optional" json:"apiGroup" yaml:"apiGroup"`
	// Namespace of the referenced object.
	//
	// If the object kind is non-namespace,
	// such as "User" or "Group", and this value is not empty the Authorizer
	// should report an error.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

// Sysctl defines a kernel parameter to be set.
type Sysctl struct {
	// Name of a property to set.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Value of a property to set.
	Value *string `field:"required" json:"value" yaml:"value"`
}

// Taint effects.
type TaintEffect string

const (
	// This means that no pod will be able to schedule onto the node unless it has a matching toleration.
	TaintEffect_NO_SCHEDULE TaintEffect = "NO_SCHEDULE"
	// This is a "preference" or "soft" version of `NO_SCHEDULE` -- the system will try to avoid placing a pod that does not tolerate the taint on the node, but it is not required.
	TaintEffect_PREFER_NO_SCHEDULE TaintEffect = "PREFER_NO_SCHEDULE"
	// This affects pods that are already running on the node as follows:.
	//
	// - Pods that do not tolerate the taint are evicted immediately.
	// - Pods that tolerate the taint without specifying `duration` remain bound forever.
	// - Pods that tolerate the taint with a specified `duration` remain bound for
	//    the specified amount of time.
	TaintEffect_NO_EXECUTE TaintEffect = "NO_EXECUTE"
)

// A node that is matched by taint selectors.
type TaintedNode interface {
	TaintSelector() *[]NodeTaintQuery
}

// The jsii proxy struct for TaintedNode
type jsiiProxy_TaintedNode struct {
	_ byte // padding
}

func (j *jsiiProxy_TaintedNode) TaintSelector() *[]NodeTaintQuery {
	var returns *[]NodeTaintQuery
	_jsii_.Get(
		j,
		"taintSelector",
		&returns,
	)
	return returns
}


func NewTaintedNode(taintSelector *[]NodeTaintQuery) TaintedNode {
	_init_.Initialize()

	j := jsiiProxy_TaintedNode{}

	_jsii_.Create(
		"cdk8s-plus-24.TaintedNode",
		[]interface{}{taintSelector},
		&j,
	)

	return &j
}

func NewTaintedNode_Override(t TaintedNode, taintSelector *[]NodeTaintQuery) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.TaintedNode",
		[]interface{}{taintSelector},
		t,
	)
}

// Options for `Probe.fromTcpSocket()`.
type TcpSocketProbeOptions struct {
	// Minimum consecutive failures for the probe to be considered failed after having succeeded.
	//
	// Defaults to 3. Minimum value is 1.
	FailureThreshold *float64 `field:"optional" json:"failureThreshold" yaml:"failureThreshold"`
	// Number of seconds after the container has started before liveness probes are initiated.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	//
	InitialDelaySeconds cdk8s.Duration `field:"optional" json:"initialDelaySeconds" yaml:"initialDelaySeconds"`
	// How often (in seconds) to perform the probe.
	//
	// Default to 10 seconds. Minimum value is 1.
	PeriodSeconds cdk8s.Duration `field:"optional" json:"periodSeconds" yaml:"periodSeconds"`
	// Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1.
	//
	// Must be 1 for liveness and startup. Minimum value is 1.
	SuccessThreshold *float64 `field:"optional" json:"successThreshold" yaml:"successThreshold"`
	// Number of seconds after which the probe times out.
	//
	// Defaults to 1 second. Minimum value is 1.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	//
	TimeoutSeconds cdk8s.Duration `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
	// The host name to connect to on the container.
	Host *string `field:"optional" json:"host" yaml:"host"`
	// The TCP port to connect to on the container.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

// Create a secret for storing a TLS certificate and its associated key.
// See: https://kubernetes.io/docs/concepts/configuration/secret/#tls-secrets
//
type TlsSecret interface {
	Secret
	// The group portion of the API version (e.g. "authorization.k8s.io").
	ApiGroup() *string
	// The underlying cdk8s API object.
	// See: base.Resource.apiObject
	//
	ApiObject() cdk8s.ApiObject
	// The object's API version (e.g. "authorization.k8s.io/v1").
	ApiVersion() *string
	// Whether or not the secret is immutable.
	Immutable() *bool
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
	// Adds a string data field to the secert.
	AddStringData(key *string, value *string)
	// Return the IApiResource this object represents.
	AsApiResource() IApiResource
	// Return the non resource url this object represents.
	AsNonApiResource() *string
	// Gets a string data by key or undefined.
	GetStringData(key *string) *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for TlsSecret
type jsiiProxy_TlsSecret struct {
	jsiiProxy_Secret
}

func (j *jsiiProxy_TlsSecret) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TlsSecret) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TlsSecret) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TlsSecret) Immutable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"immutable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TlsSecret) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TlsSecret) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TlsSecret) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TlsSecret) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TlsSecret) Permissions() ResourcePermissions {
	var returns ResourcePermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TlsSecret) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TlsSecret) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}


func NewTlsSecret(scope constructs.Construct, id *string, props *TlsSecretProps) TlsSecret {
	_init_.Initialize()

	j := jsiiProxy_TlsSecret{}

	_jsii_.Create(
		"cdk8s-plus-24.TlsSecret",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewTlsSecret_Override(t TlsSecret, scope constructs.Construct, id *string, props *TlsSecretProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.TlsSecret",
		[]interface{}{scope, id, props},
		t,
	)
}

// Imports a secret from the cluster as a reference.
func TlsSecret_FromSecretName(scope constructs.Construct, id *string, name *string) ISecret {
	_init_.Initialize()

	var returns ISecret

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.TlsSecret",
		"fromSecretName",
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
func TlsSecret_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.TlsSecret",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TlsSecret) AddStringData(key *string, value *string) {
	_jsii_.InvokeVoid(
		t,
		"addStringData",
		[]interface{}{key, value},
	)
}

func (t *jsiiProxy_TlsSecret) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		t,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TlsSecret) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TlsSecret) GetStringData(key *string) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringData",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TlsSecret) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options for `TlsSecret`.
type TlsSecretProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// If set to true, ensures that data stored in the Secret cannot be updated (only object metadata can be modified).
	//
	// If not set to true, the field can be modified at any time.
	Immutable *bool `field:"optional" json:"immutable" yaml:"immutable"`
	// The TLS cert.
	TlsCert *string `field:"required" json:"tlsCert" yaml:"tlsCert"`
	// The TLS key.
	TlsKey *string `field:"required" json:"tlsKey" yaml:"tlsKey"`
}

// Available topology domains.
type Topology interface {
	Key() *string
}

// The jsii proxy struct for Topology
type jsiiProxy_Topology struct {
	_ byte // padding
}

func (j *jsiiProxy_Topology) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}


// Custom key for the node label that the system uses to denote the topology domain.
func Topology_Custom(key *string) Topology {
	_init_.Initialize()

	var returns Topology

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Topology",
		"custom",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func Topology_HOSTNAME() Topology {
	_init_.Initialize()
	var returns Topology
	_jsii_.StaticGet(
		"cdk8s-plus-24.Topology",
		"HOSTNAME",
		&returns,
	)
	return returns
}

func Topology_REGION() Topology {
	_init_.Initialize()
	var returns Topology
	_jsii_.StaticGet(
		"cdk8s-plus-24.Topology",
		"REGION",
		&returns,
	)
	return returns
}

func Topology_ZONE() Topology {
	_init_.Initialize()
	var returns Topology
	_jsii_.StaticGet(
		"cdk8s-plus-24.Topology",
		"ZONE",
		&returns,
	)
	return returns
}

// Represents a user.
type User interface {
	constructs.Construct
	ISubject
	ApiGroup() *string
	Kind() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
	// Return the subject configuration.
	// See: ISubect.toSubjectConfiguration()
	//
	ToSubjectConfiguration() *SubjectConfiguration
}

// The jsii proxy struct for User
type jsiiProxy_User struct {
	internal.Type__constructsConstruct
	jsiiProxy_ISubject
}

func (j *jsiiProxy_User) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Reference a user in the cluster by name.
func User_FromName(scope constructs.Construct, id *string, name *string) User {
	_init_.Initialize()

	var returns User

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.User",
		"fromName",
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
func User_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.User",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) ToSubjectConfiguration() *SubjectConfiguration {
	var returns *SubjectConfiguration

	_jsii_.Invoke(
		u,
		"toSubjectConfiguration",
		nil, // no parameters
		&returns,
	)

	return returns
}

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
		"cdk8s-plus-24.Volume",
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
		"cdk8s-plus-24.Volume",
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
		"cdk8s-plus-24.Volume",
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
		"cdk8s-plus-24.Volume",
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
		"cdk8s-plus-24.Volume",
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
		"cdk8s-plus-24.Volume",
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
		"cdk8s-plus-24.Volume",
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
		"cdk8s-plus-24.Volume",
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
		"cdk8s-plus-24.Volume",
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

// Mount a volume from the pod to the container.
type VolumeMount struct {
	// Determines how mounts are propagated from the host to container and the other way around.
	//
	// When not set, MountPropagationNone is used.
	//
	// Mount propagation allows for sharing volumes mounted by a Container to
	// other Containers in the same Pod, or even to other Pods on the same node.
	Propagation MountPropagation `field:"optional" json:"propagation" yaml:"propagation"`
	// Mounted read-only if true, read-write otherwise (false or unspecified).
	//
	// Defaults to false.
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Path within the volume from which the container's volume should be mounted.).
	SubPath *string `field:"optional" json:"subPath" yaml:"subPath"`
	// Expanded path within the volume from which the container's volume should be mounted.
	//
	// Behaves similarly to SubPath but environment variable references
	// $(VAR_NAME) are expanded using the container's environment. Defaults to ""
	// (volume's root).
	//
	// `subPathExpr` and `subPath` are mutually exclusive.
	SubPathExpr *string `field:"optional" json:"subPathExpr" yaml:"subPathExpr"`
	// Path within the container at which the volume should be mounted.
	//
	// Must not
	// contain ':'.
	Path *string `field:"required" json:"path" yaml:"path"`
	// The volume to mount.
	Volume Volume `field:"required" json:"volume" yaml:"volume"`
}

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
	DockerRegistryAuth() DockerConfigSecret
	HostAliases() *[]*HostAlias
	InitContainers() *[]Container
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

func (j *jsiiProxy_Workload) DockerRegistryAuth() DockerConfigSecret {
	var returns DockerConfigSecret
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

func (j *jsiiProxy_Workload) InitContainers() *[]Container {
	var returns *[]Container
	_jsii_.Get(
		j,
		"initContainers",
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
	_jsii_.InvokeVoid(
		w,
		"addHostAlias",
		[]interface{}{hostAlias},
	)
}

func (w *jsiiProxy_Workload) AddInitContainer(cont *ContainerProps) Container {
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

// Properties for `Workload`.
type WorkloadProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// Indicates whether a service account token should be automatically mounted.
	// See: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/#use-the-default-service-account-to-access-the-api-server
	//
	AutomountServiceAccountToken *bool `field:"optional" json:"automountServiceAccountToken" yaml:"automountServiceAccountToken"`
	// List of containers belonging to the pod.
	//
	// Containers cannot currently be
	// added or removed. There must be at least one container in a Pod.
	//
	// You can add additionnal containers using `podSpec.addContainer()`
	Containers *[]*ContainerProps `field:"optional" json:"containers" yaml:"containers"`
	// DNS settings for the pod.
	// See: https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/
	//
	Dns *PodDnsProps `field:"optional" json:"dns" yaml:"dns"`
	// A secret containing docker credentials for authenticating to a registry.
	DockerRegistryAuth DockerConfigSecret `field:"optional" json:"dockerRegistryAuth" yaml:"dockerRegistryAuth"`
	// HostAlias holds the mapping between IP and hostnames that will be injected as an entry in the pod's hosts file.
	HostAliases *[]*HostAlias `field:"optional" json:"hostAliases" yaml:"hostAliases"`
	// List of initialization containers belonging to the pod.
	//
	// Init containers are executed in order prior to containers being started.
	// If any init container fails, the pod is considered to have failed and is handled according to its restartPolicy.
	// The name for an init container or normal container must be unique among all containers.
	// Init containers may not have Lifecycle actions, Readiness probes, Liveness probes, or Startup probes.
	// The resourceRequirements of an init container are taken into account during scheduling by finding the highest request/limit
	// for each resource type, and then using the max of of that value or the sum of the normal containers.
	// Limits are applied to init containers in a similar fashion.
	//
	// Init containers cannot currently be added ,removed or updated.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/
	//
	InitContainers *[]*ContainerProps `field:"optional" json:"initContainers" yaml:"initContainers"`
	// Restart policy for all containers within the pod.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy
	//
	RestartPolicy RestartPolicy `field:"optional" json:"restartPolicy" yaml:"restartPolicy"`
	// SecurityContext holds pod-level security attributes and common container settings.
	SecurityContext *PodSecurityContextProps `field:"optional" json:"securityContext" yaml:"securityContext"`
	// A service account provides an identity for processes that run in a Pod.
	//
	// When you (a human) access the cluster (for example, using kubectl), you are
	// authenticated by the apiserver as a particular User Account (currently this
	// is usually admin, unless your cluster administrator has customized your
	// cluster). Processes in containers inside pods can also contact the
	// apiserver. When they do, they are authenticated as a particular Service
	// Account (for example, default).
	// See: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
	//
	ServiceAccount IServiceAccount `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// List of volumes that can be mounted by containers belonging to the pod.
	//
	// You can also add volumes later using `podSpec.addVolume()`
	// See: https://kubernetes.io/docs/concepts/storage/volumes
	//
	Volumes *[]Volume `field:"optional" json:"volumes" yaml:"volumes"`
	// The pod metadata of this workload.
	PodMetadata *cdk8s.ApiObjectMetadata `field:"optional" json:"podMetadata" yaml:"podMetadata"`
	// Automatically allocates a pod label selector for this workload and add it to the pod metadata.
	//
	// This ensures this workload manages pods created by
	// its pod template.
	Select *bool `field:"optional" json:"select" yaml:"select"`
}

// Controls the pod scheduling strategy of this workload.
//
// It offers some additional API's on top of the core pod scheduling.
type WorkloadScheduling interface {
	PodScheduling
	Instance() AbstractPod
	// Assign this pod a specific node by name.
	//
	// The scheduler ignores the Pod, and the kubelet on the named node
	// tries to place the Pod on that node. Overrules any affinity rules of the pod.
	//
	// Some limitations of static assignment are:
	//
	// - If the named node does not exist, the Pod will not run, and in some
	//    cases may be automatically deleted.
	// - If the named node does not have the resources to accommodate the Pod,
	//    the Pod will fail and its reason will indicate why, for example OutOfmemory or OutOfcpu.
	// - Node names in cloud environments are not always predictable or stable.
	//
	// Will throw is the pod is already assigned to named node.
	//
	// Under the hood, this method utilizes the `nodeName` property.
	Assign(node NamedNode)
	// Attract this pod to a node matched by selectors. You can select a node by using `Node.labeled()`.
	//
	// Attracting to multiple nodes (i.e invoking this method multiple times) acts as
	// an OR condition, meaning the pod will be assigned to either one of the nodes.
	//
	// Under the hood, this method utilizes the `nodeAffinity` property.
	// See: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#node-affinity
	//
	Attract(node LabeledNode, options *PodSchedulingAttractOptions)
	// Co-locate this pod with a scheduling selection.
	//
	// A selection can be one of:
	//
	// - An instance of a `Pod`.
	// - An instance of a `Workload` (e.g `Deployment`, `StatefulSet`).
	// - An un-managed pod that can be selected via `Pods.select()`.
	//
	// Co-locating with multiple selections ((i.e invoking this method multiple times)) acts as
	// an AND condition. meaning the pod will be assigned to a node that satisfies all
	// selections (i.e runs at least one pod that satisifies each selection).
	//
	// Under the hood, this method utilizes the `podAffinity` property.
	// See: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#inter-pod-affinity-and-anti-affinity
	//
	Colocate(selector IPodSelector, options *PodSchedulingColocateOptions)
	// Seperate this pod from a scheduling selection.
	//
	// A selection can be one of:
	//
	// - An instance of a `Pod`.
	// - An instance of a `Workload` (e.g `Deployment`, `StatefulSet`).
	// - An un-managed pod that can be selected via `Pods.select()`.
	//
	// Seperating from multiple selections acts as an AND condition. meaning the pod
	// will not be assigned to a node that satisfies all selections (i.e runs at least one pod that satisifies each selection).
	//
	// Under the hood, this method utilizes the `podAntiAffinity` property.
	// See: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#inter-pod-affinity-and-anti-affinity
	//
	Separate(selector IPodSelector, options *PodSchedulingSeparateOptions)
	// Spread the pods in this workload by the topology key.
	//
	// A spread is a separation of the pod from itself and is used to
	// balance out pod replicas across a given topology.
	Spread(options *WorkloadSchedulingSpreadOptions)
	// Allow this pod to tolerate taints matching these tolerations.
	//
	// You can put multiple taints on the same node and multiple tolerations on the same pod.
	// The way Kubernetes processes multiple taints and tolerations is like a filter: start with
	// all of a node's taints, then ignore the ones for which the pod has a matching toleration;
	// the remaining un-ignored taints have the indicated effects on the pod. In particular:
	//
	// - if there is at least one un-ignored taint with effect NoSchedule then Kubernetes will
	//    not schedule the pod onto that node
	// - if there is no un-ignored taint with effect NoSchedule but there is at least one un-ignored
	//    taint with effect PreferNoSchedule then Kubernetes will try to not schedule the pod onto the node
	// - if there is at least one un-ignored taint with effect NoExecute then the pod will be evicted from
	//    the node (if it is already running on the node), and will not be scheduled onto the node (if it is
	//    not yet running on the node).
	//
	// Under the hood, this method utilizes the `tolerations` property.
	// See: https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/
	//
	Tolerate(node TaintedNode)
}

// The jsii proxy struct for WorkloadScheduling
type jsiiProxy_WorkloadScheduling struct {
	jsiiProxy_PodScheduling
}

func (j *jsiiProxy_WorkloadScheduling) Instance() AbstractPod {
	var returns AbstractPod
	_jsii_.Get(
		j,
		"instance",
		&returns,
	)
	return returns
}


func NewWorkloadScheduling(instance AbstractPod) WorkloadScheduling {
	_init_.Initialize()

	j := jsiiProxy_WorkloadScheduling{}

	_jsii_.Create(
		"cdk8s-plus-24.WorkloadScheduling",
		[]interface{}{instance},
		&j,
	)

	return &j
}

func NewWorkloadScheduling_Override(w WorkloadScheduling, instance AbstractPod) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.WorkloadScheduling",
		[]interface{}{instance},
		w,
	)
}

func (w *jsiiProxy_WorkloadScheduling) Assign(node NamedNode) {
	_jsii_.InvokeVoid(
		w,
		"assign",
		[]interface{}{node},
	)
}

func (w *jsiiProxy_WorkloadScheduling) Attract(node LabeledNode, options *PodSchedulingAttractOptions) {
	_jsii_.InvokeVoid(
		w,
		"attract",
		[]interface{}{node, options},
	)
}

func (w *jsiiProxy_WorkloadScheduling) Colocate(selector IPodSelector, options *PodSchedulingColocateOptions) {
	_jsii_.InvokeVoid(
		w,
		"colocate",
		[]interface{}{selector, options},
	)
}

func (w *jsiiProxy_WorkloadScheduling) Separate(selector IPodSelector, options *PodSchedulingSeparateOptions) {
	_jsii_.InvokeVoid(
		w,
		"separate",
		[]interface{}{selector, options},
	)
}

func (w *jsiiProxy_WorkloadScheduling) Spread(options *WorkloadSchedulingSpreadOptions) {
	_jsii_.InvokeVoid(
		w,
		"spread",
		[]interface{}{options},
	)
}

func (w *jsiiProxy_WorkloadScheduling) Tolerate(node TaintedNode) {
	_jsii_.InvokeVoid(
		w,
		"tolerate",
		[]interface{}{node},
	)
}

// Options for `WorkloadScheduling.spread`.
type WorkloadSchedulingSpreadOptions struct {
	// Which topology to spread on.
	Topology Topology `field:"optional" json:"topology" yaml:"topology"`
	// Indicates the spread is optional, with this weight score.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

