package main

import (
	"fmt"
	"time"
)

func doSomething(channel chan string) {
	for i := 0; i < 10; i++ {
		fmt.Println("doSomething: ", i)
		time.Sleep(1 * time.Second)
	}
	channel <- "doSomething is done"
}

func main() {
	channel := make(chan string)
	go doSomething(channel)
	go doSomething(channel)
	// time.Sleep(8 * time.Second)
	response := <-channel
	fmt.Println(response)
	close(channel)
}
