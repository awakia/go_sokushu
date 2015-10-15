package models

import "strconv"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Constructor
func NewUser(name string, age int) *User {
	return &User{
		Name: name,
		Age:  age,
	}
}

func (u *User) String() string {
	return u.Name + "(" + strconv.Itoa(u.Age) + ")"
}
