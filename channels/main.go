package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println(msg)

	buffer_chan := make(chan string, 2)

	buffer_chan <- "Buffered"
	buffer_chan <- "Channel"

	fmt.Println(<-buffer_chan)
	fmt.Println(<-buffer_chan)

}
