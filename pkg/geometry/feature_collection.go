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

func (ctx *FeatureCollection) Transform(target int) error {
	for _, f := range ctx.Features {
		err := f.Geometry.Transform(target)
		if err != nil {
			return err
		}
	}
	return nil
}
