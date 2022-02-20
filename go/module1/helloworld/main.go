package main

import (
	"fmt"
	"time"
)

func main() {
	go println("a")
	go println("b")
	go println("c")

	ch := make(chan int, 6)

	go func() {
		fmt.Println("hello from goroutine")
		ch <- 0
		ch <- 1
		ch <- 2
		ch <- 3
		ch <- 4
		ch <- 5
		close(ch)
	}()

	for i := range ch {
		fmt.Printf("receiving:%v \n", i)

	}

	go producer(c)

	go consumer(c)

	selectChannel()
	time.Sleep(time.Second * 10)

}

//观察者模式

var c = make(chan int)

//定义生产者为只发送通道
func producer(ch chan<- int) {
	ch <- 1
}

// 消费者只接收通道
func consumer(ch <-chan int) {
	fmt.Println("handle:", <-ch)
}

func selectChannel() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	//为协程定义超时时间
	timer := time.NewTimer(time.Second * 5)

	go func() {
		ch1 <- 1
	}()
	go func() {
		ch2 <- 2
	}()
	select {
	case v := <-ch1:
		fmt.Printf("1:%v\n", v)
	case v := <-ch2:
		fmt.Printf("2: %v\n", v)
	case <-timer.C:
		fmt.Printf("timeout")
	}

	time.Sleep(time.Second * 5)
}
