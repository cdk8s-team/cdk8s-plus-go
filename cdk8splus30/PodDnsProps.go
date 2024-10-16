package cdk8splus30


// Properties for `PodDns`.
type PodDnsProps struct {
	// Specifies the hostname of the Pod.
	// Default: - Set to a system-defined value.
	//
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// If true the pod's hostname will be configured as the pod's FQDN, rather than the leaf name (the default).
	//
	// In Linux containers, this means setting the FQDN in the hostname field of the kernel (the nodename field of struct utsname).
	// In Windows containers, this means setting the registry value of hostname for the registry
	// key HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Services\Tcpip\Parameters to FQDN.
	// If a pod does not have FQDN, this has no effect.
	// Default: false.
	//
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
	// Default: DnsPolicy.CLUSTER_FIRST
	//
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
	// Default: - No subdomain.
	//
	Subdomain *string `field:"optional" json:"subdomain" yaml:"subdomain"`
}

