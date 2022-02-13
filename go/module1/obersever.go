package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type IProducer interface {
	Send(value int)
}

type Producer struct {
	// only send
	ch    chan<- int
	topic string
}

// Send 生产者
func (prc *Producer) Send(number int) {
	prc.ch <- number
}

type IConsumer interface {
	handle()
	Start()
}

type Consumer struct {
	//only receive
	ch    <-chan int
	close <-chan bool
	name  string
}

func (consumer *Consumer) handle() {
	go func(ch <-chan int) {
		ticker := time.NewTicker(time.Second)
		for _ = range ticker.C {
			message := <-consumer.ch
			fmt.Println("receive：", message)
		}

	}(consumer.ch)
}

func (consumer *Consumer) Start() {
	consumer.handle()
}

func main() {
	ch := make(chan int, 10)
	defer close(ch)
	producer := Producer{ch: ch, topic: "cn message"}
	consumer := Consumer{ch: ch, name: "cn group"}

	go func(ch chan<- int) {
		ticker := time.NewTicker(time.Second)
		for _ = range ticker.C {
			rand.Seed(time.Now().UnixNano())
			message := rand.Intn(999)
			producer.Send(message)
			fmt.Println("message:", message)
		}
	}(ch)

	consumer.Start()
	input := bufio.NewReader(os.Stdin)
	for true {
		text, _ := input.ReadString('\n')
		if strings.Contains(text, "exit") {
			return
		}
	}
}
