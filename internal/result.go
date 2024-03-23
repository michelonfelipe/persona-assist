package internal

type Result struct {
	Persona
	Attack
	Enemy // If it's an All attack, this would be disconsidered
}
