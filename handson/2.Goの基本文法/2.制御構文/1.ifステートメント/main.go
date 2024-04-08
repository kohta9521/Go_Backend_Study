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
	if err != nil {
		fmt.Println("ERROR!!")
		return
	}
	fmt.Print(x + "は、")
	if n % 2 == 0 {
		fmt.Println("偶数です。")
	} else {
		fmt.Println("奇数です。")
	}

}

func Input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg + ": ")
	scanner.Scan()
	return scanner.Text()
}