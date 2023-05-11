package cdk8splus25


// HostAlias holds the mapping between IP and hostnames that will be injected as an entry in the pod's /etc/hosts file.
type HostAlias struct {
	// Hostnames for the chosen IP address.
	Hostnames *[]*string `field:"required" json:"hostnames" yaml:"hostnames"`
	// IP address of the host file entry.
	Ip *string `field:"required" json:"ip" yaml:"ip"`
}

