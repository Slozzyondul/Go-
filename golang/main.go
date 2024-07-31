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
	goroutines1 "golang/goroutines"
	iflo "golang/ifLo"
	"golang/ifelse"
	interfaces1 "golang/interfaces"
	"golang/intro"
	maps1 "golang/maps"
	methods1 "golang/methods"
	nakedreturn "golang/naked_return"
	"golang/pointers"
	range1 "golang/range"
	slices1 "golang/slices"
	"golang/structs"
	switch1 "golang/switch"
	typeparameters "golang/type_parameters"
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
	slices1.NilSlices()
	slices1.MakingSlices()
	slices1.SlicesofSlices()
	slices1.AppendSlice()
	range1.Range()
	range1.Range1()
	slices1.SliceExercise()
	//slices1.DecodeImage()
	maps1.Maps1()
	maps1.MapLiterals()
	maps1.MapsLiterals1()
	maps1.MutatingMaps()
	maps1.ExerciseWordcount()
	functions.FUnctionValues()
	functions.Closure()
	functions.Fibonacci()
	methods1.Methods1()
	methods1.Methods2()
	methods1.Methods3()
	methods1.Methods4()
	methods1.Methods5()
	methods1.Methods6()
	methods1.IndirectionalReverse()
	methods1.PointerReceiver()
	interfaces1.Interface1()
	//interfaces1.TypeInterfaces()
	interfaces1.InterfaceValues()
	interfaces1.NilInterfaces()
	interfaces1.EmptyInterface()
	interfaces1.TypeAssertionInterface()
	interfaces1.SwitchInterface()
	interfaces1.StringerInterface()
	interfaces1.IpAddressInterface()
	interfaces1.Error_interface()
	interfaces1.Newtonmethod1()
	interfaces1.ReaderInterface()
	interfaces1.ReaderInterfaceInfinityLoop()
	interfaces1.RotReaderInterface()
	interfaces1.ImageInterface()
	interfaces1.ImageInterface1()
	typeparameters.TypeParameters()
	typeparameters.GenericType()
	goroutines1.Goroutine1()
	goroutines1.ChannelsWithgoroutines()
	//goroutines1.GoroutineDeadlock()
	goroutines1.GoroutineWithoutDeadlock()
	goroutines1.SenderReceiver()
	goroutines1.SelectGoroutine()
}
