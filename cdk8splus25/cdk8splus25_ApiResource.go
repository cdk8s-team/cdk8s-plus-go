// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-25 synthesizes Kubernetes manifests for Kubernetes 1.25.0
package cdk8splus25

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus25/v2/jsii"
)

// Represents information about an API resource type.
type ApiResource interface {
	IApiEndpoint
	IApiResource
	// The group portion of the API version (e.g. `authorization.k8s.io`).
	ApiGroup() *string
	// The name of the resource type as it appears in the relevant API endpoint.
	//
	// Example:
	//   - "pods" or "pods/log"
	//
	// See: https://kubernetes.io/docs/reference/access-authn-authz/rbac/#referring-to-resources
	//
	ResourceType() *string
	// Return the IApiResource this object represents.
	AsApiResource() IApiResource
	// Return the non resource url this object represents.
	AsNonApiResource() *string
}

// The jsii proxy struct for ApiResource
type jsiiProxy_ApiResource struct {
	jsiiProxy_IApiEndpoint
	jsiiProxy_IApiResource
}

func (j *jsiiProxy_ApiResource) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiResource) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}


// API resource information for a custom resource type.
func ApiResource_Custom(options *ApiResourceOptions) ApiResource {
	_init_.Initialize()

	if err := validateApiResource_CustomParameters(options); err != nil {
		panic(err)
	}
	var returns ApiResource

	_jsii_.StaticInvoke(
		"cdk8s-plus-25.ApiResource",
		"custom",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func ApiResource_API_SERVICES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"API_SERVICES",
		&returns,
	)
	return returns
}

func ApiResource_BINDINGS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"BINDINGS",
		&returns,
	)
	return returns
}

func ApiResource_CERTIFICATE_SIGNING_REQUESTS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"CERTIFICATE_SIGNING_REQUESTS",
		&returns,
	)
	return returns
}

func ApiResource_CLUSTER_ROLE_BINDINGS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"CLUSTER_ROLE_BINDINGS",
		&returns,
	)
	return returns
}

func ApiResource_CLUSTER_ROLES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"CLUSTER_ROLES",
		&returns,
	)
	return returns
}

func ApiResource_COMPONENT_STATUSES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"COMPONENT_STATUSES",
		&returns,
	)
	return returns
}

func ApiResource_CONFIG_MAPS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"CONFIG_MAPS",
		&returns,
	)
	return returns
}

func ApiResource_CONTROLLER_REVISIONS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"CONTROLLER_REVISIONS",
		&returns,
	)
	return returns
}

func ApiResource_CRON_JOBS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"CRON_JOBS",
		&returns,
	)
	return returns
}

func ApiResource_CSI_DRIVERS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"CSI_DRIVERS",
		&returns,
	)
	return returns
}

func ApiResource_CSI_NODES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"CSI_NODES",
		&returns,
	)
	return returns
}

func ApiResource_CSI_STORAGE_CAPACITIES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"CSI_STORAGE_CAPACITIES",
		&returns,
	)
	return returns
}

func ApiResource_CUSTOM_RESOURCE_DEFINITIONS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"CUSTOM_RESOURCE_DEFINITIONS",
		&returns,
	)
	return returns
}

func ApiResource_DAEMON_SETS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"DAEMON_SETS",
		&returns,
	)
	return returns
}

func ApiResource_DEPLOYMENTS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"DEPLOYMENTS",
		&returns,
	)
	return returns
}

func ApiResource_ENDPOINT_SLICES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"ENDPOINT_SLICES",
		&returns,
	)
	return returns
}

func ApiResource_ENDPOINTS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"ENDPOINTS",
		&returns,
	)
	return returns
}

func ApiResource_EVENTS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"EVENTS",
		&returns,
	)
	return returns
}

func ApiResource_FLOW_SCHEMAS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"FLOW_SCHEMAS",
		&returns,
	)
	return returns
}

func ApiResource_HORIZONTAL_POD_AUTOSCALERS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"HORIZONTAL_POD_AUTOSCALERS",
		&returns,
	)
	return returns
}

func ApiResource_INGRESS_CLASSES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"INGRESS_CLASSES",
		&returns,
	)
	return returns
}

func ApiResource_INGRESSES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"INGRESSES",
		&returns,
	)
	return returns
}

func ApiResource_JOBS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"JOBS",
		&returns,
	)
	return returns
}

