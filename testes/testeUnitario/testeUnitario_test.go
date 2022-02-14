package testeUnitario

import "testing"

// Por convenção o nome de um arquivo de teste é o nome do arquivo q será testado
// com o '_test' no final

const erroPadraoMedia = "Valor esperado %v, mas o resultado encontrado foi %v."

// Por convenção tbm, as funções de teste começão com 'Test' junto com o nome da função
// a ser testada
// A função de teste recebe uma estrutura auxiliar para ajudar nos teste. Um Helper
func TestMedia(t *testing.T) {
	// Define um valor esperado para o retorno da função
	valorEsperado := 7.28
	// Chama a função a ser testada
	valor := Media(7.2, 9.9, 6.1, 5.9)

	// Compara o valor retornado da função com o valor esperado
	if valor != valorEsperado {
		t.Errorf(erroPadraoMedia, valorEsperado, valor)
	}
}
