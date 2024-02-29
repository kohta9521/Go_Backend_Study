package main

import "fmt"

// 型
// 配列型

func main() {
	// 明示的宣言
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	// 配列の初期化
	var arr2 [3]string = [3]string{"A", "B", "C"}
	fmt.Println(arr2)

	// 暗黙的な定義
	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	arr4 := [...]string{"D", "E", "F"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	// 値の取り出し
	fmt.Println(arr2[1])
	fmt.Println(arr2[2])
	fmt.Println(arr2[2-1])

	// 値の更新
	arr2[2] = "C"
	fmt.Println(arr2)

	var arr5 [4]int
	// arr5 = arr1
	fmt.Println(arr5)

	// 配列の要素数を数える

	fmt.Println(len(arr2))
}