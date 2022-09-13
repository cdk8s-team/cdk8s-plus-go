package k8s


// ServiceSpec describes the attributes that a user creates on a service.
type ServiceSpec struct {
	// allocateLoadBalancerNodePorts defines if NodePorts will be automatically allocated for services with type LoadBalancer.
	//
	// Default is "true". It may be set to "false" if the cluster load-balancer does not rely on NodePorts.  If the caller requests specific NodePorts (by specifying a value), those requests will be respected, regardless of this field. This field may only be set for services with type LoadBalancer and will be cleared if the type is changed to any other type. This field is beta-level and is only honored by servers that enable the ServiceLBNodePortControl feature.
	AllocateLoadBalancerNodePorts *bool `field:"optional" json:"allocateLoadBalancerNodePorts" yaml:"allocateLoadBalancerNodePorts"`
	// clusterIP is the IP address of the service and is usually assigned randomly.
	//
	// If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be blank) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above).  Valid values are "None", empty string (""), or a valid IP address. Setting this to "None" makes a "headless service" (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required.  Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a Service of type ExternalName, creation will fail. This field will be wiped when updating a Service to type ExternalName. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	ClusterIp *string `field:"optional" json:"clusterIp" yaml:"clusterIp"`
	// ClusterIPs is a list of IP addresses assigned to this service, and are usually assigned randomly.
	//
	// If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be empty) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above).  Valid values are "None", empty string (""), or a valid IP address.  Setting this to "None" makes a "headless service" (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required.  Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a Service of type ExternalName, creation will fail. This field will be wiped when updating a Service to type ExternalName.  If this field is not specified, it will be initialized from the clusterIP field.  If this field is specified, clients must ensure that clusterIPs[0] and clusterIP have the same value.
	//
	// Unless the "IPv6DualStack" feature gate is enabled, this field is limited to one value, which must be the same as the clusterIP field.  If the feature gate is enabled, this field may hold a maximum of two entries (dual-stack IPs, in either order).  These IPs must correspond to the values of the ipFamilies field. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	ClusterIPs *[]*string `field:"optional" json:"clusterIPs" yaml:"clusterIPs"`
	// externalIPs is a list of IP addresses for which nodes in the cluster will also accept traffic for this service.
	//
	// These IPs are not managed by Kubernetes.  The user is responsible for ensuring that traffic arrives at a node with this IP.  A common example is external load-balancers that are not part of the Kubernetes system.
	ExternalIPs *[]*string `field:"optional" json:"externalIPs" yaml:"externalIPs"`
	// externalName is the external reference that discovery mechanisms will return as an alias for this service (e.g. a DNS CNAME record). No proxying will be involved.  Must be a lowercase RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) and requires `type` to be "ExternalName".
	ExternalName *string `field:"optional" json:"externalName" yaml:"externalName"`
	// externalTrafficPolicy denotes if this Service desires to route external traffic to node-local or cluster-wide endpoints.
	//
	// "Local" preserves the client source IP and avoids a second hop for LoadBalancer and Nodeport type services, but risks potentially imbalanced traffic spreading. "Cluster" obscures the client source IP and may cause a second hop to another node, but should have good overall load-spreading.
	ExternalTrafficPolicy *string `field:"optional" json:"externalTrafficPolicy" yaml:"externalTrafficPolicy"`
	// healthCheckNodePort specifies the healthcheck nodePort for the service.
	//
	// This only applies when type is set to LoadBalancer and externalTrafficPolicy is set to Local. If a value is specified, is in-range, and is not in use, it will be used.  If not specified, a value will be automatically allocated.  External systems (e.g. load-balancers) can use this port to determine if a given node holds endpoints for this service or not.  If this field is specified when creating a Service which does not need it, creation will fail. This field will be wiped when updating a Service to no longer need it (e.g. changing type).
	HealthCheckNodePort *float64 `field:"optional" json:"healthCheckNodePort" yaml:"healthCheckNodePort"`
	// InternalTrafficPolicy specifies if the cluster internal traffic should be routed to all endpoints or node-local endpoints only.
	//
	// "Cluster" routes internal traffic to a Service to all endpoints. "Local" routes traffic to node-local endpoints only, traffic is dropped if no node-local endpoints are ready. The default value is "Cluster".
	InternalTrafficPolicy *string `field:"optional" json:"internalTrafficPolicy" yaml:"internalTrafficPolicy"`
	// IPFamilies is a list of IP families (e.g. IPv4, IPv6) assigned to this service, and is gated by the "IPv6DualStack" feature gate.  This field is usually assigned automatically based on cluster configuration and the ipFamilyPolicy field. If this field is specified manually, the requested family is available in the cluster, and ipFamilyPolicy allows it, it will be used; otherwise creation of the service will fail.  This field is conditionally mutable: it allows for adding or removing a secondary IP family, but it does not allow changing the primary IP family of the Service.  Valid values are "IPv4" and "IPv6".  This field only applies to Services of types ClusterIP, NodePort, and LoadBalancer, and does apply to "headless" services.  This field will be wiped when updating a Service to type ExternalName.
	//
	// This field may hold a maximum of two entries (dual-stack families, in either order).  These families must correspond to the values of the clusterIPs field, if specified. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field.
	IpFamilies *[]*string `field:"optional" json:"ipFamilies" yaml:"ipFamilies"`
	// IPFamilyPolicy represents the dual-stack-ness requested or required by this Service, and is gated by the "IPv6DualStack" feature gate.
	//
	// If there is no value provided, then this field will be set to SingleStack. Services can be "SingleStack" (a single IP family), "PreferDualStack" (two IP families on dual-stack configured clusters or a single IP family on single-stack clusters), or "RequireDualStack" (two IP families on dual-stack configured clusters, otherwise fail). The ipFamilies and clusterIPs fields depend on the value of this field.  This field will be wiped when updating a service to type ExternalName.
	IpFamilyPolicy *string `field:"optional" json:"ipFamilyPolicy" yaml:"ipFamilyPolicy"`
	// loadBalancerClass is the class of the load balancer implementation this Service belongs to.
	//
	// If specified, the value of this field must be a label-style identifier, with an optional prefix, e.g. "internal-vip" or "example.com/internal-vip". Unprefixed names are reserved for end-users. This field can only be set when the Service type is 'LoadBalancer'. If not set, the default load balancer implementation is used, today this is typically done through the cloud provider integration, but should apply for any default implementation. If set, it is assumed that a load balancer implementation is watching for Services with a matching class. Any default load balancer implementation (e.g. cloud providers) should ignore Services that set this field. This field can only be set when creating or updating a Service to type 'LoadBalancer'. Once set, it can not be changed. This field will be wiped when a service is updated to a non 'LoadBalancer' type.
	LoadBalancerClass *string `field:"optional" json:"loadBalancerClass" yaml:"loadBalancerClass"`
	// Only applies to Service Type: LoadBalancer LoadBalancer will get created with the IP specified in this field.
	//
	// This feature depends on whether the underlying cloud-provider supports specifying the loadBalancerIP when a load balancer is created. This field will be ignored if the cloud-provider does not support the feature.
	LoadBalancerIp *string `field:"optional" json:"loadBalancerIp" yaml:"loadBalancerIp"`
	// If specified and supported by the platform, this will restrict traffic through the cloud-provider load-balancer will be restricted to the specified client IPs.
	//
	// This field will be ignored if the cloud-provider does not support the feature." More info: https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/
	LoadBalancerSourceRanges *[]*string `field:"optional" json:"loadBalancerSourceRanges" yaml:"loadBalancerSourceRanges"`
	// The list of ports that are exposed by this service.
	//
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	Ports *[]*ServicePort `field:"optional" json:"ports" yaml:"ports"`
	// publishNotReadyAddresses indicates that any agent which deals with endpoints for this Service should disregard any indications of ready/not-ready.
	//
	// The primary use case for setting this field is for a StatefulSet's Headless Service to propagate SRV DNS records for its Pods for the purpose of peer discovery. The Kubernetes controllers that generate Endpoints and EndpointSlice resources for Services interpret this to mean that all endpoints are considered "ready" even if the Pods themselves are not. Agents which consume only Kubernetes generated endpoints through the Endpoints or EndpointSlice resources can safely assume this behavior.
	PublishNotReadyAddresses *bool `field:"optional" json:"publishNotReadyAddresses" yaml:"publishNotReadyAddresses"`
	// Route service traffic to pods with label keys and values matching this selector.
	//
	// If empty or not present, the service is assumed to have an external process managing its endpoints, which Kubernetes will not modify. Only applies to types ClusterIP, NodePort, and LoadBalancer. Ignored if type is ExternalName. More info: https://kubernetes.io/docs/concepts/services-networking/service/
	Selector *map[string]*string `field:"optional" json:"selector" yaml:"selector"`
	// Supports "ClientIP" and "None".
	//
	// Used to maintain session affinity. Enable client IP based session affinity. Must be ClientIP or None. Defaults to None. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	SessionAffinity *string `field:"optional" json:"sessionAffinity" yaml:"sessionAffinity"`
	// sessionAffinityConfig contains the configurations of session affinity.
	SessionAffinityConfig *SessionAffinityConfig `field:"optional" json:"sessionAffinityConfig" yaml:"sessionAffinityConfig"`
	// type determines how the Service is exposed.
	//
	// Defaults to ClusterIP. Valid options are ExternalName, ClusterIP, NodePort, and LoadBalancer. "ClusterIP" allocates a cluster-internal IP address for load-balancing to endpoints. Endpoints are determined by the selector or if that is not specified, by manual construction of an Endpoints object or EndpointSlice objects. If clusterIP is "None", no virtual IP is allocated and the endpoints are published as a set of endpoints rather than a virtual IP. "NodePort" builds on ClusterIP and allocates a port on every node which routes to the same endpoints as the clusterIP. "LoadBalancer" builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the same endpoints as the clusterIP. "ExternalName" aliases this service to the specified externalName. Several other fields do not apply to ExternalName services. More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types
	Type *string `field:"optional" json:"type" yaml:"type"`
}

