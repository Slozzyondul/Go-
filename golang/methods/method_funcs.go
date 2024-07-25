package methods1

import (
	"fmt"
	"math"
)

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Methods2() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
