package main

import (
	c "trabgb/commands"
	r "trabgb/receptors"
)

func main() {
	hero := &r.Hero{}

	attackCommand := &c.AttackCommand{
		Character: hero,
	}

	attackCommand.Exec()

	defendCommand := &c.DefendCommand{
		Character: hero,
	}

	defendCommand.Exec()

}
