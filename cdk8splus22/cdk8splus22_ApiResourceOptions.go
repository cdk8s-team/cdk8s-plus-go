// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-22 synthesizes Kubernetes manifests for Kubernetes 1.22.0
package cdk8splus22


// Options for `ApiResource`.
type ApiResourceOptions struct {
	// The group portion of the API version (e.g. `authorization.k8s.io`).
	ApiGroup *string `field:"required" json:"apiGroup" yaml:"apiGroup"`
	// The name of the resource type as it appears in the relevant API endpoint.
	//
	// Example:
	//   - "pods" or "pods/log"
	//
	// See: https://kubernetes.io/docs/reference/access-authn-authz/rbac/#referring-to-resources
	//
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
}

