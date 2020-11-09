package invocators

import (
	c "trabgb/commands"
)

// Button - estrutura (ou objeto dependendo da lang) que invoca os Command
type Button struct {
	Command c.Command
}

// Press - invoca o command
func (b *Button) Press() {
	b.Command.Exec()
}
