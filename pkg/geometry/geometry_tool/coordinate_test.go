package geometry_tool

import (
	"fmt"
	"github.com/nodejayes/geolib/pkg/geometry"
	"testing"
)

var GetCoordinatesTest = []struct {
	name        string
	srid        int
	geometry    string
	c1          []float64
	c2          [][]float64
	c3          [][][]float64
	c4          [][][][]float64
	expectedErr string
}{
	{
		name:     "Get Point Coordinate",
		geometry: "Point",
		srid:     4326,
		c1:       []float64{1, 1},
	},
	{
		name:     "Get Line Coordinate",
		geometry: "Line",
		srid:     4326,
		c2:       [][]float64{{1, 1}, {2, 2}},
	},
	{
		name:     "Get Polygon Coordinate",
		geometry: "Polygon",
		srid:     4326,
		c3:       [][][]float64{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}},
	},
	{
		name:     "Get MultiPoint Coordinate",
		geometry: "MultiPoint",
		srid:     4326,
		c2:       [][]float64{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}},
	},
	{
		name:     "Get MultiLine Coordinate",
		geometry: "MultiLine",
		srid:     4326,
		c3:       [][][]float64{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}},
	},
	{
		name:     "Get MultiPolygon Coordinate",
		geometry: "MultiPolygon",
		srid:     4326,
		c4:       [][][][]float64{{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}}},
	},
}

func TestGetCoordinates(t *testing.T) {
	for _, tt := range GetCoordinatesTest {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			switch tt.geometry {
			case "Point":
				g := geometry.NewPoint(tt.c1, tt.srid)
				var tmp []float64
				err = GetCoordinates(g, &tmp)
				if err == nil && (tmp[0] != tt.c1[0] || tmp[1] != tt.c1[1]) {
					t.Errorf("Expect tmp to be %v: %v", tt.c1, tmp)
					return
				}
			case "Line":
				g := geometry.NewLine(tt.c2, tt.srid)
				var tmp [][]float64
				err = GetCoordinates(g, &tmp)
				for idx, p := range tt.c2 {
					if err == nil && (tmp[idx][0] != p[0] || tmp[idx][1] != p[1]) {
						t.Errorf("Expect tmp to be %v: %v", tt.c1, tmp)
						return
					}
				}
			case "Polygon":
				g := geometry.NewPolygon(tt.c3, tt.srid)
				var tmp [][][]float64
				err = GetCoordinates(g, &tmp)
				for idx1, ring := range tt.c3 {
					for idx2, p := range ring {
						if err == nil && (tmp[idx1][idx2][0] != p[0] || tmp[idx1][idx2][1] != p[1]) {
							t.Errorf("Expect tmp to be %v: %v", tt.c1, tmp)
							return
						}
					}
				}
			case "MultiPoint":
				g := geometry.NewMultiPoint(tt.c2, tt.srid)
				var tmp [][]float64
				err = GetCoordinates(g, &tmp)
				for idx, p := range tt.c2 {
					if err == nil && (tmp[idx][0] != p[0] || tmp[idx][1] != p[1]) {
						t.Errorf("Expect tmp to be %v: %v", tt.c1, tmp)
						return
					}
				}
			case "MultiLine":
				g := geometry.NewMultiLine(tt.c3, tt.srid)
				var tmp [][][]float64
				err = GetCoordinates(g, &tmp)
				for idx1, ring := range tt.c3 {
					for idx2, p := range ring {
						if err == nil && (tmp[idx1][idx2][0] != p[0] || tmp[idx1][idx2][1] != p[1]) {
							t.Errorf("Expect tmp to be %v: %v", tt.c1, tmp)
							return
						}
					}
				}
			case "MultiPolygon":
				g := geometry.NewMultiPolygon(tt.c4, tt.srid)
				var tmp [][][][]float64
				err = GetCoordinates(g, &tmp)
				for idx1, poly := range tt.c4 {
					for idx2, ring := range poly {
						for idx3, p := range ring {
							if err == nil && (tmp[idx1][idx2][idx3][0] != p[0] || tmp[idx1][idx2][idx3][1] != p[1]) {
								t.Errorf("Expect tmp to be %v: %v", tt.c1, tmp)
								return
							}
						}
					}
				}
			}

			if tt.expectedErr == "" && err != nil {
				t.Errorf("Expect err to be nil: %v", err.Error())
				return
			}
			if err == nil && tt.expectedErr != "" {
				t.Errorf("Expect err to be %v: nil", tt.expectedErr)
				return
			}
			if err != nil && tt.expectedErr != err.Error() {
				t.Errorf("Expect err to be %v: %v", tt.expectedErr, err.Error())
				return
			}
		})
	}
}

func ExampleGetCoordinates() {
	p := geometry.NewPoint([]float64{1, 2}, 4326)
	var pointCoordinates []float64
	_ = GetCoordinates(p, &pointCoordinates)
	fmt.Printf("%v", pointCoordinates)
	// Output: [1 2]
}
