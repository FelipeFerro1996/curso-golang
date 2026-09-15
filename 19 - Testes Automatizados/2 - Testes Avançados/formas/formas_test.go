package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {

	t.Run("Retângulo", func(t *testing.T) {
		ret := Retangulo{12, 10}
		areaEsperada := float64(120)
		areaRetornada := ret.Area()
		if areaEsperada != areaRetornada {
			t.Fatalf("A area retornada %f, é diferente da area esperada %f", areaRetornada, areaEsperada)
		}
	})

	t.Run("Circulo", func(t *testing.T) {

		circ := Circulo{10}
		areaEsperada := (math.Pi * 100)
		areaRetornada := circ.Area()
		if areaEsperada != areaRetornada {
			t.Fatalf("A area retornada %f, é diferente da area esperada %f", areaRetornada, areaEsperada)
		}

	})

}
