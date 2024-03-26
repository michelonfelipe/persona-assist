package data

import "github.com/michelonfelipe/persona-assist/internal"

func Orpheus() internal.Persona {
	return internal.Persona{
		Name: "Orpheus",
		Attacks: []internal.Attack{
			Bash(), Agi(), Agilao(),
		},
		Reactions: internal.Reactions{internal.Darkness: internal.Weak, internal.Electricity: internal.Weak},
	}
}

func Angel() internal.Persona {
	return internal.Persona{
		Name: "Angel",
		Attacks: []internal.Attack{
			SingleShot(), Garu(),
		},
		Reactions: internal.Reactions{internal.Darkness: internal.Weak},
	}
}

func Valkyrie() internal.Persona {
	return internal.Persona{
		Name: "Valkyrie",
		Attacks: []internal.Attack{
			SingleShot(), Garu(),
		},
		Reactions: internal.Reactions{internal.Fire: internal.Weak},
	}
}

func Penthesilea() internal.Persona {
	return internal.Persona{
		Name: "Penthesilea",
		Attacks: []internal.Attack{
			Bufu(), Mabufu(),
		},
		Reactions: internal.Reactions{internal.Fire: internal.Weak},
	}
}

func Hermes() internal.Persona {
	return internal.Persona{
		Name: "Hermes",
		Attacks: []internal.Attack{
			Cleave(), Agi(),
		},
		Reactions: internal.Reactions{internal.Wind: internal.Weak},
	}
}

func JackLantern() internal.Persona {
	return internal.Persona{
		Name: "Jack o' Lantern",
		Attacks: []internal.Attack{
			Agi(),
		},
		Reactions: internal.Reactions{internal.Ice: internal.Weak},
	}
}

func BlueSigil() internal.Persona {
	return internal.Persona{
		Name: "Blue Sigil",
		Attacks: []internal.Attack{
			Cleave(),
		},
		Reactions: internal.Reactions{internal.Strike: internal.Weak},
	}
}
