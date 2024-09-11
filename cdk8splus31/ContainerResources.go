package cdk8splus31


// CPU and memory compute resources.
type ContainerResources struct {
	Cpu *CpuResources `field:"optional" json:"cpu" yaml:"cpu"`
	EphemeralStorage *EphemeralStorageResources `field:"optional" json:"ephemeralStorage" yaml:"ephemeralStorage"`
	Memory *MemoryResources `field:"optional" json:"memory" yaml:"memory"`
}

