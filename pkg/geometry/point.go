package geometry

import (
	"github.com/nodejayes/geolib/pkg/reference_system"
)

type Point struct {
	Type        string                           `json:"type"`
	Coordinates Coordinate1D                     `json:"coordinates"`
	CRS         reference_system.ReferenceSystem `json:"crs"`
}

func NewPoint(coordinates Coordinate1D, crs reference_system.ReferenceSystem) Point {
	return Point{
		Type:        PointType,
		Coordinates: coordinates,
		CRS:         crs,
	}
}
