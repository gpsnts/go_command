package commands

import (
	r "trabgb/receptors"
)

// MoveUpCommand é a implementação do Command para o Character movimentar-se para cima
type MoveUpCommand struct {
	Character r.Character
}

// Exec executa a ação do Command (Mover para cima)
func (c *MoveUpCommand) Exec() {
	c.Character.MoveRight()
}
