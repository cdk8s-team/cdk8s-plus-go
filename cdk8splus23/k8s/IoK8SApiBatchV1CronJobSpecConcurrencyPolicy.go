package k8s


// Specifies how to treat concurrent executions of a Job.
//
// Valid values are: - "Allow" (default): allows CronJobs to run concurrently; - "Forbid": forbids concurrent runs, skipping next run if previous run hasn't finished yet; - "Replace": cancels currently running job and replaces it with a new one
//
// Possible enum values:
// - `"Allow"` allows CronJobs to run concurrently.
// - `"Forbid"` forbids concurrent runs, skipping next run if previous hasn't finished yet.
// - `"Replace"` cancels currently running job and replaces it with a new one.
type IoK8SApiBatchV1CronJobSpecConcurrencyPolicy string

const (
	// Allow.
	IoK8SApiBatchV1CronJobSpecConcurrencyPolicy_ALLOW IoK8SApiBatchV1CronJobSpecConcurrencyPolicy = "ALLOW"
	// Forbid.
	IoK8SApiBatchV1CronJobSpecConcurrencyPolicy_FORBID IoK8SApiBatchV1CronJobSpecConcurrencyPolicy = "FORBID"
	// Replace.
	IoK8SApiBatchV1CronJobSpecConcurrencyPolicy_REPLACE IoK8SApiBatchV1CronJobSpecConcurrencyPolicy = "REPLACE"
)

