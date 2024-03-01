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

func (u User) OutputUserDetails() {
	fmt.Println(u.nome, u.sobrenome, u.dataDeNascimento)
}

func (u *User) ClearUserName() {
	u.nome = ""
	u.sobrenome = ""
}

func NewUser(nome, sobrenome, dataDeNascimento string) (*User, error) {
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
