package k8s


// Operator represents a key's relationship to the value.
//
// Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.
//
// Possible enum values:
// - `"Equal"`
// - `"Exists"`.
type IoK8SApiCoreV1TolerationOperator string

const (
	// Equal.
	IoK8SApiCoreV1TolerationOperator_EQUAL IoK8SApiCoreV1TolerationOperator = "EQUAL"
	// Exists.
	IoK8SApiCoreV1TolerationOperator_EXISTS IoK8SApiCoreV1TolerationOperator = "EXISTS"
)

