package main

import "fmt"

// Mydata is structure.
type Mydata struct {
	Name string
	Data []int
}

func (md Mydata) PrintData() {
	fmt.Println("*** Mydata ***")
	fmt.Println("Name: ", md.Name)
	fmt.Println("Data:  ", md.Data)
	fmt.Println("****************")
}

func main() {
	taro := Mydata{
		"Hanako", []int{98, 78, 57, 88},
	}
	taro.PrintData()
}