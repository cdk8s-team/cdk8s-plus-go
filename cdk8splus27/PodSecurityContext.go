package cdk8splus27

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus27/v2/jsii"
)

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

	if err := validateNewPodSecurityContextParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_PodSecurityContext{}

	_jsii_.Create(
		"cdk8s-plus-27.PodSecurityContext",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewPodSecurityContext_Override(p PodSecurityContext, props *PodSecurityContextProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-27.PodSecurityContext",
		[]interface{}{props},
		p,
	)
}

