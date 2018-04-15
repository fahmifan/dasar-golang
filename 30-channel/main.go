package main

import (
	"fmt"
	"runtime"
)

func printMsg(what chan string) {
	fmt.Println(<-what)
}

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	for _, each := range []string{"wicj", "hung", "born"} {
		go func(who string) {
			var data = fmt.Sprintf("hello %s", who)
			messages <- data
		}(each)
	}

	// printMsg(messages)

	for i := 0; i < 3; i++ {
		printMsg(messages)
	}
}
