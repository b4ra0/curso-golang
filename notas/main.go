package main

import (
	"bufio"
	"errors"
	"example.com/notas/nota"
	"fmt"
	"os"
	"strings"
)

func main() {
	titulo, err := getUserInput("Titulo da nota:")
	handleError(err)
	conteudo, err := getUserInput("Conteúdo da nota:")
	handleError(err)
	userNota := nota.New(titulo, conteudo)
	userNota.Display()
	err = userNota.Salvar()
	handleError(err)
	fmt.Println("Nota salva com sucesso")
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func getUserInput(prompt string) (string, error) {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	handleError(err)

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	if text == "" {
		return "", errors.New("Input inválido.")
	}

	return text, nil
}
