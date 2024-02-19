package checks

import "github.com/michelonfelipe/persona-assist/internal"

type SimpleCheck struct{}

func (s SimpleCheck) Check(party internal.Party, enemies []internal.Persona) *internal.Result {
	for _, enemy := range enemies {
		for _, member := range party {
			for _, persona := range member.Personas {
				for _, attack := range persona.Attacks {
					for _, weakness := range enemy.Weaknesses {
						if attack.Type == weakness {
							return &internal.Result{
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
