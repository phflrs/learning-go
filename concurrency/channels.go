/*
A channel is a built-in Go type used to send and receive
values between goroutines.
*/
package main

import (
	"fmt"
)

func main() {
    ch := make(chan string) // create a channel of type string

    // Start a goroutine that sends a message
		go func() {
				ch <- "Hello from goroutine "
		}()


    // Main goroutine receives the message
    msg := <-ch
    fmt.Println(msg)
}
