package main

import "fmt"

func echoString(content interface{}) {
	result, ok := content.(string)
	if ok {
		fmt.Println(result)
	} else {
		fmt.Println(ok)
	}
}

func EchoRun() {
	//a := "ddd"
	a := 11
	echoString(a)
}
