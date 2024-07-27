package interfaces1

import (
	"errors"
	"fmt"
	"math"
)

// SqrtImprovedWithError calculates the square root of x using Newton's method
// and stops when the change in z is very small. It returns an error if x is negative.
func SqrtImprovedWithError(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("cannot calculate the square root of a negative number")
	}
	if x == 0 {
		return 0, nil
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
	return z, nil
}

func Newtonmethod1() {
	values := []float64{16, 2, 3, 4, 5, 10, 100, -1000}

	fmt.Println("Using fixed 10 iterations:")
	for _, v := range values {
		fmt.Printf("Calculating sqrt(%v)\n", v)
		result := math.Sqrt(v)
		fmt.Printf("Approximate: %v, math.Sqrt: %v\n\n", result, math.Sqrt(v))
	}

	fmt.Println("Using convergence criteria:")
	for _, v := range values {
		fmt.Printf("Calculating sqrt(%v)\n", v)
		result, err := SqrtImprovedWithError(v)
		if err != nil {
			fmt.Printf("Error: %v\n\n", err)
		} else {
			fmt.Printf("Approximate: %v, math.Sqrt: %v\n\n", result, math.Sqrt(v))
		}
	}
}
