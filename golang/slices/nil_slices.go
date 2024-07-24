package slices1

import "fmt"

func NilSlices() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if len(s) == 0 {
		fmt.Println("nil!")
	}
}
