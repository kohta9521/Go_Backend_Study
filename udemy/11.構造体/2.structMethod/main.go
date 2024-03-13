package main

import "fmt"

// 構造体
// struct method

type User struct {
	Name string
	Age  int
	// X, Y int
}

func (u User) SayName() {
	fmt.Println("Hello " + u.Name)
}

func (u User) SetName(name string) {
	u.Name = name
}

func (u *User) SetName2(name string) {
	u.Name = name
}

func main() {
    user1 := User{Name: "kohta", Age: 20}
	user1.SayName()

	user1.SetName("A")
	user1.SayName()

	user1.SetName2("B")
	user1.SayName()

}