package data

import "github.com/michelonfelipe/persona-assist/internal"

func Orpheus() internal.Persona {
	return internal.Persona{
		Name: "Orpheus",
		Attacks: []internal.Attack{
			Bash(), Agi(), Agilao(),
		},
		Weaknesses: []internal.AttackType{internal.Darkness, internal.Electricity},
	}
}

func Angel() internal.Persona {
	return internal.Persona{
		Name: "Angel",
		Attacks: []internal.Attack{
			SingleShot(), Garu(),
		},
		Weaknesses: []internal.AttackType{internal.Darkness},
	}
}

func Valkyrie() internal.Persona {
	return internal.Persona{
		Name: "Valkyrie",
		Attacks: []internal.Attack{
			SingleShot(), Garu(),
		},
		Weaknesses: []internal.AttackType{internal.Fire},
	}
}

func Penthesilea() internal.Persona {
	return internal.Persona{
		Name: "Penthesilea",
		Attacks: []internal.Attack{
			Bufu(), Mabufu(),
		},
		Weaknesses: []internal.AttackType{internal.Fire},
	}
}

func Hermes() internal.Persona {
	return internal.Persona{
		Name: "Hermes",
		Attacks: []internal.Attack{
			Cleave(), Agi(),
		},
		Weaknesses: []internal.AttackType{internal.Wind},
	}
}

func JackLantern() internal.Persona {
	return internal.Persona{
		Name: "Jack o' Lantern",
		Attacks: []internal.Attack{
			Agi(),
		},
		Weaknesses: []internal.AttackType{internal.Ice},
	}
}
