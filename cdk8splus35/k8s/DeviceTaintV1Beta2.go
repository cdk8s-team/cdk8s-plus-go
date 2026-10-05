package k8s

import (
	"time"
)

// The device this taint is attached to has the "effect" on any claim which does not tolerate the taint and, through the claim, to pods using the claim.
type DeviceTaintV1Beta2 struct {
	// The effect of the taint on claims that do not tolerate the taint and through such claims on the pods using them.
	//
	// Valid effects are None, NoSchedule and NoExecute. PreferNoSchedule as used for nodes is not valid here. More effects may get added in the future. Consumers must treat unknown effects like None.
	Effect *string `field:"required" json:"effect" yaml:"effect"`
	// The taint key to be applied to a device.
	//
	// Must be a label name.
	Key *string `field:"required" json:"key" yaml:"key"`
	// TimeAdded represents the time at which the taint was added.
	//
	// Added automatically during create or update if not set.
	TimeAdded *time.Time `field:"optional" json:"timeAdded" yaml:"timeAdded"`
	// The taint value corresponding to the taint key.
	//
	// Must be a label value.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

