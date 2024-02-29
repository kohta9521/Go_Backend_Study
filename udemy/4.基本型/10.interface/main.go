package main

import "fmt"

// 型
// interface

func main() {
	var x interface{}
	fmt.Println(x)

	x = 1
	fmt.Println(x)

	x = 3.14
	fmt.Println(x)

	// 演算の対処としては利用不可
}