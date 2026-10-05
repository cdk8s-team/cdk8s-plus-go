package k8s


// WorkloadList contains a list of Workload resources.
type KubeWorkloadListV1Alpha1Props struct {
	// Items is the list of Workloads.
	Items *[]*KubeWorkloadV1Alpha1Props `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

