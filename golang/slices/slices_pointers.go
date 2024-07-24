package slices1

import "fmt"

func SlicesPointers() {
	names := [4]string{
		"ondul",
		"Paul",
		"George",
		"solo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
