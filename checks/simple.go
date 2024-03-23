package checks

import (
	"github.com/michelonfelipe/persona-assist/internal"
)

type SimpleCheck struct{}

func (s SimpleCheck) Check(member internal.Member, enemiesByWeakness internal.EnemiesByWeakness) (result *internal.Result) {
	for _, persona := range member.Personas {
		for _, attack := range persona.Attacks {
			if len(enemiesByWeakness[attack.Type]) == 0 {
				continue
			}

			if attack.Target == internal.All {
				continue
			}

			if result == nil || result.Attack.Cost > attack.Cost {
				result = &internal.Result{
					Persona: persona,
					Attack:  attack,
					Enemy:   enemiesByWeakness[attack.Type][0],
				}
			}
		}
	}

	return
}
