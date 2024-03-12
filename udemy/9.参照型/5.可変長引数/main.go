package main

import "fmt"

// スライス
// 可変長引数

func Sum(s ...int) int  {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func main() {
	fmt.Println(Sum(1, 2, 3, 4, 5))

	fmt.Println(Sum())
	sl := []int{1, 2, 3}
	fmt.Println(Sum(sl...))
}