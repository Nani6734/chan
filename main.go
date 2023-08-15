package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a channel of integers
	ch := make(chan int)

	// Goroutine: Send data into the channel
	go func() {
		fmt.Println("Goroutine: Sending data...")
		ch <- 42 // Send value 42 into the channel
		fmt.Println("Goroutine: Data sent")

	}()

	// Pause to allow the goroutine to run
	time.Sleep(1 * time.Second)

	// Receive data from the channel
	data := <-ch
	fmt.Println("Main: Received data:", data)

}
