package k8s


// addressType specifies the type of address carried by this EndpointSlice.
//
// All addresses in this slice must be the same type. This field is immutable after creation. The following address types are currently supported: * IPv4: Represents an IPv4 Address. * IPv6: Represents an IPv6 Address. * FQDN: Represents a Fully Qualified Domain Name.
//
// Possible enum values:
// - `"FQDN"` represents a FQDN.
// - `"IPv4"` represents an IPv4 Address.
// - `"IPv6"` represents an IPv6 Address.
type IoK8SApiDiscoveryV1EndpointSliceAddressType string

const (
	// FQDN.
	IoK8SApiDiscoveryV1EndpointSliceAddressType_FQDN IoK8SApiDiscoveryV1EndpointSliceAddressType = "FQDN"
	// IPv4.
	IoK8SApiDiscoveryV1EndpointSliceAddressType_I_PV4 IoK8SApiDiscoveryV1EndpointSliceAddressType = "I_PV4"
	// IPv6.
	IoK8SApiDiscoveryV1EndpointSliceAddressType_I_PV6 IoK8SApiDiscoveryV1EndpointSliceAddressType = "I_PV6"
)

