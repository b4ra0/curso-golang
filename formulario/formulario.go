package main

import (
	"exemple.com/structs/user"
	"fmt"
)

func main() {
	firstName := getUserData("Digite seu nome: ")
	lastName := getUserData("Digite seu sobrenome: ")
	birthdate := getUserData("Digite sua data de nascimento (DD/MM/AAAA): ")

	var appUser *user.User

	appUser, err := user.New(firstName, lastName, birthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("lucas.baraom@gmail.com", "teste987")

	admin.OutputDetails()
	admin.ClearName()
	admin.OutputDetails()

	appUser.OutputDetails()
	appUser.ClearName()
	appUser.OutputDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
