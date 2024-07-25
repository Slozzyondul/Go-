package maps1

import "fmt"

var n = map[string]Vertex{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

func MapLiterals() {
	fmt.Println(n)
}
