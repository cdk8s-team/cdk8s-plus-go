package k8s


// WorkloadList contains a list of Workload resources.
type KubeWorkloadListV1Beta1Props struct {
	// items is the list of Workloads.
	Items *[]*KubeWorkloadV1Beta1Props `field:"required" json:"items" yaml:"items"`
	// metadata is the standard list metadata.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

