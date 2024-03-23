package assist

import (
	"testing"

	"github.com/michelonfelipe/persona-assist/data"
	"github.com/michelonfelipe/persona-assist/internal"
	"github.com/stretchr/testify/assert"
)

func TestNoResult(t *testing.T) {
	party := internal.Party{data.Protagonist()}
	enemies := []internal.Enemy{data.Angel()}

	result := Assist(party, enemies)

	assert.Nil(t, result[0], "When enemies have no weakness to the current party, the result should be nil")
}

func TestAllPartyResults(t *testing.T) {
	party := internal.Party{data.Protagonist(), data.Junpei(), data.Mitsuru()}
	enemies := []internal.Enemy{data.Valkyrie(), data.JackLantern()}

	result := Assist(party, enemies)
	expected := []*internal.Result{
		{
			Persona: data.Orpheus(),
			Attack:  data.Agi(),
			Enemy:   enemies[0],
		},
		{
			Persona: data.Hermes(),
			Attack:  data.Agi(),
			Enemy:   enemies[0],
		},
		{
			Persona: data.Penthesilea(),
			Attack:  data.Bufu(),
			Enemy:   enemies[1],
		},
	}

	assert.EqualValues(t, expected, result, "Each result item should represent the best attack for the party member with same index")
}

func TestSimpleWeakness(t *testing.T) {
	party := internal.Party{data.Protagonist()}
	enemies := []internal.Enemy{data.Valkyrie()}

	result := Assist(party, enemies)
	expected := &internal.Result{
		Persona: data.Orpheus(),
		Attack:  data.Agi(),
		Enemy:   enemies[0],
	}

	assert.EqualValues(t, expected, result[0], "If there is only one enemy, it should return the attack that matches the enemy weakness, and has the least cost")
}

func TestAllAttack(t *testing.T) {
	party := internal.Party{data.Mitsuru()}
	enemies := []internal.Enemy{data.JackLantern(), data.JackLantern(), data.JackLantern()}

	result := Assist(party, enemies)
	expected := &internal.Result{
		Persona: data.Penthesilea(),
		Attack:  data.Mabufu(),
		Enemy:   enemies[0],
	}

	assert.EqualValues(t, expected, result[0], "If the cost of a All attack is less than N Single attacks, the All attack should be chosen")
}

func TestEnoughPE(t *testing.T) {
	junpei := data.Junpei()
	junpei.PE = 0
	party := internal.Party{junpei}
	enemies := []internal.Enemy{data.Valkyrie()}

	result := Assist(party, enemies)

	assert.Nil(t, result[0], "If the player has not enough PE to cast an attack, it should not be chosen")
}

func TestEnoughHealth(t *testing.T) {
	protagonist := data.Protagonist()
	protagonist.CurrentHealth = protagonist.MaxHealth / 10
	party := internal.Party{protagonist}
	enemies := []internal.Enemy{data.BlueSigil()}

	result := Assist(party, enemies)

	assert.Nil(t, result[0], "If the player would end up with less than 30%% of their max health, a health cost attack should not be choosen")
}

func TestChooseHealthCostAttack(t *testing.T) {
	party := internal.Party{data.Protagonist()}
	enemies := []internal.Enemy{data.BlueSigil()}

	result := Assist(party, enemies)
	expected := &internal.Result{
		Persona: data.Orpheus(),
		Attack:  data.Bash(),
		Enemy:   enemies[0],
	}
	assert.EqualValues(t, expected, result[0], "Should choose a health based attack if viable")
}
