package main

import (
	"fmt"
)

func daysOfWeek(number int) string {
	switch number {
	case 1: 
		return "Domingo"
	case 2: 
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sábado"
	default:
		return "Invalid Day"
	}
}

func main() {
	day := daysOfWeek(4);
	fmt.Println(day)

	day = daysOfWeek(400);
	fmt.Println(day)
}