package reference_system

import (
	"github.com/onsi/gomega"
	"testing"
)

func GetWgs84() *ReferenceSystem {
	return &ReferenceSystem{
		Type: systemType,
		Properties: ReferenceSystemProperties{
			Name: systemHeader + "4326",
		},
	}
}

func GetPseudoMercator() *ReferenceSystem {
	return &ReferenceSystem{
		Type: systemType,
		Properties: ReferenceSystemProperties{
			Name: systemHeader + "3857",
		},
	}
}

func TestReferenceSystem_TransformPoints(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	points := [][]float64{
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
		{13.6, 59.6},
	}
	transformed, err := GetWgs84().TransformPoints(GetPseudoMercator().GetSrId(), points)
	g.Expect(err).To(gomega.BeNil())
	g.Expect(transformed).ToNot(gomega.BeNil())
	g.Expect(points[0]).To(gomega.Equal([]float64{13.6, 59.6}))
	g.Expect(transformed[0]).To(gomega.Equal([]float64{1513945.0747885206, 8311215.713002436}))
}
