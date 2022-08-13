// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24


// CPU and memory compute resources.
type ContainerResources struct {
	Cpu *CpuResources `field:"optional" json:"cpu" yaml:"cpu"`
	Memory *MemoryResources `field:"optional" json:"memory" yaml:"memory"`
}

