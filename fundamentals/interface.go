package main

import (
	"fmt"
)

// interface
type tank interface {
	// methods
	Tarea() float64
	Volume() float64
}
type myValue struct {
	redius float64
	height float64
}

// Implementing methods of
// the tank interface
func (m myValue) Tarea() float64 {
	return (2 * m.redius * m.height) + (2 * 3.14 * m.redius * m.redius)
}

func (m myValue) Volume() float64 {
	return (3.14 * m.redius * m.redius * m.height)
}

// main method
func main() {
	// accessing elements of the tank interface
	var t tank
	t = myValue{10, 14}
	fmt.Println("Area of tank: ", t.Tarea())
	fmt.Println("Volumen of tank: ", t.Volume())
	fmt.Println(true)
}
