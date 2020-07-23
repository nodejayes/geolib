package vector

import (
	"github.com/onsi/gomega"
	"testing"
)

var A = []float64{0, 0}
var B = []float64{5, 1}
var C = []float64{3, 2}

func TestNewTriangle(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle(A, B, C)
	g.Expect(data).ToNot(gomega.BeNil())
}

func TestTriangle_Method_A(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle(A, B, C)
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.A()).To(gomega.Equal(A))
}

func TestTriangle_Method_B(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle(A, B, C)
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.B()).To(gomega.Equal(B))
}

func TestTriangle_Method_C(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle(A, B, C)
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.C()).To(gomega.Equal(C))
}

func TestTriangle_Method_EdgeAB(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle(A, B, C)
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.EdgeAB()).To(gomega.Equal(5.0990195135927845))
}

func TestTriangle_Method_EdgeAC(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle(A, B, C)
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.EdgeAC()).To(gomega.Equal(3.605551275463989))
}

func TestTriangle_Method_EdgeCB(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle(A, B, C)
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.EdgeCB()).To(gomega.Equal(2.23606797749979))
}

func TestTriangle_Method_AngleAlpha(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle(A, B, C)
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.AngleAlpha()).To(gomega.Equal(22.38013524215307))
}

func TestTriangle_Method_AngleBeta(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle(A, B, C)
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.AngleBeta()).To(gomega.Equal(37.874983972971805))
}

func TestTriangle_Method_AngleGamma(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle(A, B, C)
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.AngleGamma()).To(gomega.Equal(119.74488231457212))
}

func TestTriangle_AngleSumMustHave180Degrees(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle(A, B, C)
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.AngleGamma() + data.AngleBeta() + data.AngleAlpha()).To(gomega.Equal(180.000001529697))
}

func TestTriangle_Method_Diameter(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle(A, B, C)
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.Diameter()).To(gomega.Equal(14.308909704985036))
}

func TestTriangle_Method_Radius(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle(A, B, C)
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.Radius()).To(gomega.Equal(data.Diameter() / 2))
}

func TestTriangle_Scope(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle(A, B, C)
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.Scope()).To(gomega.Equal(10.940638766556564))
}

func TestTriangle_HeightFrom(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle(A, B, C)
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.HeightFrom("A")).To(gomega.Equal(3.130495169777232))
	g.Expect(data.HeightFrom("B")).To(gomega.Equal(1.9414506856750204))
	g.Expect(data.HeightFrom("C")).To(gomega.Equal(1.3728129465925845))
	g.Expect(data.HeightFrom("")).To(gomega.Equal(0.0))
}

func TestTriangle_ShortestDistance_Negative(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle([]float64{
		386130, 5692254,
	}, []float64{
		386408.967053963,
		5692913.35517462,
	}, []float64{
		386106.870525168,
		5692632.47844243,
	})
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.EdgeAB()).To(gomega.Equal(715.941243046222))
	g.Expect(data.EdgeAC()).To(gomega.Equal(379.18452498739))
	g.Expect(data.EdgeCB()).To(gomega.Equal(412.49733501657795))
	g.Expect(data.HeightFrom("A")).To(gomega.Equal(292.9317229151599))
	g.Expect(data.ShortestDistance("A")).To(gomega.Equal(379.18452498739))
}

func TestTriangle_ShortestDistance_Positive(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle([]float64{
		386130, 5692254,
	}, []float64{
		386106.870525168,
		5692632.47844243,
	}, []float64{
		386211.636432768,
		5692238.19060928,
	})
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.EdgeAB()).To(gomega.Equal(379.18452498739))
	g.Expect(data.EdgeAC()).To(gomega.Equal(83.15313577986396))
	g.Expect(data.EdgeCB()).To(gomega.Equal(407.96910515995535))
	g.Expect(data.HeightFrom("A")).To(gomega.Equal(74.83891952375997))
	g.Expect(data.ShortestDistance("A")).To(gomega.Equal(74.83891952375997))
}

func TestTriangle_ShortestDistance_Negative1(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle([]float64{
		386130, 5692254,
	}, []float64{
		386211.636432768,
		5692238.19060928,
	}, []float64{
		386259.794025127,
		5692052.3088027,
	})
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.EdgeAB()).To(gomega.Equal(83.15313577986396))
	g.Expect(data.EdgeAC()).To(gomega.Equal(239.84542527868626))
	g.Expect(data.EdgeCB()).To(gomega.Equal(192.01874835387844))
	g.Expect(data.HeightFrom("A")).To(gomega.Equal(75.06238598507127))
	g.Expect(data.ShortestDistance("A")).To(gomega.Equal(83.15313577986396))
}

func TestTriangle_ShortestDistance_Negative2(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle([]float64{
		386130, 5692254,
	}, []float64{
		386211.636432768,
		5692238.19060928,
	}, []float64{
		386259.794025127,
		5692052.3088027,
	})
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.EdgeAB()).To(gomega.Equal(83.15313577986396))
	g.Expect(data.EdgeAC()).To(gomega.Equal(239.84542527868626))
	g.Expect(data.EdgeCB()).To(gomega.Equal(192.01874835387844))
	g.Expect(data.HeightFrom("A")).To(gomega.Equal(75.06238598507127))
	g.Expect(data.ShortestDistance("A")).To(gomega.Equal(83.15313577986396))
}
