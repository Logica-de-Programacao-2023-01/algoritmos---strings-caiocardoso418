package main

import (
	"fmt"
)

func containsOnlyNumbers(in string) bool {
	for _, char := range in {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

func main() {
	var in string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&in)

	if containsOnlyNumbers(in) {
		fmt.Println("A string tem só números.")
	} else {
		fmt.Println("A string não tem só números.")
	}
}
