package main

import (
	"fmt"
	"time"
)

func main() {
	ready := make(chan bool)

	go func() {
		time.Sleep(5 * time.Second)
		ready <- true
	}()

	// block until ready receive true
	select {
	case <-ready:
		fmt.Println("ready")
	}
}
