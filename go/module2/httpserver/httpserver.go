package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

type HandlerFunc func(ResponseWriter, *http.Request)

//custom route
var route = make(map[string]HandlerFunc)

func init() {
	route["/healthz"] = health
}

func main() {
	//init port
	var port = flag.String("port", "808", "setting http server default port")
	flag.Parse()

	glog.V(2).Infof("start http server")
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandle)
	//mux.HandleFunc("/healthz", health)
	//mux.HandleFunc("/debug/pprof/", pprof.Index)
	//mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	//mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	//mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

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

// ResponseWriter Decorator patterns
// Implementing http statusCode record
type ResponseWriter struct {
	rw         http.ResponseWriter
	statusCode int
}

func (w *ResponseWriter) Header() http.Header {
	return w.rw.Header()
}

func (w *ResponseWriter) Write(b []byte) (int, error) {
	return w.rw.Write(b)
}

func (w *ResponseWriter) WriteHeader(statusCode int) {
	w.rw.WriteHeader(statusCode)
}

func rootHandle(rw http.ResponseWriter, request *http.Request) {
	writer := ResponseWriter{rw: rw, statusCode: 200}
	for k, v := range request.Header {
		writer.Header().Set(k, v[0])
	}
	version := os.Getenv("VERSION")
	writer.Header().Set("Version", version)

	//self route
	handlerFunc := route[request.RequestURI]
	if handlerFunc != nil {
		handlerFunc(writer, request)
	}

	body, err := io.ReadAll(request.Body)
	if err == nil && body != nil {
		fmt.Println("body:", string(body))
	}
	var ip = GetCurrentIp(request)
	fmt.Println("IP:", ip)
	fmt.Println("status code:", writer.statusCode)
}

func GetCurrentIp(r *http.Request) interface{} {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

func health(w ResponseWriter, _ *http.Request) {
	w.WriteHeader(200)
	_, err := w.Write([]byte("200"))
	if err != nil {
		fmt.Println(err)
	}
}
