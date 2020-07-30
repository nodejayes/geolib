package geometry

import (
	"testing"
)

func TestPoint_GetCoordinates(t *testing.T) {
	t.Run("Error on invalid coordinate type", func(t *testing.T) {
		expected := "wrong type given expect []float64"
		pt := NewPoint([]float64{1, 2}, 4326)
		var coords [][]float64
		err := pt.GetCoordinates(&coords)
		if err == nil {
			t.Errorf("Point Coordinates was written into [][]float64 expect to return a error")
			return
		}
		if err.Error() != expected {
			t.Errorf("Wrong Error Message\n%v \nexpect: %v", err.Error(), expected)
		}
	})
}
