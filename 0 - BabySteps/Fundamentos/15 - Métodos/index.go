package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

//O método que espera como um usuário com o nome chamado salvar
/*O %s é o valor que vai ser substituido pelo u.nome*/
func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Usuário 1", 20}
	/*O salvar vai receber os dados do usuário sem precisar 
	passar como parâmetro na função salvar*/
	usuario1.salvar()

	usuario2 := usuario{"Davi", 30}
	/*O salvar vai receber os dados do usuário sem precisar 
	passar como parâmetro na função salvar*/
	usuario2.salvar()

	maiorDeIdade := usuario2.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}