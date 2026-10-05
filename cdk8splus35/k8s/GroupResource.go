package k8s


// GroupResource specifies a Group and a Resource, but does not force a version.
//
// This is useful for identifying concepts during lookup stages without having partially valid types.
type GroupResource struct {
	Group *string `field:"required" json:"group" yaml:"group"`
	Resource *string `field:"required" json:"resource" yaml:"resource"`
}

