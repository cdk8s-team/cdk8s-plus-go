package cdk8splus30

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Properties for `CronJob`.
type CronJobProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// Indicates whether a service account token should be automatically mounted.
	// See: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/#use-the-default-service-account-to-access-the-api-server
	//
	// Default: false.
	//
	AutomountServiceAccountToken *bool `field:"optional" json:"automountServiceAccountToken" yaml:"automountServiceAccountToken"`
	// List of containers belonging to the pod.
	//
	// Containers cannot currently be
	// added or removed. There must be at least one container in a Pod.
	//
	// You can add additionnal containers using `podSpec.addContainer()`
	// Default: - No containers. Note that a pod spec must include at least one container.
	//
	Containers *[]*ContainerProps `field:"optional" json:"containers" yaml:"containers"`
	// DNS settings for the pod.
	// See: https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/
	//
	// Default:  policy: DnsPolicy.CLUSTER_FIRST
	// hostnameAsFQDN: false.
	//
	Dns *PodDnsProps `field:"optional" json:"dns" yaml:"dns"`
	// A secret containing docker credentials for authenticating to a registry.
	// Default: - No auth. Images are assumed to be publicly available.
	//
	DockerRegistryAuth ISecret `field:"optional" json:"dockerRegistryAuth" yaml:"dockerRegistryAuth"`
	// HostAlias holds the mapping between IP and hostnames that will be injected as an entry in the pod's hosts file.
	HostAliases *[]*HostAlias `field:"optional" json:"hostAliases" yaml:"hostAliases"`
	// Host network for the pod.
	// Default: false.
	//
	HostNetwork *bool `field:"optional" json:"hostNetwork" yaml:"hostNetwork"`
	// List of initialization containers belonging to the pod.
	//
	// Init containers are executed in order prior to containers being started.
	// If any init container fails, the pod is considered to have failed and is handled according to its restartPolicy.
	// The name for an init container or normal container must be unique among all containers.
	// Init containers may not have Lifecycle actions, Readiness probes, Liveness probes, or Startup probes.
	// The resourceRequirements of an init container are taken into account during scheduling by finding the highest request/limit
	// for each resource type, and then using the max of of that value or the sum of the normal containers.
	// Limits are applied to init containers in a similar fashion.
	//
	// Init containers cannot currently be added ,removed or updated.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/
	//
	// Default: - No init containers.
	//
	InitContainers *[]*ContainerProps `field:"optional" json:"initContainers" yaml:"initContainers"`
	// Isolates the pod.
	//
	// This will prevent any ingress or egress connections to / from this pod.
	// You can however allow explicit connections post instantiation by using the `.connections` property.
	// Default: false.
	//
	Isolate *bool `field:"optional" json:"isolate" yaml:"isolate"`
	// Restart policy for all containers within the pod.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy
	//
	// Default: RestartPolicy.ALWAYS
	//
	RestartPolicy RestartPolicy `field:"optional" json:"restartPolicy" yaml:"restartPolicy"`
	// SecurityContext holds pod-level security attributes and common container settings.
	// Default:   fsGroupChangePolicy: FsGroupChangePolicy.FsGroupChangePolicy.ALWAYS
	// ensureNonRoot: true.
	//
	SecurityContext *PodSecurityContextProps `field:"optional" json:"securityContext" yaml:"securityContext"`
	// A service account provides an identity for processes that run in a Pod.
	//
	// When you (a human) access the cluster (for example, using kubectl), you are
	// authenticated by the apiserver as a particular User Account (currently this
	// is usually admin, unless your cluster administrator has customized your
	// cluster). Processes in containers inside pods can also contact the
	// apiserver. When they do, they are authenticated as a particular Service
	// Account (for example, default).
	// See: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
	//
	// Default: - No service account.
	//
	ServiceAccount IServiceAccount `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// Grace period until the pod is terminated.
	// Default: Duration.seconds(30)
	//
	TerminationGracePeriod cdk8s.Duration `field:"optional" json:"terminationGracePeriod" yaml:"terminationGracePeriod"`
	// List of volumes that can be mounted by containers belonging to the pod.
	//
	// You can also add volumes later using `podSpec.addVolume()`
	// See: https://kubernetes.io/docs/concepts/storage/volumes
	//
	// Default: - No volumes.
	//
	Volumes *[]Volume `field:"optional" json:"volumes" yaml:"volumes"`
	// The pod metadata of this workload.
	PodMetadata *cdk8s.ApiObjectMetadata `field:"optional" json:"podMetadata" yaml:"podMetadata"`
	// Automatically allocates a pod label selector for this workload and add it to the pod metadata.
	//
	// This ensures this workload manages pods created by
	// its pod template.
	// Default: true.
	//
	Select *bool `field:"optional" json:"select" yaml:"select"`
	// Automatically spread pods across hostname and zones.
	// See: https://kubernetes.io/docs/concepts/scheduling-eviction/topology-spread-constraints/#internal-default-constraints
	//
	// Default: false.
	//
	Spread *bool `field:"optional" json:"spread" yaml:"spread"`
	// Specifies the duration the job may be active before the system tries to terminate it.
	// Default: - If unset, then there is no deadline.
	//
	ActiveDeadline cdk8s.Duration `field:"optional" json:"activeDeadline" yaml:"activeDeadline"`
	// Specifies the number of retries before marking this job failed.
	// Default: - If not set, system defaults to 6.
	//
	BackoffLimit *float64 `field:"optional" json:"backoffLimit" yaml:"backoffLimit"`
	// Limits the lifetime of a Job that has finished execution (either Complete or Failed).
	//
	// If this field is set, after the Job finishes, it is eligible to
	// be automatically deleted. When the Job is being deleted, its lifecycle
	// guarantees (e.g. finalizers) will be honored. If this field is set to zero,
	// the Job becomes eligible to be deleted immediately after it finishes. This
	// field is alpha-level and is only honored by servers that enable the
	// `TTLAfterFinished` feature.
	// Default: - If this field is unset, the Job won't be automatically deleted.
	//
	TtlAfterFinished cdk8s.Duration `field:"optional" json:"ttlAfterFinished" yaml:"ttlAfterFinished"`
	// Specifies the time in which the job would run again.
	//
	// This is defined as a cron expression in the CronJob resource.
	Schedule cdk8s.Cron `field:"required" json:"schedule" yaml:"schedule"`
	// Specifies the concurrency policy for the job.
	// Default: ConcurrencyPolicy.Forbid
	//
	ConcurrencyPolicy ConcurrencyPolicy `field:"optional" json:"concurrencyPolicy" yaml:"concurrencyPolicy"`
	// Specifies the number of failed jobs history retained.
	//
	// This would retain the Job and the associated Pod resource and can be useful for debugging.
	// Default: 1.
	//
	FailedJobsRetained *float64 `field:"optional" json:"failedJobsRetained" yaml:"failedJobsRetained"`
	// Kubernetes attempts to start cron jobs at its schedule time, but this is not guaranteed.
	//
	// This deadline specifies
	// how much time can pass after a schedule point, for which kubernetes can still start the job.
	// For example, if this is set to 100 seconds, kubernetes is allowed to start the job at a maximum 100 seconds after
	// the scheduled time.
	//
	// Note that the Kubernetes CronJobController checks for things every 10 seconds, for this reason, a deadline below 10
	// seconds is not allowed, as it may cause your job to never be scheduled.
	//
	// In addition, kubernetes will stop scheduling jobs if more than 100 schedules were missed (for any reason).
	// This property also controls what time interval should kubernetes consider when counting for missed schedules.
	//
	// For example, suppose a CronJob is set to schedule a new Job every one minute beginning at 08:30:00,
	// and its `startingDeadline` field is not set. If the CronJob controller happens to be down from 08:29:00 to 10:21:00,
	// the job will not start as the number of missed jobs which missed their schedule is greater than 100.
	// However, if `startingDeadline` is set to 200 seconds, kubernetes will only count 3 missed schedules, and thus
	// start a new execution at 10:22:00.
	// Default: Duration.seconds(10)
	//
	StartingDeadline cdk8s.Duration `field:"optional" json:"startingDeadline" yaml:"startingDeadline"`
	// Specifies the number of successful jobs history retained.
	//
	// This would retain the Job and the associated Pod resource and can be useful for debugging.
	// Default: 3.
	//
	SuccessfulJobsRetained *float64 `field:"optional" json:"successfulJobsRetained" yaml:"successfulJobsRetained"`
	// Specifies if the cron job should be suspended.
	//
	// Only applies to future executions, current ones are remained untouched.
	// Default: false.
	//
	Suspend *bool `field:"optional" json:"suspend" yaml:"suspend"`
	// Specifies the timezone for the job.
	//
	// This helps aligining the schedule to follow the specified timezone.
	// See: {@link https://en.wikipedia.org/wiki/List_of_tz_database_time_zones} for list of valid timezone values.
	//
	// Default: - Timezone of kube-controller-manager process.
	//
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

