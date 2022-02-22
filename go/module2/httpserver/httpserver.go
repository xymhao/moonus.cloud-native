package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"log"
	"net/http"
	"net/http/pprof"
	"os"
	"strings"
)

func main() {
	//init port
	var port = flag.String("port", "808", "setting http server default port")
	flag.Parse()

	glog.V(2).Infof("start http server")
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandle)
	mux.HandleFunc("/health", health)
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	//start server
	go func() {
		err := http.ListenAndServe(":"+*port, mux)
		if err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Println("start up http server port:", *port)
	fmt.Println("enter ‘exit’ or 'ctrl+c' to stop server")
	input := bufio.NewReader(os.Stdin)
	for true {
		text, _ := input.ReadString('\n')
		if strings.Contains(text, "exit") {
			fmt.Println("stop server")
			return
		}
	}
}

func rootHandle(writer http.ResponseWriter, request *http.Request) {
	for k, v := range request.Header {
		writer.Header().Set(k, v[0])
	}
	version := os.Getenv("VERSION")
	writer.Header().Set("Version", version)

	for k, v := range request.Header {
		writer.Write([]byte(fmt.Sprintf("%v=%v\n", k, v)))
	}
	statusCode := 200
	writer.WriteHeader(statusCode)
	fmt.Println("IP:", request.RequestURI)
	fmt.Println("status code:", statusCode)
}

func health(w http.ResponseWriter, request *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("200"))
}
