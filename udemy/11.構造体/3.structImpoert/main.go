package main

import "fmt"

// struct
// 埋め込み

type T struct {
	User
}

type User struct {
	Name string
	Age  int
}

func (u *User) SetName() {
	u.Name = "A"

}

func main() {
	t := T{User: User{Name: "kohta", Age: 10}}
	fmt.Println(t)

	fmt.Println(t.User.Name)

	fmt.Println(t.Name)

	t.User.SetName()
	fmt.Println(t)

}