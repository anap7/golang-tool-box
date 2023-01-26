// Unity test

/*	Comandos de teste
	go test ./... = Busca todos os arquivos dentro do projeto que possuam o Test para rodar
	go test -v = Faz o teste com mais detalhes
	go test --cover =  Pegar toda cobertura de test
	go test --coverprofile coverage.txt = Cria um arquivo dos resultados dos testes
	
	Os comandos abaixo devem ter um arquivo criado antes de executar

	go tool cover --func=coverage.txt = Registra no arquivo de texto do coverage e qual parte do código deve ser testada ainda 
	go tool cover --html=coverage.txt = Registra um arquivo HTML do coverage e qual parte e linhas do código devem ser testada ainda
*/
package address_test

import (
	. "intro-test/address"
	"testing"
)

type testScenario struct {
	addressTypeReceived string
	addressTypeExpected string
}
// Obrigatoriamente começa com TestNomeDoTest
func TestAddressType(t *testing.T) {
	//Para rodar os testes em paralelo
	t.Parallel()

	testScenario := []testScenario {
		{ "Rua ABC", "Rua" },
		{ "Avenida Show", "Avenida" },
		{ "Estrada da Felicidade", "Estrada" },
		{ "Rodovia Show", "Rodovia" },
		{ "rua ABC", "Rua" },
		{ "ESTRADA COSTEIRA", "Estrada" },
		{ "estrada da liberdade", "Estrada" },
		//{ "Praça das bandeiras", "Invalid" },
		//{ "", "Invalid" },
	}
	
		for _, scenario := range testScenario {
			receivedAddressType := AddressType(scenario.addressTypeReceived)
			if receivedAddressType != scenario.addressTypeExpected {
				t.Errorf("The received type is different than expected! It expect %s and received %s", 
				scenario.addressTypeExpected,
				receivedAddressType, 
			)
		}
	}
}

func TestAny(t *testing.T) {
	//Para rodar os testes em paralelo
	t.Parallel()

	if 1 > 2 {
		t.Errorf("The test broken!")
	}
}