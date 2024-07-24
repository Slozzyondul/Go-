package structs

import "fmt"

func StructsField() {
	v := Vertex{4, 2}
	v.X = 4
	v.Y = 2
	fmt.Println(v.X, v.Y)
}
