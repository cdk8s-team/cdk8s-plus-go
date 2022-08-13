// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Properties for `Service`.
type ServiceProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// The IP address of the service and is usually assigned randomly by the master.
	//
	// If an address is specified manually and is not in use by others, it
	// will be allocated to the service; otherwise, creation of the service will
	// fail. This field can not be changed through updates. Valid values are
	// "None", empty string (""), or a valid IP address. "None" can be specified
	// for headless services when proxying is not required. Only applies to types
	// ClusterIP, NodePort, and LoadBalancer. Ignored if type is ExternalName.
	// See: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	//
	ClusterIP *string `field:"optional" json:"clusterIP" yaml:"clusterIP"`
	// A list of IP addresses for which nodes in the cluster will also accept traffic for this service.
	//
	// These IPs are not managed by Kubernetes. The user
	// is responsible for ensuring that traffic arrives at a node with this IP. A
	// common example is external load-balancers that are not part of the
	// Kubernetes system.
	ExternalIPs *[]*string `field:"optional" json:"externalIPs" yaml:"externalIPs"`
	// The externalName to be used when ServiceType.EXTERNAL_NAME is set.
	ExternalName *string `field:"optional" json:"externalName" yaml:"externalName"`
	// A list of CIDR IP addresses, if specified and supported by the platform, will restrict traffic through the cloud-provider load-balancer to the specified client IPs.
	//
	// More info: https://kubernetes.io/docs/tasks/access-application-cluster/configure-cloud-provider-firewall/
	LoadBalancerSourceRanges *[]*string `field:"optional" json:"loadBalancerSourceRanges" yaml:"loadBalancerSourceRanges"`
	// The ports this service binds to.
	//
	// If the selector of the service is a managed pod / workload,
	// its ports will are automatically extracted and used as the default value.
	// Otherwise, no ports are bound.
	Ports *[]*ServicePort `field:"optional" json:"ports" yaml:"ports"`
	// Which pods should the service select and route to.
	//
	// You can pass one of the following:
	//
	// - An instance of `Pod` or any workload resource (e.g `Deployment`, `StatefulSet`, ...)
	// - Pods selected by the `Pods.select` function. Note that in this case only labels can be specified.
	//
	// Example:
	//   // select the pods of a specific deployment
	//   const backend = new kplus.Deployment(this, 'Backend', ...);
	//   new kplus.Service(this, 'Service', { selector: backend });
	//
	//   // select all pods labeled with the `tier=backend` label
	//   const backend = kplus.Pod.labeled({ tier: 'backend' });
	//   new kplus.Service(this, 'Service', { selector: backend });
	//
	Selector IPodSelector `field:"optional" json:"selector" yaml:"selector"`
	// Determines how the Service is exposed.
	//
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types
	Type ServiceType `field:"optional" json:"type" yaml:"type"`
}

