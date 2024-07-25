package maps1

import "fmt"

var o = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func MapsLiterals1() {
	fmt.Println(o)
}
