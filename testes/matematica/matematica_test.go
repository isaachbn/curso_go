package matematica

import "testing"

const error = "Valor esperado %v, mas o resultado encontrato foi %v."

func TestMedia(t *testing.T) {
	t.Parallel()
	valorEsperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)

	if valor != valorEsperado {
		t.Error(error, valorEsperado, valor)
	}
}
