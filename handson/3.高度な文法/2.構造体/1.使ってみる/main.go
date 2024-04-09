package main

import "fmt"

// 構造体
var mydata struct {
	Name string
	Data []int
}

func main() {
	mydata.Name = "Taro"
	mydata.Data = []int{10, 20, 30}
	fmt.Println(mydata)
}