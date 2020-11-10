package commands

import (
	r "trabgb/receptors"
)

// MoveDownCommand é a implementação do Command para o Character movimentar-se para baixo
type MoveDownCommand struct {
	Character r.Character
}

// Exec executa a ação do Command (Mover para baixo)
func (c *MoveDownCommand) Exec() {
	c.Character.MoveDown()
}
