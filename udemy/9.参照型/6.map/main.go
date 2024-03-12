package main

import "fmt"

// map


func main() {
	var m = map[string]int{"A": 100, "B": 200}
	fmt.Println(m)

	m2 := map[string]int{"A": 100, "B": 200}
	fmt.Println(m2)

	m3 := map[int]string{
		1: "A",
		2: "B",
		3: "C",
	}
	fmt.Println(m3)

	m4 := make(map[int]string)
	fmt.Println(m4)
	m4[1] = "Japan"
	m4[2] = "USA"
	m4[3] = "China"
	fmt.Println(m4)

	fmt.Println(m["A"])
	fmt.Println(m4[2])
	fmt.Println(m4[3])
	s, ok := m4[1]
	if !ok {
		fmt.Println("error")
	}
	fmt.Println(s, ok)

	// delete
	delete(m4, 2)
	fmt.Println(m4)

	// len
	fmt.Println(len(m4))
}