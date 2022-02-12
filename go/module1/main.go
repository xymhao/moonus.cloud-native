package main

import "fmt"

//go 语言特性学习

func main() {
	b := []int{1, 2, 3}
	fmt.Printf("%v", b)
	des := []string{"I", "am", "stupid", "and", "weak"}
	des = TransferKey(des)
	fmt.Println(des) //[I am smart and strong]
	println("SliceDemo")
	SliceDemo()
	println("ReflectDemo")
	ReflectDemo()
	println("DemoServiceType")
	DemoServiceType()
}
