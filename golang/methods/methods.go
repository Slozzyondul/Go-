package methods1

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Methods1() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
