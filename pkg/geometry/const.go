package geometry

const (
	PointType             = "Point"
	MultiPointType        = "MultiPoint"
	LineType              = "LineString"
	MultiLineType         = "MultiLineString"
	PolygonType           = "Polygon"
	MultiPolygonType      = "MultiPolygon"
	FeatureType           = "Feature"
	FeatureCollectionType = "FeatureCollection"
)

type Coordinate1D = []float64
type Coordinate2D = [][]float64
type Coordinate3D = [][][]float64
type Coordinate4D = [][][][]float64
