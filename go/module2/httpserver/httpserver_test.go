package main

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"io"
	"net/http"
	"testing"
	"time"
)

func TestHealth(t *testing.T) {
	go main()
	time.Sleep(time.Second * 2)
	resp, _ := http.Get("http://localhost:808/healthz")
	read, err := io.ReadAll(resp.Body)
	fmt.Println(read, err)
	fmt.Println("response content :", string(read))
	assert.Equal(t, "200", string(read))
}
