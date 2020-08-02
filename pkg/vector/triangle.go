package vector

import (
	"github.com/nodejayes/geolib/pkg/definitions"
	"github.com/nodejayes/geolib/pkg/vector/vector_math"
	"math"
)

// a Planar Triangle in 2D Space
type Triangle struct {
	a []float64
	b []float64
	c []float64
}

func NewTriangle(a []float64, b []float64, c []float64) *Triangle {
	return &Triangle{a, b, c}
}

// get Point A
func (ctx *Triangle) GetA() []float64 {
	return ctx.a
}

// get Point B
func (ctx *Triangle) GetB() []float64 {
	return ctx.b
}

// get Point C
func (ctx *Triangle) GetC() []float64 {
	return ctx.c
}

// get the Length of all Edges
func (ctx *Triangle) Scope() float64 {
	return ctx.EdgeAB() + ctx.EdgeAC() + ctx.EdgeCB()
}

// get the Height from a Point to the opposite Edge
// the Height can be outside of the Triangle
func (ctx *Triangle) HeightFrom(point string) float64 {
	switch point {
	case "A":
		return ctx.EdgeAB() * math.Sin(ctx.AngleBeta()*definitions.DegToRad)
	case "B":
		return ctx.EdgeCB() * math.Sin(ctx.AngleGamma()*definitions.DegToRad)
	case "C":
		return ctx.EdgeAC() * math.Sin(ctx.AngleAlpha()*definitions.DegToRad)
	default:
		return 0.0
	}
}

// get the shortest Distance to the opposite Edge that is not outside the Triangle
func (ctx *Triangle) ShortestDistance(point string) float64 {
	h := ctx.HeightFrom(point)
	switch point {
	case "A":
		return snapHeight(ctx.EdgeAB(), ctx.EdgeAC(), ctx.EdgeCB(), h)
	case "B":
		return snapHeight(ctx.EdgeAB(), ctx.EdgeCB(), ctx.EdgeCB(), h)
	case "C":
		return snapHeight(ctx.EdgeCB(), ctx.EdgeAC(), ctx.EdgeCB(), h)
	default:
		return 0.0
	}
}

func (ctx *Triangle) Area() float64 {
	return 0.0
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
