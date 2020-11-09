package commands

import (
	d "trabgb/receptors"
)

// OnCommand - Estrutura do Command "On"
type OnCommand struct {
	Device d.Device
}

// Exec em OnCommand é a implementação de Command (interface) Exec
func (com *OnCommand) Exec() {
	com.Device.On()
}
