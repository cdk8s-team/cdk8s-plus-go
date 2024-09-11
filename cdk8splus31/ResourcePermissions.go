package cdk8splus31

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus31/v2/jsii"
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

	if err := validateNewResourcePermissionsParameters(instance); err != nil {
		panic(err)
	}
	j := jsiiProxy_ResourcePermissions{}

	_jsii_.Create(
		"cdk8s-plus-31.ResourcePermissions",
		[]interface{}{instance},
		&j,
	)

	return &j
}

func NewResourcePermissions_Override(r ResourcePermissions, instance Resource) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-31.ResourcePermissions",
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

