package main

import "fmt"


func main() {
	var s, b int
	fmt.Scanf("%d %d", &s, &b)
	i := s * b
	if i % 2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}