package internal

type Party []Member

type Member struct {
	Personas      []Persona
	PE            int
	CurrentHealth int
	MaxHealth     int
}

type Reaction int

const (
	Weak Reaction = iota
	Resists
	Absorbs
	Reflects
)

type Reactions map[AttackType]Reaction

type Persona struct {
	Name    string
	Attacks []Attack
	Reactions
}

type Enemy = Persona

type EnemiesByWeakness map[AttackType][]Enemy
