package k8s


// ParamKind is a tuple of Group Kind and Version.
type ParamKindV1Beta1 struct {
	// apiVersion is the API group version the resources belong to.
	//
	// In format of "group/version". Required.
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// kind is the API kind the resources belong to.
	//
	// Required.
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
}

