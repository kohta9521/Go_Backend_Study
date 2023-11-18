package main

import "fmt"

func main() {
	var n int = 100
	fmt.Println(n)

	x := 200
	fmt.Println(x)

	// 文字列の連結
	s := "Hello, " + "world!"
	fmt.Println(s)

	// const (定数宣言)
	const y = 1
	fmt.Println(y)

	// 列挙型
	const (
		Appple = iota
		Orange
		Banana
	)
}
