package cdk8splus24


// Represents the rules mapping the paths under a specified host to the related backend services.
//
// Incoming requests are first evaluated for a host match,
// then routed to the backend associated with the matching path.
type IngressRule struct {
	// Backend defines the referenced service endpoint to which the traffic will be forwarded to.
	Backend IngressBackend `field:"required" json:"backend" yaml:"backend"`
	// Host is the fully qualified domain name of a network host, as defined by RFC 3986.
	//
	// Note the following deviations from the "host" part of the URI as
	// defined in the RFC: 1. IPs are not allowed. Currently an IngressRuleValue
	// can only apply to the IP in the Spec of the parent Ingress. 2. The `:`
	// delimiter is not respected because ports are not allowed. Currently the
	// port of an Ingress is implicitly :80 for http and :443 for https. Both
	// these may change in the future. Incoming requests are matched against the
	// host before the IngressRuleValue.
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Path is an extended POSIX regex as defined by IEEE Std 1003.1, (i.e this follows the egrep/unix syntax, not the perl syntax) matched against the path of an incoming request. Currently it can contain characters disallowed from the conventional "path" part of a URL as defined by RFC 3986. Paths must begin with a '/'.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Specify how the path is matched against request paths.
	//
	// By default, path
	// types will be matched by prefix.
	// See: https://kubernetes.io/docs/concepts/services-networking/ingress/#path-types
	//
	PathType HttpIngressPathType `field:"optional" json:"pathType" yaml:"pathType"`
}

