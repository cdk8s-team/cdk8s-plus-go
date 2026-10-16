package k8s


// ResourcePoolStatusRequestSpec defines the filters for the pool status request.
type ResourcePoolStatusRequestSpecV1Alpha3 struct {
	// Driver specifies the DRA driver name to filter pools.
	//
	// Only pools from ResourceSlices with this driver will be included. Must be a DNS subdomain (e.g., "gpu.example.com").
	Driver *string `field:"required" json:"driver" yaml:"driver"`
	// DefaultPartitionTypeAttribute optionally names a device attribute (by its fully qualified name, e.g. "gpu.example.com/profile") to use as the default grouping attribute for partitionable devices whose slice has not declared one themselves.
	//
	// A slice's own PartitionTypeAttribute always takes precedence. This default applies only to devices whose slice does not declare one, so that a request can still get an accurate partitionSummary from a driver that has not been updated to declare it. When neither the slice nor this default names an attribute, a partitionable pool reports no partitionSummary.
	//
	// Must include the domain qualifier.
	DefaultPartitionTypeAttribute *string `field:"optional" json:"defaultPartitionTypeAttribute" yaml:"defaultPartitionTypeAttribute"`
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

