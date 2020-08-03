package vector

// Triangle a Planar Area of 3 Vectors in 2D Space
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
