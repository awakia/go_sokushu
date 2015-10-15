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
		1,
		name,
		age,
	}
}

func main() {
	fmt.Printf("user: %v\n", NewUser("@awakia", 29))
}
