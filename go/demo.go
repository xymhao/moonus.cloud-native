package main

import (
	"fmt"
	"reflect"
)

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

// TransferKey replace key 'stupid' to smart
// replace 'weak' to strong
func TransferKey(words []string) []string {
	for i := range words {
		if words[i] == "stupid" {
			words[i] = "smart"
		}

		if words[i] == "weak" {
			words[i] = "strong"
		}
	}
	return words
}

func SliceDemo() {
	myArray := [5]int{1, 2, 3, 4, 5}
	mySlice := myArray[1:3]
	fmt.Printf("%v", mySlice)

	fullSlice := myArray[:]
	fmt.Printf("%v", fullSlice)
}

// DeleteItem delete index item for slice
func DeleteItem(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

type MyType struct {
	Name string `json:"name,omitempty" protobuf:"bytes,10,opt,name"`
	Age  int    `json:"age"`
}

// ReflectDemo demo of go reflect, to get tag from struct field
func ReflectDemo() {
	mt := MyType{Name: "xym"}
	mtType := reflect.TypeOf(mt)
	name := mtType.Field(0)
	tag := name.Tag.Get("protobuf")
	println(tag)
}

// ServiceType 类型重命名
// 由于go 语言没有枚举，通过重命名可以实现枚举
type ServiceType string

// ServiceTypeClusterIp
const (
	ServiceTypeClusterIp ServiceType = "ClusterIp"
	ServiceTypeNodePort  ServiceType = "NodePort"
)

func DemoServiceType() {
	printServiceType(ServiceTypeNodePort)
	printServiceType(ServiceTypeClusterIp)
}

func printServiceType(serviceType ServiceType) {
	fmt.Printf("%v\n", serviceType)
}

func conditionDemo(serviceType ServiceType) {
	switch serviceType {
	case ServiceTypeClusterIp:
		fmt.Printf("i am ip")
	case ServiceTypeNodePort:
		fmt.Printf("i am port")
	}
}
