package main

import (
	"fmt"
	"strings"
)

func isCamelCase(in string) bool {
	return in != "" && in == strings.ToLower(in)
}

func countWords(in string) int {
	return strings.Count(in, "") - 1
}

func main() {
	var in string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&in)

	if isCamelCase(in) {
		fmt.Println("A string está em camelCase.")
	} else {
		fmt.Println("A string não está em camelCase.")
	}

	words := countWords(in)
	fmt.Printf("A string possui %d palavra(s).\n", words)
}
