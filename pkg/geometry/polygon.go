package geometry

import "github.com/nodejayes/geolib/pkg/reference_system"

type Polygon struct {
	Type        string                           `json:"type"`
	Coordinates Coordinate3D                     `json:"coordinates"`
	CRS         reference_system.ReferenceSystem `json:"crs"`
}

func NewPolygon(coords Coordinate3D, crs reference_system.ReferenceSystem) *Polygon {
	return &Polygon{
		Type:        PolygonType,
		Coordinates: coords,
		CRS:         crs,
	}
}

func (ctx *Polygon) GetCoordinates(data interface{}) {
	switch d := data.(type) {
	case *Coordinate3D:
		*d = append(*d, ctx.Coordinates...)
	}
}

func (ctx *Polygon) Transform(target int) error {
	var tmp Coordinate3D
	for _, ring := range ctx.Coordinates {
		transformed, err := ctx.CRS.TransformPoints(target, ring)
		if err != nil {
			return err
		}
		tmp = append(tmp, transformed)
	}
	ctx.Coordinates = tmp
	ctx.CRS = reference_system.New(target)
	return nil
}
