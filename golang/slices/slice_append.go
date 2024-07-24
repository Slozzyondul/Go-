package slices1

import "fmt"

func AppendSlice() {
	var s []int
	PrintSliceAppend(s)

	// append works on nil slices.
	s = append(s, 0)
	PrintSliceAppend(s)

	// The slice grows as needed.
	s = append(s, 1)
	PrintSliceAppend(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	PrintSliceAppend(s)
}

func PrintSliceAppend(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
