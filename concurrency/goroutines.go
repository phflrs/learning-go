package main

import (
	"fmt"
	"time"
)

func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}

func main() {

    f("direct")

		/* runs asynchronously, doesn't stop the execution */ 
    go f("goroutine")

		/* runs asynchronously, doesn't stop the execution */
    go func(msg string) {
        fmt.Println(msg)
    }("going") // calling anonymous func with parameter "going"

    time.Sleep(time.Second)
    fmt.Println("done")
}