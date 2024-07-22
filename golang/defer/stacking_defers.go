package defer1

import "fmt"

func StackData() {
	fmt.Println("counting")

	for i := 0; i <= 10; i++ {
		defer fmt.Println(i)

	}

}
