package k8s


// EvictionRequest defines a request that should ideally result in a graceful eviction of a .spec.target (e.g. termination of a pod).
//
// The evictionrequest-controller observes intents of all EvictionRequests and transforms them into Evictions.
// - .spec.requester is set as a label on the Eviction for easier lookup.
// - Each target can have a set of responders assigned to it. Eviction objects are observed by
// these responders, who implement the eviction logic and update the Eviction's status with
// progress.
//
// There is many-to-many relationship between EvictionRequests and Evictions in general. And many-to-one if the target is a  pod.
//
// If all requesters withdraw their eviction intent for a common target, the eviction will be canceled. Deleting an EvictionRequest also counts as a withdrawal. Once all EvictionRequest of a target are removed, the corresponding Evictions are eventually garbage collected.
type KubeEvictionRequestV1Alpha1Props struct {
	// spec defines the eviction request specification.
	//
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *EvictionRequestSpecV1Alpha1 `field:"required" json:"spec" yaml:"spec"`
	// metadata is the standard object metadata;
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

