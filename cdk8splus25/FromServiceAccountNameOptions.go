package cdk8splus25


type FromServiceAccountNameOptions struct {
	// The name of the namespace the service account belongs to.
	// Default: "default".
	//
	NamespaceName *string `field:"optional" json:"namespaceName" yaml:"namespaceName"`
}

