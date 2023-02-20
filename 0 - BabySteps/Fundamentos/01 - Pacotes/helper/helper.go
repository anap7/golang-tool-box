package helper

import "fmt"

/*GO não é uma linguagem orientada a objetos, então não existe private, public e etc. O que indica se a função é pública e privada é se a função está coma inicial em caixa alta ou baixa
 - Caixa alta (Uppercase) = Public
 - Caixa baixa (Lowercase) = Private*/
func WriteFromFilePointUppercase() {
	fmt.Println("Writing from helper file, with uppercase")
	//Como o arquivo helper-two.go está dentro do mesmo pacote, não é necessário importar
	writeFromFilePointLowercase()
}