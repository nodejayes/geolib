package geometry

import (
	"github.com/nodejayes/geolib/pkg/reference_system"
	"github.com/onsi/gomega"
	"testing"
)

func TestMultiPoint_GetCoordinates(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	coords := Coordinate2D{{1, 2}, {3, 4}}
	mp := NewMultiPoint(coords, reference_system.New(4326))
	var c Coordinate2D
	mp.GetCoordinates(&c)
	g.Expect(c).To(gomega.Equal(coords))
}
