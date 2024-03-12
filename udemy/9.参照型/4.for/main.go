package main

import "fmt"

// スライス
// for

func main() {
	sl := []string{"A", "B", "C"}
	fmt.Println(sl)

	// for _, v := range sl {
	// 	fmt.Println(v)
	// }

	// 古典的for
	for i := 0; i < len(sl); i ++ {
		fmt.Println(sl[i])
	}
}