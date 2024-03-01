package nota

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Nota struct {
	Titulo    string    `json:"titulo"`
	Conteudo  string    `json:"conteudo"`
	CreatedAt time.Time `json:"created_at"`
}

func (nota Nota) Display() {
	fmt.Printf("O título da sua nota \"%v\" tem o seguinte conteúdo: \n\n%v\n", nota.Titulo, nota.Conteudo)
}

func (nota Nota) Salvar() error {
	filename := strings.ReplaceAll(nota.Titulo, " ", "_")
	filename = strings.ToLower(filename) + ".json"
	jsonFile, err := json.Marshal(nota)

	if err != nil {
		return err
	}

	return os.WriteFile(filename, jsonFile, 0644)
}

func New(titulo, conteudo string) Nota {
	nota := Nota{
		Titulo:    titulo,
		Conteudo:  conteudo,
		CreatedAt: time.Now(),
	}
	return nota
}
