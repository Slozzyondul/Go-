package goroutines1

import (
	"fmt"
	"time"
)

//Goroutines run in the same address space, so access to shared memory must be synchronized.

func say(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func Goroutine1() {
	go say("solo")
	say("hello")
}
