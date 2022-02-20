package main

import (
	"fmt"
	"sync"
)

func main() {
	//defer 入栈
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")

	//result ： 3 2 1
	fmt.Println("loopFunc")
	loopFunc()

	//time.Sleep(time.Second * 5)
}

func loopErrorFunc() {
	lock := sync.Mutex{}
	for i := 0; i < 3; i++ {
		lock.Lock()
		//fatal error: all goroutines are asleep - deadlock!
		//函数退出才执行
		defer lock.Unlock()
	}
}

func loopFunc() {
	group := sync.WaitGroup{}
	group.Add(3)

	lock := sync.Mutex{}
	for i := 0; i < 3; i++ {
		//通过闭包
		go func(i int) {
			lock.Lock()
			fmt.Printf("loop:%d\n", i)
			lock.Unlock()
			group.Done()
		}(i)
	}

	defer group.Wait()

}
