package main

import "fmt"

// 型
// 浮動小数点型

func main() {
	var fl64 float64 = 2.4
	fmt.Println(fl64)

	// 暗黙的な定義
	fl := 3.2

	// intと違い環境依存ではない
	fmt.Println(fl64 + fl)

	// 型を調べるとどちらともfloat64
	fmt.Printf("%T, %T\n", fl64, fl)

	// 型指定
	var fl32 float32 = 1.2
	// float32は基本的にはあまり使用しない
	fmt.Printf("%T\n", fl32)


	// 型変換
	fmt.Printf("%T\n", float64(fl32))

	// 演算が不能な特殊な型？
	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf) // +Inf

	ninf := -1.0 / zero
	fmt.Println(ninf) // -Inf

	nan := zero / zero
	fmt.Println(nan) // NaN(Not a Number)
}
