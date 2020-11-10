package commands

import (
	r "trabgb/receptors"
)

// AttackCommand é a implementação do Command para o Character atacar
type AttackCommand struct {
	Character r.Character
}

// Exec executa a ação do Command (Ataque)
func (c *AttackCommand) Exec() {
	c.Character.Attack()
}
