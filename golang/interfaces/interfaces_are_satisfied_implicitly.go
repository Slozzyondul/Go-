package interfaces1

//import "fmt"

//A type implements an interface by implementing its methods.
//There is no explicit declaration of intent, no "implements" keyword.

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
// func (t T) M() {
// 	fmt.Println(t.S)
// }

// func TypeInterfaces() {
// 	var i I = T{"slozzy"}
// 	i.M()
// }
