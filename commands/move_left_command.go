package commands

import (
	r "trabgb/receptors"
)

// MoveLeftCommand é a implementação do Command para o Character movimentar-se para esquerda
type MoveLeftCommand struct {
	Character r.Character
}

// Exec executa a ação do Command (Mover para esquerda)
func (c *MoveLeftCommand) Exec() {
	c.Character.MoveLeft()
}
