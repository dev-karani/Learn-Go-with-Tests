package main

import (
	"fmt"
	"time"
)

func main() {
	//channels
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)

	//buffered channels
	bufferedChan := make(chan string, 2)

	bufferedChan <- "go"
	bufferedChan <- "by examples is fun"

	fmt.Println(<-bufferedChan)
	fmt.Println(<-bufferedChan)

	// channel synchronization
	done := make(chan bool, 1)
	go worker(done)

	<-done

	// channel directions
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

}

func worker(done chan bool) {
	fmt.Println("...working....")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
