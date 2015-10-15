package db

import "strconv"

type User struct {
	ID   int
	Name string
	Age  int
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
