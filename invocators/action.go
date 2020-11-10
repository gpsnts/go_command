package invocators

import (
	c "trabgb/commands"
)

// Action é responsável por chamar os Commands associados
type Action struct {
	Command c.Command
}

// Act age para invocar os Commands
func (a *Action) Act() {
	a.Command.Exec()
}
