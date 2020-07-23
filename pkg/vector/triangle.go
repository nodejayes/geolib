package vector

import (
	"github.com/nodejayes/geolib/pkg/definitions"
	"github.com/nodejayes/geolib/pkg/vector_math"
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
func (ctx *Triangle) A() []float64 {
	return ctx.a
}

// get Point B
func (ctx *Triangle) B() []float64 {
	return ctx.b
}

// get Point C
func (ctx *Triangle) C() []float64 {
	return ctx.c
}

// get the Edge from Point C to B
func (ctx *Triangle) EdgeCB() float64 {
	return vector_math.Pythagoras(ctx.b[0], ctx.b[1], ctx.c[0], ctx.c[1])
}

// get the Edge from Point A to C
func (ctx *Triangle) EdgeAC() float64 {
	return vector_math.Pythagoras(ctx.c[0], ctx.c[1], ctx.a[0], ctx.a[1])
}

// get the Edge from Point A to B
func (ctx *Triangle) EdgeAB() float64 {
	return vector_math.Pythagoras(ctx.b[0], ctx.b[1], ctx.a[0], ctx.a[1])
}

// get the Angle opposite of Edge from Point C to B
func (ctx *Triangle) AngleAlpha() float64 {
	return vector_math.CosineTheorem(ctx.EdgeCB(), ctx.EdgeAC(), ctx.EdgeAB(), true)
}

// get the Angle opposite of Edge from Point A to C
func (ctx *Triangle) AngleBeta() float64 {
	return vector_math.CosineTheorem(ctx.EdgeAC(), ctx.EdgeAB(), ctx.EdgeCB(), true)
}

// get the Angle opposite of Edge from Point A to B
func (ctx *Triangle) AngleGamma() float64 {
	return vector_math.CosineTheorem(ctx.EdgeAB(), ctx.EdgeAC(), ctx.EdgeCB(), true)
}

// get the Diameter of the Circle that intersects all 3 Points
func (ctx *Triangle) Diameter() float64 {
	return ctx.EdgeAB() / math.Sin(ctx.AngleGamma())
}

// get the Radius of the Circle that intersects all 3 Points
func (ctx *Triangle) Radius() float64 {
	return ctx.Diameter() / 2
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
