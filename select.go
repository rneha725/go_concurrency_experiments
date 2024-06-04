package main

import (
	"fmt"
	"time"
)

/*
fanIn using select statement.
there is an infinite select statement, to timeout, if any message takes more than 1ms to deliver, the loop breaks
*/

func selectFanIn() {
	channel := fanIn()
	channel1 := generator("channel1")
	channel2 := generator("channel2")
	for {
		select {
		case msg := <-channel1:
			fmt.Println(msg)
		case msg := <-channel2:
			fmt.Println(msg)
		case msg := <-channel:
			fmt.Println(msg)
		case <-time.After(1 * time.Millisecond):
			fmt.Println("Program ends")
			return
		}
	}
}
