package k8s


// CompositeDisruptionMode defines how individual entities within a composite pod group can be disrupted.
//
// Exactly one mode must be set.
type CompositeDisruptionModeV1Beta1 struct {
	// all specifies that all children groups can only be disrupted together.
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// single specifies that children groups can be disrupted independently from each other.
	Single interface{} `field:"optional" json:"single" yaml:"single"`
}

