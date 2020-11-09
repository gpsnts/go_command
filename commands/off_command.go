package commands

import (
	d "trabgb/receptors"
)

// OffCommand - Estrutura do Command "Off"
type OffCommand struct {
	Device d.Device
}

// Exec em OnCommand é a implementação de Command (interface) Exec
func (com *OffCommand) Exec() {
	com.Device.Off()
}
