package main

import (
	"fmt"
)

/*
Channels are the pipes that connect concurrent goroutines.
You can send values into channels from one goroutine and
receive those values into another goroutine.
*/

func main() {

	message := make(chan string)

	go func() { message <- "ping" }()

	msg := <-message
	fmt.Println(msg)

}

/*
By default sends and receives block until both the sender and receiver are ready.
channel的发送和接受默认是阻塞的。
*/
