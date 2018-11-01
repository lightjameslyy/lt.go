package main

import (
	"fmt"
	"time"
)

/*
Go’s select lets you wait on multiple channel operations.
Combining goroutines and channels with select is a powerful feature of Go.
*/

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// 如果没有 default，那么 代码块会被阻塞，指导有一个 case 通过评估；否则一直阻塞
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

}
