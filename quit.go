package main

import (
	"fmt"
	"time"
)

/*
A quit functionality can be implemented using a channel from goroutine to the main thread. We can adds acknowledgement to
the goroutine also by waiting on the read on this channel.
*/

func quit() {
	quitChannel := make(chan bool)
	channel := make(chan string)

	go func() {
		channel <- "message"
		quitChannel <- true
		fmt.Printf("got ack for quitChannel %d", <-quitChannel)
	}()

	for {
		select {
		case msg := <-channel:
			fmt.Println(msg)
		case <-quitChannel:
			fmt.Println("received quit by the channel")
			time.Sleep(1 * time.Second)
			quitChannel <- false
			return
		}
	}
}
