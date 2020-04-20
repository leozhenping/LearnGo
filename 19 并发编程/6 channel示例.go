package main

import (
	"fmt"
	"time"
)

var channel = make(chan int)

func printer(s string) {

	for _, str := range s {
		fmt.Printf("%c", str)
		time.Sleep(100 * time.Millisecond)
	}

}

func person1() {
	printer("hello")
	channel <- 8
}

func person2() {
	<- channel
	printer("world")

}

func main() {
	go person1()
	go person2()

	for {
		;
	}
}
