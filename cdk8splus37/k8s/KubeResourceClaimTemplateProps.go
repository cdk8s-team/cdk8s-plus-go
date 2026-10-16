package k8s


// ResourceClaimTemplate is used to produce ResourceClaim objects.
type KubeResourceClaimTemplateProps struct {
	// Standard object metadata.
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
	// Describes the ResourceClaim that is to be generated.
	//
	// This field is immutable. A ResourceClaim will get created by the control plane for a Pod when needed and then not get updated anymore.
	Spec *ResourceClaimTemplateSpec `field:"optional" json:"spec" yaml:"spec"`
}

