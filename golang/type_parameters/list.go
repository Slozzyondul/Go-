package typeparameters

import (
	"fmt"
)

// Go also supports generic types.
// A type can be parameterized with a type parameter, which could be useful for implementing generic data structures.
// List represents a singly-linked list that holds values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

// Add appends a new element to the end of the list.
func (l *List[T]) Add(val T) {
	if l.next == nil {
		l.next = &List[T]{val: val}
	} else {
		l.next.Add(val)
	}
}

// Remove removes the first occurrence of the specified value from the list using a comparator function.
func (l *List[T]) Remove(val T, eq func(a, b T) bool) bool {
	if l.next == nil {
		return false
	}
	if eq(l.next.val, val) {
		l.next = l.next.next
		return true
	}
	return l.next.Remove(val, eq)
}

// Find returns true if the specified value is in the list using a comparator function.
func (l *List[T]) Find(val T, eq func(a, b T) bool) bool {
	if l.next == nil {
		return false
	}
	if eq(l.next.val, val) {
		return true
	}
	return l.next.Find(val, eq)
}

// Display prints all the elements in the list.
func (l *List[T]) Display() {
	if l.next == nil {
		return
	}
	fmt.Println(l.next.val)
	l.next.Display()
}

// Length returns the number of elements in the list.
func (l *List[T]) Length() int {
	if l.next == nil {
		return 0
	}
	return 1 + l.next.Length()
}

func GenericType() {
	// Example usage
	list := &List[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Display()
	fmt.Println("Length:", list.Length())
	// Custom equality function for integers
	eq := func(a, b int) bool {
		return a == b
	}

	fmt.Println("Find 2:", list.Find(2, eq))
	fmt.Println("Remove 2:", list.Remove(2, eq))
	list.Display()
	fmt.Println("Length:", list.Length())
}
