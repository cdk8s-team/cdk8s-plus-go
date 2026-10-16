package k8s


// WorkloadList contains a list of Workload resources.
type KubeWorkloadListV1Alpha3Props struct {
	// Items is the list of Workloads.
	Items *[]*KubeWorkloadV1Alpha3Props `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

