package k8s


// NonResourceAttributes includes the authorization attributes available for non-resource requests to the Authorizer interface.
type NonResourceAttributes struct {
	// path is the URL path of the request.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// verb is the standard HTTP verb.
	Verb *string `field:"optional" json:"verb" yaml:"verb"`
}

