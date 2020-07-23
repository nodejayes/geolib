package geometry

import (
	"github.com/nodejayes/geolib/pkg/reference_system"
	"github.com/onsi/gomega"
	"testing"
)

func TestPolygon_GetCoordinates(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	coords := Coordinate3D{{{1, 2}, {3, 4}, {5, 6}, {1, 2}}}
	mp := NewPolygon(coords, reference_system.New(4326))
	var c Coordinate3D
	mp.GetCoordinates(&c)
	g.Expect(c).To(gomega.Equal(coords))
}
