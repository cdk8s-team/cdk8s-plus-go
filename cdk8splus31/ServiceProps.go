package cdk8splus31

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
	// Default: - Automatically assigned.
	//
	ClusterIP *string `field:"optional" json:"clusterIP" yaml:"clusterIP"`
	// A list of IP addresses for which nodes in the cluster will also accept traffic for this service.
	//
	// These IPs are not managed by Kubernetes. The user
	// is responsible for ensuring that traffic arrives at a node with this IP. A
	// common example is external load-balancers that are not part of the
	// Kubernetes system.
	// Default: - No external IPs.
	//
	ExternalIPs *[]*string `field:"optional" json:"externalIPs" yaml:"externalIPs"`
	// The externalName to be used when ServiceType.EXTERNAL_NAME is set.
	// Default: - No external name.
	//
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
	// Default: - either the selector ports, or none.
	//
	Ports *[]*ServicePort `field:"optional" json:"ports" yaml:"ports"`
	// The publishNotReadyAddresses indicates that any agent which deals with endpoints for this Service should disregard any indications of ready/not-ready.
	//
	// More info: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#servicespec-v1-core
	// Default: - false.
	//
	PublishNotReadyAddresses *bool `field:"optional" json:"publishNotReadyAddresses" yaml:"publishNotReadyAddresses"`
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
	// Default: - unset, the service is assumed to have an external process managing
	// its endpoints, which Kubernetes will not modify.
	//
	Selector IPodSelector `field:"optional" json:"selector" yaml:"selector"`
	// Determines how the Service is exposed.
	//
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types
	// Default: ServiceType.ClusterIP
	//
	Type ServiceType `field:"optional" json:"type" yaml:"type"`
}

