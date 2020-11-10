package commands

import (
	r "trabgb/receptors"
)

// MoveRightCommand é a implementação do Command para o Character movimentar-se para direita
type MoveRightCommand struct {
	Character r.Character
}

// Exec executa a ação do Command (Mover para direita)
func (c *MoveRightCommand) Exec() {
	c.Character.MoveRight()
}
