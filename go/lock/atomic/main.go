package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	//add()
	atomicAdd()

	p := swap()
	fmt.Println("swap:", p.name)
}

func atomicAdd() {
	var i = int32(1)
	go func() {
		addInt32 := atomic.AddInt32(&i, 2)
		//time.Sleep(1)
		swapInt32 := atomic.CompareAndSwapInt32(&i, 3, 4)
		fmt.Println("go1", addInt32)
		fmt.Println("swap:", swapInt32)
	}()

	go func() {
		addInt32 := atomic.AddInt32(&i, 1)
		fmt.Println("go2", addInt32)
	}()

	go func() {
		addInt32 := atomic.AddInt32(&i, 5)
		fmt.Println("go3", addInt32)
	}()

	time.Sleep(time.Second)
}

func add() {
	var i = int32(1)
	go func() {
		addInt32 := i + 2
		fmt.Println(addInt32)
	}()

	go func() {
		addInt32 := i + 1
		fmt.Println(addInt32)
	}()

	go func() {
		addInt32 := i + 5
		fmt.Println(addInt32)
	}()

	time.Sleep(time.Second)
}

type people struct {
	name string
}

func swap() people {
	var i = people{name: "p1"}
	j := &i
	fmt.Println("j", j)
	fmt.Println("i", i)
	i = people{name: "p2"}
	return *j
}
