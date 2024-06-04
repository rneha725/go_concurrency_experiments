package main

import "fmt"

/*
fanIn using select statement
*/

func selectFanIn() {
	channel1 := generator("first")
	channel2 := generator("second")
	for {
		select {
		case msg := <-channel1:
			fmt.Println(msg)
		case msg := <-channel2:
			fmt.Println(msg)
		}
	}
}
