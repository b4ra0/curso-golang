package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

type Todo struct {
	Texto string `json:"texto"`
}

func (todo Todo) Display() {
	fmt.Println(todo.Texto)
}

func (todo Todo) Salvar() error {
	jsonFile, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile("todo.json", jsonFile, 0644)
}

func New(texto string) Todo {
	todo := Todo{
		Texto: texto,
	}
	return todo
}
