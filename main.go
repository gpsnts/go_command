package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	c "trabgb/commands"
	i "trabgb/invocators"
	r "trabgb/receptors"
)

func main() {
	// "Instância" de Hero
	hero := &r.Hero{}

	// Commands
	attackCommand := &c.AttackCommand{
		Character: hero,
	}

	defendCommand := &c.DefendCommand{
		Character: hero,
	}

	moveLeftCommand := &c.MoveLeftCommand{
		Character: hero,
	}

	moveRightCommand := &c.MoveRightCommand{
		Character: hero,
	}

	moveUpCommand := &c.MoveUpCommand{
		Character: hero,
	}

	moveDownCommand := &c.MoveDownCommand{
		Character: hero,
	}

	// Event-loop
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("--> ")
		text, _ := reader.ReadString('\n')

		// LF parser (*NIX-like)
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("0", text) == 0 {
			break
		}

		var action *i.Action

		switch {
		case strings.Compare("z", text) == 0:
			action = &i.Action{
				Command: attackCommand,
			}
		case strings.Compare("x", text) == 0:
			action = &i.Action{
				Command: defendCommand,
			}
		case strings.Compare("w", text) == 0:
			action = &i.Action{
				Command: moveUpCommand,
			}
		case strings.Compare("s", text) == 0:
			action = &i.Action{
				Command: moveDownCommand,
			}
		case strings.Compare("a", text) == 0:
			action = &i.Action{
				Command: moveLeftCommand,
			}
		case strings.Compare("d", text) == 0:
			action = &i.Action{
				Command: moveRightCommand,
			}
		case strings.Compare("exit", text) == 0:
			return
		}

		if action != nil {
			fmt.Println("Ação feita:")
			action.Act()
		}
	}
}
