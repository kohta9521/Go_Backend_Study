package main

import "fmt"

// if
// 条件分岐

func main() {
	a := 0
	if a == 2 {
		fmt.Println("two")
	} else if ( a ==1 ) {
		fmt.Println("one")
	} else {
		fmt.Println("i dont know")
	}

	// 簡易
	if b := 100; b == 100 {
		fmt.Println("B is one hundred")
	}

	// x := 0
	// if x == 2; true {
	// 	fmt.Println(x)
	// }
	// fmt.Println(x)
}