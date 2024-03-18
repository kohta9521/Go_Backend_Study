package main

import (
	"fmt"
	"os"
)

// os


func main() {
	// Exit
	// os.Exit(1)
	// fmt.Println("Start")


	// defer func() {
	// 	fmt.Println("defer")
	// }()
	// os.Exit(0)

	// _, err := os.Open("A.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// Args
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])
	fmt.Println(os.Args[3])

	// os.Argsの要素数を表示
	fmt.Printf("length=%d\n", len(os.Args))
	// os.Argsの中身を全て表示
	for i, v := range os.Args {
		fmt.Println(i, v)
	}
}