package main

import (
	"fmt"
	typeconversions "golang/Type_conversions"
	"golang/arrays"
	basictype "golang/basic_type"
	"golang/constants"
	defer1 "golang/defer"
	forloop "golang/forLo"
	"golang/functions"
	iflo "golang/ifLo"
	"golang/ifelse"
	"golang/intro"
	nakedreturn "golang/naked_return"
	"golang/pointers"
	slices1 "golang/slices"
	"golang/structs"
	switch1 "golang/switch"
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
	//exercise.Newtonmethod()
	switch1.Switch()
	switch1.EvaluationSwitch()
	switch1.NoConditionSwitch()
	defer1.Defer()
	defer1.StackData()
	pointers.Pointers()
	structs.Structs()
	structs.StructsField()
	structs.StructPointers()
	structs.StructLiterals()
	arrays.Arrays()
	slices1.Slices()
	slices1.SlicesPointers()
	slices1.SliceLiterals()
	slices1.SliceBounds()
	slices1.LengthCapacity()
}
