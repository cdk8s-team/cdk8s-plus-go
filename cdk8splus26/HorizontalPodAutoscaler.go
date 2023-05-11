package cdk8splus26

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus26/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// A HorizontalPodAutoscaler scales a workload up or down in response to a metric change.
//
// This allows your services to scale up when demand is high and scale down
// when they are no longer needed.
//
//
// Typical use cases for HorizontalPodAutoscaler:
//
// * When Memory usage is above 70%, scale up the number of replicas to meet the demand.
// * When CPU usage is below 30%, scale down the number of replicas to save resources.
// * When a service is experiencing a spike in traffic, scale up the number of replicas
//    to meet the demand. Then, when the traffic subsides, scale down the number of
//    replicas to save resources.
//
// The autoscaler uses the following algorithm to determine the number of replicas to scale:
//
// `desiredReplicas = ceil[currentReplicas * ( currentMetricValue / desiredMetricValue )]`
//
// HorizontalPodAutoscaler's can be used to with any `Scalable` workload:
// * Deployment
// * StatefulSet
//
// **Targets that already have a replica count defined:**
//
// Remove any replica counts from the target resource before associating with a
// HorizontalPodAutoscaler. If this isn't done, then any time a change to that object is applied,
// Kubernetes will scale the current number of Pods to the value of the target.replicas key. This
// may not be desired and could lead to unexpected behavior.
//
// Example:
//   const backend = new kplus.Deployment(this, 'Backend', ...);
//
//   const hpa = new kplus.HorizontalPodAutoscaler(chart, 'Hpa', {
//    target: backend,
//    maxReplicas: 10,
//    scaleUp: {
//      policies: [
//        {
//          replicas: kplus.Replicas.absolute(3),
//          duration: Duration.minutes(5),
//        },
//      ],
//    },
//   });
//
// See: https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/#implicit-maintenance-mode-deactivation
//
type HorizontalPodAutoscaler interface {
	Resource
	// The group portion of the API version (e.g. "authorization.k8s.io").
	ApiGroup() *string
	// The underlying cdk8s API object.
	// See: base.Resource.apiObject
	//
	ApiObject() cdk8s.ApiObject
	// The object's API version (e.g. "authorization.k8s.io/v1").
	ApiVersion() *string
	// The object kind (e.g. "Deployment").
	Kind() *string
	// The maximum number of replicas that can be scaled up to.
	MaxReplicas() *float64
	Metadata() cdk8s.ApiObjectMetadataDefinition
	// The metric conditions that trigger a scale up or scale down.
	Metrics() *[]Metric
	// The minimum number of replicas that can be scaled down to.
	MinReplicas() *float64
	// The name of this API object.
	Name() *string
	// The tree node.
	Node() constructs.Node
	Permissions() ResourcePermissions
	// The unique, namespace-global, name of an object inside the Kubernetes cluster.
	//
	// If this is omitted, the ApiResource should represent all objects of the given type.
	ResourceName() *string
	// The name of a resource type as it appears in the relevant API endpoint.
	ResourceType() *string
	// The scaling behavior when scaling down.
	ScaleDown() *ScalingRules
	// The scaling behavior when scaling up.
	ScaleUp() *ScalingRules
	// The workload to scale up or down.
	Target() IScalable
	// Return the IApiResource this object represents.
	AsApiResource() IApiResource
	// Return the non resource url this object represents.
	AsNonApiResource() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for HorizontalPodAutoscaler
type jsiiProxy_HorizontalPodAutoscaler struct {
	jsiiProxy_Resource
}

func (j *jsiiProxy_HorizontalPodAutoscaler) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscaler) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscaler) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscaler) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscaler) MaxReplicas() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxReplicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscaler) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscaler) Metrics() *[]Metric {
	var returns *[]Metric
	_jsii_.Get(
		j,
		"metrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscaler) MinReplicas() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minReplicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscaler) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscaler) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscaler) Permissions() ResourcePermissions {
	var returns ResourcePermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscaler) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscaler) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscaler) ScaleDown() *ScalingRules {
	var returns *ScalingRules
	_jsii_.Get(
		j,
		"scaleDown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscaler) ScaleUp() *ScalingRules {
	var returns *ScalingRules
	_jsii_.Get(
		j,
		"scaleUp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscaler) Target() IScalable {
	var returns IScalable
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}


func NewHorizontalPodAutoscaler(scope constructs.Construct, id *string, props *HorizontalPodAutoscalerProps) HorizontalPodAutoscaler {
	_init_.Initialize()

	if err := validateNewHorizontalPodAutoscalerParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_HorizontalPodAutoscaler{}

	_jsii_.Create(
		"cdk8s-plus-26.HorizontalPodAutoscaler",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewHorizontalPodAutoscaler_Override(h HorizontalPodAutoscaler, scope constructs.Construct, id *string, props *HorizontalPodAutoscalerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-26.HorizontalPodAutoscaler",
		[]interface{}{scope, id, props},
		h,
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
func HorizontalPodAutoscaler_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHorizontalPodAutoscaler_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-plus-26.HorizontalPodAutoscaler",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizontalPodAutoscaler) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		h,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizontalPodAutoscaler) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizontalPodAutoscaler) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

