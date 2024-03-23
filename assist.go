package assist

import (
	"github.com/michelonfelipe/persona-assist/checks"
	"github.com/michelonfelipe/persona-assist/internal"
)

type Checkable interface {
	Check(party internal.Party, weaknessByEnemy [][]internal.AttackType) *internal.Result
}

func Assist(party internal.Party, enemies []internal.Persona) *internal.Result {
	checks := []Checkable{checks.SimpleCheck{}}
	weaknessByEnemy := enemiesWeakness(enemies)

	for _, checkable := range checks {
		result := checkable.Check(party, weaknessByEnemy)
		if result != nil {
			return result
		}
	}

	return nil
}

func enemiesWeakness(enemies []internal.Persona) (result [][]internal.AttackType) {
	for _, enemy := range enemies {
		result = append(result, enemy.Weaknesses)
	}

	return
}
