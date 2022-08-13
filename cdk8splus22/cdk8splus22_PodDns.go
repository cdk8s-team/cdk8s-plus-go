// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-22 synthesizes Kubernetes manifests for Kubernetes 1.22.0
package cdk8splus22

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus22/v2/jsii"
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
		"cdk8s-plus-22.PodDns",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewPodDns_Override(p PodDns, props *PodDnsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-22.PodDns",
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

