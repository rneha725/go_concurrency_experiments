package main

/*
Fan-in:
channel1----▼
			|-----> channel -->
channel2----▲

Why do we need fan-in?
Following simple reading from two channel will create two blocking calls for the receiver function, it will look
something like this:

f() {
	forLoop()
		print(from channel1) //blocking
		print(from channel2) //blocking
}

We might not want to wait for the receiver on these channels, and want to process as soon as any of these channels get a
value. This can be done using a third channel. What we can do is, we can launch two goroutines, both of these will write
to this third channel. In the receiver, we will read from this single channel, as soon as any of the goroutines communicates.
This has been implemented in this module.
*/

import (
	"fmt"
	"time"
)

func runFanIn() {
	channel := fanIn()
	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

func generator(message string) <-chan string {
	channel := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			channel <- fmt.Sprintf("%s: %s", "message: ", message)
		}

	}()

	return channel
}

func fanIn() <-chan string {
	channel := make(chan string)
	channel1 := generator("first")
	channel2 := generator("second")

	go func() {
		for i := 0; i < 5; i++ {
			channel <- <-channel1
			time.Sleep(1 * time.Millisecond)
		}
	}()
	go func() {
		for i := 0; i < 5; i++ {
			channel <- <-channel2
		}
	}()
	return channel
}
