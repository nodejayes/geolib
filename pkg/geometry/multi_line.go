package geometry

import "github.com/nodejayes/geolib/pkg/reference_system"

type MultiLine struct {
	Type        string                           `json:"type"`
	Coordinates Coordinate3D                     `json:"coordinates"`
	CRS         reference_system.ReferenceSystem `json:"crs"`
}

func NewMultiLine(coordinates Coordinate3D, crs reference_system.ReferenceSystem) MultiLine {
	return MultiLine{
		Type:        MultiLineType,
		Coordinates: coordinates,
		CRS:         crs,
	}
}
