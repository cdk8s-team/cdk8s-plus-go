package k8s


// ResourcePoolStatusRequestSpec defines the filters for the pool status request.
type ResourcePoolStatusRequestSpecV1Alpha3 struct {
	// Driver specifies the DRA driver name to filter pools.
	//
	// Only pools from ResourceSlices with this driver will be included. Must be a DNS subdomain (e.g., "gpu.example.com").
	Driver *string `field:"required" json:"driver" yaml:"driver"`
	// Limit optionally specifies the maximum number of pools to return in the status.
	//
	// If more pools match the filter criteria, the response will be truncated (i.e., len(status.pools) < status.poolCount).
	//
	// Default: 100 Minimum: 1 Maximum: 1000.
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// PoolName optionally filters to a specific pool name.
	//
	// If not specified, all pools from the specified driver are included. When specified, must be a non-empty valid resource pool name (DNS subdomains separated by "/").
	PoolName *string `field:"optional" json:"poolName" yaml:"poolName"`
}

