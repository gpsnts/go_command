package receptors

import "fmt"

// Hero é a classe que representa o Heroi no programa
type Hero struct {
	position string
}

// Attack simboliza o "ataque" do Heroi
func (h *Hero) Attack() {
	fmt.Println("Atacou!")
}

// Defend simboliza a defesa do Heroi
func (h *Hero) Defend() {
	fmt.Println("Defendeu!")
}

// MoveUp simboliza o movimento para cima
func (h *Hero) MoveUp() {
	h.position = "Cima"
}

// MoveDown simboliza o movimento para baixo
func (h *Hero) MoveDown() {
	h.position = "Baixo"
}

// MoveLeft simboliza o movimento para esquerda
func (h *Hero) MoveLeft() {
	h.position = "Esquerda"
}

// MoveRight simboliza o movimento para direita
func (h *Hero) MoveRight() {
	h.position = "Direita"
}
