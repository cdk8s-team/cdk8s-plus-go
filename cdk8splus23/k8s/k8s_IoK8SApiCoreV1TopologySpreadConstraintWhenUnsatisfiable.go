package k8s


// WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfy the spread constraint.
//
// - DoNotSchedule (default) tells the scheduler not to schedule it. - ScheduleAnyway tells the scheduler to schedule the pod in any location,
// but giving higher precedence to topologies that would help reduce the
// skew.
// A constraint is considered "Unsatisfiable" for an incoming pod if and only if every possible node assignment for that pod would violate "MaxSkew" on some topology. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 3/1/1: | zone1 | zone2 | zone3 | | P P P |   P   |   P   | If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduled to zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfies MaxSkew(1). In other words, the cluster can still be imbalanced, but scheduler won't make it *more* imbalanced. It's a required field.
//
// Possible enum values:
// - `"DoNotSchedule"` instructs the scheduler not to schedule the pod when constraints are not satisfied.
// - `"ScheduleAnyway"` instructs the scheduler to schedule the pod even if constraints are not satisfied.
type IoK8SApiCoreV1TopologySpreadConstraintWhenUnsatisfiable string

const (
	// DoNotSchedule.
	IoK8SApiCoreV1TopologySpreadConstraintWhenUnsatisfiable_DO_NOT_SCHEDULE IoK8SApiCoreV1TopologySpreadConstraintWhenUnsatisfiable = "DO_NOT_SCHEDULE"
	// ScheduleAnyway.
	IoK8SApiCoreV1TopologySpreadConstraintWhenUnsatisfiable_SCHEDULE_ANYWAY IoK8SApiCoreV1TopologySpreadConstraintWhenUnsatisfiable = "SCHEDULE_ANYWAY"
)

