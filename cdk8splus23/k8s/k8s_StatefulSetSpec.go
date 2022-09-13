package k8s


// A StatefulSetSpec is the specification of a StatefulSet.
type StatefulSetSpec struct {
	// selector is a label query over pods that should match the replica count.
	//
	// It must match the pod template's labels. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
	Selector *LabelSelector `field:"required" json:"selector" yaml:"selector"`
	// serviceName is the name of the service that governs this StatefulSet.
	//
	// This service must exist before the StatefulSet, and is responsible for the network identity of the set. Pods get DNS/hostnames that follow the pattern: pod-specific-string.serviceName.default.svc.cluster.local where "pod-specific-string" is managed by the StatefulSet controller.
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	// template is the object that describes the pod that will be created if insufficient replicas are detected.
	//
	// Each pod stamped out by the StatefulSet will fulfill this Template, but have a unique identity from the rest of the StatefulSet.
	Template *PodTemplateSpec `field:"required" json:"template" yaml:"template"`
	// Minimum number of seconds for which a newly created pod should be ready without any of its container crashing for it to be considered available.
	//
	// Defaults to 0 (pod will be considered available as soon as it is ready) This is an alpha field and requires enabling StatefulSetMinReadySeconds feature gate.
	MinReadySeconds *float64 `field:"optional" json:"minReadySeconds" yaml:"minReadySeconds"`
	// persistentVolumeClaimRetentionPolicy describes the lifecycle of persistent volume claims created from volumeClaimTemplates.
	//
	// By default, all persistent volume claims are created as needed and retained until manually deleted. This policy allows the lifecycle to be altered, for example by deleting persistent volume claims when their stateful set is deleted, or when their pod is scaled down. This requires the StatefulSetAutoDeletePVC feature gate to be enabled, which is alpha.  +optional
	PersistentVolumeClaimRetentionPolicy *StatefulSetPersistentVolumeClaimRetentionPolicy `field:"optional" json:"persistentVolumeClaimRetentionPolicy" yaml:"persistentVolumeClaimRetentionPolicy"`
	// podManagementPolicy controls how pods are created during initial scale up, when replacing pods on nodes, or when scaling down.
	//
	// The default policy is `OrderedReady`, where pods are created in increasing order (pod-0, then pod-1, etc) and the controller will wait until each pod is ready before continuing. When scaling down, the pods are removed in the opposite order. The alternative policy is `Parallel` which will create pods in parallel to match the desired scale without waiting, and on scale down will delete all pods at once.
	//
	// Possible enum values:
	// - `"OrderedReady"` will create pods in strictly increasing order on scale up and strictly decreasing order on scale down, progressing only when the previous pod is ready or terminated. At most one pod will be changed at any time.
	// - `"Parallel"` will create and delete pods as soon as the stateful set replica count is changed, and will not wait for pods to be ready or complete termination.
	PodManagementPolicy IoK8SApiAppsV1StatefulSetSpecPodManagementPolicy `field:"optional" json:"podManagementPolicy" yaml:"podManagementPolicy"`
	// replicas is the desired number of replicas of the given Template.
	//
	// These are replicas in the sense that they are instantiations of the same Template, but individual replicas also have a consistent identity. If unspecified, defaults to 1.
	Replicas *float64 `field:"optional" json:"replicas" yaml:"replicas"`
	// revisionHistoryLimit is the maximum number of revisions that will be maintained in the StatefulSet's revision history.
	//
	// The revision history consists of all revisions not represented by a currently applied StatefulSetSpec version. The default value is 10.
	RevisionHistoryLimit *float64 `field:"optional" json:"revisionHistoryLimit" yaml:"revisionHistoryLimit"`
	// updateStrategy indicates the StatefulSetUpdateStrategy that will be employed to update Pods in the StatefulSet when a revision is made to Template.
	UpdateStrategy *StatefulSetUpdateStrategy `field:"optional" json:"updateStrategy" yaml:"updateStrategy"`
	// volumeClaimTemplates is a list of claims that pods are allowed to reference.
	//
	// The StatefulSet controller is responsible for mapping network identities to claims in a way that maintains the identity of a pod. Every claim in this list must have at least one matching (by name) volumeMount in one container in the template. A claim in this list takes precedence over any volumes in the template, with the same name.
	VolumeClaimTemplates *[]*KubePersistentVolumeClaimProps `field:"optional" json:"volumeClaimTemplates" yaml:"volumeClaimTemplates"`
}

