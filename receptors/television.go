package device

import "fmt"

// Tv - estrutura que coordena um estado para uma "TV" de acordo com o Command
type Tv struct {
	isRunning bool
}

// On - seta o estado da TV para ligada
func (t *Tv) On() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

// Off - seta o estado da TV para desligada
func (t *Tv) Off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}
