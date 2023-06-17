package main

import (
	"fmt"
	"strings"
)

func replaceLetter(input string, oldLetter, newLetter string) string {
	return strings.ReplaceAll(input, oldLetter, newLetter)
}

func main() {
	var in, velha, nova string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&in)

	fmt.Print("Digite a letra a ser substituída: ")
	fmt.Scanln(&velha)

	fmt.Print("Digite a nova letra: ")
	fmt.Scanln(&nova)

	newString := replaceLetter(in, velha, nova)
	fmt.Println("Nova string:", newString)
}
