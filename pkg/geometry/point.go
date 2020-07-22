package geometry

import (
	"github.com/nodejayes/geolib/pkg/reference_system"
)

type Point struct {
	Type        string                           `json:"type"`
	Coordinates Coordinate1D                     `json:"coordinates"`
	CRS         reference_system.ReferenceSystem `json:"crs"`
}

func NewPoint(coordinates Coordinate1D, crs reference_system.ReferenceSystem) *Point {
	return &Point{
		Type:        PointType,
		Coordinates: coordinates,
		CRS:         crs,
	}
}

func (ctx *Point) Transform(target int) error {
	transformed, err := ctx.CRS.TransformPoints(target, [][]float64{ctx.Coordinates})
	if err != nil {
		return err
	}
	ctx.CRS = reference_system.New(target)
	ctx.Coordinates = transformed[0]
	return nil
}