func ApiResource_LEASES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"LEASES",
		&returns,
	)
	return returns
}

func ApiResource_LIMIT_RANGES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"LIMIT_RANGES",
		&returns,
	)
	return returns
}

func ApiResource_LOCAL_SUBJECT_ACCESS_REVIEWS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"LOCAL_SUBJECT_ACCESS_REVIEWS",
		&returns,
	)
	return returns
}

func ApiResource_MUTATING_WEBHOOK_CONFIGURATIONS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"MUTATING_WEBHOOK_CONFIGURATIONS",
		&returns,
	)
	return returns
}

func ApiResource_NAMESPACES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"NAMESPACES",
		&returns,
	)
	return returns
}

func ApiResource_NETWORK_POLICIES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"NETWORK_POLICIES",
		&returns,
	)
	return returns
}

func ApiResource_NODES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"NODES",
		&returns,
	)
	return returns
}

func ApiResource_PERSISTENT_VOLUME_CLAIMS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"PERSISTENT_VOLUME_CLAIMS",
		&returns,
	)
	return returns
}

func ApiResource_PERSISTENT_VOLUMES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"PERSISTENT_VOLUMES",
		&returns,
	)
	return returns
}

func ApiResource_POD_DISRUPTION_BUDGETS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"POD_DISRUPTION_BUDGETS",
		&returns,
	)
	return returns
}

func ApiResource_POD_TEMPLATES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"POD_TEMPLATES",
		&returns,
	)
	return returns
}

func ApiResource_PODS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"PODS",
		&returns,
	)
	return returns
}

func ApiResource_PRIORITY_CLASSES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"PRIORITY_CLASSES",
		&returns,
	)
	return returns
}

func ApiResource_PRIORITY_LEVEL_CONFIGURATIONS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"PRIORITY_LEVEL_CONFIGURATIONS",
		&returns,
	)
	return returns
}

func ApiResource_REPLICA_SETS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"REPLICA_SETS",
		&returns,
	)
	return returns
}

func ApiResource_REPLICATION_CONTROLLERS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"REPLICATION_CONTROLLERS",
		&returns,
	)
	return returns
}

func ApiResource_RESOURCE_QUOTAS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"RESOURCE_QUOTAS",
		&returns,
	)
	return returns
}

func ApiResource_ROLE_BINDINGS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"ROLE_BINDINGS",
		&returns,
	)
	return returns
}

func ApiResource_ROLES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"ROLES",
		&returns,
	)
	return returns
}

func ApiResource_RUNTIME_CLASSES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"RUNTIME_CLASSES",
		&returns,
	)
	return returns
}

func ApiResource_SECRETS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"SECRETS",
		&returns,
	)
	return returns
}

func ApiResource_SELF_SUBJECT_ACCESS_REVIEWS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"SELF_SUBJECT_ACCESS_REVIEWS",
		&returns,
	)
	return returns
}

func ApiResource_SELF_SUBJECT_RULES_REVIEWS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"SELF_SUBJECT_RULES_REVIEWS",
		&returns,
	)
	return returns
}

func ApiResource_SERVICE_ACCOUNTS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"SERVICE_ACCOUNTS",
		&returns,
	)
	return returns
}

func ApiResource_SERVICES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"SERVICES",
		&returns,
	)
	return returns
}

func ApiResource_STATEFUL_SETS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"STATEFUL_SETS",
		&returns,
	)
	return returns
}

func ApiResource_STORAGE_CLASSES() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"STORAGE_CLASSES",
		&returns,
	)
	return returns
}

func ApiResource_SUBJECT_ACCESS_REVIEWS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"SUBJECT_ACCESS_REVIEWS",
		&returns,
	)
	return returns
}

func ApiResource_TOKEN_REVIEWS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"TOKEN_REVIEWS",
		&returns,
	)
	return returns
}

func ApiResource_VALIDATING_WEBHOOK_CONFIGURATIONS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"VALIDATING_WEBHOOK_CONFIGURATIONS",
		&returns,
	)
	return returns
}

func ApiResource_VOLUME_ATTACHMENTS() ApiResource {
	_init_.Initialize()
	var returns ApiResource
	_jsii_.StaticGet(
		"cdk8s-plus-25.ApiResource",
		"VOLUME_ATTACHMENTS",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ApiResource) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		a,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiResource) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

