package main

import "fmt"

type person interface {
	say() string
}

type student struct {
}

func (st student) say() string {
	return "i am student"
}

type teacher struct {
}

func (t teacher) say() string {
	return "i am teacher"
}

func main() {
	st := student{}
	t := teacher{}
	all := []person{st, t}

	for i := range all {
		fmt.Println(all[i].say())
	}
}
