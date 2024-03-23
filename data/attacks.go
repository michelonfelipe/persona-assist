package data

import "github.com/michelonfelipe/persona-assist/internal"

var Bash = internal.Attack{
	Name:   "Bash",
	Type:   internal.Strike,
	Target: internal.Single,
	Cost:   0,
}

var Agi = internal.Attack{
	Name:   "Agi",
	Type:   internal.Fire,
	Target: internal.Single,
	Cost:   3,
}
var Agilao = internal.Attack{
	Name:   "Agilao",
	Type:   internal.Fire,
	Target: internal.Single,
	Cost:   6,
}

var SingleShot = internal.Attack{
	Name:   "Single Shot",
	Type:   internal.Pierce,
	Target: internal.Single,
	Cost:   0,
}

var Garu = internal.Attack{
	Name:   "Garu",
	Type:   internal.Wind,
	Target: internal.Single,
	Cost:   3,
}

var Cleave = internal.Attack{
	Name:   "Cleave",
	Type:   internal.Slash,
	Target: internal.Single,
	Cost:   0,
}

var Bufu = internal.Attack{
	Name:   "Bufu",
	Type:   internal.Ice,
	Target: internal.Single,
	Cost:   4,
}

var Mabufu = internal.Attack{
	Name:   "Mabufu",
	Type:   internal.Ice,
	Target: internal.All,
	Cost:   8,
}
