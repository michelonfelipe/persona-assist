package assist

import (
	"github.com/michelonfelipe/persona-assist/internal"
)

func check(member internal.Member, enemiesByWeakness internal.EnemiesByWeakness) (result *internal.Result) {
	for _, persona := range member.Personas {
		for _, attack := range persona.Attacks {
			if len(enemiesByWeakness[attack.Type]) == 0 || // any enemy have a weakness to the attack
				member.PE < attack.Cost { // member has enough PE to use the attack
				continue
			}

			if result == nil || worthChangeAttack(result, attack, enemiesByWeakness) {
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

func worthChangeAttack(result *internal.Result, newAttack internal.Attack, enemiesByWeakness internal.EnemiesByWeakness) bool {
	// if they are have the same target, pick cheaper one
	if newAttack.Target == result.Attack.Target {
		return newAttack.Cost < result.Attack.Cost
	}

	// choose the best cost effective attack
	toBeAffectedEnemies := len(enemiesByWeakness[newAttack.Type])

	if result.Attack.Target == internal.Single {
		return newAttack.Cost < result.Attack.Cost*toBeAffectedEnemies
	}

	return newAttack.Cost*toBeAffectedEnemies < result.Attack.Cost
}
