package main

import (
	"fmt"
	"time"
)

func main() {
	doneChan := make(chan int)
	nTasks := 10

	go func() {
		for i := 0; i < nTasks; i++ {
			fmt.Printf("task %d running...\n", i)
			time.Sleep(200 * time.Millisecond)
			doneChan <- i
		}
	}()

	for i := 0; i < nTasks; i++ {
		fmt.Printf("task %d done\n", <-doneChan)
	}
}
