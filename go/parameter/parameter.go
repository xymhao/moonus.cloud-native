package main

import (
	"flag"
	"fmt"
	"os"
)

var myVariable = 0

func main() {
	//参数解析demo
	//获取程序执行的入参
	fmt.Println("os args", os.Args)

	//解析 --name
	name := flag.String("name", "world", "specify")
	flag.Parse()
	fmt.Println(*name)
}

// 初始化函数
func init() {

}
