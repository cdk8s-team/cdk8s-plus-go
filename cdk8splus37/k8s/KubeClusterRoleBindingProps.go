package k8s


// ClusterRoleBinding references a ClusterRole, but not contain it.
//
// It can reference a ClusterRole in the global namespace, and adds who information via Subject.
type KubeClusterRoleBindingProps struct {
	// roleRef can only reference a ClusterRole in the global namespace.
	//
	// If the RoleRef cannot be resolved, the Authorizer must return an error. This field is immutable.
	RoleRef *RoleRef `field:"required" json:"roleRef" yaml:"roleRef"`
	// metadata is the standard object's metadata.
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
	// subjects holds references to the objects the role applies to.
	Subjects *[]*Subject `field:"optional" json:"subjects" yaml:"subjects"`
}

