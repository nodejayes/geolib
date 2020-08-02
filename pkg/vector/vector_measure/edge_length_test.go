package vector_measure

import (
	"fmt"
	"testing"
)

type TestTriangle struct {
	A []float64
	B []float64
	C []float64
}

func (t TestTriangle) GetA() []float64 {
	return t.A
}

func (t TestTriangle) GetB() []float64 {
	return t.B
}

func (t TestTriangle) GetC() []float64 {
	return t.C
}

var EdgeLengthTests = []struct {
	name           string
	triangle       TestTriangle
	edge           EdgeKey
	expectedLength float64
}{
	{
		name: "Test EdgeAB",
		triangle: TestTriangle{
			A: []float64{
				386130, 5692254,
			},
			B: []float64{
				386408.967053963, 5692913.35517462,
			},
			C: []float64{
				386106.870525168, 5692632.47844243,
			},
		},
		edge:           EdgeAB,
		expectedLength: 715.941243046222,
	},
	{
		name: "Test EdgeAC",
		triangle: TestTriangle{
			A: []float64{
				386130, 5692254,
			},
			B: []float64{
				386408.967053963, 5692913.35517462,
			},
			C: []float64{
				386106.870525168, 5692632.47844243,
			},
		},
		edge:           EdgeAC,
		expectedLength: 379.18452498739,
	},
	{
		name: "Test EdgeCB",
		triangle: TestTriangle{
			A: []float64{
				386130, 5692254,
			},
			B: []float64{
				386408.967053963, 5692913.35517462,
			},
			C: []float64{
				386106.870525168, 5692632.47844243,
			},
		},
		edge:           EdgeCB,
		expectedLength: 412.49733501657795,
	},
	{
		name: "Test EdgeAB A: 0,0 B: 5,1 C: 3,2",
		triangle: TestTriangle{
			A: []float64{0, 0},
			B: []float64{5, 1},
			C: []float64{3, 2},
		},
		edge:           EdgeAB,
		expectedLength: 5.0990195135927845,
	},
	{
		name: "Test EdgeAC A: 0,0 B: 5,1 C: 3,2",
		triangle: TestTriangle{
			A: []float64{0, 0},
			B: []float64{5, 1},
			C: []float64{3, 2},
		},
		edge:           EdgeAC,
		expectedLength: 3.605551275463989,
	},
	{
		name: "Test EdgeAC A: 0,0 B: 5,1 C: 3,2",
		triangle: TestTriangle{
			A: []float64{0, 0},
			B: []float64{5, 1},
			C: []float64{3, 2},
		},
		edge:           EdgeCB,
		expectedLength: 2.23606797749979,
	},
}

func TestEdgeLength(t *testing.T) {
	for _, tt := range EdgeLengthTests {
		t.Run(tt.name, func(t *testing.T) {
			l, err := EdgeLength(tt.triangle, tt.edge)
			if err != nil {
				t.Errorf("expect err to be nil: %v", err.Error())
				return
			}
			if l != tt.expectedLength {
				t.Errorf("expect l to be %v: %v", tt.expectedLength, l)
				return
			}
		})
	}
}

func ExampleEdgeLength() {
	t := TestTriangle{
		A: []float64{
			386130, 5692254,
		},
		B: []float64{
			386408.967053963, 5692913.35517462,
		},
		C: []float64{
			386106.870525168, 5692632.47844243,
		},
	}
	abLength, _ := EdgeLength(t, EdgeAB)
	fmt.Println(abLength)
	// Output: 715.941243046222
}
