package slices1

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// Create a 2D slice with the given dimensions
	picture := make([][]uint8, dy)
	for y := range picture {
		picture[y] = make([]uint8, dx)
		for x := range picture[y] {
			// Fill each element with the value (x + y) / 2
			picture[y][x] = uint8((x + y) / 2)
		}
	}
	return picture
}

func SliceExercise() {
	pic.Show(Pic)
}
