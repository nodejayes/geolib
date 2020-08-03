package vector

import (
	"fmt"
)

func ExampleNewTriangle() {
	t := NewTriangle([]float64{0, 0}, []float64{5, 1}, []float64{3, 2})
	fmt.Printf("%v", t)
	// Output: &{[0 0] [5 1] [3 2]}
}

func ExampleTriangle_GetA() {
	t := NewTriangle([]float64{0, 0}, []float64{5, 1}, []float64{3, 2})
	fmt.Printf("%v", t.GetA())
	// Output: [0 0]
}

func ExampleTriangle_GetB() {
	t := NewTriangle([]float64{0, 0}, []float64{5, 1}, []float64{3, 2})
	fmt.Printf("%v", t.GetB())
	// Output: [5 1]
}

func ExampleTriangle_GetC() {
	t := NewTriangle([]float64{0, 0}, []float64{5, 1}, []float64{3, 2})
	fmt.Printf("%v", t.GetC())
	// Output: [3 2]
}
