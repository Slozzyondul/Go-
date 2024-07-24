package slices1

import "fmt"

func LengthCapacity() {
	s := []int{2, 3, 5, 7, 11, 13}
	PrintSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	PrintSlice(s)

	// Extend its length.
	s = s[:4]
	PrintSlice(s)

	// Drop its first two values.
	s = s[2:]
	PrintSlice(s)
}

func PrintSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
