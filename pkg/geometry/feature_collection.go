package geometry

// FeatureCollection represent many Features.
type FeatureCollection struct {
	Type     string     `json:"type"`
	Features []*Feature `json:"features"`
}

// NewFeatureCollection create a Collection of Features from a Slice of Feature Pointers.
func NewFeatureCollection(features []*Feature) *FeatureCollection {
	return &FeatureCollection{
		Type:     FeatureCollectionType,
		Features: features,
	}
}

// Transform convert each Geometry of Feature in the FeatureCollection.
func (ctx *FeatureCollection) Transform(target int) error {
	for _, f := range ctx.Features {
		err := f.Transform(target)
		if err != nil {
			return err
		}
	}
	return nil
}
