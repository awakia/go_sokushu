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

func (u *User) String() string {
	return u.Name + "(" + strconv.Itoa(u.Age) + ")"
}
