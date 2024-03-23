package internal

type Party []Member

type Member struct {
	Personas      []Persona
	PE            int
	CurrentHealth int
	MaxHealth     int
}

type Persona struct {
	Name       string
	Attacks    []Attack
	Weaknesses []AttackType
}

type Enemy = Persona

type EnemiesByWeakness map[AttackType][]Enemy
