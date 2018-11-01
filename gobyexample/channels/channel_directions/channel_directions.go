package main

import "fmt"

/*
When using channels as function parameters,
you can specify if a channel is meant to only send or receive values.
This specificity increases the type-safety of the program.
channel做参数时可以指定它只能发送（或接受）值，这样可以提高类型安全。
*/

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
