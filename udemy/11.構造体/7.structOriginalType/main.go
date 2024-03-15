package main

import "fmt"

// struct
// original type

type MyInt int

func main() {
	var m1 MyInt
	fmt.Println(m1)
	fmt.Printf("%T\n", m1)

	// i := 100
	// fmt.Println(m1 + i)

}