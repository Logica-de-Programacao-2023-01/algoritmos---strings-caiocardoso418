package main

import (
	"fmt"
)

func findUniqueLetters(input string) string {
	letterCount := make(map[rune]int)
	for _, char := range input {
		letterCount[char]++
	}

	uniqueLetters := ""
	for _, char := range input {
		if letterCount[char] == 1 {
			uniqueLetters += string(char)
		}
	}

	return uniqueLetters
}

func main() {
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	unique := findUniqueLetters(input)
	fmt.Println("Letras únicas:", unique)
}
