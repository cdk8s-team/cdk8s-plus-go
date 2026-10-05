package k8s


// DeviceClaimConfiguration is used for configuration parameters in DeviceClaim.
type DeviceClaimConfigurationV1Beta1 struct {
	// Opaque provides driver-specific configuration parameters.
	Opaque *OpaqueDeviceConfigurationV1Beta1 `field:"optional" json:"opaque" yaml:"opaque"`
	// Requests lists the names of requests where the configuration applies. If empty, it applies to all requests.
	//
	// References to subrequests must include the name of the main request and may include the subrequest using the format <main request>[/<subrequest>]. If just the main request is given, the configuration applies to all subrequests.
	Requests *[]*string `field:"optional" json:"requests" yaml:"requests"`
}

