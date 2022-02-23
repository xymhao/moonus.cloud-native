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
	mux.HandleFunc("/healthz", health)
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

type MonitorResponseWriter struct {
	rw         http.ResponseWriter
	statusCode int
}

func (w *MonitorResponseWriter) Header() http.Header {
	return w.rw.Header()
}

func (w *MonitorResponseWriter) Write(b []byte) (int, error) {
	return w.rw.Write(b)
}

func (w *MonitorResponseWriter) WriteHeader(statusCode int) {
	w.rw.WriteHeader(statusCode)
}

func rootHandle(rw http.ResponseWriter, request *http.Request) {
	writer := MonitorResponseWriter{rw: rw, statusCode: 200}
	for k, v := range request.Header {
		writer.Header().Set(k, v[0])
	}
	version := os.Getenv("VERSION")
	writer.Header().Set("Version", version)

	for k, v := range request.Header {
		writer.Write([]byte(fmt.Sprintf("%v=%v\n", k, v)))
	}
	for _, i := range writer.Header() {
		fmt.Println(i)
	}
	fmt.Println("IP:", request.RemoteAddr)
	fmt.Println("status code:", writer.statusCode)
}

func health(w http.ResponseWriter, request *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("200"))
}
