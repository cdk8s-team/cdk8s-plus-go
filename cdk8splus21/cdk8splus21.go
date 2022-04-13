// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-21 synthesizes Kubernetes manifests for Kubernetes 1.21.0
package cdk8splus21

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus21/jsii"

	"github.com/aws/constructs-go/constructs/v3"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus21/internal"
)

// Options to add a deployment to a service.
type AddDeploymentOptions struct {
	// The name of this port within the service.
	//
	// This must be a DNS_LABEL. All
	// ports within a ServiceSpec must have unique names. This maps to the 'Name'
	// field in EndpointPort objects. Optional if only one ServicePort is defined
	// on this service.
	Name *string `json:"name" yaml:"name"`
	// The port on each node on which this service is exposed when type=NodePort or LoadBalancer.
	//
	// Usually assigned by the system. If specified, it will be
	// allocated to the service if unused or else creation of the service will
	// fail. Default is to auto-allocate a port if the ServiceType of this Service
	// requires one.
	// See: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport
	//
	NodePort *float64 `json:"nodePort" yaml:"nodePort"`
	// The IP protocol for this port.
	//
	// Supports "TCP", "UDP", and "SCTP". Default is TCP.
	Protocol Protocol `json:"protocol" yaml:"protocol"`
	// The port number the service will redirect to.
	TargetPort *float64 `json:"targetPort" yaml:"targetPort"`
	// The port number the service will bind to.
	Port *float64 `json:"port" yaml:"port"`
}

// Options for `configmap.addDirectory()`.
type AddDirectoryOptions struct {
	// Glob patterns to exclude when adding files.
	Exclude *[]*string `json:"exclude" yaml:"exclude"`
	// A prefix to add to all keys in the config map.
	KeyPrefix *string `json:"keyPrefix" yaml:"keyPrefix"`
}

// Options for `Probe.fromCommand()`.
type CommandProbeOptions struct {
	// Minimum consecutive failures for the probe to be considered failed after having succeeded.
	//
	// Defaults to 3. Minimum value is 1.
	FailureThreshold *float64 `json:"failureThreshold" yaml:"failureThreshold"`
	// Number of seconds after the container has started before liveness probes are initiated.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	//
	InitialDelaySeconds cdk8s.Duration `json:"initialDelaySeconds" yaml:"initialDelaySeconds"`
	// How often (in seconds) to perform the probe.
	//
	// Default to 10 seconds. Minimum value is 1.
	PeriodSeconds cdk8s.Duration `json:"periodSeconds" yaml:"periodSeconds"`
	// Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1.
	//
	// Must be 1 for liveness and startup. Minimum value is 1.
	SuccessThreshold *float64 `json:"successThreshold" yaml:"successThreshold"`
	// Number of seconds after which the probe times out.
	//
	// Defaults to 1 second. Minimum value is 1.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	//
	TimeoutSeconds cdk8s.Duration `json:"timeoutSeconds" yaml:"timeoutSeconds"`
}

