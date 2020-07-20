package geometry

type FeatureCollection struct {
	Type     string     `json:"type"`
	Features []*Feature `json:"features"`
}

func NewFeatureCollection(features []*Feature) *FeatureCollection {
	return &FeatureCollection{
		Type:     FeatureCollectionType,
		Features: features,
	}
}
