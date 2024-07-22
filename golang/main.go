package main

import (
	"fmt"
	typeconversions "golang/Type_conversions"
	basictype "golang/basic_type"
	"golang/constants"
	"golang/exercise"
	forloop "golang/forLo"
	"golang/functions"
	iflo "golang/ifLo"
	"golang/ifelse"
	"golang/intro"
	nakedreturn "golang/naked_return"
	"golang/variables"
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
	nakedreturn.Split()
	variables.Variables()
	variables.Variablesinitializers()
	variables.Shortvariable()
	basictype.Basic()
	basictype.Zero()
	typeconversions.Type()
	typeconversions.Conversions()
	constants.Constants()
	constants.Numericconstants()
	forloop.Forloop()
	forloop.Forloop1()
	forloop.Forloop2()
	//forloop.Forloop3()
	iflo.Ifstatement()
	iflo.Ifstatement1()
	ifelse.Ifelse()
	exercise.Newtonmethod()
}
