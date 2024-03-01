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

	appUser, err := user.NewUser(firstName, lastName, birthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.OutputUserDetails()

	appUser.ClearUserName()

	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
