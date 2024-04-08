package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)



func main() {
	x := Input("type a number")
	fmt.Print(x + "月は、")
	switch n, err := strconv.Atoi(x); n {
		case 0:
			fmt.Println("整数値が得られません。")
			fmt.Println(err)
		case 1, 2, 12:
			fmt.Println("冬です。")
		case 3, 4, 5:
			fmt.Println("春です。")
		case 6, 7, 8:
			fmt.Println("夏です。")
		case 9, 10, 11:
			fmt.Println("秋です。")
		default:
			fmt.Println("月の値ではありません。")
	}
}

func Input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg + ": ")
	scanner.Scan()
	return scanner.Text()
}