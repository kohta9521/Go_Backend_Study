package main

import (
	"fmt"
	"strings"
)


func main() {
	var s string
	fmt.Scan(&s)

	digits := strings.Split(s, "")

	count := 0
	for _, digit := range digits {
		if digit == "1" {
			count ++
		}
	}

	fmt.Println(count)
}