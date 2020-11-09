package main

import (
	c "trabgb/commands"
	i "trabgb/invocators"
	d "trabgb/receptors"
)

func main() {
	tv := &d.Tv{}

	onCommand := &c.OnCommand{
		Device: tv,
	}

	offCommand := &c.OffCommand{
		Device: tv,
	}

	onButton := &i.Button{
		Command: onCommand,
	}
	onButton.Press()

	offButton := &i.Button{
		Command: offCommand,
	}
	offButton.Press()
}
