package geometry

import (
	"errors"
	"github.com/nodejayes/geolib/pkg/reference_system"
)

type Point struct {
	Type        string                           `json:"type"`
	Coordinates Coordinate1D                     `json:"coordinates"`
	CRS         reference_system.ReferenceSystem `json:"crs"`
}

func NewPoint(coordinates Coordinate1D, srid int) *Point {
	return &Point{
		Type:        PointType,
		Coordinates: coordinates,
		CRS:         reference_system.New(srid),
	}
}

func (ctx *Point) GetCoordinates(data interface{}) error {
	switch d := data.(type) {
	case *Coordinate1D:
		*d = append(*d, ctx.Coordinates...)
		return nil
	}
	return errors.New("wrong type given expect []float64")
}

func (ctx *Point) Transform(target int) error {
	transformed, err := ctx.CRS.TransformPoints(target, Coordinate2D{ctx.Coordinates})
	if err != nil {
		return err
	}
	ctx.CRS = reference_system.New(target)
	ctx.Coordinates = transformed[0]
	return nil
}
