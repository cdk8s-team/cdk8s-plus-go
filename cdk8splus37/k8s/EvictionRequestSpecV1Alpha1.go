package k8s


// EvictionRequestSpec is a specification of an EvictionRequest.
type EvictionRequestSpecV1Alpha1 struct {
	// intent specifies the action that should be taken for the specified target.
	//
	// - Eviction means that the requester is interested in the eviction of the target. - Withdrawn means that the requester is no longer interested in the eviction of the target.
	// If all requesters' intents are withdrawn for a common target, the eviction will be canceled.
	// Cancellation consequences:
	// - Inactive responders will never run.
	// - Active responders are expected to cancel the eviction.
	// - Completed or Interrupted responders should not take any action.
	Intent *string `field:"required" json:"intent" yaml:"intent"`
	// requester allows you to identify the entity, that requested the eviction of the target.
	//
	// It must be a valid domain-prefixed key (such as "acme.io/foo"). Domain names *.k8s.io and *.kubernetes.io are reserved. This field is required and immutable.
	Requester *string `field:"required" json:"requester" yaml:"requester"`
	// target contains a reference to an object (e.g. a pod) that should be evicted. This field is required and immutable.
	Target *EvictionRequestTargetV1Alpha1 `field:"required" json:"target" yaml:"target"`
}

