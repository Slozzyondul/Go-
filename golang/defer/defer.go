package defer1

//A defer statement defers the execution of a function until the surrounding function returns.

import "fmt"

func Defer() {
	defer fmt.Println("solo, hold on for the next command to run before running")

	fmt.Println("hello")
}
