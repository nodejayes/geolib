package geometry

import (
	"errors"
	"github.com/nodejayes/geolib/pkg/reference_system"
)

// Polygon represent a Geometric Area in a Coordinate System.
type Polygon struct {
	Type        string                           `json:"type"`
	Coordinates Coordinate3D                     `json:"coordinates"`
	CRS         reference_system.ReferenceSystem `json:"crs"`
}

// NewPolygon create a new Polygon from Coordinates and SrId of a Coordinate System.
func NewPolygon(coords Coordinate3D, srid int) *Polygon {
	return &Polygon{
		Type:        PolygonType,
		Coordinates: coords,
		CRS:         reference_system.New(srid),
	}
}

// GetCoordinates write the Coordinates from the Polygon int the given data Variable
func (ctx *Polygon) GetCoordinates(data interface{}) error {
	switch d := data.(type) {
	case Coordinate3D:
	case *Coordinate3D:
		*d = append(*d, ctx.Coordinates...)
		return nil
	}
	return errors.New("wrong type given expect [][][]float64")
}

// Transform converts the current Polygon Coordinates into another Reference System and change the CRS Property to it.
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
