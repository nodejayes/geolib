package vector_math

import (
	"github.com/onsi/gomega"
	"testing"
)

func TestPythagoras(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(Pythagoras(1, 2, 3, 4)).To(gomega.Equal(2.8284271247461903))
}
