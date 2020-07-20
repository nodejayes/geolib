package geometry

import "github.com/nodejayes/geolib/pkg/reference_system"

type Polygon struct {
	Type        string                           `json:"type"`
	Coordinates Coordinate3D                     `json:"coordinates"`
	CRS         reference_system.ReferenceSystem `json:"crs"`
}

func NewPolygon(coords Coordinate3D, crs reference_system.ReferenceSystem) Polygon {
	return Polygon{
		Type:        PolygonType,
		Coordinates: coords,
		CRS:         crs,
	}
}
