package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)


func main() {
	x := Input("type a number")
	fmt.Print(x + "は、")
	if n, err := strconv.Atoi(x); err == nil {
		if n & 2 == 0 {
			fmt.Println("偶数です。")
		} else {
			fmt.Println("奇数です。")
		}
	} else {
		fmt.Println("整数ではありません。")
	}
	
}

func Input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg + ": ")
	scanner.Scan()
	return scanner.Text()
}