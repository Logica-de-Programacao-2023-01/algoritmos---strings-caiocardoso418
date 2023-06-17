package main

import (
	"fmt"
)

func numeros(input string) bool {
	for _, char := range input {
		if char >= '0' && char <= '9' {
			return true
		}
	}
	return false
}

func main() {
	var in string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&in)

	if numeros(in) {
		fmt.Println("A string um número.")
	} else {
		fmt.Println("não contém números.")
	}
}
