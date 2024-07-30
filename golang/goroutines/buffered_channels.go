package goroutines1

// Provide the buffer length as the second argument to make to initialize a buffered channel
//ch := make(chan int, 100)
//Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.

import "fmt"

func GoroutineDeadlock() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// Overfilling the buffer
	ch <- 3 // This will cause a deadlock

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
