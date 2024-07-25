package methods1

import (
	"fmt"
	"math"
)

func (v Vertex) Abs1() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Methods4() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
