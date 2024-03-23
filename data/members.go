package data

import "github.com/michelonfelipe/persona-assist/internal"

func Protagonist() internal.Member {
	return internal.Member{
		Personas:      []internal.Persona{Orpheus(), Angel(), Valkyrie()},
		PE:            100,
		MaxHealth:     300,
		CurrentHealth: 300,
	}
}

func Mitsuru() internal.Member {
	return internal.Member{
		Personas:      []internal.Persona{Penthesilea()},
		PE:            100,
		MaxHealth:     300,
		CurrentHealth: 300,
	}
}

func Junpei() internal.Member {
	return internal.Member{
		Personas:      []internal.Persona{Hermes()},
		PE:            100,
		MaxHealth:     300,
		CurrentHealth: 300,
	}
}
