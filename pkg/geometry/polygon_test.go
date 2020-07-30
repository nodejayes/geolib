package geometry

import (
	"encoding/json"
	"fmt"
	"testing"
)

func ExampleNewPolygon() {
	polygon := NewPolygon(Coordinate3D{{{1, 2}, {3, 4}, {5, 6}, {1, 2}}}, 4326)
	stream, _ := json.Marshal(polygon)
	fmt.Println(string(stream))
	// Output: {"type":"Polygon","coordinates":[[[1,2],[3,4],[5,6],[1,2]]],"crs":{"type":"name","properties":{"name":"EPSG:4326"}}}
}

func ExamplePolygon_GetCoordinates() {
	polygon := NewPolygon(Coordinate3D{{{1, 2}, {3, 4}, {5, 6}, {1, 2}}}, 4326)
	var c [][][]float64
	// read into c the coordinates
	// only input the Reference !!!
	_ = polygon.GetCoordinates(&c)
	fmt.Println(c)
	// Output: [[[1 2] [3 4] [5 6] [1 2]]]
}

func ExamplePolygon_Transform() {
	polygon := NewPolygon(Coordinate3D{{{1, 2}, {3, 4}, {5, 6}, {1, 2}}}, 4326)
	// transform the Polygon into PseudoMercator
	_ = polygon.Transform(3857)
	stream, _ := json.Marshal(polygon)
	fmt.Println(string(stream))
	// Output: {"type":"Polygon","coordinates":[[[111319.49079327357,222684.20850554318],[333958.4723798207,445640.10965602525],[556597.4539663679,669141.0570442441],[111319.49079327357,222684.20850554318]]],"crs":{"type":"name","properties":{"name":"EPSG:3857"}}}
}

var CoordinateTest = []struct {
	testName        string
	geometry        func(coords interface{}, srid int) IGeometrier
	coordinates     interface{}
	srid            int
	c               interface{}
	expectedErr     string
	nilErr          string
	wrongMessageErr string
}{
	{
		testName:    "Error on invalid coordinate type",
		coordinates: Coordinate1D{1, 2},
		srid:        4326,
		geometry: func(coords interface{}, srid int) IGeometrier {
			switch c := coords.(type) {
			case []float64:
				return NewPoint(c, srid)
			}
			return nil
		},
		c:               [][]float64{},
		expectedErr:     "wrong type given expect []float64",
		nilErr:          "Point Coordinates was written into [][]float64 expect to return a error",
		wrongMessageErr: "Wrong Error Message\n%v \nexpect: %v",
	},
	{
		testName:    "Error on invalid coordinate type",
		coordinates: Coordinate2D{{1, 2}, {3, 4}, {5, 6}, {1, 2}},
		srid:        4326,
		geometry: func(coords interface{}, srid int) IGeometrier {
			switch c := coords.(type) {
			case [][]float64:
				return NewLine(c, srid)
			}
			return nil
		},
		c:               []float64{},
		expectedErr:     "wrong type given expect [][]float64",
		nilErr:          "Line Coordinates was written into []float64 expect to return a error",
		wrongMessageErr: "Wrong Error Message\n%v \nexpect: %v",
	},
	{
		testName:    "Error on invalid coordinate type",
		coordinates: Coordinate3D{{{1, 2}, {3, 4}, {5, 6}, {1, 2}}},
		srid:        4326,
		geometry: func(coords interface{}, srid int) IGeometrier {
			switch c := coords.(type) {
			case [][][]float64:
				return NewPolygon(c, srid)
			}
			return nil
		},
		c:               []float64{},
		expectedErr:     "wrong type given expect [][][]float64",
		nilErr:          "Polygon Coordinates was written into []float64 expect to return a error",
		wrongMessageErr: "Wrong Error Message\n%v \nexpect: %v",
	},
}

func TestPolygon_GetCoordinates(t *testing.T) {
	for _, tt := range CoordinateTest {
		t.Run(tt.testName, func(t *testing.T) {
			geom := tt.geometry(tt.coordinates, tt.srid)
			err := geom.GetCoordinates(&tt.c)
			if err == nil {
				t.Errorf(tt.nilErr)
				return
			}
			if err.Error() != tt.expectedErr {
				t.Errorf(tt.wrongMessageErr, err.Error(), tt.expectedErr)
				return
			}
		})
	}
}
