package geometry

import "github.com/nodejayes/geolib/pkg/reference_system"

type MultiPoint struct {
	Type        string                           `json:"type"`
	Coordinates Coordinate2D                     `json:"coordinates"`
	CRS         reference_system.ReferenceSystem `json:"crs"`
}

func NewMultiPoint(coordinates Coordinate2D, crs reference_system.ReferenceSystem) MultiPoint {
	return MultiPoint{
		Type:        MultiPointType,
		Coordinates: coordinates,
		CRS:         crs,
	}
}
