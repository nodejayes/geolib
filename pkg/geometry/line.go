package geometry

import "github.com/nodejayes/geolib/pkg/reference_system"

type Line struct {
	Type        string                           `json:"type"`
	Coordinates Coordinate2D                     `json:"coordinates"`
	CRS         reference_system.ReferenceSystem `json:"crs"`
}

func NewLine(coordinates Coordinate2D, crs reference_system.ReferenceSystem) Line {
	return Line{
		Type:        LineType,
		Coordinates: coordinates,
		CRS:         crs,
	}
}
