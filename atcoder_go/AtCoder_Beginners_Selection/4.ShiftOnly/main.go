package main

import (
	"fmt"
	"math"
)


func main() {
	var N int
	fmt.Scan(&N)

	numbers := make([]int, N)
	for i := 0; i < N; i ++ {
		fmt.Scan(&numbers[i])
	}

	minOperations := math.MaxInt32

	for _, number := range numbers {
		operations := 0
		for number % 2 == 0 {
			number /= 2
			operations++
		}
		if operations < minOperations {
			minOperations = operations
		}
	}

	fmt.Println(minOperations)
}