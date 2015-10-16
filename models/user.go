package models

import (
	"strconv"

	"github.com/awakia/go_old/db"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Constructor
func NewUser(name string, age int) *User {
	user := &User{
		Name: name,
		Age:  age,
	}
	db.Get().Create(user)
	return user
}

func GetUser(id int) *User {
	var user *User
	db.Get().Where("id = ?", id).First(user)
	return user
}

func AllUsers() []*User {
	var users []*User
	db.Get().Find(&users)
	return users
}

func (u *User) String() string {
	return u.Name + "(" + strconv.Itoa(u.Age) + ")"
}
