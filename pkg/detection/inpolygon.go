package detection

import (
	"github.com/nodejayes/geolib/pkg/geometry"
)

// BooleanPointInPolygon Takes a {@link Point} and a {@link Polygon} and determines if the point
// resides inside the polygon. The polygon can be convex or concave. The function accounts for holes.
func PointInPolygon(point *geometry.Point, polygon *geometry.Polygon, ignoreBoundary bool) bool {
	bbox := getBoundingBox(polygon)

	// Quick elimination if point is not inside bbox
	if InBBox(point, bbox) == false {
		return false
	}

	// normalize to multipolygon
	polys := [][][][]float64{polygon.Coordinates}
	insidePoly := false
	for i := 0; i < len(polys) && !insidePoly; i++ {
		// check if it is in the outer ring first
		if InRing(point, geometry.NewLine(polys[i][0], point.CRS), ignoreBoundary) {
			inHole := false
			k := 1
			// check for the point in any of the holes
			for k < len(polys[i]) && !inHole {
				if InRing(point, geometry.NewLine(polys[i][k], point.CRS), !ignoreBoundary) {
					inHole = true
				}
				k++
			}
			if !inHole {
				insidePoly = true
			}
		}
	}
	return insidePoly
}

func InRing(ptGeom *geometry.Point, ringGeom *geometry.Line, ignoreBoundary bool) bool {
	pt := ptGeom.Coordinates
	ring := ringGeom.Coordinates
	ringLen := len(ring)
	isInside := false
	if ring[0][0] == ring[ringLen-1][0] && ring[0][1] == ring[ringLen-1][1] {
		ring = ring[0 : ringLen-1]
		ringLen = len(ring)
	}
	i := 0
	j := ringLen - 1

	for i < ringLen {
		xi := ring[i][0]
		yi := ring[i][1]
		xj := ring[j][0]
		yj := ring[j][1]
		onBoundary := (pt[1]*(xi-xj)+yi*(xj-pt[0])+yj*(pt[0]-xi) == 0) &&
			((xi-pt[0])*(xj-pt[0]) <= 0) && ((yi-pt[1])*(yj-pt[1]) <= 0)
		if onBoundary {
			return !ignoreBoundary
		}
		intersect := ((yi > pt[1]) != (yj > pt[1])) &&
			(pt[0] < (xj-xi)*(pt[1]-yi)/(yj-yi)+xi)
		if intersect {
			isInside = !isInside
		}
		j = i
		i++
	}
	return isInside
}

// InBBox check if a Point is in a Bounding Box
func InBBox(pt *geometry.Point, bbox *geometry.BoundingBox) bool {
	return bbox.MinX <= pt.Coordinates[0] &&
		bbox.MaxX <= pt.Coordinates[1] &&
		bbox.MinY >= pt.Coordinates[0] &&
		bbox.MaxY >= pt.Coordinates[1]
}

func getBoundingBox(polygonGeom *geometry.Polygon) *geometry.BoundingBox {
	polygon := polygonGeom.Coordinates
	gn := len(polygon)
	if gn == 0 {
		panic("invalid Polygon")
	}

	box := []float64{polygon[0][0][0], polygon[0][0][0], polygon[0][0][1], polygon[0][0][1]}
	for i := 0; i < gn; i++ {
		// Polygons
		for j := 0; j < len(polygon[i]); j++ {
			// Vertices
			if polygon[i][j][0] < box[0] {
				box[0] = polygon[i][j][0]
			}

			if polygon[i][j][0] > box[1] {
				box[1] = polygon[i][j][0]
			}

			if polygon[i][j][1] < box[2] {
				box[2] = polygon[i][j][1]
			}

			if polygon[i][j][1] > box[3] {
				box[3] = polygon[i][j][1]
			}
		}
	}
	return geometry.NewBoundingBox(box[0], box[1], box[2], box[3], polygonGeom.CRS)
}
