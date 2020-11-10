package commands

import (
	r "trabgb/receptors"
)

// DefendCommand é a implementação do Command para o Character defender-se
type DefendCommand struct {
	Character r.Character
}

// Exec executa a ação do Command (Defesa)
func (c *DefendCommand) Exec() {
	c.Character.Defend()
}
