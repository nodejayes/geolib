package geometry

import "errors"

// ITransformer transform into another Coordinate System
type ITransformer interface {
	Transform(target int) error
}

// IGeometrier returns Coordinates and implement ITransformer
type IGeometrier interface {
	ITransformer
	GetCoordinates(data interface{}) error
}

// Feature represent any Geometry of Type Point, Line, Polygon, MultiPoint, MultiLine and MultiPolygon with Properties
type Feature struct {
	Type       string                 `json:"type"`
	Geometry   interface{}            `json:"geometry"`
	Properties map[string]interface{} `json:"properties"`
}

// NewFeature create a new Feature from a Geometry Type and some Properties
func NewFeature(geometry IGeometrier, properties map[string]interface{}) *Feature {
	return &Feature{
		Type:       FeatureType,
		Geometry:   geometry,
		Properties: properties,
	}
}

// Transform convert the Geometry of the Feature into the given SrId of a Coordinate System.
func (ctx *Feature) Transform(target int) error {
	switch g := ctx.Geometry.(type) {
	case Point:
	case Line:
	case Polygon:
	case MultiPoint:
	case MultiLine:
	case MultiPolygon:
	case *Point:
	case *Line:
	case *Polygon:
	case *MultiPoint:
	case *MultiLine:
	case *MultiPolygon:
		return g.Transform(target)
	}
	return errors.New("not supported Geometry")
}
