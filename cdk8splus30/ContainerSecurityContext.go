package cdk8splus30

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2/jsii"
)

// Container security attributes and settings.
type ContainerSecurityContext interface {
	AllowPrivilegeEscalation() *bool
	Capabilities() *ContainerSecutiryContextCapabilities
	EnsureNonRoot() *bool
	Group() *float64
	Privileged() *bool
	ReadOnlyRootFilesystem() *bool
	SeccompProfile() *SeccompProfile
	User() *float64
}

// The jsii proxy struct for ContainerSecurityContext
type jsiiProxy_ContainerSecurityContext struct {
	_ byte // padding
}

func (j *jsiiProxy_ContainerSecurityContext) AllowPrivilegeEscalation() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"allowPrivilegeEscalation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerSecurityContext) Capabilities() *ContainerSecutiryContextCapabilities {
	var returns *ContainerSecutiryContextCapabilities
	_jsii_.Get(
		j,
		"capabilities",
		&returns,
	)
	return returns
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

func (j *jsiiProxy_ContainerSecurityContext) SeccompProfile() *SeccompProfile {
	var returns *SeccompProfile
	_jsii_.Get(
		j,
		"seccompProfile",
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

	if err := validateNewContainerSecurityContextParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerSecurityContext{}

	_jsii_.Create(
		"cdk8s-plus-30.ContainerSecurityContext",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewContainerSecurityContext_Override(c ContainerSecurityContext, props *ContainerSecurityContextProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-30.ContainerSecurityContext",
		[]interface{}{props},
		c,
	)
}

