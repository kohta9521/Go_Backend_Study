package main

import "fmt"

// 変数

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

func main() {
	// 明示的な定義
	// var 変数名 型 = 値
	var i int = 100
	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	var i3 int
	var s3 string
	fmt.Println(i3, s3)

	i3 = 300
	s3 = "Go is cool lang"
	fmt.Println(i3, s3)

	// 値の更新
	i = 150
	fmt.Println(i)

	// 暗黙的な定義
	// 変数名 := 値
	i4 := 400
	fmt.Println(i4)

	i4 = 450
	fmt.Println(i4)

	// i4 = "Hello"

	// 暗黙的な定義は関数内でしか使用不可
	// 明示的な定義と暗黙的な定義は混在可能
	// しかし普段は明示的な定義を使用する
	// 他の人のことも考える
	// 関数内での変数は関数ないのでみ使用可能
	outer()

	// fmt.Println(s4)

	// go内では定義したものは必ず使用しないといけない
	// var s5 string = "Not Use"
}
