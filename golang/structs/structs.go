package structs

// structs ara a collection of fields

import "fmt"

type Vertex struct {
	X int
	Y int
}

func Structs() {
	fmt.Println(Vertex{1, 2})
}
