package assist

import (
	"github.com/michelonfelipe/persona-assist/internal"
)

func check(member internal.Member, enemies []internal.Enemy) (result *internal.Result) {
	for _, persona := range member.Personas {
		for _, attack := range persona.Attacks {
			toBeAffectedEnemies := toBeAffectedEnemies(attack, enemies)

			if !canUseAttack(toBeAffectedEnemies, member, attack) {
				continue
			}

			if result == nil || worthChangeAttack(result, attack, toBeAffectedEnemies) {
				result = &internal.Result{
					Persona: persona,
					Attack:  attack,
					Enemy:   toBeAffectedEnemies[0],
				}
			}
		}
	}

	return
}

func toBeAffectedEnemies(attack internal.Attack, enemies []internal.Enemy) []internal.Enemy {
	result := make([]internal.Enemy, 0)

	for index, enemy := range enemies {
		if enemy.Reactions[attack.Type] == internal.Weak {
			result = append(result, enemies[index])
		}
	}

	return result
}

func canUseAttack(enemies []internal.Enemy, member internal.Member, attack internal.Attack) bool {
	anyEnemiesAreWeakToAttack := len(enemies) > 0
	hasEnoughPE := attack.CostType == internal.PE && member.PE >= attack.Cost
	hasEnoughHealth := attack.CostType == internal.HP && member.CurrentHealth*attack.Cost/10 > member.MaxHealth/3

	return anyEnemiesAreWeakToAttack && (hasEnoughPE || hasEnoughHealth)
}

func worthChangeAttack(result *internal.Result, newAttack internal.Attack, enemies []internal.Enemy) bool {
	// if they are have the same target, pick cheaper one
	if newAttack.Target == result.Attack.Target {
		return newAttack.Cost < result.Attack.Cost
	}

	// choose the best cost effective attack
	toBeAffectedEnemies := len(enemies)

	if result.Attack.Target == internal.Single {
		return newAttack.Cost < result.Attack.Cost*toBeAffectedEnemies
	}

	return newAttack.Cost*toBeAffectedEnemies < result.Attack.Cost
}
