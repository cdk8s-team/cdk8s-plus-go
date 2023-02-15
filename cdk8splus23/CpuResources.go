// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-23 synthesizes Kubernetes manifests for Kubernetes 1.23.0
package cdk8splus23


// CPU request and limit.
type CpuResources struct {
	Limit Cpu `field:"optional" json:"limit" yaml:"limit"`
	Request Cpu `field:"optional" json:"request" yaml:"request"`
}

