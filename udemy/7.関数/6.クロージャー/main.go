package main

import "fmt"

// 関数
// クロージャーの実装

func Later() func(string) string {
	var store string
	// storeは初期値""を持つ
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func main() {
	f := Later()
	fmt.Println(f("Hello"))
	fmt.Println(f("my"))
	fmt.Println(f("name"))
	fmt.Println(f("is"))
	fmt.Println(f("go"))
	fmt.Println(f("lang"))
}