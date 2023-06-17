package main

import (
	"fmt"
	"strings"
)

func main() {
	var in string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&in)

	if ehPalindromo(in) {
		fmt.Println("A string é um palíndromo.")
	} else {
		fmt.Println("A string não é um palíndromo.")
	}
}

func ehPalindromo(s string) bool {
	s = strings.ReplaceAll(strings.ToLower(s), " ", "")

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}
