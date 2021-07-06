package main

import (
	_ "fmt"
	"github.com/jairoandre/simple-auth/model"
)

func addUser(user model.User) model.User {
	rawPassword := user.Password
	user.Password, _ = HashPassword(rawPassword)
	att, _ := model.InsertUser(&user)
	return *att
}

func findAllUsers() []*model.User {
	users, _ := model.FindAllUsers()
	return users
}