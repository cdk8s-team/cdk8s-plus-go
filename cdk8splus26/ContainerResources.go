// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-26 synthesizes Kubernetes manifests for Kubernetes 1.26.0
package cdk8splus26


// CPU and memory compute resources.
type ContainerResources struct {
	Cpu *CpuResources `field:"optional" json:"cpu" yaml:"cpu"`
	EphemeralStorage *EphemeralStorageResources `field:"optional" json:"ephemeralStorage" yaml:"ephemeralStorage"`
	Memory *MemoryResources `field:"optional" json:"memory" yaml:"memory"`
}

