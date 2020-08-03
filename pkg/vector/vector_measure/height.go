package vector_measure

import (
	"errors"
	"github.com/nodejayes/geolib/pkg/definitions"
	"github.com/nodejayes/geolib/pkg/vector/vector_math"
	"math"
)

// HeightFrom get the Height from a Point to the opposite Edge
// the Height can be outside of the Triangle
func HeightFrom(t I3Facer, point PointKey) (float64, error) {
	switch point {
	case Alpha:
		lAB, err := EdgeLength(t, EdgeAB)
		aB, err := AngleFrom(t, Beta)
		if err != nil {
			return 0, err
		}
		return lAB * math.Sin(aB*definitions.DegToRad), nil
	case Beta:
		lAB, err := EdgeLength(t, EdgeCB)
		aB, err := AngleFrom(t, Gamma)
		if err != nil {
			return 0, err
		}
		return lAB * math.Sin(aB*definitions.DegToRad), nil
	case Gamma:
		lAB, err := EdgeLength(t, EdgeAC)
		aB, err := AngleFrom(t, Alpha)
		if err != nil {
			return 0, err
		}
		return lAB * math.Sin(aB*definitions.DegToRad), nil
	}
	return 0, errors.New("invalid Point Key " + point + " only support Alpha, Beta or Gamma")
}

// ShortestDistance get the shortest Distance to the opposite Edge that is not outside the Triangle
func ShortestDistance(t I3Facer, point string) (float64, error) {
	h, err := HeightFrom(t, point)
	lAB, err := EdgeLength(t, EdgeAB)
	lAC, err := EdgeLength(t, EdgeAC)
	lCB, err := EdgeLength(t, EdgeCB)
	if err != nil {
		return 0, err
	}
	switch point {
	case Alpha:
		return snapHeight(lAB, lAC, lCB, h), nil
	case Beta:
		return snapHeight(lAB, lCB, lCB, h), nil
	case Gamma:
		return snapHeight(lCB, lAC, lCB, h), nil
	}
	return 0, errors.New("invalid Point Key " + point + " only support Alpha, Beta or Gamma")
}

func snapHeight(l1, l2, l3, h float64) float64 {
	side1 := l3 - vector_math.PythagorasLength(l1, h)
	side2 := l3 - vector_math.PythagorasLength(l2, h)
	if side1 < 0 || side2 < 0 {
		if side1 < side2 {
			return l2
		} else {
			return l1
		}
	}
	return h
}
