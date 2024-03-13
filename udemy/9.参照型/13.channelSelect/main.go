package main

import "fmt"

// channel
// select

func main() {
	ch1 := make(chan int, 2)
	ch2 := make(chan string, 2)

	ch2 <- "A"
	ch1 <- 1
	// v1 := <- ch1
	// v2 := <- ch2
	// fmt.Println(v1)
	// fmt.Println(v2)

	// 複数のchannelを扱う場合
	select {
	case v1 := <- ch1:
		fmt.Println(v1 + 1000)
	case v2 := <- ch2:
		fmt.Println(v2 + "!?")
	default:
		fmt.Println("どちらでもない")
	}

	// 活用例
	ch3 := make(chan int)
	ch4 := make(chan int)
	ch5 := make(chan int)

	// reciever
	go func() {
		for {
			i := <-ch3
			ch4 <- i * 2
		}
	}()

	// reciever
	go func() {
		for {
			i2 := <-ch4
			ch5 <- i2 - 1
		}
	}()

	n := 0
	for {
		select {
		case ch3 <- n:
			n ++
		case i3 := <- ch5:
			fmt.Println("recieved", i3)
		}
	}
}