package data

import "github.com/michelonfelipe/persona-assist/internal"

var Orpheus = internal.Persona{
	Name: "Orpheus",
	Attacks: []internal.Attack{
		Bash, Agi, Agilao,
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

var Penthesilea = internal.Persona{
	Name: "Penthesilea",
	Attacks: []internal.Attack{
		Bufu, Mabufu,
	},
	Weaknesses: []internal.AttackType{internal.Fire},
}

var Hermes = internal.Persona{
	Name: "Hermes",
	Attacks: []internal.Attack{
		Cleave, Agi,
	},
	Weaknesses: []internal.AttackType{internal.Wind},
}

var JackLantern = internal.Persona{
	Name: "Jack o' Lantern",
	Attacks: []internal.Attack{
		Agi,
	},
	Weaknesses: []internal.AttackType{internal.Ice},
}
