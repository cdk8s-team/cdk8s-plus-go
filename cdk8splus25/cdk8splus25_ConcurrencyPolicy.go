// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-25 synthesizes Kubernetes manifests for Kubernetes 1.25.0
package cdk8splus25


// Concurrency policy for CronJobs.
type ConcurrencyPolicy string

const (
	// This policy allows to run job concurrently.
	ConcurrencyPolicy_ALLOW ConcurrencyPolicy = "ALLOW"
	// This policy does not allow to run job concurrently.
	//
	// It does not let a new job to be scheduled if the previous one is not finished yet.
	ConcurrencyPolicy_FORBID ConcurrencyPolicy = "FORBID"
	// This policy replaces the currently running job if a new job is being scheduled.
	ConcurrencyPolicy_REPLACE ConcurrencyPolicy = "REPLACE"
)

