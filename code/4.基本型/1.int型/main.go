package main

import "fmt"

// 型
// int

func main() {
	var i int = 100
	fmt.Println(i)

	var i2 int64 = 200
	fmt.Println(i + int(i2))

	// 型変換
	// 現在の型を調べる方法
	fmt.Printf("%T\n", i2)

	// %Tは型を表示する 書式指定子と呼ばれている

	fmt.Printf("%T\n", int32(i2))
}
