package main

import "fmt"

// pointer

func Double(i int) {
	i *= 2
}

func DoubleV2(i *int) {
	*i *= 2
}

func DoubleV3(s []int) {
	for i, v := range s {
		s[i] = v * 2
	}
}

func main() {
	var n int = 100
	fmt.Println(n)
	fmt.Println(&n)

	// Double(n)
	// fmt.Println(n)

	// var p *int = &n
	// fmt.Println(p)
	// fmt.Println(*p)

	// *p = 300
	// fmt.Println(n)
	// n = 200
	// fmt.Println(*p)

	DoubleV2(&n)
	fmt.Println(n)

	// DoubleV2(p)
	// fmt.Println(*p)

	var sl []int = []int{1, 2, 3}
	DoubleV3(sl)
	fmt.Println(sl)


}

