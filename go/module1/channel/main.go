package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan int, 10)
	done := make(chan bool)

	defer close(message)

	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-done:
				fmt.Println("consumer done and stop.")
				return
			default:
				fmt.Println(<-message)
			}

		}

	}()

	//producer
	for i := 0; i < 10; i++ {
		message <- i
	}
	close(message)
	time.Sleep(20 * time.Second)

	close(done)
	time.Sleep(time.Second)

	fmt.Println("complete")
}
