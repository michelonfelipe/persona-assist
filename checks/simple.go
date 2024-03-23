package checks

import (
	"sort"

	"github.com/michelonfelipe/persona-assist/internal"
)

type SimpleCheck struct{}

func (s SimpleCheck) Check(party internal.Party, weaknessByEnemy [][]internal.AttackType) *internal.Result {
	for enemyIndex, enemyWeaknesses := range weaknessByEnemy {
		for _, member := range party {
			for _, persona := range member.Personas {
				for _, attack := range attacksSortedByCost(persona.Attacks) {
					if attack.Target == internal.All {
						continue
					}

					for _, weakness := range enemyWeaknesses {
						if attack.Type == weakness {
							return &internal.Result{
								Member:      member,
								Persona:     persona,
								Attack:      attack,
								TargetIndex: enemyIndex,
							}
						}
					}
				}
			}
		}
	}

	return nil
}

func attacksSortedByCost(a []internal.Attack) []internal.Attack {
	sorted := make([]internal.Attack, len(a))
	copy(sorted, a)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Cost < sorted[j].Cost
	})

	return sorted
}
