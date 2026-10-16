package k8s


// Role is a namespaced, logical grouping of PolicyRules that can be referenced as a unit by a RoleBinding.
type KubeRoleProps struct {
	// metadata is the standard object's metadata.
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
	// rules holds all the PolicyRules for this Role.
	Rules *[]*PolicyRule `field:"optional" json:"rules" yaml:"rules"`
}

