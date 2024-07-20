package forloop

import (
	"fmt"
)

func Forloop() {
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("Sum of numbers from 0 to 9 is:", sum)
}
