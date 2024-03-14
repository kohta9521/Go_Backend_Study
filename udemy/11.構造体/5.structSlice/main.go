package main

import "fmt"

// struct
// slice


type User struct {
	Name string
	Age  int
}

type Users []*User

// type Users struct {
// 	Users []*Users
// }

func main() {
	user1 := User{Name: "user1", Age: 10}
	user2 := User{Name: "user2", Age: 20}
	user3 := User{Name: "user3", Age: 30}
	user4 := User{Name: "user4", Age: 40}

	users := Users{}
	users = append(users, &user1)
	users = append(users, &user2)
	users = append(users, &user3, &user4)
	
	for _, u := range users {
		fmt.Println(*u)
	}

	// useres2 := make([]*Users, 0)
	// users2 = appned(users2, &users)
}