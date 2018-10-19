package main

import "fmt"

/*
By default channels are unbuffered.
Buffered channels accept a limited number of values
without a corresponding receiver for those values.
*/

func main() {

	message := make(chan string, 2)

	message <- "buffered"
	message <- "channel"

	fmt.Println(<-message)
	fmt.Println(<-message)
}
