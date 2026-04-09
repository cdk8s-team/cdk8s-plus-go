package cdk8splus34


type ContainerSecutiryContextCapabilities struct {
	// Added capabilities.
	Add *[]Capability `field:"optional" json:"add" yaml:"add"`
	// Removed capabilities.
	Drop *[]Capability `field:"optional" json:"drop" yaml:"drop"`
}

