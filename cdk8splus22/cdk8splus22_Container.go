// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-22 synthesizes Kubernetes manifests for Kubernetes 1.22.0
package cdk8splus22

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus22/v2/jsii"
)

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
	// Deprecated: - use `portNumber`.
	Port() *float64
	// The port number that was configured for this container.
	//
	// If undefined, either the container doesn't expose a port, or its
	// port configuration is stored in the `ports` field.
	PortNumber() *float64
	// Ports exposed by this containers.
	//
	// Returns a copy, use `addPort` to modify.
	Ports() *[]*ContainerPort
	// Compute resources (CPU and memory requests and limits) required by the container.
	// See: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	//
	Resources() *ContainerResources
	// The security context of the container.
	SecurityContext() ContainerSecurityContext
	// The working directory inside the container.
	WorkingDir() *string
	// Add a port to expose from this container.
	AddPort(port *ContainerPort)
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

func (j *jsiiProxy_Container) PortNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Container) Ports() *[]*ContainerPort {
	var returns *[]*ContainerPort
	_jsii_.Get(
		j,
		"ports",
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

	if err := validateNewContainerParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Container{}

	_jsii_.Create(
		"cdk8s-plus-22.Container",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewContainer_Override(c Container, props *ContainerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-22.Container",
		[]interface{}{props},
		c,
	)
}

func (c *jsiiProxy_Container) AddPort(port *ContainerPort) {
	if err := c.validateAddPortParameters(port); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPort",
		[]interface{}{port},
	)
}

func (c *jsiiProxy_Container) Mount(path *string, storage IStorage, options *MountOptions) {
	if err := c.validateMountParameters(path, storage, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"mount",
		[]interface{}{path, storage, options},
	)
}

