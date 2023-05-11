package cdk8splus26


// Options for exposing a service using an ingress.
type ExposeServiceViaIngressOptions struct {
	// The ingress to add rules to.
	Ingress Ingress `field:"optional" json:"ingress" yaml:"ingress"`
	// The type of the path.
	PathType HttpIngressPathType `field:"optional" json:"pathType" yaml:"pathType"`
}

