package main

import "fmt"

// 構造体
// struct

type User struct {
	Name string
	Age  int
	// X, Y int
}

func UpdateUser(user User) {
	user.Name = "A"
	user.Age = 1000
}

func UpdateUser2(user *User) {
	user.Name = "A"
	user.Age = 1000
}

func main() {
	var user1 User
	fmt.Println(user1)
	user1.Name = "Kohta"
	user1.Age = 20
	fmt.Println(user1)

	user2 := User{}
	fmt.Println(user2)
	user2.Name = "Sato"
	user2.Age = 30
	fmt.Println(user2)

	user3 := User{Name: "Tanaka", Age: 40}
	fmt.Println(user3)

	user4 := User{"user4", 40}
	fmt.Println(user4)

	// user5 := User{}

	user6 := User{Name: "user6"}
	fmt.Println(user6)

	// pointer
	user7 := new(User)
	fmt.Println(user7)

	// 関数の引数として構造体を渡す
	user8 := &User{}
	fmt.Println(user8)


	UpdateUser(user1)
	UpdateUser2(user8)

	fmt.Println(user1)
	fmt.Println(user8)
}
