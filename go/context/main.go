package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	baseCtx := context.Background()
	ctx := context.WithValue(baseCtx, "a", "b")
	go func(c context.Context) {
		value := c.Value("a")
		fmt.Println(value)
	}(ctx)

	//设置超时时间
	ctx2, cancelFunc := context.WithTimeout(baseCtx, 3*time.Second)
	defer cancelFunc()
	go func(c context.Context) {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-c.Done():
				fmt.Println("interrupt... context timeout")
				return
			default:
				fmt.Println("enter default")
			}
		}
	}(ctx2)

	select {
	case <-ctx2.Done():
		fmt.Println("main done")
	}
	time.Sleep(time.Second * 5)
}
