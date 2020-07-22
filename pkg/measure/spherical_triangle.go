package measure

import (
	"github.com/nodejayes/geolib/pkg/definitions"
	"math"
)

func SphericalTriangle(p1 []float64, p2 []float64) float64 {
	lat1 := p1[1] * definitions.DegToRad
	lon1 := p1[0] * definitions.DegToRad
	lat2 := p2[1] * definitions.DegToRad
	lon2 := p2[0] * definitions.DegToRad
	return (definitions.EarthRadius / 1000) * math.Acos(math.Sin(lat1)*math.Sin(lat2)+math.Cos(lat1)*math.Cos(lat2)*math.Cos(lon2-lon1))
}

func HaversineDistance(p1 []float64, p2 []float64) float64 {
	lat1 := p1[1]
	lon1 := p1[0]
	lat2 := p2[1]
	lon2 := p2[0]

	x1 := lat2 - lat1
	dLat := x1 * definitions.DegToRad
	x2 := lon2 - lon1
	dLon := x2 * definitions.DegToRad
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1*definitions.DegToRad)*math.Cos(lat2*definitions.DegToRad)*
			math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return (definitions.EarthRadius / 1000) * c
}

func Pythagoras(ax, ay, bx, by float64) float64 {
	return PythagorasLength(ax-bx, ay-by)
}

func PythagorasLength(l1, l2 float64) float64 {
	return math.Sqrt(math.Pow(l1, 2) + math.Pow(l2, 2))
}
