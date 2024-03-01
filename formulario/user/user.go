package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	nome             string
	sobrenome        string
	dataDeNascimento string
	createdAt        time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func (u User) OutputDetails() {
	fmt.Println(u.nome, u.sobrenome, u.dataDeNascimento)
}

func (u *User) ClearName() {
	u.nome = ""
	u.sobrenome = ""
}

func New(nome, sobrenome, dataDeNascimento string) (*User, error) {
	if nome == "" || sobrenome == "" || dataDeNascimento == "" {
		return nil, errors.New("Nome, sobrenome ou data de nascimento não digitada")
	}
	return &User{
		nome:             nome,
		sobrenome:        sobrenome,
		dataDeNascimento: dataDeNascimento,
		createdAt:        time.Now(),
	}, nil
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			nome:             "Admin",
			sobrenome:        "Admin",
			dataDeNascimento: "---",
			createdAt:        time.Now(),
		},
	}
}
