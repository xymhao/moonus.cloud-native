package main

import (
	"errors"
	"fmt"
)

func main() {
	err := fmt.Errorf("i am error")
	fmt.Println(err)

	var error = errors.New("not found")
	if error == nil {
		fmt.Println(err)
	}

	//recover 类似于try catch
	defer func() {
		fmt.Println("defer func is called45 ")
		if err2 := recover(); err2 != nil {
			fmt.Println(err)
		}
	}()

	panic("a panic is triggered")
}

type StatusError struct {
	ErrStatus int
	Message   string
}

// Error implements the error interface
func (e StatusError) Error() string {
	return e.Message
}
