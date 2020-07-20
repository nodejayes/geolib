package geometry

import (
	"github.com/nodejayes/geolib/pkg/reference_system"
	"math"
)

type BoundingBox struct {
	MinX float64                          `json:"minx"`
	MaxX float64                          `json:"maxx"`
	MinY float64                          `json:"miny"`
	MaxY float64                          `json:"maxy"`
	Crs  reference_system.ReferenceSystem `json:"crs"`
}

func NewBoundingBox(minX, maxX, minY, maxY float64, crs reference_system.ReferenceSystem) *BoundingBox {
	return &BoundingBox{
		MinX: minX,
		MaxX: maxX,
		MinY: minY,
		MaxY: maxY,
		Crs:  crs,
	}
}

func (bb *BoundingBox) GetCoordinates() *Coordinate3D {
	return &Coordinate3D{
		{
			{bb.MinX, bb.MinY},
			{bb.MinX, bb.MaxY},
			{bb.MaxX, bb.MaxY},
			{bb.MaxX, bb.MinY},
			{bb.MinX, bb.MinY},
		},
	}
}

func (bb *BoundingBox) Add(boundingBox *BoundingBox) {
	if boundingBox.MinX < bb.MinX {
		bb.MinX = boundingBox.MinX
	}
	if boundingBox.MaxX > bb.MaxX {
		bb.MaxX = boundingBox.MaxX
	}
	if boundingBox.MinY < bb.MinY {
		bb.MinY = boundingBox.MinY
	}
	if boundingBox.MaxY > bb.MaxY {
		bb.MaxY = boundingBox.MaxY
	}
}

func (bb *BoundingBox) SnapToGrid(width int) {
	bb.MinX = math.Floor((math.Floor(bb.MinX)-1)/float64(width)) * float64(width)
	bb.MinY = math.Floor((math.Floor(bb.MinY)-1)/float64(width)) * float64(width)
	bb.MaxX = (math.Floor((math.Floor(bb.MaxX)+1)/float64(width)) * float64(width)) + float64(width)
	bb.MaxY = (math.Floor((math.Floor(bb.MaxY)+1)/float64(width)) * float64(width)) + float64(width)
}
