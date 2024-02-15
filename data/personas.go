package data

import "github.com/michelonfelipe/persona-assist/internal"

var Orpheus = internal.Persona{
	Name: "Orpheus",
	Attacks: []internal.Attack{
		Bash, Agi,
	},
	Weaknesses: []internal.AttackType{internal.Darkness, internal.Electricity},
}

var Angel = internal.Persona{
	Name: "Angel",
	Attacks: []internal.Attack{
		SingleShot, Garu,
	},
	Weaknesses: []internal.AttackType{internal.Darkness},
}

var Valkyrie = internal.Persona{
	Name: "Valkyrie",
	Attacks: []internal.Attack{
		SingleShot, Garu,
	},
	Weaknesses: []internal.AttackType{internal.Fire},
}
