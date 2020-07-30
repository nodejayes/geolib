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

func TestPolygon_GetCoordinates(t *testing.T) {
	t.Run("Error on invalid coordinate type", func(t *testing.T) {
		polygon := NewPolygon(Coordinate3D{{{1, 2}, {3, 4}, {5, 6}, {1, 2}}}, 4326)
		var c []float64
		err := polygon.GetCoordinates(&c)
		if err == nil {
			t.Errorf("Error is nil")
			return
		}
		if err.Error() != "wrong type given expect [][][]float64" {
			t.Errorf("Wrong Error Message\n%v \nexpect: wrong type given expect [][][]float64", err.Error())
		}
	})
}
