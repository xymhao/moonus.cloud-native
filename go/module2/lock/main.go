package main

import (
	"fmt"
	"sync"
	"time"
)

//线程安全，CSP的通讯模型

type Status int

const (
	Open  Status = 1
	Close Status = 0
)

type toilet struct {
	sync.Mutex
	status Status
}

func main() {
	go rLock()
	go wLock()
	go lock()
	time.Sleep(time.Second)
}

// read 不互斥，会输出1，2，3
func rLock() {
	lock := sync.RWMutex{}
	for i := 0; i < 3; i++ {
		lock.RLock()
		defer lock.RUnlock()
		fmt.Println("rlock", i)
	}
}
func wLock() {
	lock := sync.RWMutex{}
	for i := 0; i < 3; i++ {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("wLock", i)
	}
}

func lock() {
	lock := sync.Mutex{}
	for i := 0; i < 3; i++ {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("wLock", i)
	}
}
