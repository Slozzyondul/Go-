package maps1

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func Maps1() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	m["solo"] = Vertex{
		27, 1996,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m["solo"])
}
