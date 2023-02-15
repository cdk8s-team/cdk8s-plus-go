// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24


// CPU request and limit.
type CpuResources struct {
	Limit Cpu `field:"optional" json:"limit" yaml:"limit"`
	Request Cpu `field:"optional" json:"request" yaml:"request"`
}

