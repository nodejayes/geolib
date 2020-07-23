package geometry

import "github.com/nodejayes/geolib/pkg/reference_system"

type MultiLine struct {
	Type        string                           `json:"type"`
	Coordinates Coordinate3D                     `json:"coordinates"`
	CRS         reference_system.ReferenceSystem `json:"crs"`
}

func NewMultiLine(coordinates Coordinate3D, crs reference_system.ReferenceSystem) *MultiLine {
	return &MultiLine{
		Type:        MultiLineType,
		Coordinates: coordinates,
		CRS:         crs,
	}
}

func (ctx *MultiLine) GetCoordinates(data interface{}) {
	switch d := data.(type) {
	case *Coordinate3D:
		*d = append(*d, ctx.Coordinates...)
	}
}

func (ctx *MultiLine) Transform(target int) error {
	var tmp Coordinate3D
	for _, line := range ctx.Coordinates {
		transformed, err := ctx.CRS.TransformPoints(target, line)
		if err != nil {
			return err
		}
		tmp = append(tmp, transformed)
	}
	ctx.Coordinates = tmp
	ctx.CRS = reference_system.New(target)
	return nil
}
