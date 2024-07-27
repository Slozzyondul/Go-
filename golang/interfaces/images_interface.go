package interfaces1

import (
	"fmt"
	"image"
)

func ImageInterface() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())

	println(m)
}
