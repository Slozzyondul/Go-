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

func Forloop1() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("Sum of numbers from 1 to 1000 is:", sum)
}

//while version in Go

func Forloop2() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("using while to solve similar problem as earlier:", sum)
}
