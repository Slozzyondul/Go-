package interfaces1

//An empty interface may hold values of any type. (Every type implements at least zero methods.)

import "fmt"

func EmptyInterface() {
	var i interface{}
	describe1(i)

	i = 42
	describe1(i)

	i = "hello"
	describe1(i)
}

func describe1(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
