package main

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"testing"
	"time"
)

func TestHandleMessageConsumerCompleted(t *testing.T) {
	messageChan := make(chan string, 10)
	h := Handle{index: 0, message: messageChan, flag: false, done: make(chan bool)}
	count := 5
	for i := 0; i < count; i++ {
		messageChan <- fmt.Sprint(i)
	}
	go func() {
		time.Sleep(time.Second * 3)
		h.stop()

	}()
	h.HandleMessage()
	close(h.done)

	time.Sleep(time.Second * 6)
	//consumer complete
	assert.Equal(t, 0, len(messageChan))
}

func TestChannelDistribution(t *testing.T) {
	channels := initChannels(3)

	producer := Producer{Channels: channels, topic: "test"}
	producer.Producer(3, "3-demo")
	producer.Producer(3, "3-demo")
	producer.Producer(3, "3-demo")

	producer.Producer(2, "2-demo")
	producer.Producer(2, "2-demo")

	producer.Producer(1, "1-demo")

	assert.Equal(t, 3, len(channels[0]))
	assert.Equal(t, 1, len(channels[1]))
	assert.Equal(t, 2, len(channels[2]))
}
