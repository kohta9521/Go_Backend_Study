package main

import "fmt"

// interface
// custom error

// type error interface {
// 	Error() string
// }

type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{Message: "カスタムエラーが発生しました。", ErrCode: 1234}
}

func main() {
	err := RaiseError()
	fmt.Println(err.Error())
}
