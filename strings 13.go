package main

import (
	"fmt"
)

func main() {
	var in string
	fmt.Print("Digite uma string numérica: ")
	fmt.Scanln(&in)

	if ehSequenciaCrescente(in) {
		fmt.Println("A string é uma sequência numérica crescente.")
	} else {
		fmt.Println("A string não é uma sequência numérica crescente.")
	}
}

func ehSequenciaCrescente(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] <= s[i-1] {
			return false
		}
	}

	return true
}
