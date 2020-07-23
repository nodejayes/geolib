package measure

import (
	"github.com/nodejayes/geolib/pkg/geometry"
	"github.com/nodejayes/geolib/pkg/reference_system"
	"github.com/onsi/gomega"
	"testing"
)

var shortP1 = []float64{8.413210, 49.99170}
var shortP2 = []float64{8.421820, 50.00490}
var longP1 = []float64{13.37770, 52.51640}
var longP2 = []float64{-9.177944, 38.69267}

func TestSphericalTriangle(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	// near QGIS Measurement (1.592723 km) maybe QGIS uses the Bessel Ellipsoid
	g.Expect(SphericalTriangle(shortP1, shortP2)).To(gomega.Equal(1.5933197255112694))
	g.Expect(HaversineDistance(shortP1, shortP2)).To(gomega.Equal(1.5933197236818772))
	shortP1Geom := geometry.NewPoint(shortP1, reference_system.New(4326))
	shortP2Geom := geometry.NewPoint(shortP2, reference_system.New(4326))
	err := shortP1Geom.Transform(3857)
	g.Expect(err).To(gomega.BeNil())
	err = shortP2Geom.Transform(3857)
	g.Expect(err).To(gomega.BeNil())
	g.Expect(
		Planar(shortP1Geom.Coordinates, shortP2Geom.Coordinates)).
		To(gomega.Equal(2478.731094304433))
	// near QGIS Measurement (2318.216871 km)
	g.Expect(SphericalTriangle(longP1, longP2)).To(gomega.Equal(2317.581202880789))
	g.Expect(HaversineDistance(longP1, longP2)).To(gomega.Equal(2317.5812028807877))
}
