package model

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users = make([]*User, 0)
var id = 0

func SampleUser() User {
	return User{1, "Pedro", "pedro@gmail.com", "33232323"}
}

func InsertUser(user *User) (*User, error) {
	id += 1
	user.ID = id
	users = append(users, user)
	return user, nil
}

func FindAllUsers() ([]*User, error) {
	return users, nil
}