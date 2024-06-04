package main

/*
This is how sequencing is implemented, we are trying to communicate using two types of channels. One is where goroutines
are writing and the main thread is reading. The other one is used by goroutines to wait for main thread's writes.

channel1 -> routine 1 writes to it
channel2 -> routine 2 writes to it
waitChannel -> main thread writes to it.

The logic is that a goroutine should wait till main thread's actions for sending a value in its channel.
*/

import "fmt"

type Message struct {
	msg         string
	waitChannel chan bool
}

func runSequencing() {
	waitChannel := make(chan bool)
	channel1 := generator1("first", waitChannel)
	channel2 := generator1("second", waitChannel)

	for i := 0; i < 5; i++ {
		msg1 := <-channel1
		fmt.Println(msg1.msg)
		msg2 := <-channel2
		fmt.Println(msg2.msg)
		msg1.waitChannel <- true
		msg2.waitChannel <- true
	}

}

func generator1(msg string, waitChannel chan bool) <-chan Message {
	channel := make(chan Message)
	go func() {
		for i := 0; i < 5; i++ {
			channel <- Message{msg, waitChannel}
			<-waitChannel
		}
	}()

	return channel
}
