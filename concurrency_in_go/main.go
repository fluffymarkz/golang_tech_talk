package main

import (
	"fmt"
)

func main() {
	// Create the channel
	messages := make(chan string)

	/* messages <- "ping" */

	// run a go routine to send the message
	go func() { messages <- "ping" }()

	// when you receive the message store it in msg variable
	msg := <-messages

	// print received message
	fmt.Println(msg)
}
