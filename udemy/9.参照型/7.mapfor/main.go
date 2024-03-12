package main

// map
// for

func main() {
	m := map[string]int{
		"Apple": 100,
		"Banana": 200,
	}
	for k, v := range m {
		println(k, v)
	}
}