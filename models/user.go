package main

import "fmt"

type User struct {
	ID   int
	Name string
	Age  int
}

// Constructor
func NewUser(name string, age int) User {
	return User{
		Name: name,
		Age:  age,
	}
}

func main() {
	user := NewUser("@awakia", 29)
	fmt.Printf("user: %v\n", user.ID)

}
