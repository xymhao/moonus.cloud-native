package main

import (
	"flag"
	"github.com/golang/glog"
	"log"
	"net/http"
)

//将标准输出转至文件
//日志分级
//自带时间戳，方便调试
func main() {
	flag.Set("v", "4")
	glog.V(2).Infof("start http server...")
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandle)
	err := http.ListenAndServe(":8890", mux)

	if err != nil {
		log.Fatal(err)
	}
}

func rootHandle(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("123"))
}
