package vector_measure

import (
	"errors"
	"fmt"
	"github.com/nodejayes/geolib/pkg/vector/vector_math"
)

type PointKey = string

const (
	Alpha PointKey = "A"
	Beta  PointKey = "B"
	Gamma PointKey = "C"
)

func AngleFrom(t I3Facer, point PointKey) (float64, error) {
	lAB, err := EdgeLength(t, EdgeAB)
	if err != nil {
		return 0, err
	}
	lAC, err := EdgeLength(t, EdgeAC)
	if err != nil {
		return 0, err
	}
	lCB, err := EdgeLength(t, EdgeCB)
	if err != nil {
		return 0, err
	}
	switch point {
	case Alpha:
		return vector_math.CosineTheorem(lCB, lAC, lAB, true), nil
	case Beta:
		return vector_math.CosineTheorem(lAC, lAB, lCB, true), nil
	case Gamma:
		return vector_math.CosineTheorem(lAB, lAC, lCB, true), nil
	}
	return 0, errors.New(fmt.Sprintf("invalid point key %v only support Alpha, Beta or Gamma", point))
}
