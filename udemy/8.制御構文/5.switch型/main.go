package main

import "fmt"

// switch
// 型スイッチ

func anything(a interface{}) {
	fmt.Println(a)
}

func main() {
	// 型アサーション
	anything("aaa")
	anything(1)

	var x interface{} = 3
	i, isInt := x.(int)
	fmt.Println(i, isInt)
	fmt.Println(i + 2)
	// fmt.Println(x + 2)

	f, isFloat64 := x.(float64)
	fmt.Println(f, isFloat64)

	if x == nil {
		fmt.Println("None")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i, "x is int")
	} else if s, isString := x.(string); isString {
		fmt.Println(s, "x is string")
	} else {
		fmt.Println("I don't know")
	}

	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("I don't know")
	}

	switch v := x.(type) {
	case bool:
		fmt.Println(v, "bool")
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	default:
		fmt.Println("I don't know")
	}
}