package main

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"sync/atomic"
	"testing"
)

func TestHandleMessageCount(t *testing.T) {
	messageChan := make(chan string, 10)
	cancelChan := make(chan bool)
	record := int32(0)
	h := Handle{index: 0, message: messageChan, test: func() {
		atomic.AddInt32(&record, 1)
	}, cancel: cancelChan}
	count := 5
	for i := 0; i < count; i++ {
		messageChan <- fmt.Sprint(i)
	}
	close(messageChan)
	h.HandleMessage()
	assert.Equal(t, record, int32(count))
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
