package assist

import (
	"testing"

	"github.com/michelonfelipe/persona-assist/data"
	"github.com/michelonfelipe/persona-assist/internal"
	"github.com/stretchr/testify/assert"
)

func TestNoResult(t *testing.T) {
	party := internal.Party{data.Protagonist}
	enemies := []internal.Persona{data.Angel}

	result := Assist(party, enemies)

	assert.Nil(t, result, "When enemies have no weakness to the current party, the result should be nil")
}

func TestSimpleWeakness(t *testing.T) {
	party := internal.Party{data.Protagonist}
	enemies := []internal.Persona{data.Valkyrie}

	result := Assist(party, enemies)
	expected := internal.Result{
		Member:  data.Protagonist,
		Persona: data.Orpheus,
		Attack:  data.Agi,
		Target:  data.Valkyrie,
	}

	assert.EqualValues(t, expected, *result, "If there is only one enemy, it should return the an attack that matches the enemy weakness")
}