// ConfigMap holds configuration data for pods to consume.
type ConfigMap interface {
	Resource
	IConfigMap
	// The underlying cdk8s API object.
	// See: base.Resource.apiObject
	//
	ApiObject() cdk8s.ApiObject
	// The binary data associated with this config map.
	//
	// Returns a copy. To add data records, use `addBinaryData()` or `addData()`.
	BinaryData() *map[string]*string
	// The data associated with this config map.
	//
	// Returns an copy. To add data records, use `addData()` or `addBinaryData()`.
	Data() *map[string]*string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	// The name of this API object.
	Name() *string
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
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if there the construct is valid.
	// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
	// instead of overriding this method.
	OnValidate() *[]*string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for ConfigMap
type jsiiProxy_ConfigMap struct {
	jsiiProxy_Resource
	jsiiProxy_IConfigMap
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


func NewConfigMap(scope constructs.Construct, id *string, props *ConfigMapProps) ConfigMap {
	_init_.Initialize()

	j := jsiiProxy_ConfigMap{}

	_jsii_.Create(
		"cdk8s-plus-21.ConfigMap",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewConfigMap_Override(c ConfigMap, scope constructs.Construct, id *string, props *ConfigMapProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-21.ConfigMap",
		[]interface{}{scope, id, props},
		c,
	)
}

// Represents a ConfigMap created elsewhere.
func ConfigMap_FromConfigMapName(name *string) IConfigMap {
	_init_.Initialize()

	var returns IConfigMap

	_jsii_.StaticInvoke(
		"cdk8s-plus-21.ConfigMap",
		"fromConfigMapName",
		[]interface{}{name},
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

func (c *jsiiProxy_ConfigMap) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigMap) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_ConfigMap) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
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
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata" yaml:"metadata"`
	// BinaryData contains the binary data.
	//
	// Each key must consist of alphanumeric characters, '-', '_' or '.'.
	// BinaryData can contain byte sequences that are not in the UTF-8 range. The
	// keys stored in BinaryData must not overlap with the ones in the Data field,
	// this is enforced during validation process.
	//
	// You can also add binary data using `configMap.addBinaryData()`.
	BinaryData *map[string]*string `json:"binaryData" yaml:"binaryData"`
	// Data contains the configuration data.
	//
	// Each key must consist of alphanumeric characters, '-', '_' or '.'. Values
	// with non-UTF-8 byte sequences must use the BinaryData field. The keys
	// stored in Data must not overlap with the keys in the BinaryData field, this
	// is enforced during validation process.
	//
	// You can also add data using `configMap.addData()`.
	Data *map[string]*string `json:"data" yaml:"data"`
}

// Options for the ConfigMap-based volume.
type ConfigMapVolumeOptions struct {
	// Mode bits to use on created files by default.
	//
	// Must be a value between 0 and
	// 0777. Defaults to 0644. Directories within the path are not affected by
	// this setting. This might be in conflict with other options that affect the
	// file mode, like fsGroup, and the result can be other mode bits set.
	DefaultMode *float64 `json:"defaultMode" yaml:"defaultMode"`
	// If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value.
	//
	// If specified, the listed keys will be projected
	// into the specified paths, and unlisted keys will not be present. If a key
	// is specified which is not present in the ConfigMap, the volume setup will
	// error unless it is marked optional. Paths must be relative and may not
	// contain the '..' path or start with '..'.
	Items *map[string]*PathMapping `json:"items" yaml:"items"`
	// The volume name.
	Name *string `json:"name" yaml:"name"`
	// Specify whether the ConfigMap or its keys must be defined.
	Optional *bool `json:"optional" yaml:"optional"`
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
	// The environment variables for this container.
	//
	// Returns a copy. To add environment variables use `addEnv()`.
	Env() *map[string]EnvValue
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
	Resources() *Resources
	// The security context of the container.
	SecurityContext() ContainerSecurityContext
	// The working directory inside the container.
	WorkingDir() *string
	// Add an environment value to the container.
	//
	// The variable value can come
	// from various dynamic sources such a secrets of config maps.
	// See: EnvValue.fromXXX
	//
	AddEnv(name *string, value EnvValue)
	// Mount a volume to a specific path so that it is accessible by the container.
	//
	// Every pod that is configured to use this container will autmoatically have access to the volume.
	Mount(path *string, volume Volume, options *MountOptions)
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

func (j *jsiiProxy_Container) Env() *map[string]EnvValue {
	var returns *map[string]EnvValue
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

func (j *jsiiProxy_Container) Resources() *Resources {
	var returns *Resources
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
		"cdk8s-plus-21.Container",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewContainer_Override(c Container, props *ContainerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-21.Container",
		[]interface{}{props},
		c,
	)
}

func (c *jsiiProxy_Container) AddEnv(name *string, value EnvValue) {
	_jsii_.InvokeVoid(
		c,
		"addEnv",
		[]interface{}{name, value},
	)
}

func (c *jsiiProxy_Container) Mount(path *string, volume Volume, options *MountOptions) {
	_jsii_.InvokeVoid(
		c,
		"mount",
		[]interface{}{path, volume, options},
	)
}

// Properties for creating a container.
type ContainerProps struct {
	// Docker image name.
	Image *string `json:"image" yaml:"image"`
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
	Args *[]*string `json:"args" yaml:"args"`
	// Entrypoint array.
	//
	// Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment.
	// If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME).
	// Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated.
	// More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	Command *[]*string `json:"command" yaml:"command"`
	// List of environment variables to set in the container.
	//
	// Cannot be updated.
	Env *map[string]EnvValue `json:"env" yaml:"env"`
	// Image pull policy for this container.
	ImagePullPolicy ImagePullPolicy `json:"imagePullPolicy" yaml:"imagePullPolicy"`
	// Periodic probe of container liveness.
	//
	// Container will be restarted if the probe fails.
	Liveness Probe `json:"liveness" yaml:"liveness"`
	// Name of the container specified as a DNS_LABEL.
	//
	// Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.
	Name *string `json:"name" yaml:"name"`
	// Number of port to expose on the pod's IP address.
	//
	// This must be a valid port number, 0 < x < 65536.
	Port *float64 `json:"port" yaml:"port"`
	// Determines when the container is ready to serve traffic.
	Readiness Probe `json:"readiness" yaml:"readiness"`
	// Compute resources (CPU and memory requests and limits) required by the container.
	// See: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	//
	Resources *Resources `json:"resources" yaml:"resources"`
	// SecurityContext defines the security options the container should be run with.
	//
	// If set, the fields override equivalent fields of the pod's security context.
	// See: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
	//
	SecurityContext *ContainerSecurityContextProps `json:"securityContext" yaml:"securityContext"`
	// StartupProbe indicates that the Pod has successfully initialized.
	//
	// If specified, no other probes are executed until this completes successfully.
	Startup Probe `json:"startup" yaml:"startup"`
	// Pod volumes to mount into the container's filesystem.
	//
	// Cannot be updated.
	VolumeMounts *[]*VolumeMount `json:"volumeMounts" yaml:"volumeMounts"`
	// Container's working directory.
	//
	// If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.
	WorkingDir *string `json:"workingDir" yaml:"workingDir"`
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
		"cdk8s-plus-21.ContainerSecurityContext",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewContainerSecurityContext_Override(c ContainerSecurityContext, props *ContainerSecurityContextProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-21.ContainerSecurityContext",
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
	EnsureNonRoot *bool `json:"ensureNonRoot" yaml:"ensureNonRoot"`
	// The GID to run the entrypoint of the container process.
	Group *float64 `json:"group" yaml:"group"`
	// Run container in privileged mode.
	//
	// Processes in privileged containers are essentially equivalent to root on the host.
	Privileged *bool `json:"privileged" yaml:"privileged"`
	// Whether this container has a read-only root filesystem.
	ReadOnlyRootFilesystem *bool `json:"readOnlyRootFilesystem" yaml:"readOnlyRootFilesystem"`
	// The UID to run the entrypoint of the container process.
	User *float64 `json:"user" yaml:"user"`
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
		"cdk8s-plus-21.Cpu",
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
		"cdk8s-plus-21.Cpu",
		"units",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

// CPU request and limit.
type CpuResources struct {
	Limit Cpu `json:"limit" yaml:"limit"`
	Request Cpu `json:"request" yaml:"request"`
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
	Resource
	IPodTemplate
	// The underlying cdk8s API object.
	// See: base.Resource.apiObject
	//
	ApiObject() cdk8s.ApiObject
	// The containers belonging to the pod.
	//
	// Use `addContainer` to add containers.
	Containers() *[]Container
	// An optional list of hosts and IPs that will be injected into the pod's hosts file if specified.
	//
	// This is only valid for non-hostNetwork pods.
	HostAliases() *[]*HostAlias
	// The init containers belonging to the pod.
	//
	// Use `addInitContainer` to add init containers.
	InitContainers() *[]Container
	// The labels this deployment will match against in order to select pods.
	//
	// Returns a a copy. Use `selectByLabel()` to add labels.
	LabelSelector() *map[string]*string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	// The name of this API object.
	Name() *string
	// Provides read/write access to the underlying pod metadata of the resource.
	PodMetadata() cdk8s.ApiObjectMetadataDefinition
	// Number of desired pods.
	Replicas() *float64
	// Restart policy for all containers within the pod.
	RestartPolicy() RestartPolicy
	SecurityContext() PodSecurityContext
	// The service account used to run this pod.
	ServiceAccount() IServiceAccount
	// The volumes associated with this pod.
	//
	// Use `addVolume` to add volumes.
	Volumes() *[]Volume
	// Add a container to the pod.
	AddContainer(container *ContainerProps) Container
	AddHostAlias(hostAlias *HostAlias)
	// Add an init container to the pod.
	AddInitContainer(container *ContainerProps) Container
	// Add a volume to the pod.
	AddVolume(volume Volume)
	// Expose a deployment via an ingress.
	//
	// This will first expose the deployment with a service, and then expose the service via an ingress.
	ExposeViaIngress(path *string, options *ExposeDeploymentViaIngressOptions) IngressV1Beta1
	// Expose a deployment via a service.
	//
	// This is equivalent to running `kubectl expose deployment <deployment-name>`.
	ExposeViaService(options *ExposeDeploymentViaServiceOptions) Service
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if there the construct is valid.
	// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
	// instead of overriding this method.
	OnValidate() *[]*string
	// Configure a label selector to this deployment.
	//
	// Pods that have the label will be selected by deployments configured with this spec.
	SelectByLabel(key *string, value *string)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Deployment
type jsiiProxy_Deployment struct {
	jsiiProxy_Resource
	jsiiProxy_IPodTemplate
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

func (j *jsiiProxy_Deployment) Containers() *[]Container {
	var returns *[]Container
	_jsii_.Get(
		j,
		"containers",
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

func (j *jsiiProxy_Deployment) LabelSelector() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelSelector",
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

func (j *jsiiProxy_Deployment) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
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

func (j *jsiiProxy_Deployment) Replicas() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicas",
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
		"cdk8s-plus-21.Deployment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDeployment_Override(d Deployment, scope constructs.Construct, id *string, props *DeploymentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-21.Deployment",
		[]interface{}{scope, id, props},
		d,
	)
}

func (d *jsiiProxy_Deployment) AddContainer(container *ContainerProps) Container {
	var returns Container

	_jsii_.Invoke(
		d,
		"addContainer",
		[]interface{}{container},
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

func (d *jsiiProxy_Deployment) AddInitContainer(container *ContainerProps) Container {
	var returns Container

	_jsii_.Invoke(
		d,
		"addInitContainer",
		[]interface{}{container},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Deployment) AddVolume(volume Volume) {
	_jsii_.InvokeVoid(
		d,
		"addVolume",
		[]interface{}{volume},
	)
}

func (d *jsiiProxy_Deployment) ExposeViaIngress(path *string, options *ExposeDeploymentViaIngressOptions) IngressV1Beta1 {
	var returns IngressV1Beta1

	_jsii_.Invoke(
		d,
		"exposeViaIngress",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Deployment) ExposeViaService(options *ExposeDeploymentViaServiceOptions) Service {
	var returns Service

	_jsii_.Invoke(
		d,
		"exposeViaService",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Deployment) OnPrepare() {
	_jsii_.InvokeVoid(
		d,
		"onPrepare",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Deployment) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (d *jsiiProxy_Deployment) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Deployment) SelectByLabel(key *string, value *string) {
	_jsii_.InvokeVoid(
		d,
		"selectByLabel",
		[]interface{}{key, value},
	)
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

// Properties for initialization of `Deployment`.
type DeploymentProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata" yaml:"metadata"`
	// List of containers belonging to the pod.
	//
	// Containers cannot currently be
	// added or removed. There must be at least one container in a Pod.
	//
	// You can add additionnal containers using `podSpec.addContainer()`
	Containers *[]*ContainerProps `json:"containers" yaml:"containers"`
	// HostAlias holds the mapping between IP and hostnames that will be injected as an entry in the pod's hosts file.
	HostAliases *[]*HostAlias `json:"hostAliases" yaml:"hostAliases"`
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
	InitContainers *[]*ContainerProps `json:"initContainers" yaml:"initContainers"`
	// Restart policy for all containers within the pod.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy
	//
	RestartPolicy RestartPolicy `json:"restartPolicy" yaml:"restartPolicy"`
	// SecurityContext holds pod-level security attributes and common container settings.
	SecurityContext *PodSecurityContextProps `json:"securityContext" yaml:"securityContext"`
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
	ServiceAccount IServiceAccount `json:"serviceAccount" yaml:"serviceAccount"`
	// List of volumes that can be mounted by containers belonging to the pod.
	//
	// You can also add volumes later using `podSpec.addVolume()`
	// See: https://kubernetes.io/docs/concepts/storage/volumes
	//
	Volumes *[]Volume `json:"volumes" yaml:"volumes"`
	// The pod metadata.
	PodMetadata *cdk8s.ApiObjectMetadata `json:"podMetadata" yaml:"podMetadata"`
	// Automatically allocates a pod selector for this deployment.
	//
	// If this is set to `false` you must define your selector through
	// `deployment.podMetadata.addLabel()` and `deployment.selectByLabel()`.
	DefaultSelector *bool `json:"defaultSelector" yaml:"defaultSelector"`
	// Number of desired pods.
	Replicas *float64 `json:"replicas" yaml:"replicas"`
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
	Medium EmptyDirMedium `json:"medium" yaml:"medium"`
	// Total amount of local storage required for this EmptyDir volume.
	//
	// The size
	// limit is also applicable for memory medium. The maximum usage on memory
	// medium EmptyDir would be the minimum value between the SizeLimit specified
	// here and the sum of memory limits of all containers in a pod.
	SizeLimit cdk8s.Size `json:"sizeLimit" yaml:"sizeLimit"`
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
		"cdk8s-plus-21.EnvValue",
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
		"cdk8s-plus-21.EnvValue",
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
		"cdk8s-plus-21.EnvValue",
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
		"cdk8s-plus-21.EnvValue",
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
		"cdk8s-plus-21.EnvValue",
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
		"cdk8s-plus-21.EnvValue",
		"fromValue",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Options to specify an envionment variable value from a ConfigMap key.
type EnvValueFromConfigMapOptions struct {
	// Specify whether the ConfigMap or its key must be defined.
	Optional *bool `json:"optional" yaml:"optional"`
}

// Options to specify an environment variable value from a field reference.
type EnvValueFromFieldRefOptions struct {
	// Version of the schema the FieldPath is written in terms of.
	ApiVersion *string `json:"apiVersion" yaml:"apiVersion"`
	// The key to select the pod label or annotation.
	Key *string `json:"key" yaml:"key"`
}

// Options to specify an environment variable value from the process environment.
type EnvValueFromProcessOptions struct {
	// Specify whether the key must exist in the environment.
	//
	// If this is set to true, and the key does not exist, an error will thrown.
	Required *bool `json:"required" yaml:"required"`
}

// Options to specify an environment variable value from a resource.
type EnvValueFromResourceOptions struct {
	// The container to select the value from.
	Container Container `json:"container" yaml:"container"`
	// The output format of the exposed resource.
	Divisor *string `json:"divisor" yaml:"divisor"`
}

// Options to specify an environment variable value from a Secret.
type EnvValueFromSecretOptions struct {
	// Specify whether the Secret or its key must be defined.
	Optional *bool `json:"optional" yaml:"optional"`
}

// Options for exposing a deployment via an ingress.
type ExposeDeploymentViaIngressOptions struct {
	// The name of the service to expose.
	//
	// This will be set on the Service.metadata and must be a DNS_LABEL
	Name *string `json:"name" yaml:"name"`
	// The port that the service should serve on.
	Port *float64 `json:"port" yaml:"port"`
	// The IP protocol for this port.
	//
	// Supports "TCP", "UDP", and "SCTP". Default is TCP.
	Protocol Protocol `json:"protocol" yaml:"protocol"`
	// The type of the exposed service.
	ServiceType ServiceType `json:"serviceType" yaml:"serviceType"`
	// The port number the service will redirect to.
	TargetPort *float64 `json:"targetPort" yaml:"targetPort"`
	// The ingress to add rules to.
	Ingress IngressV1Beta1 `json:"ingress" yaml:"ingress"`
}

// Options for exposing a deployment via a service.
type ExposeDeploymentViaServiceOptions struct {
	// The name of the service to expose.
	//
	// This will be set on the Service.metadata and must be a DNS_LABEL
	Name *string `json:"name" yaml:"name"`
	// The port that the service should serve on.
	Port *float64 `json:"port" yaml:"port"`
	// The IP protocol for this port.
	//
	// Supports "TCP", "UDP", and "SCTP". Default is TCP.
	Protocol Protocol `json:"protocol" yaml:"protocol"`
	// The type of the exposed service.
	ServiceType ServiceType `json:"serviceType" yaml:"serviceType"`
	// The port number the service will redirect to.
	TargetPort *float64 `json:"targetPort" yaml:"targetPort"`
}

// Options for exposing a service using an ingress.
type ExposeServiceViaIngressOptions struct {
	// The ingress to add rules to.
	Ingress IngressV1Beta1 `json:"ingress" yaml:"ingress"`
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

// HostAlias holds the mapping between IP and hostnames that will be injected as an entry in the pod's /etc/hosts file.
type HostAlias struct {
	// Hostnames for the chosen IP address.
	Hostnames *[]*string `json:"hostnames" yaml:"hostnames"`
	// IP address of the host file entry.
	Ip *string `json:"ip" yaml:"ip"`
}

// Options for `Probe.fromHttpGet()`.
type HttpGetProbeOptions struct {
	// Minimum consecutive failures for the probe to be considered failed after having succeeded.
	//
	// Defaults to 3. Minimum value is 1.
	FailureThreshold *float64 `json:"failureThreshold" yaml:"failureThreshold"`
	// Number of seconds after the container has started before liveness probes are initiated.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	//
	InitialDelaySeconds cdk8s.Duration `json:"initialDelaySeconds" yaml:"initialDelaySeconds"`
	// How often (in seconds) to perform the probe.
	//
	// Default to 10 seconds. Minimum value is 1.
	PeriodSeconds cdk8s.Duration `json:"periodSeconds" yaml:"periodSeconds"`
	// Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1.
	//
	// Must be 1 for liveness and startup. Minimum value is 1.
	SuccessThreshold *float64 `json:"successThreshold" yaml:"successThreshold"`
	// Number of seconds after which the probe times out.
	//
	// Defaults to 1 second. Minimum value is 1.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	//
	TimeoutSeconds cdk8s.Duration `json:"timeoutSeconds" yaml:"timeoutSeconds"`
	// The TCP port to use when sending the GET request.
	Port *float64 `json:"port" yaml:"port"`
}

// Represents a config map.
type IConfigMap interface {
	IResource
}

// The jsii proxy for IConfigMap
type jsiiProxy_IConfigMap struct {
	jsiiProxy_IResource
}

// Represents a resource that can be configured with a kuberenets pod spec. (e.g `Deployment`, `Job`, `Pod`, ...).
//
// Use the `PodSpec` class as an implementation helper.
type IPodSpec interface {
	// Add a container to the pod.
	AddContainer(container *ContainerProps) Container
	// Add an init container to the pod.
	AddInitContainer(container *ContainerProps) Container
	// Add a volume to the pod.
	AddVolume(volume Volume)
	// The containers belonging to the pod.
	//
	// Use `addContainer` to add containers.
	Containers() *[]Container
	// An optional list of hosts and IPs that will be injected into the pod's hosts file if specified.
	//
	// This is only valid for non-hostNetwork pods.
	HostAliases() *[]*HostAlias
	// The init containers belonging to the pod.
	//
	// Use `addInitContainer` to add init containers.
	InitContainers() *[]Container
	// Restart policy for all containers within the pod.
	RestartPolicy() RestartPolicy
	// The service account used to run this pod.
	ServiceAccount() IServiceAccount
	// The volumes associated with this pod.
	//
	// Use `addVolume` to add volumes.
	Volumes() *[]Volume
}

// The jsii proxy for IPodSpec
type jsiiProxy_IPodSpec struct {
	_ byte // padding
}

func (i *jsiiProxy_IPodSpec) AddContainer(container *ContainerProps) Container {
	var returns Container

	_jsii_.Invoke(
		i,
		"addContainer",
		[]interface{}{container},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IPodSpec) AddInitContainer(container *ContainerProps) Container {
	var returns Container

	_jsii_.Invoke(
		i,
		"addInitContainer",
		[]interface{}{container},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IPodSpec) AddVolume(volume Volume) {
	_jsii_.InvokeVoid(
		i,
		"addVolume",
		[]interface{}{volume},
	)
}

func (j *jsiiProxy_IPodSpec) Containers() *[]Container {
	var returns *[]Container
	_jsii_.Get(
		j,
		"containers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPodSpec) HostAliases() *[]*HostAlias {
	var returns *[]*HostAlias
	_jsii_.Get(
		j,
		"hostAliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPodSpec) InitContainers() *[]Container {
	var returns *[]Container
	_jsii_.Get(
		j,
		"initContainers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPodSpec) RestartPolicy() RestartPolicy {
	var returns RestartPolicy
	_jsii_.Get(
		j,
		"restartPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPodSpec) ServiceAccount() IServiceAccount {
	var returns IServiceAccount
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPodSpec) Volumes() *[]Volume {
	var returns *[]Volume
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}

// Represents a resource that can be configured with a kuberenets pod template. (e.g `Deployment`, `Job`, ...).
//
// Use the `PodTemplate` class as an implementation helper.
type IPodTemplate interface {
	IPodSpec
	// Provides read/write access to the underlying pod metadata of the resource.
	PodMetadata() cdk8s.ApiObjectMetadataDefinition
}

// The jsii proxy for IPodTemplate
type jsiiProxy_IPodTemplate struct {
	jsiiProxy_IPodSpec
}

func (j *jsiiProxy_IPodTemplate) PodMetadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"podMetadata",
		&returns,
	)
	return returns
}

// Represents a resource.
type IResource interface {
	// The Kubernetes name of this resource.
	Name() *string
}

// The jsii proxy for IResource
type jsiiProxy_IResource struct {
	_ byte // padding
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
type IngressV1Beta1 interface {
	Resource
	// The underlying cdk8s API object.
	// See: base.Resource.apiObject
	//
	ApiObject() cdk8s.ApiObject
	Metadata() cdk8s.ApiObjectMetadataDefinition
	// The name of this API object.
	Name() *string
	// Defines the default backend for this ingress.
	//
	// A default backend capable of
	// servicing requests that don't match any rule.
	AddDefaultBackend(backend IngressV1Beta1Backend)
	// Specify a default backend for a specific host name.
	//
	// This backend will be used as a catch-all for requests
	// targeted to this host name (the `Host` header matches this value).
	AddHostDefaultBackend(host *string, backend IngressV1Beta1Backend)
	// Adds an ingress rule applied to requests to a specific host and a specific HTTP path (the `Host` header matches this value).
	AddHostRule(host *string, path *string, backend IngressV1Beta1Backend)
	// Adds an ingress rule applied to requests sent to a specific HTTP path.
	AddRule(path *string, backend IngressV1Beta1Backend)
	// Adds rules to this ingress.
	AddRules(rules ...*IngressV1Beta1Rule)
	AddTls(tls *[]*IngressV1Beta1Tls)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	OnValidate() *[]*string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for IngressV1Beta1
type jsiiProxy_IngressV1Beta1 struct {
	jsiiProxy_Resource
}

func (j *jsiiProxy_IngressV1Beta1) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IngressV1Beta1) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IngressV1Beta1) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


func NewIngressV1Beta1(scope constructs.Construct, id *string, props *IngressV1Beta1Props) IngressV1Beta1 {
	_init_.Initialize()

	j := jsiiProxy_IngressV1Beta1{}

	_jsii_.Create(
		"cdk8s-plus-21.IngressV1Beta1",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewIngressV1Beta1_Override(i IngressV1Beta1, scope constructs.Construct, id *string, props *IngressV1Beta1Props) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-21.IngressV1Beta1",
		[]interface{}{scope, id, props},
		i,
	)
}

func (i *jsiiProxy_IngressV1Beta1) AddDefaultBackend(backend IngressV1Beta1Backend) {
	_jsii_.InvokeVoid(
		i,
		"addDefaultBackend",
		[]interface{}{backend},
	)
}

func (i *jsiiProxy_IngressV1Beta1) AddHostDefaultBackend(host *string, backend IngressV1Beta1Backend) {
	_jsii_.InvokeVoid(
		i,
		"addHostDefaultBackend",
		[]interface{}{host, backend},
	)
}

func (i *jsiiProxy_IngressV1Beta1) AddHostRule(host *string, path *string, backend IngressV1Beta1Backend) {
	_jsii_.InvokeVoid(
		i,
		"addHostRule",
		[]interface{}{host, path, backend},
	)
}

func (i *jsiiProxy_IngressV1Beta1) AddRule(path *string, backend IngressV1Beta1Backend) {
	_jsii_.InvokeVoid(
		i,
		"addRule",
		[]interface{}{path, backend},
	)
}

func (i *jsiiProxy_IngressV1Beta1) AddRules(rules ...*IngressV1Beta1Rule) {
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

func (i *jsiiProxy_IngressV1Beta1) AddTls(tls *[]*IngressV1Beta1Tls) {
	_jsii_.InvokeVoid(
		i,
		"addTls",
		[]interface{}{tls},
	)
}

func (i *jsiiProxy_IngressV1Beta1) OnPrepare() {
	_jsii_.InvokeVoid(
		i,
		"onPrepare",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IngressV1Beta1) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		i,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (i *jsiiProxy_IngressV1Beta1) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IngressV1Beta1) ToString() *string {
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
type IngressV1Beta1Backend interface {
}

// The jsii proxy struct for IngressV1Beta1Backend
type jsiiProxy_IngressV1Beta1Backend struct {
	_ byte // padding
}

// A Kubernetes `Service` to use as the backend for this path.
func IngressV1Beta1Backend_FromService(service Service, options *ServiceIngressV1BetaBackendOptions) IngressV1Beta1Backend {
	_init_.Initialize()

	var returns IngressV1Beta1Backend

	_jsii_.StaticInvoke(
		"cdk8s-plus-21.IngressV1Beta1Backend",
		"fromService",
		[]interface{}{service, options},
		&returns,
	)

	return returns
}

// Properties for `Ingress`.
type IngressV1Beta1Props struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata" yaml:"metadata"`
	// The default backend services requests that do not match any rule.
	//
	// Using this option or the `addDefaultBackend()` method is equivalent to
	// adding a rule with both `path` and `host` undefined.
	DefaultBackend IngressV1Beta1Backend `json:"defaultBackend" yaml:"defaultBackend"`
	// Routing rules for this ingress.
	//
	// Each rule must define an `IngressBackend` that will receive the requests
	// that match this rule. If both `host` and `path` are not specifiec, this
	// backend will be used as the default backend of the ingress.
	//
	// You can also add rules later using `addRule()`, `addHostRule()`,
	// `addDefaultBackend()` and `addHostDefaultBackend()`.
	Rules *[]*IngressV1Beta1Rule `json:"rules" yaml:"rules"`
	// TLS settings for this ingress.
	//
	// Using this option tells the ingress controller to expose a TLS endpoint.
	// Currently the Ingress only supports a single TLS port, 443. If multiple
	// members of this list specify different hosts, they will be multiplexed on
	// the same port according to the hostname specified through the SNI TLS
	// extension, if the ingress controller fulfilling the ingress supports SNI.
	Tls *[]*IngressV1Beta1Tls `json:"tls" yaml:"tls"`
}

// Represents the rules mapping the paths under a specified host to the related backend services.
//
// Incoming requests are first evaluated for a host match,
// then routed to the backend associated with the matching path.
type IngressV1Beta1Rule struct {
	// Backend defines the referenced service endpoint to which the traffic will be forwarded to.
	Backend IngressV1Beta1Backend `json:"backend" yaml:"backend"`
	// Host is the fully qualified domain name of a network host, as defined by RFC 3986.
	//
	// Note the following deviations from the "host" part of the URI as
	// defined in the RFC: 1. IPs are not allowed. Currently an IngressRuleValue
	// can only apply to the IP in the Spec of the parent Ingress. 2. The `:`
	// delimiter is not respected because ports are not allowed. Currently the
	// port of an Ingress is implicitly :80 for http and :443 for https. Both
	// these may change in the future. Incoming requests are matched against the
	// host before the IngressRuleValue.
	Host *string `json:"host" yaml:"host"`
	// Path is an extended POSIX regex as defined by IEEE Std 1003.1, (i.e this follows the egrep/unix syntax, not the perl syntax) matched against the path of an incoming request. Currently it can contain characters disallowed from the conventional "path" part of a URL as defined by RFC 3986. Paths must begin with a '/'.
	Path *string `json:"path" yaml:"path"`
}

// Represents the TLS configuration mapping that is passed to the ingress controller for SSL termination.
type IngressV1Beta1Tls struct {
	// Hosts are a list of hosts included in the TLS certificate.
	//
	// The values in
	// this list must match the name/s used in the TLS Secret.
	Hosts *[]*string `json:"hosts" yaml:"hosts"`
	// Secret is the secret that contains the certificate and key used to terminate SSL traffic on 443.
	//
	// If the SNI host in a listener conflicts with
	// the "Host" header field used by an IngressRule, the SNI host is used for
	// termination and value of the Host header is used for routing.
	Secret ISecret `json:"secret" yaml:"secret"`
}

// A Job creates one or more Pods and ensures that a specified number of them successfully terminate.
//
// As pods successfully complete,
// the Job tracks the successful completions. When a specified number of successful completions is reached, the task (ie, Job) is complete.
// Deleting a Job will clean up the Pods it created. A simple case is to create one Job object in order to reliably run one Pod to completion.
// The Job object will start a new Pod if the first Pod fails or is deleted (for example due to a node hardware failure or a node reboot).
// You can also use a Job to run multiple Pods in parallel.
type Job interface {
	Resource
	IPodTemplate
	// Duration before job is terminated.
	//
	// If undefined, there is no deadline.
	ActiveDeadline() cdk8s.Duration
	// The underlying cdk8s API object.
	// See: base.Resource.apiObject
	//
	ApiObject() cdk8s.ApiObject
	// Number of retries before marking failed.
	BackoffLimit() *float64
	// The containers belonging to the pod.
	//
	// Use `addContainer` to add containers.
	Containers() *[]Container
	// An optional list of hosts and IPs that will be injected into the pod's hosts file if specified.
	//
	// This is only valid for non-hostNetwork pods.
	HostAliases() *[]*HostAlias
	// The init containers belonging to the pod.
	//
	// Use `addInitContainer` to add init containers.
	InitContainers() *[]Container
	Metadata() cdk8s.ApiObjectMetadataDefinition
	// The name of this API object.
	Name() *string
	// Provides read/write access to the underlying pod metadata of the resource.
	PodMetadata() cdk8s.ApiObjectMetadataDefinition
	// Restart policy for all containers within the pod.
	RestartPolicy() RestartPolicy
	SecurityContext() PodSecurityContext
	// The service account used to run this pod.
	ServiceAccount() IServiceAccount
	// TTL before the job is deleted after it is finished.
	TtlAfterFinished() cdk8s.Duration
	// The volumes associated with this pod.
	//
	// Use `addVolume` to add volumes.
	Volumes() *[]Volume
	// Add a container to the pod.
	AddContainer(container *ContainerProps) Container
	AddHostAlias(hostAlias *HostAlias)
	// Add an init container to the pod.
	AddInitContainer(container *ContainerProps) Container
	// Add a volume to the pod.
	AddVolume(volume Volume)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if there the construct is valid.
	// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
	// instead of overriding this method.
	OnValidate() *[]*string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Job
type jsiiProxy_Job struct {
	jsiiProxy_Resource
	jsiiProxy_IPodTemplate
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

func (j *jsiiProxy_Job) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
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

func (j *jsiiProxy_Job) Containers() *[]Container {
	var returns *[]Container
	_jsii_.Get(
		j,
		"containers",
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

func (j *jsiiProxy_Job) PodMetadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"podMetadata",
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
		"cdk8s-plus-21.Job",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewJob_Override(j Job, scope constructs.Construct, id *string, props *JobProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-21.Job",
		[]interface{}{scope, id, props},
		j,
	)
}

func (j *jsiiProxy_Job) AddContainer(container *ContainerProps) Container {
	var returns Container

	_jsii_.Invoke(
		j,
		"addContainer",
		[]interface{}{container},
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

func (j *jsiiProxy_Job) AddInitContainer(container *ContainerProps) Container {
	var returns Container

	_jsii_.Invoke(
		j,
		"addInitContainer",
		[]interface{}{container},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) AddVolume(volume Volume) {
	_jsii_.InvokeVoid(
		j,
		"addVolume",
		[]interface{}{volume},
	)
}

func (j *jsiiProxy_Job) OnPrepare() {
	_jsii_.InvokeVoid(
		j,
		"onPrepare",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Job) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		j,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (j *jsiiProxy_Job) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		j,
		"onValidate",
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

// Properties for initialization of `Job`.
type JobProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata" yaml:"metadata"`
	// List of containers belonging to the pod.
	//
	// Containers cannot currently be
	// added or removed. There must be at least one container in a Pod.
	//
	// You can add additionnal containers using `podSpec.addContainer()`
	Containers *[]*ContainerProps `json:"containers" yaml:"containers"`
	// HostAlias holds the mapping between IP and hostnames that will be injected as an entry in the pod's hosts file.
	HostAliases *[]*HostAlias `json:"hostAliases" yaml:"hostAliases"`
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
	InitContainers *[]*ContainerProps `json:"initContainers" yaml:"initContainers"`
	// Restart policy for all containers within the pod.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy
	//
	RestartPolicy RestartPolicy `json:"restartPolicy" yaml:"restartPolicy"`
	// SecurityContext holds pod-level security attributes and common container settings.
	SecurityContext *PodSecurityContextProps `json:"securityContext" yaml:"securityContext"`
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
	ServiceAccount IServiceAccount `json:"serviceAccount" yaml:"serviceAccount"`
	// List of volumes that can be mounted by containers belonging to the pod.
	//
	// You can also add volumes later using `podSpec.addVolume()`
	// See: https://kubernetes.io/docs/concepts/storage/volumes
	//
	Volumes *[]Volume `json:"volumes" yaml:"volumes"`
	// The pod metadata.
	PodMetadata *cdk8s.ApiObjectMetadata `json:"podMetadata" yaml:"podMetadata"`
	// Specifies the duration the job may be active before the system tries to terminate it.
	ActiveDeadline cdk8s.Duration `json:"activeDeadline" yaml:"activeDeadline"`
	// Specifies the number of retries before marking this job failed.
	BackoffLimit *float64 `json:"backoffLimit" yaml:"backoffLimit"`
	// Limits the lifetime of a Job that has finished execution (either Complete or Failed).
	//
	// If this field is set, after the Job finishes, it is eligible to
	// be automatically deleted. When the Job is being deleted, its lifecycle
	// guarantees (e.g. finalizers) will be honored. If this field is set to zero,
	// the Job becomes eligible to be deleted immediately after it finishes. This
	// field is alpha-level and is only honored by servers that enable the
	// `TTLAfterFinished` feature.
	TtlAfterFinished cdk8s.Duration `json:"ttlAfterFinished" yaml:"ttlAfterFinished"`
}

// Memory request and limit.
type MemoryResources struct {
	Limit cdk8s.Size `json:"limit" yaml:"limit"`
	Request cdk8s.Size `json:"request" yaml:"request"`
}

// Options for mounts.
type MountOptions struct {
	// Determines how mounts are propagated from the host to container and the other way around.
	//
	// When not set, MountPropagationNone is used.
	//
	// Mount propagation allows for sharing volumes mounted by a Container to
	// other Containers in the same Pod, or even to other Pods on the same node.
	Propagation MountPropagation `json:"propagation" yaml:"propagation"`
	// Mounted read-only if true, read-write otherwise (false or unspecified).
	//
	// Defaults to false.
	ReadOnly *bool `json:"readOnly" yaml:"readOnly"`
	// Path within the volume from which the container's volume should be mounted.).
	SubPath *string `json:"subPath" yaml:"subPath"`
	// Expanded path within the volume from which the container's volume should be mounted.
	//
	// Behaves similarly to SubPath but environment variable references
	// $(VAR_NAME) are expanded using the container's environment. Defaults to ""
	// (volume's root).
	//
	// `subPathExpr` and `subPath` are mutually exclusive.
	SubPathExpr *string `json:"subPathExpr" yaml:"subPathExpr"`
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

// Maps a string key to a path within a volume.
type PathMapping struct {
	// The relative path of the file to map the key to.
	//
	// May not be an absolute
	// path. May not contain the path element '..'. May not start with the string
	// '..'.
	Path *string `json:"path" yaml:"path"`
	// Optional: mode bits to use on this file, must be a value between 0 and 0777.
	//
	// If not specified, the volume defaultMode will be used. This might be
	// in conflict with other options that affect the file mode, like fsGroup, and
	// the result can be other mode bits set.
	Mode *float64 `json:"mode" yaml:"mode"`
}

// Pod is a collection of containers that can run on a host.
//
// This resource is
// created by clients and scheduled onto hosts.
type Pod interface {
	Resource
	IPodSpec
	// The underlying cdk8s API object.
	// See: base.Resource.apiObject
	//
	ApiObject() cdk8s.ApiObject
	// The containers belonging to the pod.
	//
	// Use `addContainer` to add containers.
	Containers() *[]Container
	// An optional list of hosts and IPs that will be injected into the pod's hosts file if specified.
	//
	// This is only valid for non-hostNetwork pods.
	HostAliases() *[]*HostAlias
	// The init containers belonging to the pod.
	//
	// Use `addInitContainer` to add init containers.
	InitContainers() *[]Container
	Metadata() cdk8s.ApiObjectMetadataDefinition
	// The name of this API object.
	Name() *string
	// Restart policy for all containers within the pod.
	RestartPolicy() RestartPolicy
	SecurityContext() PodSecurityContext
	// The service account used to run this pod.
	ServiceAccount() IServiceAccount
	// The volumes associated with this pod.
	//
	// Use `addVolume` to add volumes.
	Volumes() *[]Volume
	// Add a container to the pod.
	AddContainer(container *ContainerProps) Container
	AddHostAlias(hostAlias *HostAlias)
	// Add an init container to the pod.
	AddInitContainer(container *ContainerProps) Container
	// Add a volume to the pod.
	AddVolume(volume Volume)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if there the construct is valid.
	// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
	// instead of overriding this method.
	OnValidate() *[]*string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Pod
type jsiiProxy_Pod struct {
	jsiiProxy_Resource
	jsiiProxy_IPodSpec
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

func (j *jsiiProxy_Pod) Containers() *[]Container {
	var returns *[]Container
	_jsii_.Get(
		j,
		"containers",
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

func (j *jsiiProxy_Pod) RestartPolicy() RestartPolicy {
	var returns RestartPolicy
	_jsii_.Get(
		j,
		"restartPolicy",
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
		"cdk8s-plus-21.Pod",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewPod_Override(p Pod, scope constructs.Construct, id *string, props *PodProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-21.Pod",
		[]interface{}{scope, id, props},
		p,
	)
}

func (p *jsiiProxy_Pod) AddContainer(container *ContainerProps) Container {
	var returns Container

	_jsii_.Invoke(
		p,
		"addContainer",
		[]interface{}{container},
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

func (p *jsiiProxy_Pod) AddInitContainer(container *ContainerProps) Container {
	var returns Container

	_jsii_.Invoke(
		p,
		"addInitContainer",
		[]interface{}{container},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pod) AddVolume(volume Volume) {
	_jsii_.InvokeVoid(
		p,
		"addVolume",
		[]interface{}{volume},
	)
}

func (p *jsiiProxy_Pod) OnPrepare() {
	_jsii_.InvokeVoid(
		p,
		"onPrepare",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pod) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		p,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (p *jsiiProxy_Pod) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"onValidate",
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

// Properties for initialization of `Pod`.
type PodProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata" yaml:"metadata"`
	// List of containers belonging to the pod.
	//
	// Containers cannot currently be
	// added or removed. There must be at least one container in a Pod.
	//
	// You can add additionnal containers using `podSpec.addContainer()`
	Containers *[]*ContainerProps `json:"containers" yaml:"containers"`
	// HostAlias holds the mapping between IP and hostnames that will be injected as an entry in the pod's hosts file.
	HostAliases *[]*HostAlias `json:"hostAliases" yaml:"hostAliases"`
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
	InitContainers *[]*ContainerProps `json:"initContainers" yaml:"initContainers"`
	// Restart policy for all containers within the pod.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy
	//
	RestartPolicy RestartPolicy `json:"restartPolicy" yaml:"restartPolicy"`
	// SecurityContext holds pod-level security attributes and common container settings.
	SecurityContext *PodSecurityContextProps `json:"securityContext" yaml:"securityContext"`
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
	ServiceAccount IServiceAccount `json:"serviceAccount" yaml:"serviceAccount"`
	// List of volumes that can be mounted by containers belonging to the pod.
	//
	// You can also add volumes later using `podSpec.addVolume()`
	// See: https://kubernetes.io/docs/concepts/storage/volumes
	//
	Volumes *[]Volume `json:"volumes" yaml:"volumes"`
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
		"cdk8s-plus-21.PodSecurityContext",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewPodSecurityContext_Override(p PodSecurityContext, props *PodSecurityContextProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-21.PodSecurityContext",
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
	EnsureNonRoot *bool `json:"ensureNonRoot" yaml:"ensureNonRoot"`
	// Modify the ownership and permissions of pod volumes to this GID.
	FsGroup *float64 `json:"fsGroup" yaml:"fsGroup"`
	// Defines behavior of changing ownership and permission of the volume before being exposed inside Pod.
	//
	// This field will only apply to volume types which support fsGroup based ownership(and permissions).
	// It will have no effect on ephemeral volume types such as: secret, configmaps and emptydir.
	FsGroupChangePolicy FsGroupChangePolicy `json:"fsGroupChangePolicy" yaml:"fsGroupChangePolicy"`
	// The GID to run the entrypoint of the container process.
	Group *float64 `json:"group" yaml:"group"`
	// Sysctls hold a list of namespaced sysctls used for the pod.
	//
	// Pods with unsupported sysctls (by the container runtime) might fail to launch.
	Sysctls *[]*Sysctl `json:"sysctls" yaml:"sysctls"`
	// The UID to run the entrypoint of the container process.
	User *float64 `json:"user" yaml:"user"`
}

// Provides read/write capabilities ontop of a `PodSpecProps`.
type PodSpec interface {
	IPodSpec
	// The containers belonging to the pod.
	//
	// Use `addContainer` to add containers.
	Containers() *[]Container
	// An optional list of hosts and IPs that will be injected into the pod's hosts file if specified.
	//
	// This is only valid for non-hostNetwork pods.
	HostAliases() *[]*HostAlias
	// The init containers belonging to the pod.
	//
	// Use `addInitContainer` to add init containers.
	InitContainers() *[]Container
	// Restart policy for all containers within the pod.
	RestartPolicy() RestartPolicy
	SecurityContext() PodSecurityContext
	// The service account used to run this pod.
	ServiceAccount() IServiceAccount
	// The volumes associated with this pod.
	//
	// Use `addVolume` to add volumes.
	Volumes() *[]Volume
	// Add a container to the pod.
	AddContainer(container *ContainerProps) Container
	AddHostAlias(hostAlias *HostAlias)
	// Add an init container to the pod.
	AddInitContainer(container *ContainerProps) Container
	// Add a volume to the pod.
	AddVolume(volume Volume)
}

// The jsii proxy struct for PodSpec
type jsiiProxy_PodSpec struct {
	jsiiProxy_IPodSpec
}

func (j *jsiiProxy_PodSpec) Containers() *[]Container {
	var returns *[]Container
	_jsii_.Get(
		j,
		"containers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpec) HostAliases() *[]*HostAlias {
	var returns *[]*HostAlias
	_jsii_.Get(
		j,
		"hostAliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpec) InitContainers() *[]Container {
	var returns *[]Container
	_jsii_.Get(
		j,
		"initContainers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpec) RestartPolicy() RestartPolicy {
	var returns RestartPolicy
	_jsii_.Get(
		j,
		"restartPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpec) SecurityContext() PodSecurityContext {
	var returns PodSecurityContext
	_jsii_.Get(
		j,
		"securityContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpec) ServiceAccount() IServiceAccount {
	var returns IServiceAccount
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpec) Volumes() *[]Volume {
	var returns *[]Volume
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}


func NewPodSpec(props *PodSpecProps) PodSpec {
	_init_.Initialize()

	j := jsiiProxy_PodSpec{}

	_jsii_.Create(
		"cdk8s-plus-21.PodSpec",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewPodSpec_Override(p PodSpec, props *PodSpecProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-21.PodSpec",
		[]interface{}{props},
		p,
	)
}

func (p *jsiiProxy_PodSpec) AddContainer(container *ContainerProps) Container {
	var returns Container

	_jsii_.Invoke(
		p,
		"addContainer",
		[]interface{}{container},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSpec) AddHostAlias(hostAlias *HostAlias) {
	_jsii_.InvokeVoid(
		p,
		"addHostAlias",
		[]interface{}{hostAlias},
	)
}

func (p *jsiiProxy_PodSpec) AddInitContainer(container *ContainerProps) Container {
	var returns Container

	_jsii_.Invoke(
		p,
		"addInitContainer",
		[]interface{}{container},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSpec) AddVolume(volume Volume) {
	_jsii_.InvokeVoid(
		p,
		"addVolume",
		[]interface{}{volume},
	)
}

// Properties of a `PodSpec`.
type PodSpecProps struct {
	// List of containers belonging to the pod.
	//
	// Containers cannot currently be
	// added or removed. There must be at least one container in a Pod.
	//
	// You can add additionnal containers using `podSpec.addContainer()`
	Containers *[]*ContainerProps `json:"containers" yaml:"containers"`
	// HostAlias holds the mapping between IP and hostnames that will be injected as an entry in the pod's hosts file.
	HostAliases *[]*HostAlias `json:"hostAliases" yaml:"hostAliases"`
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
	InitContainers *[]*ContainerProps `json:"initContainers" yaml:"initContainers"`
	// Restart policy for all containers within the pod.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy
	//
	RestartPolicy RestartPolicy `json:"restartPolicy" yaml:"restartPolicy"`
	// SecurityContext holds pod-level security attributes and common container settings.
	SecurityContext *PodSecurityContextProps `json:"securityContext" yaml:"securityContext"`
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
	ServiceAccount IServiceAccount `json:"serviceAccount" yaml:"serviceAccount"`
	// List of volumes that can be mounted by containers belonging to the pod.
	//
	// You can also add volumes later using `podSpec.addVolume()`
	// See: https://kubernetes.io/docs/concepts/storage/volumes
	//
	Volumes *[]Volume `json:"volumes" yaml:"volumes"`
}

// Provides read/write capabilities ontop of a `PodTemplateProps`.
type PodTemplate interface {
	PodSpec
	IPodTemplate
	// The containers belonging to the pod.
	//
	// Use `addContainer` to add containers.
	Containers() *[]Container
	// An optional list of hosts and IPs that will be injected into the pod's hosts file if specified.
	//
	// This is only valid for non-hostNetwork pods.
	HostAliases() *[]*HostAlias
	// The init containers belonging to the pod.
	//
	// Use `addInitContainer` to add init containers.
	InitContainers() *[]Container
	// Provides read/write access to the underlying pod metadata of the resource.
	PodMetadata() cdk8s.ApiObjectMetadataDefinition
	// Restart policy for all containers within the pod.
	RestartPolicy() RestartPolicy
	SecurityContext() PodSecurityContext
	// The service account used to run this pod.
	ServiceAccount() IServiceAccount
	// The volumes associated with this pod.
	//
	// Use `addVolume` to add volumes.
	Volumes() *[]Volume
	// Add a container to the pod.
	AddContainer(container *ContainerProps) Container
	AddHostAlias(hostAlias *HostAlias)
	// Add an init container to the pod.
	AddInitContainer(container *ContainerProps) Container
	// Add a volume to the pod.
	AddVolume(volume Volume)
}

// The jsii proxy struct for PodTemplate
type jsiiProxy_PodTemplate struct {
	jsiiProxy_PodSpec
	jsiiProxy_IPodTemplate
}

func (j *jsiiProxy_PodTemplate) Containers() *[]Container {
	var returns *[]Container
	_jsii_.Get(
		j,
		"containers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodTemplate) HostAliases() *[]*HostAlias {
	var returns *[]*HostAlias
	_jsii_.Get(
		j,
		"hostAliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodTemplate) InitContainers() *[]Container {
	var returns *[]Container
	_jsii_.Get(
		j,
		"initContainers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodTemplate) PodMetadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"podMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodTemplate) RestartPolicy() RestartPolicy {
	var returns RestartPolicy
	_jsii_.Get(
		j,
		"restartPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodTemplate) SecurityContext() PodSecurityContext {
	var returns PodSecurityContext
	_jsii_.Get(
		j,
		"securityContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodTemplate) ServiceAccount() IServiceAccount {
	var returns IServiceAccount
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodTemplate) Volumes() *[]Volume {
	var returns *[]Volume
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}


func NewPodTemplate(props *PodTemplateProps) PodTemplate {
	_init_.Initialize()

	j := jsiiProxy_PodTemplate{}

	_jsii_.Create(
		"cdk8s-plus-21.PodTemplate",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewPodTemplate_Override(p PodTemplate, props *PodTemplateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-21.PodTemplate",
		[]interface{}{props},
		p,
	)
}

func (p *jsiiProxy_PodTemplate) AddContainer(container *ContainerProps) Container {
	var returns Container

	_jsii_.Invoke(
		p,
		"addContainer",
		[]interface{}{container},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodTemplate) AddHostAlias(hostAlias *HostAlias) {
	_jsii_.InvokeVoid(
		p,
		"addHostAlias",
		[]interface{}{hostAlias},
	)
}

func (p *jsiiProxy_PodTemplate) AddInitContainer(container *ContainerProps) Container {
	var returns Container

	_jsii_.Invoke(
		p,
		"addInitContainer",
		[]interface{}{container},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodTemplate) AddVolume(volume Volume) {
	_jsii_.InvokeVoid(
		p,
		"addVolume",
		[]interface{}{volume},
	)
}

// Properties of a `PodTemplate`.
//
// Adds metadata information on top of the spec.
type PodTemplateProps struct {
	// List of containers belonging to the pod.
	//
	// Containers cannot currently be
	// added or removed. There must be at least one container in a Pod.
	//
	// You can add additionnal containers using `podSpec.addContainer()`
	Containers *[]*ContainerProps `json:"containers" yaml:"containers"`
	// HostAlias holds the mapping between IP and hostnames that will be injected as an entry in the pod's hosts file.
	HostAliases *[]*HostAlias `json:"hostAliases" yaml:"hostAliases"`
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
	InitContainers *[]*ContainerProps `json:"initContainers" yaml:"initContainers"`
	// Restart policy for all containers within the pod.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy
	//
	RestartPolicy RestartPolicy `json:"restartPolicy" yaml:"restartPolicy"`
	// SecurityContext holds pod-level security attributes and common container settings.
	SecurityContext *PodSecurityContextProps `json:"securityContext" yaml:"securityContext"`
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
	ServiceAccount IServiceAccount `json:"serviceAccount" yaml:"serviceAccount"`
	// List of volumes that can be mounted by containers belonging to the pod.
	//
	// You can also add volumes later using `podSpec.addVolume()`
	// See: https://kubernetes.io/docs/concepts/storage/volumes
	//
	Volumes *[]Volume `json:"volumes" yaml:"volumes"`
	// The pod metadata.
	PodMetadata *cdk8s.ApiObjectMetadata `json:"podMetadata" yaml:"podMetadata"`
}

// Probe describes a health check to be performed against a container to determine whether it is alive or ready to receive traffic.
type Probe interface {
}

// The jsii proxy struct for Probe
type jsiiProxy_Probe struct {
	_ byte // padding
}

func NewProbe_Override(p Probe) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-21.Probe",
		nil, // no parameters
		p,
	)
}

// Defines a probe based on a command which is executed within the container.
func Probe_FromCommand(command *[]*string, options *CommandProbeOptions) Probe {
	_init_.Initialize()

	var returns Probe

	_jsii_.StaticInvoke(
		"cdk8s-plus-21.Probe",
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
		"cdk8s-plus-21.Probe",
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
		"cdk8s-plus-21.Probe",
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
	FailureThreshold *float64 `json:"failureThreshold" yaml:"failureThreshold"`
	// Number of seconds after the container has started before liveness probes are initiated.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	//
	InitialDelaySeconds cdk8s.Duration `json:"initialDelaySeconds" yaml:"initialDelaySeconds"`
	// How often (in seconds) to perform the probe.
	//
	// Default to 10 seconds. Minimum value is 1.
	PeriodSeconds cdk8s.Duration `json:"periodSeconds" yaml:"periodSeconds"`
	// Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1.
	//
	// Must be 1 for liveness and startup. Minimum value is 1.
	SuccessThreshold *float64 `json:"successThreshold" yaml:"successThreshold"`
	// Number of seconds after which the probe times out.
	//
	// Defaults to 1 second. Minimum value is 1.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	//
	TimeoutSeconds cdk8s.Duration `json:"timeoutSeconds" yaml:"timeoutSeconds"`
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
	IResource
	// The underlying cdk8s API object.
	ApiObject() cdk8s.ApiObject
	Metadata() cdk8s.ApiObjectMetadataDefinition
	// The name of this API object.
	Name() *string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if there the construct is valid.
	// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
	// instead of overriding this method.
	OnValidate() *[]*string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Resource
type jsiiProxy_Resource struct {
	internal.Type__constructsConstruct
	jsiiProxy_IResource
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


// Creates a new construct node.
func NewResource_Override(r Resource, scope constructs.Construct, id *string, options *constructs.ConstructOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-21.Resource",
		[]interface{}{scope, id, options},
		r,
	)
}

func (r *jsiiProxy_Resource) OnPrepare() {
	_jsii_.InvokeVoid(
		r,
		"onPrepare",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resource) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		r,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (r *jsiiProxy_Resource) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"onValidate",
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

// Initialization properties for resources.
type ResourceProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata" yaml:"metadata"`
}

// CPU and memory compute resources.
type Resources struct {
	Cpu *CpuResources `json:"cpu" yaml:"cpu"`
	Memory *MemoryResources `json:"memory" yaml:"memory"`
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
	// The underlying cdk8s API object.
	// See: base.Resource.apiObject
	//
	ApiObject() cdk8s.ApiObject
	Metadata() cdk8s.ApiObjectMetadataDefinition
	// The name of this API object.
	Name() *string
	// Adds a string data field to the secert.
	AddStringData(key *string, value *string)
	// Gets a string data by key or undefined.
	GetStringData(key *string) *string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if there the construct is valid.
	// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
	// instead of overriding this method.
	OnValidate() *[]*string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Secret
type jsiiProxy_Secret struct {
	jsiiProxy_Resource
	jsiiProxy_ISecret
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


func NewSecret(scope constructs.Construct, id *string, props *SecretProps) Secret {
	_init_.Initialize()

	j := jsiiProxy_Secret{}

	_jsii_.Create(
		"cdk8s-plus-21.Secret",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewSecret_Override(s Secret, scope constructs.Construct, id *string, props *SecretProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-21.Secret",
		[]interface{}{scope, id, props},
		s,
	)
}

// Imports a secret from the cluster as a reference.
func Secret_FromSecretName(name *string) ISecret {
	_init_.Initialize()

	var returns ISecret

	_jsii_.StaticInvoke(
		"cdk8s-plus-21.Secret",
		"fromSecretName",
		[]interface{}{name},
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

func (s *jsiiProxy_Secret) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Secret) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_Secret) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
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

type SecretProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata" yaml:"metadata"`
	// stringData allows specifying non-binary secret data in string form.
	//
	// It is
	// provided as a write-only convenience method. All keys and values are merged
	// into the data field on write, overwriting any existing values. It is never
	// output when reading from the API.
	StringData *map[string]*string `json:"stringData" yaml:"stringData"`
	// Optional type associated with the secret.
	//
	// Used to facilitate programmatic
	// handling of secret data by various controllers.
	Type *string `json:"type" yaml:"type"`
}

// Represents a specific value in JSON secret.
type SecretValue struct {
	// The JSON key.
	Key *string `json:"key" yaml:"key"`
	// The secret.
	Secret ISecret `json:"secret" yaml:"secret"`
}

// Options for the Secret-based volume.
type SecretVolumeOptions struct {
	// Mode bits to use on created files by default.
	//
	// Must be a value between 0 and
	// 0777. Defaults to 0644. Directories within the path are not affected by
	// this setting. This might be in conflict with other options that affect the
	// file mode, like fsGroup, and the result can be other mode bits set.
	DefaultMode *float64 `json:"defaultMode" yaml:"defaultMode"`
	// If unspecified, each key-value pair in the Data field of the referenced secret will be projected into the volume as a file whose name is the key and content is the value.
	//
	// If specified, the listed keys will be projected
	// into the specified paths, and unlisted keys will not be present. If a key
	// is specified which is not present in the secret, the volume setup will
	// error unless it is marked optional. Paths must be relative and may not
	// contain the '..' path or start with '..'.
	Items *map[string]*PathMapping `json:"items" yaml:"items"`
	// The volume name.
	Name *string `json:"name" yaml:"name"`
	// Specify whether the secret or its keys must be defined.
	Optional *bool `json:"optional" yaml:"optional"`
}

// An abstract way to expose an application running on a set of Pods as a network service.
//
// With Kubernetes you don't need to modify your application to use an unfamiliar service discovery mechanism.
// Kubernetes gives Pods their own IP addresses and a single DNS name for a set of Pods, and can load-balance across them.
//
// For example, consider a stateless image-processing backend which is running with 3 replicas. Those replicas are fungible—frontends do not care which backend they use.
// While the actual Pods that compose the backend set may change, the frontend clients should not need to be aware of that,
// nor should they need to keep track of the set of backends themselves.
// The Service abstraction enables this decoupling.
//
// If you're able to use Kubernetes APIs for service discovery in your application, you can query the API server for Endpoints,
// that get updated whenever the set of Pods in a Service changes. For non-native applications, Kubernetes offers ways to place a network port
// or load balancer in between your application and the backend Pods.
type Service interface {
	Resource
	// The underlying cdk8s API object.
	// See: base.Resource.apiObject
	//
	ApiObject() cdk8s.ApiObject
	// The IP address of the service and is usually assigned randomly by the master.
	ClusterIP() *string
	// The externalName to be used for EXTERNAL_NAME types.
	ExternalName() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	// The name of this API object.
	Name() *string
	// Ports for this service.
	//
	// Use `serve()` to expose additional service ports.
	Ports() *[]*ServicePort
	// Returns the labels which are used to select pods for this service.
	Selector() *map[string]*string
	// Determines how the Service is exposed.
	Type() ServiceType
	// Associate a deployment to this service.
	//
	// If not targetPort is specific in the portOptions, then requests will be routed
	// to the port exposed by the first container in the deployment's pods.
	// The deployment's `labelSelector` will be used to select pods.
	AddDeployment(deployment Deployment, options *AddDeploymentOptions)
	// Services defined using this spec will select pods according the provided label.
	AddSelector(label *string, value *string)
	// Expose a service via an ingress using the specified path.
	//
	// Returns: The `Ingress` resource that was used.
	ExposeViaIngress(path *string, options *ExposeServiceViaIngressOptions) IngressV1Beta1
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if there the construct is valid.
	// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
	// instead of overriding this method.
	OnValidate() *[]*string
	// Configure a port the service will bind to.
	//
	// This method can be called multiple times.
	Serve(port *float64, options *ServicePortOptions)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Service
type jsiiProxy_Service struct {
	jsiiProxy_Resource
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

func (j *jsiiProxy_Service) Ports() *[]*ServicePort {
	var returns *[]*ServicePort
	_jsii_.Get(
		j,
		"ports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Selector() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"selector",
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
		"cdk8s-plus-21.Service",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewService_Override(s Service, scope constructs.Construct, id *string, props *ServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-21.Service",
		[]interface{}{scope, id, props},
		s,
	)
}

func (s *jsiiProxy_Service) AddDeployment(deployment Deployment, options *AddDeploymentOptions) {
	_jsii_.InvokeVoid(
		s,
		"addDeployment",
		[]interface{}{deployment, options},
	)
}

func (s *jsiiProxy_Service) AddSelector(label *string, value *string) {
	_jsii_.InvokeVoid(
		s,
		"addSelector",
		[]interface{}{label, value},
	)
}

func (s *jsiiProxy_Service) ExposeViaIngress(path *string, options *ExposeServiceViaIngressOptions) IngressV1Beta1 {
	var returns IngressV1Beta1

	_jsii_.Invoke(
		s,
		"exposeViaIngress",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Service) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Service) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_Service) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Service) Serve(port *float64, options *ServicePortOptions) {
	_jsii_.InvokeVoid(
		s,
		"serve",
		[]interface{}{port, options},
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
	// The underlying cdk8s API object.
	// See: base.Resource.apiObject
	//
	ApiObject() cdk8s.ApiObject
	Metadata() cdk8s.ApiObjectMetadataDefinition
	// The name of this API object.
	Name() *string
	// List of secrets allowed to be used by pods running using this service account.
	//
	// Returns a copy. To add a secret, use `addSecret()`.
	Secrets() *[]ISecret
	// Allow a secret to be accessed by pods using this service account.
	AddSecret(secret ISecret)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if there the construct is valid.
	// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
	// instead of overriding this method.
	OnValidate() *[]*string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for ServiceAccount
type jsiiProxy_ServiceAccount struct {
	jsiiProxy_Resource
	jsiiProxy_IServiceAccount
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
		"cdk8s-plus-21.ServiceAccount",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewServiceAccount_Override(s ServiceAccount, scope constructs.Construct, id *string, props *ServiceAccountProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-21.ServiceAccount",
		[]interface{}{scope, id, props},
		s,
	)
}

// Imports a service account from the cluster as a reference.
func ServiceAccount_FromServiceAccountName(name *string) IServiceAccount {
	_init_.Initialize()

	var returns IServiceAccount

	_jsii_.StaticInvoke(
		"cdk8s-plus-21.ServiceAccount",
		"fromServiceAccountName",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceAccount) AddSecret(secret ISecret) {
	_jsii_.InvokeVoid(
		s,
		"addSecret",
		[]interface{}{secret},
	)
}

func (s *jsiiProxy_ServiceAccount) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceAccount) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_ServiceAccount) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
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

// Properties for initialization of `ServiceAccount`.
//
// Properties for initialization of `ServiceAccount`.
type ServiceAccountProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata" yaml:"metadata"`
	// List of secrets allowed to be used by pods running using this ServiceAccount.
	// See: https://kubernetes.io/docs/concepts/configuration/secret
	//
	Secrets *[]ISecret `json:"secrets" yaml:"secrets"`
}

// Options for setting up backends for ingress rules.
type ServiceIngressV1BetaBackendOptions struct {
	// The port to use to access the service.
	//
	// - This option will fail if the service does not expose any ports.
	// - If the service exposes multiple ports, this option must be specified.
	// - If the service exposes a single port, this option is optional and if
	//    specified, it must be the same port exposed by the service.
	Port *float64 `json:"port" yaml:"port"`
}

// Definition of a service port.
type ServicePort struct {
	// The name of this port within the service.
	//
	// This must be a DNS_LABEL. All
	// ports within a ServiceSpec must have unique names. This maps to the 'Name'
	// field in EndpointPort objects. Optional if only one ServicePort is defined
	// on this service.
	Name *string `json:"name" yaml:"name"`
	// The port on each node on which this service is exposed when type=NodePort or LoadBalancer.
	//
	// Usually assigned by the system. If specified, it will be
	// allocated to the service if unused or else creation of the service will
	// fail. Default is to auto-allocate a port if the ServiceType of this Service
	// requires one.
	// See: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport
	//
	NodePort *float64 `json:"nodePort" yaml:"nodePort"`
	// The IP protocol for this port.
	//
	// Supports "TCP", "UDP", and "SCTP". Default is TCP.
	Protocol Protocol `json:"protocol" yaml:"protocol"`
	// The port number the service will redirect to.
	TargetPort *float64 `json:"targetPort" yaml:"targetPort"`
	// The port number the service will bind to.
	Port *float64 `json:"port" yaml:"port"`
}

type ServicePortOptions struct {
	// The name of this port within the service.
	//
	// This must be a DNS_LABEL. All
	// ports within a ServiceSpec must have unique names. This maps to the 'Name'
	// field in EndpointPort objects. Optional if only one ServicePort is defined
	// on this service.
	Name *string `json:"name" yaml:"name"`
	// The port on each node on which this service is exposed when type=NodePort or LoadBalancer.
	//
	// Usually assigned by the system. If specified, it will be
	// allocated to the service if unused or else creation of the service will
	// fail. Default is to auto-allocate a port if the ServiceType of this Service
	// requires one.
	// See: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport
	//
	NodePort *float64 `json:"nodePort" yaml:"nodePort"`
	// The IP protocol for this port.
	//
	// Supports "TCP", "UDP", and "SCTP". Default is TCP.
	Protocol Protocol `json:"protocol" yaml:"protocol"`
	// The port number the service will redirect to.
	TargetPort *float64 `json:"targetPort" yaml:"targetPort"`
}

// Properties for initialization of `Service`.
type ServiceProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata" yaml:"metadata"`
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
	ClusterIP *string `json:"clusterIP" yaml:"clusterIP"`
	// A list of IP addresses for which nodes in the cluster will also accept traffic for this service.
	//
	// These IPs are not managed by Kubernetes. The user
	// is responsible for ensuring that traffic arrives at a node with this IP. A
	// common example is external load-balancers that are not part of the
	// Kubernetes system.
	ExternalIPs *[]*string `json:"externalIPs" yaml:"externalIPs"`
	// The externalName to be used when ServiceType.EXTERNAL_NAME is set.
	ExternalName *string `json:"externalName" yaml:"externalName"`
	// A list of CIDR IP addresses, if specified and supported by the platform, will restrict traffic through the cloud-provider load-balancer to the specified client IPs.
	//
	// More info: https://kubernetes.io/docs/tasks/access-application-cluster/configure-cloud-provider-firewall/
	LoadBalancerSourceRanges *[]*string `json:"loadBalancerSourceRanges" yaml:"loadBalancerSourceRanges"`
	// The port exposed by this service.
	//
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	Ports *[]*ServicePort `json:"ports" yaml:"ports"`
	// Determines how the Service is exposed.
	//
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types
	Type ServiceType `json:"type" yaml:"type"`
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
	Resource
	IPodTemplate
	// The underlying cdk8s API object.
	// See: base.Resource.apiObject
	//
	ApiObject() cdk8s.ApiObject
	// The containers belonging to the pod.
	//
	// Use `addContainer` to add containers.
	Containers() *[]Container
	// An optional list of hosts and IPs that will be injected into the pod's hosts file if specified.
	//
	// This is only valid for non-hostNetwork pods.
	HostAliases() *[]*HostAlias
	// The init containers belonging to the pod.
	//
	// Use `addInitContainer` to add init containers.
	InitContainers() *[]Container
	// The labels this statefulset will match against in order to select pods.
	//
	// Returns a a copy. Use `selectByLabel()` to add labels.
	LabelSelector() *map[string]*string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	// The name of this API object.
	Name() *string
	// Management policy to use for the set.
	PodManagementPolicy() PodManagementPolicy
	// Provides read/write access to the underlying pod metadata of the resource.
	PodMetadata() cdk8s.ApiObjectMetadataDefinition
	// Number of desired pods.
	Replicas() *float64
	// Restart policy for all containers within the pod.
	RestartPolicy() RestartPolicy
	SecurityContext() PodSecurityContext
	// The service account used to run this pod.
	ServiceAccount() IServiceAccount
	// The volumes associated with this pod.
	//
	// Use `addVolume` to add volumes.
	Volumes() *[]Volume
	// Add a container to the pod.
	AddContainer(container *ContainerProps) Container
	AddHostAlias(hostAlias *HostAlias)
	// Add an init container to the pod.
	AddInitContainer(container *ContainerProps) Container
	// Add a volume to the pod.
	AddVolume(volume Volume)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if there the construct is valid.
	// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
	// instead of overriding this method.
	OnValidate() *[]*string
	// Configure a label selector to this deployment.
	//
	// Pods that have the label will be selected by deployments configured with this spec.
	SelectByLabel(key *string, value *string)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for StatefulSet
type jsiiProxy_StatefulSet struct {
	jsiiProxy_Resource
	jsiiProxy_IPodTemplate
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

func (j *jsiiProxy_StatefulSet) Containers() *[]Container {
	var returns *[]Container
	_jsii_.Get(
		j,
		"containers",
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

func (j *jsiiProxy_StatefulSet) LabelSelector() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelSelector",
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

func (j *jsiiProxy_StatefulSet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
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

func (j *jsiiProxy_StatefulSet) RestartPolicy() RestartPolicy {
	var returns RestartPolicy
	_jsii_.Get(
		j,
		"restartPolicy",
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
		"cdk8s-plus-21.StatefulSet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewStatefulSet_Override(s StatefulSet, scope constructs.Construct, id *string, props *StatefulSetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-21.StatefulSet",
		[]interface{}{scope, id, props},
		s,
	)
}

func (s *jsiiProxy_StatefulSet) AddContainer(container *ContainerProps) Container {
	var returns Container

	_jsii_.Invoke(
		s,
		"addContainer",
		[]interface{}{container},
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

func (s *jsiiProxy_StatefulSet) AddInitContainer(container *ContainerProps) Container {
	var returns Container

	_jsii_.Invoke(
		s,
		"addInitContainer",
		[]interface{}{container},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSet) AddVolume(volume Volume) {
	_jsii_.InvokeVoid(
		s,
		"addVolume",
		[]interface{}{volume},
	)
}

func (s *jsiiProxy_StatefulSet) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSet) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_StatefulSet) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSet) SelectByLabel(key *string, value *string) {
	_jsii_.InvokeVoid(
		s,
		"selectByLabel",
		[]interface{}{key, value},
	)
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

// Properties for initialization of `StatefulSet`.
type StatefulSetProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata" yaml:"metadata"`
	// List of containers belonging to the pod.
	//
	// Containers cannot currently be
	// added or removed. There must be at least one container in a Pod.
	//
	// You can add additionnal containers using `podSpec.addContainer()`
	Containers *[]*ContainerProps `json:"containers" yaml:"containers"`
	// HostAlias holds the mapping between IP and hostnames that will be injected as an entry in the pod's hosts file.
	HostAliases *[]*HostAlias `json:"hostAliases" yaml:"hostAliases"`
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
	InitContainers *[]*ContainerProps `json:"initContainers" yaml:"initContainers"`
	// Restart policy for all containers within the pod.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy
	//
	RestartPolicy RestartPolicy `json:"restartPolicy" yaml:"restartPolicy"`
	// SecurityContext holds pod-level security attributes and common container settings.
	SecurityContext *PodSecurityContextProps `json:"securityContext" yaml:"securityContext"`
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
	ServiceAccount IServiceAccount `json:"serviceAccount" yaml:"serviceAccount"`
	// List of volumes that can be mounted by containers belonging to the pod.
	//
	// You can also add volumes later using `podSpec.addVolume()`
	// See: https://kubernetes.io/docs/concepts/storage/volumes
	//
	Volumes *[]Volume `json:"volumes" yaml:"volumes"`
	// The pod metadata.
	PodMetadata *cdk8s.ApiObjectMetadata `json:"podMetadata" yaml:"podMetadata"`
	// Service to associate with the statefulset.
	Service Service `json:"service" yaml:"service"`
	// Automatically allocates a pod selector for this statefulset.
	//
	// If this is set to `false` you must define your selector through
	// `statefulset.podMetadata.addLabel()` and `statefulset.selectByLabel()`.
	DefaultSelector *bool `json:"defaultSelector" yaml:"defaultSelector"`
	// Pod management policy to use for this statefulset.
	PodManagementPolicy PodManagementPolicy `json:"podManagementPolicy" yaml:"podManagementPolicy"`
	// Number of desired pods.
	Replicas *float64 `json:"replicas" yaml:"replicas"`
}

// Sysctl defines a kernel parameter to be set.
type Sysctl struct {
	// Name of a property to set.
	Name *string `json:"name" yaml:"name"`
	// Value of a property to set.
	Value *string `json:"value" yaml:"value"`
}

// Options for `Probe.fromTcpSocket()`.
type TcpSocketProbeOptions struct {
	// Minimum consecutive failures for the probe to be considered failed after having succeeded.
	//
	// Defaults to 3. Minimum value is 1.
	FailureThreshold *float64 `json:"failureThreshold" yaml:"failureThreshold"`
	// Number of seconds after the container has started before liveness probes are initiated.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	//
	InitialDelaySeconds cdk8s.Duration `json:"initialDelaySeconds" yaml:"initialDelaySeconds"`
	// How often (in seconds) to perform the probe.
	//
	// Default to 10 seconds. Minimum value is 1.
	PeriodSeconds cdk8s.Duration `json:"periodSeconds" yaml:"periodSeconds"`
	// Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1.
	//
	// Must be 1 for liveness and startup. Minimum value is 1.
	SuccessThreshold *float64 `json:"successThreshold" yaml:"successThreshold"`
	// Number of seconds after which the probe times out.
	//
	// Defaults to 1 second. Minimum value is 1.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	//
	TimeoutSeconds cdk8s.Duration `json:"timeoutSeconds" yaml:"timeoutSeconds"`
	// The host name to connect to on the container.
	Host *string `json:"host" yaml:"host"`
	// The TCP port to connect to on the container.
	Port *float64 `json:"port" yaml:"port"`
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
	Name() *string
}

// The jsii proxy struct for Volume
type jsiiProxy_Volume struct {
	_ byte // padding
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


func NewVolume(name *string, config interface{}) Volume {
	_init_.Initialize()

	j := jsiiProxy_Volume{}

	_jsii_.Create(
		"cdk8s-plus-21.Volume",
		[]interface{}{name, config},
		&j,
	)

	return &j
}

func NewVolume_Override(v Volume, name *string, config interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-21.Volume",
		[]interface{}{name, config},
		v,
	)
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
func Volume_FromConfigMap(configMap IConfigMap, options *ConfigMapVolumeOptions) Volume {
	_init_.Initialize()

	var returns Volume

	_jsii_.StaticInvoke(
		"cdk8s-plus-21.Volume",
		"fromConfigMap",
		[]interface{}{configMap, options},
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
func Volume_FromEmptyDir(name *string, options *EmptyDirVolumeOptions) Volume {
	_init_.Initialize()

	var returns Volume

	_jsii_.StaticInvoke(
		"cdk8s-plus-21.Volume",
		"fromEmptyDir",
		[]interface{}{name, options},
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
func Volume_FromSecret(secret ISecret, options *SecretVolumeOptions) Volume {
	_init_.Initialize()

	var returns Volume

	_jsii_.StaticInvoke(
		"cdk8s-plus-21.Volume",
		"fromSecret",
		[]interface{}{secret, options},
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
	Propagation MountPropagation `json:"propagation" yaml:"propagation"`
	// Mounted read-only if true, read-write otherwise (false or unspecified).
	//
	// Defaults to false.
	ReadOnly *bool `json:"readOnly" yaml:"readOnly"`
	// Path within the volume from which the container's volume should be mounted.).
	SubPath *string `json:"subPath" yaml:"subPath"`
	// Expanded path within the volume from which the container's volume should be mounted.
	//
	// Behaves similarly to SubPath but environment variable references
	// $(VAR_NAME) are expanded using the container's environment. Defaults to ""
	// (volume's root).
	//
	// `subPathExpr` and `subPath` are mutually exclusive.
	SubPathExpr *string `json:"subPathExpr" yaml:"subPathExpr"`
	// Path within the container at which the volume should be mounted.
	//
	// Must not
	// contain ':'.
	Path *string `json:"path" yaml:"path"`
	// The volume to mount.
	Volume Volume `json:"volume" yaml:"volume"`
}

