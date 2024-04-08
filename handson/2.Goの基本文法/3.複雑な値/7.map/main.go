package main

import "fmt"

func main() {
	m := map[string]int {
		"a": 100,
		"b": 200,
		"C": 300,
	}
	m["total"] = m["a"] + m["b"] + m["C"]
	fmt.Println(m)
}