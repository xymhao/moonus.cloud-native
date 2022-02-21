package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", health)
	err := http.ListenAndServe(":8980", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func health(w http.ResponseWriter, request *http.Request) {
	w.Write([]byte("ok"))
}
