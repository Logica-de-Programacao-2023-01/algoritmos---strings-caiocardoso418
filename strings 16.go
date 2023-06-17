package main

import (
	"fmt"
	"strings"
)

func isSubstring(string1, string2 string) bool {
	return strings.Contains(string1, string2)
}

func main() {
	var string1, string2 string

	fmt.Print("Digite a primeira string: ")
	fmt.Scanln(&string1)

	fmt.Print("Digite a segunda string: ")
	fmt.Scanln(&string2)

	if isSubstring(string1, string2) {
		fmt.Println("A segunda string é uma substring da primeira.")
	} else {
		fmt.Println("A segunda string não é uma substring da primeira.")
	}
}
