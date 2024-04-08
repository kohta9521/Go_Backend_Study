package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	x := Input("type a number")
	n, err := strconv.Atoi(x)
	if err == nil {
		fmt.Print(x + "は、")
	} else {
		return
	}
	switch {
		case n & 2 == 0:
			fmt.Println("偶数です。")
		case n & 2 == 1:
			fmt.Println("奇数です。")
	}
}

func Input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg + ": ")
	scanner.Scan()
	return scanner.Text()
}