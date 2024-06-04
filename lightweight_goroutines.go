package main

import "fmt"

/*
There is a daisy-chain/chinese whisperer chain to show that goroutines are lightweight.
*/

func connect(toChannel, fromChannel chan int) {
	toChannel <- 1 + <-fromChannel
}

func runDaisyChain() {
	const chainLength = 100000 //hundred thousand : not working for 100x. Slow for 10x

	rightMost := make(chan int)
	right := rightMost
	left := make(chan int)
	right = rightMost

	for i := 0; i < chainLength; i++ {
		left = make(chan int)
		go connect(left, right)
		right = left
	}

	go func(ch chan int) {
		ch <- 1
	}(rightMost)

	fmt.Println(<-left)
}
