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
		fmt.Println("1から" + x + "までの合計は、")
	} else {
		fmt.Println("整数ではありません。")
		return
	}

	t := 0
	for i := 1; i <= n; i ++ {
		t += i
	}
	fmt.Println(t, "です。")
}


func Input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg + ": ")
	scanner.Scan()
	return scanner.Text()
}