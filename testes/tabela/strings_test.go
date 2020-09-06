package strings

import (
	"strings"
	"testing"
)

const error = "%s (parte: %s) - índices: esperado (%d) <> encontrado (%d)"

func TestIndex(t *testing.T) {
	t.Parallel()
	testes := []struct {
		texto    string
		parte    string
		esperado int
	}{
		{"Isaac Henrique", "Isaac", 0},
		{"", "", 0},
		{"Opa", "opa", -1},
		{"Isaac Henrique", "a", 2},
	}

	for _, teste := range testes {
		t.Logf("Massa de teste %v", teste)
		atual := strings.Index(teste.texto, teste.parte)

		if atual != teste.esperado {
			t.Errorf(error, teste.texto, teste.parte, teste.esperado, atual)
		}
	}
}
