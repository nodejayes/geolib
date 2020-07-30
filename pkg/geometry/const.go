// this package handel's geometries from geojson
package geometry

const (
	// PointType is the GeoJSON Point Type String
	PointType = "Point"
	// MultiPointType is the GeoJSON MultiPoint Type String
	MultiPointType = "MultiPoint"
	// LineType is the GeoJSON LineString Type String
	LineType = "LineString"
	// MultiLineType is the GeoJSON MultiLineString Type String
	MultiLineType = "MultiLineString"
	// PolygonType is the GeoJSON Polygon Type String
	PolygonType = "Polygon"
	// MultiPolygonType is the GeoJSON MultiPolygon Type String
	MultiPolygonType = "MultiPolygon"
	// FeatureType is the GeoJSON Feature Type String
	FeatureType = "Feature"
	// FeatureCollectionType is the GeoJSON FeatureCollection Type String
	FeatureCollectionType = "FeatureCollection"
)

// Coordinate1D is a 1 dimensional Slice of float64 Values.
type Coordinate1D = []float64

// Coordinate2D is a 2 dimensional Slice of float64 Values.
type Coordinate2D = [][]float64

// Coordinate3D is a 3 dimensional Slice of float64 Values.
type Coordinate3D = [][][]float64

// Coordinate4D is a 4 dimensional Slice of float64 Values.
type Coordinate4D = [][][][]float64
