package main

import "fmt"

type User struct {
	ID   int
	Name string
	Age  int
}

func main() {
	user := User{
		1,
		"@awakia",
		29,
	}
	fmt.Printf("user: %v\n", user)
}
