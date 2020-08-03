package geometry_tool

import (
	"errors"
	"github.com/nodejayes/geolib/pkg/geometry"
	"reflect"
)

func copyPointer(dest, src interface{}) error {
	destVal := reflect.ValueOf(dest)
	srcVal := reflect.ValueOf(src)
	if destVal.Kind() != reflect.Ptr || srcVal.Kind() != reflect.Ptr {
		return errors.New("both inputs must be a Pointer")
	}
	srcVal.Elem().Set(destVal.Elem())
	return nil
}

// GetCoordinates returns the Coordinates of all Possible Geometries
func GetCoordinates(target interface{}, coords interface{}) error {
	point, pointOk := target.(*geometry.Point)
	_, coordOk := coords.(*[]float64)
	if pointOk && coordOk {
		err := copyPointer(&point.Coordinates, coords)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("not supported geometry type")
}
