package data

import "github.com/michelonfelipe/persona-assist/internal"

func Bash() internal.Attack {
	return internal.Attack{
		Name:     "Bash",
		Type:     internal.Strike,
		Target:   internal.Single,
		Cost:     10,
		CostType: internal.HP,
	}
}

func Agi() internal.Attack {
	return internal.Attack{
		Name:     "Agi",
		Type:     internal.Fire,
		Target:   internal.Single,
		Cost:     3,
		CostType: internal.PE,
	}
}

func Agilao() internal.Attack {
	return internal.Attack{
		Name:     "Agilao",
		Type:     internal.Fire,
		Target:   internal.Single,
		Cost:     6,
		CostType: internal.PE,
	}
}

func SingleShot() internal.Attack {
	return internal.Attack{
		Name:     "Single Shot",
		Type:     internal.Pierce,
		Target:   internal.Single,
		Cost:     6,
		CostType: internal.HP,
	}
}

func Garu() internal.Attack {
	return internal.Attack{
		Name:     "Garu",
		Type:     internal.Wind,
		Target:   internal.Single,
		Cost:     3,
		CostType: internal.PE,
	}
}

func Cleave() internal.Attack {
	return internal.Attack{
		Name:     "Cleave",
		Type:     internal.Slash,
		Target:   internal.Single,
		Cost:     10,
		CostType: internal.HP,
	}
}

func Bufu() internal.Attack {
	return internal.Attack{
		Name:     "Bufu",
		Type:     internal.Ice,
		Target:   internal.Single,
		Cost:     4,
		CostType: internal.PE,
	}
}

func Mabufu() internal.Attack {
	return internal.Attack{
		Name:     "Mabufu",
		Type:     internal.Ice,
		Target:   internal.All,
		Cost:     8,
		CostType: internal.PE,
	}
}
