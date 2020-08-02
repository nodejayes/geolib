package vector_measure

import (
	"errors"
	"fmt"
	"github.com/nodejayes/geolib/pkg/vector/vector_math"
)

type EdgeKey = string

const (
	EdgeCB EdgeKey = "CB"
	EdgeAB EdgeKey = "AB"
	EdgeAC EdgeKey = "AC"
)

type I3Facer interface {
	GetA() []float64
	GetB() []float64
	GetC() []float64
}

// EdgeLength calculates the length of a Edge from a planar Triangle
func EdgeLength(t I3Facer, edge EdgeKey) (float64, error) {
	switch edge {
	case EdgeCB:
		return vector_math.Pythagoras(t.GetB()[0], t.GetB()[1], t.GetC()[0], t.GetC()[1]), nil
	case EdgeAC:
		return vector_math.Pythagoras(t.GetC()[0], t.GetC()[1], t.GetA()[0], t.GetA()[1]), nil
	case EdgeAB:
		return vector_math.Pythagoras(t.GetB()[0], t.GetB()[1], t.GetA()[0], t.GetA()[1]), nil
	}
	return 0, errors.New(fmt.Sprintf("invalid edge %v argument only support EdgeCB, EdgeAB or EdgeAC", edge))
}
