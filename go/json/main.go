package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var obj interface{}

	err := json.Unmarshal([]byte("{\"err_no\":0,\"errmsg\":\"\",\"queryid\":\"0x9dc88613c99e5f\"}"), &obj)
	if err != nil {
		fmt.Println(err)
	}

	objMap, ok := obj.(map[string]interface{})
	if ok == true {
		for k, v := range objMap {
			fmt.Printf("key:%v value:%v\n", k, v)
		}
	}
}
