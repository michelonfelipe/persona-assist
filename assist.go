package assist

import (
	"sync"

	"github.com/michelonfelipe/persona-assist/internal"
)

func Assist(party internal.Party, enemies []internal.Enemy) []*internal.Result {
	var wg sync.WaitGroup
	results := make([]*internal.Result, len(party))
	weaknessByEnemy := enemiesWeakness(enemies)

	for index, member := range party {
		wg.Add(1)

		go func() {
			defer wg.Done()

			results[index] = check(member, weaknessByEnemy)
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
