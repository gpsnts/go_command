package commands

// Command é a interface utilizada para implementar os commands do projeto
type Command interface {
	Exec()
}
