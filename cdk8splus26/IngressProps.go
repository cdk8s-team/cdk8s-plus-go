package cdk8splus26

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Properties for `Ingress`.
type IngressProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// The default backend services requests that do not match any rule.
	//
	// Using this option or the `addDefaultBackend()` method is equivalent to
	// adding a rule with both `path` and `host` undefined.
	DefaultBackend IngressBackend `field:"optional" json:"defaultBackend" yaml:"defaultBackend"`
	// Routing rules for this ingress.
	//
	// Each rule must define an `IngressBackend` that will receive the requests
	// that match this rule. If both `host` and `path` are not specifiec, this
	// backend will be used as the default backend of the ingress.
	//
	// You can also add rules later using `addRule()`, `addHostRule()`,
	// `addDefaultBackend()` and `addHostDefaultBackend()`.
	Rules *[]*IngressRule `field:"optional" json:"rules" yaml:"rules"`
	// TLS settings for this ingress.
	//
	// Using this option tells the ingress controller to expose a TLS endpoint.
	// Currently the Ingress only supports a single TLS port, 443. If multiple
	// members of this list specify different hosts, they will be multiplexed on
	// the same port according to the hostname specified through the SNI TLS
	// extension, if the ingress controller fulfilling the ingress supports SNI.
	Tls *[]*IngressTls `field:"optional" json:"tls" yaml:"tls"`
}

