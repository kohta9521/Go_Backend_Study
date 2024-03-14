package main

import "fmt"

// struct
// constructor

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

func main() {
	user1 := NewUser("kohta", 20)
	fmt.Println(user1)

	fmt.Println(*user1)

}