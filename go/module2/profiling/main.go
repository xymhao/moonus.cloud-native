package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
)

var cpuProfile = flag.String("cpuprofile", "cpuProfile", "write cpu profile")

func main() {
	flag.Parse()
	f, err := os.Create(*cpuProfile)
	if err != nil {
		log.Fatal(err)
	}

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	var result int
	for i := 0; i < 100000000; i++ {
		result += i
	}
	log.Println("result:", result)

	//查看信息
	// go tool pprof cpuProfile

	//Type: cpu
	//Time: Feb 21, 2022 at 10:39pm (CST)
	//Duration: 214.93ms, Total samples = 30ms (13.96%)
	//Entering interactive mode (type "help" for commands, "o" for options)

}
