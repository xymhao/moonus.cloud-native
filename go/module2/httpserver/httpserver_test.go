package main

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"net/http"
	"testing"
	"time"
)

func TestHealth(t *testing.T) {
	go main()
	time.Sleep(time.Second * 2)
	resp, _ := http.Get("http://localhost:808/healthz")
	bytes := make([]byte, resp.ContentLength)
	read, err := resp.Body.Read(bytes)
	fmt.Println(read, err)
	fmt.Println("response content :", string(bytes))
	assert.Equal(t, "200", string(bytes))
}
