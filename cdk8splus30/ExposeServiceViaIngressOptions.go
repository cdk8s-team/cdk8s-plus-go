package cdk8splus30


// Options for exposing a service using an ingress.
type ExposeServiceViaIngressOptions struct {
	// The ingress to add rules to.
	// Default: - An ingress will be automatically created.
	//
	Ingress Ingress `field:"optional" json:"ingress" yaml:"ingress"`
	// The type of the path.
	// Default: HttpIngressPathType.PREFIX
	//
	PathType HttpIngressPathType `field:"optional" json:"pathType" yaml:"pathType"`
}

