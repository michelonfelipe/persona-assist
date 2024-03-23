package assist

import (
	"github.com/michelonfelipe/persona-assist/internal"
)

func check(member internal.Member, enemiesByWeakness internal.EnemiesByWeakness) (result *internal.Result) {
	for _, persona := range member.Personas {
		for _, attack := range persona.Attacks {
			if !canUseAttack(enemiesByWeakness, member, attack) {
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

func canUseAttack(e internal.EnemiesByWeakness, member internal.Member, attack internal.Attack) bool {
	anyEnemiesAreWeakToAttack := len(e[attack.Type]) > 0
	hasEnoughPE := attack.CostType == internal.PE && member.PE >= attack.Cost
	hasEnoughHealth := attack.CostType == internal.HP && member.CurrentHealth*attack.Cost/10 > member.MaxHealth/3

	return anyEnemiesAreWeakToAttack && (hasEnoughPE || hasEnoughHealth)
}
