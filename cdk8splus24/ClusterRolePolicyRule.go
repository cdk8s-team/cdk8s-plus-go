package cdk8splus24


// Policy rule of a `ClusterRole.
type ClusterRolePolicyRule struct {
	// Endpoints this rule applies to.
	//
	// Can be either api resources
	// or non api resources.
	Endpoints *[]IApiEndpoint `field:"required" json:"endpoints" yaml:"endpoints"`
	// Verbs to allow.
	//
	// (e.g ['get', 'watch'])
	Verbs *[]*string `field:"required" json:"verbs" yaml:"verbs"`
}

