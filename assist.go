package assist

import (
	"github.com/michelonfelipe/persona-assist/checks"
	"github.com/michelonfelipe/persona-assist/internal"
)

type Checkable interface {
	Check(party internal.Party, enemies []internal.Persona) *internal.Result
}

func Assist(party internal.Party, enemies []internal.Persona) *internal.Result {
	checks := []Checkable{checks.SimpleCheck{}}

	for _, checkable := range checks {
		result := checkable.Check(party, enemies)
		if result != nil {
			return result
		}
	}

	return nil
}
