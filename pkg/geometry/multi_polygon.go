package geometry

import "github.com/nodejayes/geolib/pkg/reference_system"

type MultiPolygon struct {
	Type        string                           `json:"type"`
	Coordinates Coordinate4D                     `json:"coordinates"`
	CRS         reference_system.ReferenceSystem `json:"crs"`
}

func NewMultiPolygon(coordinates Coordinate4D, crs reference_system.ReferenceSystem) MultiPolygon {
	return MultiPolygon{
		Type:        MultiPolygonType,
		Coordinates: coordinates,
		CRS:         crs,
	}
}
