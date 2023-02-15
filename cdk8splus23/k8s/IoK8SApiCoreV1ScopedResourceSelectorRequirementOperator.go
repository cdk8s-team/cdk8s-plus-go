package k8s


// Represents a scope's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist.
//
// Possible enum values:
// - `"DoesNotExist"`
// - `"Exists"`
// - `"In"`
// - `"NotIn"`.
type IoK8SApiCoreV1ScopedResourceSelectorRequirementOperator string

const (
	// DoesNotExist.
	IoK8SApiCoreV1ScopedResourceSelectorRequirementOperator_DOES_NOT_EXIST IoK8SApiCoreV1ScopedResourceSelectorRequirementOperator = "DOES_NOT_EXIST"
	// Exists.
	IoK8SApiCoreV1ScopedResourceSelectorRequirementOperator_EXISTS IoK8SApiCoreV1ScopedResourceSelectorRequirementOperator = "EXISTS"
	// In.
	IoK8SApiCoreV1ScopedResourceSelectorRequirementOperator_IN IoK8SApiCoreV1ScopedResourceSelectorRequirementOperator = "IN"
	// NotIn.
	IoK8SApiCoreV1ScopedResourceSelectorRequirementOperator_NOT_IN IoK8SApiCoreV1ScopedResourceSelectorRequirementOperator = "NOT_IN"
)

