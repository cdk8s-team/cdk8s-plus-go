package k8s


// ExemptPriorityLevelConfiguration describes the configurable aspects of the handling of exempt requests.
//
// In the mandatory exempt configuration object the values in the fields here can be modified by authorized users, unlike the rest of the `spec`.
type ExemptPriorityLevelConfigurationV1Beta2 struct {
	// `lendablePercent` prescribes the fraction of the level's NominalCL that can be borrowed by other priority levels.
	//
	// This value of this field must be between 0 and 100, inclusive, and it defaults to 0. The number of seats that other levels can borrow from this level, known as this level's LendableConcurrencyLimit (LendableCL), is defined as follows.
	//
	// LendableCL(i) = round( NominalCL(i) * lendablePercent(i)/100.0 )
	LendablePercent *float64 `field:"optional" json:"lendablePercent" yaml:"lendablePercent"`
	// `nominalConcurrencyShares` (NCS) contributes to the computation of the NominalConcurrencyLimit (NominalCL) of this level.
	//
	// This is the number of execution seats nominally reserved for this priority level. This DOES NOT limit the dispatching from this priority level but affects the other priority levels through the borrowing mechanism. The server's concurrency limit (ServerCL) is divided among all the priority levels in proportion to their NCS values:
	//
	// NominalCL(i)  = ceil( ServerCL * NCS(i) / sum_ncs ) sum_ncs = sum[priority level k] NCS(k)
	//
	// Bigger numbers mean a larger nominal concurrency limit, at the expense of every other priority level. This field has a default value of zero.
	NominalConcurrencyShares *float64 `field:"optional" json:"nominalConcurrencyShares" yaml:"nominalConcurrencyShares"`
}

