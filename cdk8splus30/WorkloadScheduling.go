package cdk8splus30

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2/jsii"
)

// Controls the pod scheduling strategy of this workload.
//
// It offers some additional API's on top of the core pod scheduling.
type WorkloadScheduling interface {
	PodScheduling
	Instance() AbstractPod
	// Assign this pod a specific node by name.
	//
	// The scheduler ignores the Pod, and the kubelet on the named node
	// tries to place the Pod on that node. Overrules any affinity rules of the pod.
	//
	// Some limitations of static assignment are:
	//
	// - If the named node does not exist, the Pod will not run, and in some
	//    cases may be automatically deleted.
	// - If the named node does not have the resources to accommodate the Pod,
	//    the Pod will fail and its reason will indicate why, for example OutOfmemory or OutOfcpu.
	// - Node names in cloud environments are not always predictable or stable.
	//
	// Will throw is the pod is already assigned to named node.
	//
	// Under the hood, this method utilizes the `nodeName` property.
	Assign(node NamedNode)
	// Attract this pod to a node matched by selectors. You can select a node by using `Node.labeled()`.
	//
	// Attracting to multiple nodes (i.e invoking this method multiple times) acts as
	// an OR condition, meaning the pod will be assigned to either one of the nodes.
	//
	// Under the hood, this method utilizes the `nodeAffinity` property.
	// See: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#node-affinity
	//
	Attract(node LabeledNode, options *PodSchedulingAttractOptions)
	// Co-locate this pod with a scheduling selection.
	//
	// A selection can be one of:
	//
	// - An instance of a `Pod`.
	// - An instance of a `Workload` (e.g `Deployment`, `StatefulSet`).
	// - An un-managed pod that can be selected via `Pods.select()`.
	//
	// Co-locating with multiple selections ((i.e invoking this method multiple times)) acts as
	// an AND condition. meaning the pod will be assigned to a node that satisfies all
	// selections (i.e runs at least one pod that satisifies each selection).
	//
	// Under the hood, this method utilizes the `podAffinity` property.
	// See: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#inter-pod-affinity-and-anti-affinity
	//
	Colocate(selector IPodSelector, options *PodSchedulingColocateOptions)
	// Seperate this pod from a scheduling selection.
	//
	// A selection can be one of:
	//
	// - An instance of a `Pod`.
	// - An instance of a `Workload` (e.g `Deployment`, `StatefulSet`).
	// - An un-managed pod that can be selected via `Pods.select()`.
	//
	// Seperating from multiple selections acts as an AND condition. meaning the pod
	// will not be assigned to a node that satisfies all selections (i.e runs at least one pod that satisifies each selection).
	//
	// Under the hood, this method utilizes the `podAntiAffinity` property.
	// See: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#inter-pod-affinity-and-anti-affinity
	//
	Separate(selector IPodSelector, options *PodSchedulingSeparateOptions)
	// Spread the pods in this workload by the topology key.
	//
	// A spread is a separation of the pod from itself and is used to
	// balance out pod replicas across a given topology.
	Spread(options *WorkloadSchedulingSpreadOptions)
	// Allow this pod to tolerate taints matching these tolerations.
	//
	// You can put multiple taints on the same node and multiple tolerations on the same pod.
	// The way Kubernetes processes multiple taints and tolerations is like a filter: start with
	// all of a node's taints, then ignore the ones for which the pod has a matching toleration;
	// the remaining un-ignored taints have the indicated effects on the pod. In particular:
	//
	// - if there is at least one un-ignored taint with effect NoSchedule then Kubernetes will
	//    not schedule the pod onto that node
	// - if there is no un-ignored taint with effect NoSchedule but there is at least one un-ignored
	//    taint with effect PreferNoSchedule then Kubernetes will try to not schedule the pod onto the node
	// - if there is at least one un-ignored taint with effect NoExecute then the pod will be evicted from
	//    the node (if it is already running on the node), and will not be scheduled onto the node (if it is
	//    not yet running on the node).
	//
	// Under the hood, this method utilizes the `tolerations` property.
	// See: https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/
	//
	Tolerate(node TaintedNode)
}

// The jsii proxy struct for WorkloadScheduling
type jsiiProxy_WorkloadScheduling struct {
	jsiiProxy_PodScheduling
}

func (j *jsiiProxy_WorkloadScheduling) Instance() AbstractPod {
	var returns AbstractPod
	_jsii_.Get(
		j,
		"instance",
		&returns,
	)
	return returns
}


func NewWorkloadScheduling(instance AbstractPod) WorkloadScheduling {
	_init_.Initialize()

	if err := validateNewWorkloadSchedulingParameters(instance); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkloadScheduling{}

	_jsii_.Create(
		"cdk8s-plus-30.WorkloadScheduling",
		[]interface{}{instance},
		&j,
	)

	return &j
}

func NewWorkloadScheduling_Override(w WorkloadScheduling, instance AbstractPod) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-30.WorkloadScheduling",
		[]interface{}{instance},
		w,
	)
}

func (w *jsiiProxy_WorkloadScheduling) Assign(node NamedNode) {
	if err := w.validateAssignParameters(node); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"assign",
		[]interface{}{node},
	)
}

func (w *jsiiProxy_WorkloadScheduling) Attract(node LabeledNode, options *PodSchedulingAttractOptions) {
	if err := w.validateAttractParameters(node, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"attract",
		[]interface{}{node, options},
	)
}

func (w *jsiiProxy_WorkloadScheduling) Colocate(selector IPodSelector, options *PodSchedulingColocateOptions) {
	if err := w.validateColocateParameters(selector, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"colocate",
		[]interface{}{selector, options},
	)
}

func (w *jsiiProxy_WorkloadScheduling) Separate(selector IPodSelector, options *PodSchedulingSeparateOptions) {
	if err := w.validateSeparateParameters(selector, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"separate",
		[]interface{}{selector, options},
	)
}

func (w *jsiiProxy_WorkloadScheduling) Spread(options *WorkloadSchedulingSpreadOptions) {
	if err := w.validateSpreadParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"spread",
		[]interface{}{options},
	)
}

func (w *jsiiProxy_WorkloadScheduling) Tolerate(node TaintedNode) {
	if err := w.validateTolerateParameters(node); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"tolerate",
		[]interface{}{node},
	)
}

