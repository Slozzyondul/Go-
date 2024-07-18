package variables

import "fmt"

var i, j = 1, 2

func Variablesinitializers() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
