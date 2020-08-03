package vector_measure

import (
	"fmt"
	"testing"
)

var AngleTest = []struct {
	name          string
	triangle      TestTriangle
	angle         PointKey
	expectedAngle float64
}{
	{
		name: "Angle Alpha",
		triangle: TestTriangle{
			A: []float64{0, 0},
			B: []float64{5, 1},
			C: []float64{3, 2},
		},
		angle:         Alpha,
		expectedAngle: 22.38013524215307,
	},
	{
		name: "Angle Beta",
		triangle: TestTriangle{
			A: []float64{0, 0},
			B: []float64{5, 1},
			C: []float64{3, 2},
		},
		angle:         Beta,
		expectedAngle: 37.874983972971805,
	},
	{
		name: "Angle Gamma",
		triangle: TestTriangle{
			A: []float64{0, 0},
			B: []float64{5, 1},
			C: []float64{3, 2},
		},
		angle:         Gamma,
		expectedAngle: 119.74488231457212,
	},
}

func TestAngleFrom(t *testing.T) {
	angleSum := 0.0
	for idx, tt := range AngleTest {
		t.Run(tt.name, func(t *testing.T) {
			val, err := AngleFrom(tt.triangle, tt.angle)
			if idx >= 0 && idx <= 2 {
				angleSum += val
			}
			if err != nil {
				t.Errorf("Expect err to be nil: %v", err.Error())
				return
			}
			if val != tt.expectedAngle {
				t.Errorf("Expect angle to be %v: %v", tt.expectedAngle, val)
			}
		})
	}
	t.Run("all Angles must have 180°", func(t *testing.T) {
		if angleSum != 180.000001529697 {
			t.Errorf("angles of a Triangle must have 180°: %v", angleSum)
		}
	})
}

func ExampleAngleFrom() {
	t := TestTriangle{
		A: []float64{0, 0},
		B: []float64{5, 1},
		C: []float64{3, 2},
	}
	angle, _ := AngleFrom(t, Alpha)
	fmt.Printf("%v", angle)
	// Output: 22.38013524215307
}
