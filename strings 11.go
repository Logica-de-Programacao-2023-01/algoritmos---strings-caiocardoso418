package main

import (
	"fmt"
	"strings"
)

func main() {
	var in string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&in)

	resultado := removerVogais(in)

	fmt.Println("Resultado: ", resultado)
}

func removerVogais(s string) string {
	vogais := "aeiouAEIOU"

	semVogais := strings.Map(func(r rune) rune {
		if strings.ContainsRune(vogais, r) {
			return -1
		}
		return r
	}, s)

	return semVogais
}
