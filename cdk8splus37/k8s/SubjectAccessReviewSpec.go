package k8s


// SubjectAccessReviewSpec is a description of the access request.
//
// Exactly one of resourceAttributes and nonResourceAttributes must be set.
type SubjectAccessReviewSpec struct {
	// extra corresponds to the user.Info.GetExtra() method from the authenticator.  Since that is input to the authorizer it needs a reflection here.
	Extra *map[string]*[]*string `field:"optional" json:"extra" yaml:"extra"`
	// groups is the groups you're testing for.
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// nonResourceAttributes describes information for a non-resource access request.
	NonResourceAttributes *NonResourceAttributes `field:"optional" json:"nonResourceAttributes" yaml:"nonResourceAttributes"`
	// resourceAttributes describes information for a resource access request.
	ResourceAttributes *ResourceAttributes `field:"optional" json:"resourceAttributes" yaml:"resourceAttributes"`
	// uid information about the requesting user.
	Uid *string `field:"optional" json:"uid" yaml:"uid"`
	// user is the user you're testing for.
	//
	// If you specify "User" but not "Groups", then is it interpreted as "What if User were not a member of any groups.
	User *string `field:"optional" json:"user" yaml:"user"`
}

