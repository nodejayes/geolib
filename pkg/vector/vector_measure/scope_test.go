package vector_measure

import (
	"fmt"
	"testing"
)

var ScopeTest = []struct {
	name          string
	triangle      TestTriangle
	expectedScope float64
}{
	{
		name: "Scope",
		triangle: TestTriangle{
			A: []float64{0, 0},
			B: []float64{5, 1},
			C: []float64{3, 2},
		},
		expectedScope: 10.940638766556564,
	},
}

func TestScope(t *testing.T) {
	for _, tt := range ScopeTest {
		t.Run(tt.name, func(t *testing.T) {
			val, err := Scope(tt.triangle)
			if err != nil {
				t.Errorf("Expect err to be nil: %v", err.Error())
				return
			}
			if val != tt.expectedScope {
				t.Errorf("Expect val to be %v: %v", tt.expectedScope, val)
				return
			}
		})
	}
}

func ExampleScope() {
	t := TestTriangle{
		A: []float64{0, 0},
		B: []float64{5, 1},
		C: []float64{3, 2},
	}
	val, _ := Scope(t)
	fmt.Printf("%v", val)
	// Output: 10.940638766556564
}
