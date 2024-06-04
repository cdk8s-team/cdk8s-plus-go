package k8s


// IPAddressList contains a list of IPAddress.
type KubeIpAddressListV1Alpha1Props struct {
	// items is the list of IPAddresses.
	Items *[]*KubeIpAddressV1Alpha1Props `field:"required" json:"items" yaml:"items"`
	// Standard object's metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

