package proj4

import (
	"github.com/onsi/gomega"
	"testing"
)

func TestReprojectPoint(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	WGS84 := "EPSG:4326"
	WebMercator := "EPSG:3857"
	x, y, z, err := ReprojectPoint(12, 50, 0, WGS84, WebMercator)
	g.Expect(err).To(gomega.BeNil())
	g.Expect(x).To(gomega.Equal(1335833.8895192828))
	g.Expect(y).To(gomega.Equal(6446275.841017158))
	g.Expect(z).To(gomega.Equal(0.0))
}

func TestReproject3dPoint(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	WGS84 := "EPSG:4326"
	WebMercator := "EPSG:3857"
	x, y, z, err := ReprojectPoint(12, 50, 5, WGS84, WebMercator)
	g.Expect(err).To(gomega.BeNil())
	g.Expect(x).To(gomega.Equal(1335833.8895192828))
	g.Expect(y).To(gomega.Equal(6446275.841017158))
	g.Expect(z).To(gomega.Equal(5.0))
}
