package k8s


// DisruptionMode defines how individual entities within a group can be disrupted.
//
// Exactly one mode can be set.
type DisruptionModeV1Beta1 struct {
	// all specifies that all children can only be disrupted together.
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// single specifies that children can be disrupted independently from each other.
	Single interface{} `field:"optional" json:"single" yaml:"single"`
}

