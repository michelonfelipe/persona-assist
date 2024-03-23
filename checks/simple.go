package checks

import (
	"github.com/michelonfelipe/persona-assist/internal"
)

type SimpleCheck struct{}

func (s SimpleCheck) Check(member internal.Member, weaknessByEnemy [][]internal.AttackType) (result *internal.Result) {
	for enemyIndex, enemyWeaknesses := range weaknessByEnemy {
		for _, persona := range member.Personas {
			for _, attack := range persona.Attacks {
				if attack.Target == internal.All {
					continue
				}

				for _, weakness := range enemyWeaknesses {
					if attack.Type != weakness {
						continue
					}

					if result == nil || result.Attack.Cost > attack.Cost {
						result = &internal.Result{
							Persona:     persona,
							Attack:      attack,
							TargetIndex: enemyIndex,
						}
					}
				}
			}
		}
	}

	return
}
