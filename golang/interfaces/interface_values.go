package interfaces1

import (
	"fmt"
	"math"
)

type F float64

func (f F) M() {
	fmt.Println(f)
}

func InterfaceValues() {
	var i I

	i = &T{"Under the hood, interface values can be thought of as a tuple of a value and a concrete type"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
