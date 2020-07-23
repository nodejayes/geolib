package geometry

import "github.com/nodejayes/geolib/pkg/reference_system"

type Line struct {
	Type        string                           `json:"type"`
	Coordinates Coordinate2D                     `json:"coordinates"`
	CRS         reference_system.ReferenceSystem `json:"crs"`
}

func NewLine(coordinates Coordinate2D, crs reference_system.ReferenceSystem) *Line {
	return &Line{
		Type:        LineType,
		Coordinates: coordinates,
		CRS:         crs,
	}
}

func (ctx *Line) GetCoordinates(data interface{}) {
	switch d := data.(type) {
	case *Coordinate2D:
		*d = append(*d, ctx.Coordinates...)
	}
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
