package interfaces1

//A type assertion provides access to an interface value's underlying concrete value.

import "fmt"

func TypeAssertionInterface() {
	var i interface{} = "hello slozzy"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// f = i.(float64) // panic
	// fmt.Println(f)
}
