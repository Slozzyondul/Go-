package interfaces1

import "fmt"

type Person struct {
	Name    string
	Age     int
	Country string
	married bool
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years) %v %v", p.Name, p.Age, p.Country, p.married)
}

func StringerInterface() {
	a := Person{"Solomon ondula", 27, "Kenya", false}
	z := Person{"Zaphod Beeblebrox", 9001, "unkwown", true}
	fmt.Println(a, z)
}
