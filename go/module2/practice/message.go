package main

import (
	"fmt"
	"math/rand"
	"time"
)

var capacity = 10

type Producer struct {
	topic string
	//readonly
	Channels []chan string
}

type Consumer struct {
	topic   string
	name    string
	handles []Handle
}

type Handle struct {
	//number of consumer instance
	index int
	//message only read
	message <-chan string
}

// Start init consumer and star handle goroutine
// 开启和channel数量相等的协程处理channel消息
func (consumer *Consumer) Start(Channels []chan string) {
	consumer.handles = make([]Handle, len(Channels))
	for i, channel := range Channels {
		consumer.handles[i].message = channel
		consumer.handles[i].index = i
		go consumer.handles[i].HandleMessage()
	}
}

// HandleMessage handle message
func (h *Handle) HandleMessage() {
	ticker := time.NewTicker(time.Millisecond * 5)
	for range ticker.C {
		msg, ok := <-h.message
		if !ok {
			return
		}
		fmt.Printf("Handle-%v, message: %v \n", h.index, msg)
	}
}

// Producer send message in target channel
func (p *Producer) Producer(tag int, key string) {
	index := tag % len(p.Channels)
	p.Channels[index] <- key
}

//console print
//---
//produce-3, value:455
//produce-1, value:382
//Handle-1, message: 382
//Handle-0, message: 455
//produce-2, value:742
//Handle-2, message: 742
//produce-1, value:630
//Handle-1, message: 630
//produce-2, value:357
//produce-3, value:549
//Handle-0, message: 549
//Handle-2, message: 357
//-----
//produce-3 => handle-0
//produce-1 => handle-1
//produce-2 => handle-2
func main() {
	channels := initChannels(3)
	topic := "cn-study"
	producer := Producer{Channels: channels, topic: topic}
	consumer := Consumer{name: "moonus", topic: topic}
	//start consumer instance
	consumer.Start(channels)

	//start three thread produce message
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		go send(producer, 1, time.Second)
		go send(producer, 2, time.Second*2)
		go send(producer, 3, time.Second*3)
	}
}

func send(producer Producer, tenantId int, duration time.Duration) {
	value := rand.Intn(999)
	fmt.Printf("produce-%v, value:%v \n", tenantId, value)
	producer.Producer(tenantId, fmt.Sprint(value))

	time.Sleep(duration)
}

func initChannels(count int) []chan string {
	var channels = make([]chan string, count)
	for i := 0; i < count; i++ {
		channels[i] = make(chan string, capacity)
	}
	return channels
}
