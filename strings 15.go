package main

import (
	"fmt"
	"strings"
)

func replaceVowels(input string) string {
	vowels := "aeiouAEIOU"
	for _, v := range vowels {
		input = strings.ReplaceAll(input, string(v), "*")
	}
	return input
}

func main() {
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	newString := replaceVowels(input)
	fmt.Println("Nova string:", newString)
}
