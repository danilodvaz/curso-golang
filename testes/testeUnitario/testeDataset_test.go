package testeUnitario

import (
	"strings"
	"testing"
)

const erroPadraoDataset = "Erro"

func TestIndex(t *testing.T) {
	// Cria vários casos de teste
	testes := []struct {
		texto    string
		parte    string
		esperado int
	}{
		{"Cod3r é show", "Cod3r", 0},
		{"", "", 0},
		{"Opa", "opa", -1},
		{"Cod3r é show", "3", 3},
	}

	for _, teste := range testes {
		// Quando o teste é executado no modo verboso, tudo do log é exibido na saída
		// go test -v
		t.Logf("Massa: %v", teste)
		atual := strings.Index(teste.texto, teste.parte)

		if atual != teste.esperado {
			t.Errorf(erroPadraoDataset)
		}
	}

}
