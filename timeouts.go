package main

/*
Reads from various channels for a specified amount of time.
*/
import (
	"fmt"
	"time"
)

func timeout() {
	channel := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			channel <- fmt.Sprintf("random %d", i)
			time.Sleep(10 * time.Millisecond)
		}
	}()

	timeout := time.After(50 * time.Millisecond)
	for {
		select {
		case msg := <-channel:
			fmt.Println(msg)
		case <-timeout:
			fmt.Println("timeout")
			return
		}
	}
}
