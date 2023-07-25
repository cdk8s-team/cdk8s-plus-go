package k8s


// PodSchedulingContext objects hold information that is needed to schedule a Pod with ResourceClaims that use "WaitForFirstConsumer" allocation mode.
//
// This is an alpha type and requires enabling the DynamicResourceAllocation feature gate.
type KubePodSchedulingContextV1Alpha2Props struct {
	// Spec describes where resources for the Pod are needed.
	Spec *PodSchedulingContextSpecV1Alpha2 `field:"required" json:"spec" yaml:"spec"`
	// Standard object metadata.
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

