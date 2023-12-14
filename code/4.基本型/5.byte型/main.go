package main

import "fmt"

// 型
// byte型

func main() {
	byteA := []byte{72, 73}
	fmt.Println(byteA)

	// 文字列への変換
	fmt.Println(string(byteA))

	// 文字列はアスキーキーコード

	c := []byte("HI")
	fmt.Println(c)

	fmt.Println(string(c))
}
