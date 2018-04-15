package main

import (
	"fmt"
	"runtime"
)

func sendMessages(ch chan<- string) {
	for i := 0; i < 20; i++ {
		ch <- fmt.Sprintf("data %d", i)
	}
	close(ch)
}

func printMessages(ch <-chan string) {
	for message := range ch {
		fmt.Println(message)
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	messages := make(chan string)
	go sendMessages(messages)
	printMessages(messages)
}
