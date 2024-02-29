package main

import "fmt"

// 定数

const Pi = 3.14

// 複数の定数
const (
	URL       = "http://google.com"
	SiteName  = "Google"
)

var Big int = 23090238402 + 1

const (
	c0 = iota
	c1
	c2
)


const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

func main() {
	fmt.Println(Pi)

	// Pi = 3
	// fmt.Println(Pi)

	fmt.Println(URL, SiteName)

	fmt.Println(A, B, C, D, E, F)

	fmt.Println(Big)

	fmt.Println(c0, c1, c2)
}