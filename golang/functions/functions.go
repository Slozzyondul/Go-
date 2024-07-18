package functions

import "fmt"

// addition
func add(x int, y int) int {
	return x + y
}

func Addition() {
	fmt.Println(add(50, 20))
}

// subtraction
func minus(x int, y int) int {
	return x - y
}

func Minus() {
	fmt.Println(minus(50, 20))
}

// swap function returns two strings
func swap(x, y string) (string, string) {
	return y, x
}

func SwapStrings() {
	a, b := swap("Hello", "World")
	fmt.Println(a, b)
}
