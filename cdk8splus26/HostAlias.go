// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-26 synthesizes Kubernetes manifests for Kubernetes 1.26.0
package cdk8splus26


// HostAlias holds the mapping between IP and hostnames that will be injected as an entry in the pod's /etc/hosts file.
type HostAlias struct {
	// Hostnames for the chosen IP address.
	Hostnames *[]*string `field:"required" json:"hostnames" yaml:"hostnames"`
	// IP address of the host file entry.
	Ip *string `field:"required" json:"ip" yaml:"ip"`
}

