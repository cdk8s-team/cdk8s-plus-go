package k8s


// The name of the scope that the selector applies to.
//
// Possible enum values:
// - `"BestEffort"` Match all pod objects that have best effort quality of service
// - `"CrossNamespacePodAffinity"` Match all pod objects that have cross-namespace pod (anti)affinity mentioned. This is a beta feature enabled by the PodAffinityNamespaceSelector feature flag.
// - `"NotBestEffort"` Match all pod objects that do not have best effort quality of service
// - `"NotTerminating"` Match all pod objects where spec.activeDeadlineSeconds is nil
// - `"PriorityClass"` Match all pod objects that have priority class mentioned
// - `"Terminating"` Match all pod objects where spec.activeDeadlineSeconds >=0
type IoK8SApiCoreV1ScopedResourceSelectorRequirementScopeName string

const (
	// BestEffort.
	IoK8SApiCoreV1ScopedResourceSelectorRequirementScopeName_BEST_EFFORT IoK8SApiCoreV1ScopedResourceSelectorRequirementScopeName = "BEST_EFFORT"
	// CrossNamespacePodAffinity.
	IoK8SApiCoreV1ScopedResourceSelectorRequirementScopeName_CROSS_NAMESPACE_POD_AFFINITY IoK8SApiCoreV1ScopedResourceSelectorRequirementScopeName = "CROSS_NAMESPACE_POD_AFFINITY"
	// NotBestEffort.
	IoK8SApiCoreV1ScopedResourceSelectorRequirementScopeName_NOT_BEST_EFFORT IoK8SApiCoreV1ScopedResourceSelectorRequirementScopeName = "NOT_BEST_EFFORT"
	// NotTerminating.
	IoK8SApiCoreV1ScopedResourceSelectorRequirementScopeName_NOT_TERMINATING IoK8SApiCoreV1ScopedResourceSelectorRequirementScopeName = "NOT_TERMINATING"
	// PriorityClass.
	IoK8SApiCoreV1ScopedResourceSelectorRequirementScopeName_PRIORITY_CLASS IoK8SApiCoreV1ScopedResourceSelectorRequirementScopeName = "PRIORITY_CLASS"
	// Terminating.
	IoK8SApiCoreV1ScopedResourceSelectorRequirementScopeName_TERMINATING IoK8SApiCoreV1ScopedResourceSelectorRequirementScopeName = "TERMINATING"
)

