package internal

type AttackType int

// Using Persona 3 Reloaded types
const (
	// Physical
	Slash AttackType = iota + 1
	Strike
	Pierce

	// Magical
	Fire
	Ice
	Electricity
	Wind
	Light
	Darkness
)

type AttackTarget int

const (
	Single AttackTarget = iota + 1
	All
)

type AttackCostType int

const (
	PE AttackCostType = iota + 1
	HP
)

type Attack struct {
	Name     string
	Type     AttackType
	Target   AttackTarget
	Cost     int
	CostType AttackCostType // PE points are straight up ints, but HP points are percentage based
}
