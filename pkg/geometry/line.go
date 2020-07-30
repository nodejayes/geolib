package geometry

import (
	"errors"
	"github.com/nodejayes/geolib/pkg/reference_system"
)

type Line struct {
	Type        string                           `json:"type"`
	Coordinates Coordinate2D                     `json:"coordinates"`
	CRS         reference_system.ReferenceSystem `json:"crs"`
}

func NewLine(coordinates Coordinate2D, srid int) *Line {
	return &Line{
		Type:        LineType,
		Coordinates: coordinates,
		CRS:         reference_system.New(srid),
	}
}

func (ctx *Line) GetCoordinates(data interface{}) error {
	switch d := data.(type) {
	case *Coordinate2D:
		*d = append(*d, ctx.Coordinates...)
		return nil
	}
	return errors.New("wrong type given expect [][]float64")
}

func (ctx *Line) Transform(target int) error {
	transformed, err := ctx.CRS.TransformPoints(target, ctx.Coordinates)
	if err != nil {
		return err
	}
	ctx.Coordinates = transformed
	ctx.CRS = reference_system.New(target)
	return nil
}
