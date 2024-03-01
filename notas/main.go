package main

import (
	"bufio"
	"errors"
	"example.com/notas/nota"
	"example.com/notas/todo"
	"fmt"
	"os"
	"strings"
)

type outputtable interface {
	Salvar() error
	Display()
}

func main() {
	texto := getTodoData()
	userTodo := todo.New(texto)
	exportarDados(userTodo)
	fmt.Println("Todo salvo com sucesso")
	titulo, conteudo := getNotaData()
	userNota := nota.New(titulo, conteudo)
	exportarDados(userNota)
	fmt.Println("Nota salva com sucesso")
}

func exportarDados(dados outputtable) {
	dados.Display()
	err := dados.Salvar()
	handleError(err)
}

func getTodoData() string {
	texto, err := getUserInput("Texto do todo:")
	handleError(err)
	return texto
}

func getNotaData() (string, string) {
	titulo, err := getUserInput("Titulo da nota:")
	handleError(err)
	conteudo, err := getUserInput("Conteúdo da nota:")
	handleError(err)
	return titulo, conteudo
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
