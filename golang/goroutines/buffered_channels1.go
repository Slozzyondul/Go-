package goroutines1

import (
	"fmt"
	"time"
)

func GoroutineWithoutDeadlock() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	// Goroutine to receive values from the channel
	go func() {
		fmt.Println(<-ch)
		fmt.Println(<-ch)
	}()

	// Overfilling the buffer
	ch <- 3 // This will now succeed because there is a goroutine to receive the value

	fmt.Println(<-ch)

	// Give the goroutine some time to finish
	time.Sleep(time.Second)
}
