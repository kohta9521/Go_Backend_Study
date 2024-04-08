package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)


func main() {
	x := Input("type a script")
	n, err := strconv.Atoi(x)
	if err != nil {
		fmt.Println("ERROR!!")
		return
	}
	P := float64(n)
	fmt.Println(int(P * 1.1))
}

func Input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg + ": ")
	scanner.Scan()
	return scanner.Text()
}