package geometry

type Feature struct {
	Type       string                 `json:"type"`
	Geometry   IGeometry              `json:"geometry"`
	Properties map[string]interface{} `json:"properties"`
}

func NewFeature(geometry IGeometry, properties map[string]interface{}) *Feature {
	return &Feature{
		Type:       FeatureType,
		Geometry:   geometry,
		Properties: properties,
	}
}

func (ctx *Feature) Transform(target int) error {
	return ctx.Geometry.Transform(target)
}
