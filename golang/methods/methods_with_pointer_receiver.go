package methods1

import (
	"fmt"
	//"math"
)

// There are two reasons to use a pointer receiver.

//The first is so that the method can modify the value that its receiver points to.

//The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.

// type Vertex struct {
// 	X, Y float64
// }

// func (v *Vertex) Scale(f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

// func (v *Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

func PointerReceiver() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
