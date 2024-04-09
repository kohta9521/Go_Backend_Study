package main

import "fmt"

func total(n int, c chan int) {
	t := 0
	for i := 1; i <= n; i ++ {
		t += i
	}
	c <- t
}

func main() {
	c := make(chan int)
	go total(100, c)
	g := <-c
	fmt.Println("total:", g)
}