package internal

type Party []Member

type Member struct {
	Personas []Persona
	PE       int
}

type Persona struct {
	Name       string
	Attacks    []Attack
	Weaknesses []AttackType
}
