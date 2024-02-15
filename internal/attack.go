package internal

type AttackType int

// Using Persona 3 Reloaded types
const (
	// Physical
	Slash AttackType = iota
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
	Single AttackTarget = iota
	All
)

type Attack struct {
	Name   string
	Type   AttackType
	Target AttackTarget
	Cost   int
}
