package detection

// BooleanPointInPolygon Takes a {@link Point} and a {@link Polygon} and determines if the point
// resides inside the polygon. The polygon can be convex or concave. The function accounts for holes.
func PointInPolygon(point []float64, polygon [][][]float64, ignoreBoundary bool) bool {
	bbox := getBoundingBox(polygon)

	// Quick elimination if point is not inside bbox
	if InBBox(point, bbox) == false {
		return false
	}

	// normalize to multipolygon
	polys := [][][][]float64{polygon}
	insidePoly := false
	for i := 0; i < len(polys) && !insidePoly; i++ {
		// check if it is in the outer ring first
		if InRing(point, polys[i][0], ignoreBoundary) {
			inHole := false
			k := 1
			// check for the point in any of the holes
			for k < len(polys[i]) && !inHole {
				if InRing(point, polys[i][k], !ignoreBoundary) {
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

func InRing(pt []float64, ring [][]float64, ignoreBoundary bool) bool {
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
func InBBox(pt []float64, bbox []float64) bool {
	return bbox[0] <= pt[0] &&
		bbox[1] <= pt[1] &&
		bbox[2] >= pt[0] &&
		bbox[3] >= pt[1]
}

func getBoundingBox(polygons [][][]float64) []float64 {
	gn := len(polygons)
	if gn == 0 {
		panic("invalid Polygon")
	}

	box := []float64{polygons[0][0][0], polygons[0][0][0], polygons[0][0][1], polygons[0][0][1]}
	for i := 0; i < gn; i++ {
		// Polygons
		for j := 0; j < len(polygons[i]); j++ {
			// Vertices
			if polygons[i][j][0] < box[0] {
				box[0] = polygons[i][j][0]
			}

			if polygons[i][j][0] > box[1] {
				box[1] = polygons[i][j][0]
			}

			if polygons[i][j][1] < box[2] {
				box[2] = polygons[i][j][1]
			}

			if polygons[i][j][1] > box[3] {
				box[3] = polygons[i][j][1]
			}
		}
	}
	return box
}
