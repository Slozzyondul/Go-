package goroutines1

//Channels are a typed conduit through which you can send and receive values with the channel operator, <-.

//Like maps and slices, channels must be created before use:
//By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func ChannelsWithgoroutines() {
	s := []int{7, 2, 8, -9, 4, 0} // Define a slice of integers

	c := make(chan int) // Create a channel of type int

	// Launch two goroutines to compute the sum of the first and second halves of the slice
	go sum(s[:len(s)/2], c) // Sum the second half of the slice
	go sum(s[len(s)/2:], c) // Sum the first half of the slice
	x, y := <-c, <-c        // Receive the results from the channel

	fmt.Println(x, y, x+y)
}
