package geometry

type Feature struct {
	Type       string                 `json:"type"`
	Geometry   interface{}            `json:"geometry"`
	Properties map[string]interface{} `json:"properties"`
}

func NewFeature(geometry interface{}, properties map[string]interface{}) *Feature {
	return &Feature{
		Type:       FeatureType,
		Geometry:   geometry,
		Properties: properties,
	}
}
