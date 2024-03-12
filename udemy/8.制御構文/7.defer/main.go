package main

import (
	"fmt"
)

// defer

func TestDefer() {
	defer fmt.Println("end")
	fmt.Println("start")
}

func RunDefer() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

func main() {
	TestDefer()

	defer func() {
		fmt.Println("1")
		fmt.Println("2")
		fmt.Println("3")
	}()

	RunDefer()

}