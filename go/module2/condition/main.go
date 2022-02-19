package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Queue struct {
	queue []string
	cond  *sync.Cond
}

func main() {
	queue := Queue{queue: []string{}, cond: sync.NewCond(&sync.Mutex{})}
	go consumer(&queue)

	for true {
		randomNumber := rand.Intn(999)
		queue.EnQueue(fmt.Sprintln(randomNumber))
		time.Sleep(time.Second)
	}
}

func consumer(q *Queue) {
	for true {
		result := q.DeQueue()
		fmt.Println("dequeue", result)
		time.Sleep(time.Second * 1)
	}
}

func (q *Queue) EnQueue(item string) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	q.queue = append(q.queue, item)
	fmt.Println("EnQueue", item)

	q.cond.Broadcast()
}

func (q *Queue) DeQueue() string {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	if len(q.queue) == 0 {
		fmt.Println("queue count empty")
		q.cond.Wait()
	}

	result := q.queue[0]
	q.queue = q.queue[1:]
	return result
}
