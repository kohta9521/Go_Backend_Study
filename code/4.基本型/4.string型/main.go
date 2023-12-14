package main

import "fmt"

// 型
// 文字列型

func main() {
	var s string = "Hello Go"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	// 数値を文字列として扱う
	var s1 string = "300"
	fmt.Println(s1)
	fmt.Printf("%T\n", s1)

	// 複数行にわたる文字列を扱う
	fmt.Println(`test
	test
		text
	`)

	// ダブルクオーテーションを文字列内で使用する方法
	fmt.Println("\"")
	fmt.Println(`"`)

	// 文字列はバイトの集合
	fmt.Println(s[0]) // 本来であれば[H]が取れるはず、、
	// 文字列への変換
}
