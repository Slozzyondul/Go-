package main

import (
	"fmt"
	"golang/functions"
	"golang/intro"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	intro.Solo()
	intro.Currentime()
	intro.Squareroot()
	functions.Addition()
	functions.Minus()
	functions.SwapStrings()

}
