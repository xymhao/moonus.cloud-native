package main

import (
	"fmt"
	_ "moonus.cloud-native/go/init/a"
)
import _ "moonus.cloud-native/go/init/b"

func init() {
	fmt.Println("main init")
}

func main() {
	fmt.Println("main function")
	a, b := getResult(1, 2)
	fmt.Printf("result:%v,%v", a, b)

	getValue(1, 2, 3, 4)
}

func getResult(x, y int) (a, b int) {
	a = x
	b = y
	return b, a
}

//可变长度参数
func getValue(a int, b ...int) {
	slices := []int{1, 2, 3}

	slices = append(
		slices,
		a,
		2,
		3,
		4,
	)

	fmt.Printf("append after : %v", slices)

}
