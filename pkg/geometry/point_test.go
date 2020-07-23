package geometry

import (
	"github.com/nodejayes/geolib/pkg/reference_system"
	"github.com/onsi/gomega"
	"testing"
)

func TestPoint_GetCoordinates(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	pt := NewPoint([]float64{1, 2}, reference_system.New(4326))
	var coords []float64
	pt.GetCoordinates(&coords)
	g.Expect(coords).To(gomega.Equal([]float64{1, 2}))
}
