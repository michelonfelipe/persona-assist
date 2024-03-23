package data

import "github.com/michelonfelipe/persona-assist/internal"

func Bash() internal.Attack {
	return internal.Attack{
		Name:   "Bash",
		Type:   internal.Strike,
		Target: internal.Single,
		Cost:   0,
	}
}

func Agi() internal.Attack {
	return internal.Attack{
		Name:   "Agi",
		Type:   internal.Fire,
		Target: internal.Single,
		Cost:   3,
	}
}
func Agilao() internal.Attack {
	return internal.Attack{
		Name:   "Agilao",
		Type:   internal.Fire,
		Target: internal.Single,
		Cost:   6,
	}
}

func SingleShot() internal.Attack {
	return internal.Attack{
		Name:   "Single Shot",
		Type:   internal.Pierce,
		Target: internal.Single,
		Cost:   0,
	}
}

func Garu() internal.Attack {
	return internal.Attack{
		Name:   "Garu",
		Type:   internal.Wind,
		Target: internal.Single,
		Cost:   3,
	}
}

func Cleave() internal.Attack {
	return internal.Attack{
		Name:   "Cleave",
		Type:   internal.Slash,
		Target: internal.Single,
		Cost:   0,
	}
}

func Bufu() internal.Attack {
	return internal.Attack{
		Name:   "Bufu",
		Type:   internal.Ice,
		Target: internal.Single,
		Cost:   4,
	}
}

func Mabufu() internal.Attack {
	return internal.Attack{
		Name:   "Mabufu",
		Type:   internal.Ice,
		Target: internal.All,
		Cost:   8,
	}
}
