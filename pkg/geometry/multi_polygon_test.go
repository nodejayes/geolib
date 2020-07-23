package geometry

import (
	"github.com/nodejayes/geolib/pkg/reference_system"
	"github.com/onsi/gomega"
	"testing"
)

func TestMultiPolygon_GetCoordinates(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	coords := Coordinate4D{{{{1, 2}, {3, 4}, {5, 6}, {1, 2}}}}
	mp := NewMultiPolygon(coords, reference_system.New(4326))
	var c Coordinate4D
	mp.GetCoordinates(&c)
	g.Expect(c).To(gomega.Equal(coords))
}
