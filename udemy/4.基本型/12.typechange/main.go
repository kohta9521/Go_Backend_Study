package main

import (
	"fmt"
)

// 型
// 型変換


func main() {
	// var i int = 1
	// fl64 := float64(i)
	// fmt.Println(fl64)
	// fmt.Printf("i = %T\n", i)
	// fmt.Printf("fl64 = %T\n", fl64)

	// i2 := int(fl64)
	// fmt.Printf("i2 = %T\n", i2)

	// fl := 10.5
	// i3 := int(fl)
	// fmt.Println(i3)
	// fmt.Printf("i3 = %T\n", i3)

	// 文字列から数値
	// var s string = "100"
	// fmt.Println(s)
	// fmt.Printf("s = %T\n", s)

	// i, err := strconv.Atoi(s)
	// fmt.Println(i)
	// fmt.Printf("i = %T\n", i)
	// fmt.Println(err)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// var i2 int = 200
	// s2 := strconv.Itoa(i2)
	// fmt.Println(s2)
	// fmt.Printf("s2 = %T\n", s2)

	// 文字列とバイト配列
	var h string = "Hello World"
	b := []byte(h)
	fmt.Println(b)

	// 文字列へ変換
	h2 := string(b)
	fmt.Println(h2)
}