package vector_math

import (
	"github.com/nodejayes/geolib/pkg/definitions"
	"math"
)

// the Cosine Theorem for Angle Calculation in a Triangle
// edge1 must be the opposite of the Angle to Calculate and we need all 3 Edges of the Triangle
func CosineTheorem(edge1, edge2, edge3 float64, inDegrees bool) float64 {
	a2 := math.Pow(edge1, 2)
	b2 := math.Pow(edge2, 2)
	c2 := math.Pow(edge3, 2)
	bc2 := -(2 * edge2 * edge3)
	cosAlpha := (a2 - (b2 + c2)) / bc2
	res := math.Acos(cosAlpha)
	if inDegrees {
		return res * definitions.RadToDeg
	}
	return res
}
