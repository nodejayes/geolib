package vector_math

import (
	"github.com/onsi/gomega"
	"testing"
)

func TestCosineTheorem(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(CosineTheorem(5.0990195135927845, 3.605551275463989, 2.23606797749979, true)).To(gomega.Equal(119.7448812969422))
	g.Expect(CosineTheorem(5.0990195135927845, 3.605551275463989, 2.23606797749979, false)).To(gomega.Equal(2.089942441041419))
}
