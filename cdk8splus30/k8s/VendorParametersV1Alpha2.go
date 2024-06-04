package k8s


// VendorParameters are opaque parameters for one particular driver.
type VendorParametersV1Alpha2 struct {
	// DriverName is the name used by the DRA driver kubelet plugin.
	DriverName *string `field:"optional" json:"driverName" yaml:"driverName"`
	// Parameters can be arbitrary setup parameters.
	//
	// They are ignored while allocating a claim.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

