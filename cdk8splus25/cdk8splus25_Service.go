// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-25 synthesizes Kubernetes manifests for Kubernetes 1.25.0
package cdk8splus25

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus25/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// An abstract way to expose an application running on a set of Pods as a network service.
//
// With Kubernetes you don't need to modify your application to use an unfamiliar service discovery mechanism.
// Kubernetes gives Pods their own IP addresses and a single DNS name for a set of Pods, and can load-balance across them.
//
// For example, consider a stateless image-processing backend which is running with 3 replicas. Those replicas are fungible—frontends do not care which backend they use.
// While the actual Pods that compose the backend set may change, the frontend clients should not need to be aware of that,
// nor should they need to keep track of the set of backends themselves.
// The Service abstraction enables this decoupling.
//
// If you're able to use Kubernetes APIs for service discovery in your application, you can query the API server for Endpoints,
// that get updated whenever the set of Pods in a Service changes. For non-native applications, Kubernetes offers ways to place a network port
// or load balancer in between your application and the backend Pods.
type Service interface {
	Resource
	// The group portion of the API version (e.g. "authorization.k8s.io").
	ApiGroup() *string
	// The underlying cdk8s API object.
	// See: base.Resource.apiObject
	//
	ApiObject() cdk8s.ApiObject
	// The object's API version (e.g. "authorization.k8s.io/v1").
	ApiVersion() *string
	// The IP address of the service and is usually assigned randomly by the master.
	ClusterIP() *string
	// The externalName to be used for EXTERNAL_NAME types.
	ExternalName() *string
	// The object kind (e.g. "Deployment").
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	// The name of this API object.
	Name() *string
	// The tree node.
	Node() constructs.Node
	Permissions() ResourcePermissions
	// Ports for this service.
	//
	// Use `bind()` to bind additional service ports.
	Ports() *[]*ServicePort
	// The unique, namespace-global, name of an object inside the Kubernetes cluster.
	//
	// If this is omitted, the ApiResource should represent all objects of the given type.
	ResourceName() *string
	// The name of a resource type as it appears in the relevant API endpoint.
	ResourceType() *string
	// Determines how the Service is exposed.
	Type() ServiceType
	// Return the IApiResource this object represents.
	AsApiResource() IApiResource
	// Return the non resource url this object represents.
	AsNonApiResource() *string
	// Configure a port the service will bind to.
	//
	// This method can be called multiple times.
	Bind(port *float64, options *ServiceBindOptions)
	// Expose a service via an ingress using the specified path.
	//
	// Returns: The `Ingress` resource that was used.
	ExposeViaIngress(path *string, options *ExposeServiceViaIngressOptions) Ingress
	// Require this service to select pods matching the selector.
	//
	// Note that invoking this method multiple times acts as an AND operator
	// on the resulting labels.
	Select(selector IPodSelector)
	// Require this service to select pods with this label.
	//
	// Note that invoking this method multiple times acts as an AND operator
	// on the resulting labels.
	SelectLabel(key *string, value *string)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Service
type jsiiProxy_Service struct {
	jsiiProxy_Resource
}

func (j *jsiiProxy_Service) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) ClusterIP() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIP",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) ExternalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Permissions() ResourcePermissions {
	var returns ResourcePermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Ports() *[]*ServicePort {
	var returns *[]*ServicePort
	_jsii_.Get(
		j,
		"ports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Type() ServiceType {
	var returns ServiceType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


func NewService(scope constructs.Construct, id *string, props *ServiceProps) Service {
	_init_.Initialize()

	if err := validateNewServiceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Service{}

	_jsii_.Create(
		"cdk8s-plus-25.Service",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewService_Override(s Service, scope constructs.Construct, id *string, props *ServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-25.Service",
		[]interface{}{scope, id, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func Service_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateService_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-25.Service",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Service) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		s,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Service) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Service) Bind(port *float64, options *ServiceBindOptions) {
	if err := s.validateBindParameters(port, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"bind",
		[]interface{}{port, options},
	)
}

func (s *jsiiProxy_Service) ExposeViaIngress(path *string, options *ExposeServiceViaIngressOptions) Ingress {
	if err := s.validateExposeViaIngressParameters(path, options); err != nil {
		panic(err)
	}
	var returns Ingress

	_jsii_.Invoke(
		s,
		"exposeViaIngress",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Service) Select(selector IPodSelector) {
	if err := s.validateSelectParameters(selector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"select",
		[]interface{}{selector},
	)
}

func (s *jsiiProxy_Service) SelectLabel(key *string, value *string) {
	if err := s.validateSelectLabelParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"selectLabel",
		[]interface{}{key, value},
	)
}

func (s *jsiiProxy_Service) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

