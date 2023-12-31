package main

import (
	"errors"
	"fmt"
)

type Data struct {
	Username string
	Password string
}

type User interface {
	GetUser() string
}
type Password interface {
	GetPass() string
}

func CekLogin(user User, pass Password) (bool, error) {
	if user.GetUser() == "admin" && pass.GetPass() == "123456" {
		return true, nil
	} else {
		return false, errors.New("Username atau Password Salah!")
	}
}

func (data Data) GetUser() string {
	return data.Username
}
func (data Data) GetPass() string {
	return data.Password
}

func main() {
	data_account := map[string]string{
		"user": "admin",
		"pass": "123456",
	}

	condition, err := CekLogin(Data{
		Username: data_account["user"],
	}, Data{Password: data_account["pass"]})

	if condition == true {
		fmt.Println("Success Login Berhasil!")
	} else {
		fmt.Println("Error", err.Error())
	}
}
