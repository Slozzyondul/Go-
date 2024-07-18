package nakedreturn

import "fmt"

// naked return example
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func Split() {
	fmt.Println(split(17))
}
