package goroutines1

import (
	"fmt"
	"time"
)

func DefaultSelection1() {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop() // Ensure the ticker is stopped to prevent resource leaks
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
