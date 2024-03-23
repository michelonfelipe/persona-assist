package assist

import (
	"sync"

	"github.com/michelonfelipe/persona-assist/checks"
	"github.com/michelonfelipe/persona-assist/internal"
)

type Checkable interface {
	Check(member internal.Member, enemiesByWeakness internal.EnemiesByWeakness) *internal.Result
}

func Assist(party internal.Party, enemies []internal.Enemy) []*internal.Result {
	var wg sync.WaitGroup
	results := make([]*internal.Result, len(party))
	checks := []Checkable{checks.SimpleCheck{}}
	weaknessByEnemy := enemiesWeakness(enemies)

	for index, member := range party {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for _, checkable := range checks {
				memberResult := checkable.Check(member, weaknessByEnemy)
				if memberResult != nil {
					results[index] = memberResult
					break
				}
			}
		}()
	}
	wg.Wait()
	return results
}

func enemiesWeakness(enemies []internal.Enemy) internal.EnemiesByWeakness {
	result := make(internal.EnemiesByWeakness)

	for _, enemy := range enemies {
		for _, weakness := range enemy.Weaknesses {
			result[weakness] = append(result[weakness], enemy)
		}
	}

	return result
}
