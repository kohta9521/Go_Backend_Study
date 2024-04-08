package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	t := 0
	x := Input("type a number")
	n, err := strconv.Atoi(x)
	if err != nil {
		goto err
	}
	for i := 1; i <= n; i ++ {
		t += i
	}
	fmt.Println("total: ", t)
	return

err:
	fmt.Println("ERRER")
}



func Input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg + ": ")
	scanner.Scan()
	return scanner.Text()
}