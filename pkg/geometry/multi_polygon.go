package geometry

import "github.com/nodejayes/geolib/pkg/reference_system"

type MultiPolygon struct {
	Type        string                           `json:"type"`
	Coordinates Coordinate4D                     `json:"coordinates"`
	CRS         reference_system.ReferenceSystem `json:"crs"`
}

func NewMultiPolygon(coordinates Coordinate4D, crs reference_system.ReferenceSystem) *MultiPolygon {
	return &MultiPolygon{
		Type:        MultiPolygonType,
		Coordinates: coordinates,
		CRS:         crs,
	}
}

func (ctx *MultiPolygon) GetCoordinates(data interface{}) {
	switch d := data.(type) {
	case *Coordinate4D:
		*d = append(*d, ctx.Coordinates...)
	}
}

func (ctx *MultiPolygon) Transform(target int) error {
	var tmp Coordinate4D
	for _, poly := range ctx.Coordinates {
		var tmp2 Coordinate3D
		for _, ring := range poly {
			transformed, err := ctx.CRS.TransformPoints(target, ring)
			if err != nil {
				return err
			}
			tmp2 = append(tmp2, transformed)
		}
		tmp = append(tmp, tmp2)
	}
	ctx.Coordinates = tmp
	ctx.CRS = reference_system.New(target)
	return nil
}
