package vector_math

import "math"

// Sentence of Pythagoras for right-angled triangles
// used by Points
func Pythagoras(ax, ay, bx, by float64) float64 {
	return PythagorasLength(ax-bx, ay-by)
}

// Sentence of Pythagoras for right-angled triangles
// used by Edge Lengths
func PythagorasLength(l1, l2 float64) float64 {
	return math.Sqrt(math.Pow(l1, 2) + math.Pow(l2, 2))
}
