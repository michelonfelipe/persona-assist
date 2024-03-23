package data

import "github.com/michelonfelipe/persona-assist/internal"

func Protagonist() internal.Member {
	return internal.Member{
		Personas: []internal.Persona{Orpheus(), Angel(), Valkyrie()},
		PE:       100,
	}
}

func Mitsuru() internal.Member {
	return internal.Member{
		Personas: []internal.Persona{Penthesilea()},
		PE:       100,
	}
}

func Junpei() internal.Member {
	return internal.Member{
		Personas: []internal.Persona{Hermes()},
		PE:       100,
	}
}
