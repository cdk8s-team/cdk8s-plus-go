package k8s


// Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.
//
// Possible enum values:
// - `"DoesNotExist"`
// - `"Exists"`
// - `"Gt"`
// - `"In"`
// - `"Lt"`
// - `"NotIn"`.
type IoK8SApiCoreV1NodeSelectorRequirementOperator string

const (
	// DoesNotExist.
	IoK8SApiCoreV1NodeSelectorRequirementOperator_DOES_NOT_EXIST IoK8SApiCoreV1NodeSelectorRequirementOperator = "DOES_NOT_EXIST"
	// Exists.
	IoK8SApiCoreV1NodeSelectorRequirementOperator_EXISTS IoK8SApiCoreV1NodeSelectorRequirementOperator = "EXISTS"
	// Gt.
	IoK8SApiCoreV1NodeSelectorRequirementOperator_GT IoK8SApiCoreV1NodeSelectorRequirementOperator = "GT"
	// In.
	IoK8SApiCoreV1NodeSelectorRequirementOperator_IN IoK8SApiCoreV1NodeSelectorRequirementOperator = "IN"
	// Lt.
	IoK8SApiCoreV1NodeSelectorRequirementOperator_LT IoK8SApiCoreV1NodeSelectorRequirementOperator = "LT"
	// NotIn.
	IoK8SApiCoreV1NodeSelectorRequirementOperator_NOT_IN IoK8SApiCoreV1NodeSelectorRequirementOperator = "NOT_IN"
)

