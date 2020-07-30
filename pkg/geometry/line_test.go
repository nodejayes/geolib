package geometry

import (
	"github.com/onsi/gomega"
	"testing"
)

func TestLine_GetCoordinates(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	coords := Coordinate2D{{1, 2}, {3, 4}}
	mp := NewLine(coords, 4326)
	var c Coordinate2D
	mp.GetCoordinates(&c)
	g.Expect(c).To(gomega.Equal(coords))
}
