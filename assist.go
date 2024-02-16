package assist

import "github.com/michelonfelipe/persona-assist/internal"

type Result struct {
	internal.Member
	internal.Persona
	internal.Attack
	Target internal.Persona
}

func Assist(party internal.Party, enemies []internal.Persona) *Result {
	for _, enemy := range enemies {
		for _, member := range party {
			for _, persona := range member.Personas {
				for _, attack := range persona.Attacks {
					for _, weakness := range enemy.Weaknesses {
						if attack.Type == weakness {
							return &Result{
								Member:  member,
								Persona: persona,
								Attack:  attack,
								Target:  enemy,
							}
						}
					}
				}
			}
		}
	}
	return nil
}
