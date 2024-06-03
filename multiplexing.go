package main

import (
	"fmt"
)

func runMultiplexing() {
	channel := multiplexing()
	for i := 0; i < 2; i++ {
		fmt.Println(<-channel)
	}
}

func createsChannelAndGoroutine(message string) <-chan string {
	channel := make(chan string)

	go func() {
		channel <- fmt.Sprintf("%s: %s", "message: ", message)
	}()

	return channel
}

func multiplexing() <-chan string {
	channel := make(chan string)
	channel1 := createsChannelAndGoroutine("first")
	channel2 := createsChannelAndGoroutine("second")
	go func() {
		channel <- <-channel1
		channel <- <-channel2
	}()
	return channel
}
