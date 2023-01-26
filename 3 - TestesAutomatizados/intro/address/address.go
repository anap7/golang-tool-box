package address

import "strings"

//AddressType: verifica se o endereço é válido
func AddressType(address string) string{
	validTypes := []string{"rua", "avenida", "estrada", "rodovia"}

	//Convertendo a string para caixa baixa
	addressInLowerCase := strings.ToLower(address)

	//Divindo a string através do espaço e pegar a primeira posição
	firstAdressWord := strings.Split(addressInLowerCase, " ")[0]

	addressIsValid := false
	
	for _, addressType := range validTypes {
		if addressType == firstAdressWord {
			addressIsValid = true
		}
	}

	if addressIsValid {
		//Deixando a primeira letra maiúscula
		return strings.Title(firstAdressWord)
	}

	return "Invalid"
}