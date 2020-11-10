package receptors

// Character é a abstração de um personagem no programa
type Character interface {
	Attack()
	Defend()
	MoveUp()
	MoveDown()
	MoveLeft()
	MoveRight()
}
