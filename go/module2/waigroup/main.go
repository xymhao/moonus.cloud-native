package main

import (
	"fmt"
	"sync"
)

func main() {
	waitByChannel()

	waitByWaitGroup()
}

func waitByWaitGroup() {
	waitGroup := sync.WaitGroup{}
	//指定次数100次
	waitGroup.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			//标记完成一次
			waitGroup.Done()
		}(i)
	}

	waitGroup.Wait()
}

//通过channel的方式，存在额外的开销
func waitByChannel() {
	c := make(chan bool, 100)

	//将每次go routines 放入channel，再遍历channel
	for i := 0; i < 100; i++ {
		go func(j int) {
			fmt.Println(j)
			c <- true
		}(i)
	}

	for i := 0; i < 100; i++ {
		<-c
	}

}
