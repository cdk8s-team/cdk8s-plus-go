package cdk8splus30


// CPU request and limit.
type CpuResources struct {
	Limit Cpu `field:"optional" json:"limit" yaml:"limit"`
	Request Cpu `field:"optional" json:"request" yaml:"request"`
}

