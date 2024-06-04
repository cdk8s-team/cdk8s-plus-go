package k8s


// DriverRequests describes all resources that are needed from one particular driver.
type DriverRequestsV1Alpha2 struct {
	// DriverName is the name used by the DRA driver kubelet plugin.
	DriverName *string `field:"optional" json:"driverName" yaml:"driverName"`
	// Requests describes all resources that are needed from the driver.
	Requests *[]*ResourceRequestV1Alpha2 `field:"optional" json:"requests" yaml:"requests"`
	// VendorParameters are arbitrary setup parameters for all requests of the claim.
	//
	// They are ignored while allocating the claim.
	VendorParameters interface{} `field:"optional" json:"vendorParameters" yaml:"vendorParameters"`
}

