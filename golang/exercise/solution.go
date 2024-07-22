package exercise

import (
	"fmt"
	"math"
)

// SqrtImproved calculates the square root of x using Newton's method
// and stops when the change in z is very small.
func SqrtImproved(x float64) float64 {
	if x == 0 {
		return 0
	}
	z := 1.0
	for {
		newZ := z - (z*z-x)/(2*z)
		if math.Abs(newZ-z) < 1e-9 {
			break
		}
		z = newZ
		fmt.Println(z)
	}
	return z
}

func Newtonmethod() {
	values := []float64{1, 2, 3, 4, 5, 10, 100, 1000}

	fmt.Println("Using fixed 10 iterations:")
	for _, v := range values {
		fmt.Printf("Calculating sqrt(%v)\n", v)
		result := math.Sqrt(v)
		fmt.Printf("Approximate: %v, math.Sqrt: %v\n\n", result, math.Sqrt(v))
	}

	fmt.Println("Using convergence criteria:")
	for _, v := range values {
		fmt.Printf("Calculating sqrt(%v)\n", v)
		result := SqrtImproved(v)
		fmt.Printf("Approximate: %v, math.Sqrt: %v\n\n", result, math.Sqrt(v))
	}
}
