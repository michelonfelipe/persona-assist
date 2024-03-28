package assist

import (
	"sync"

	"github.com/michelonfelipe/persona-assist/internal"
)

func Assist(party internal.Party, enemies []internal.Enemy) []*internal.Result {
	var wg sync.WaitGroup
	results := make([]*internal.Result, len(party))

	for index, member := range party {
		wg.Add(1)

		go func() {
			defer wg.Done()

			results[index] = check(member, enemies)
		}()
	}

	wg.Wait()
	return results
}
