package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var string1, string2 string
	fmt.Print("Digite a primeira string: ")
	fmt.Scanln(&string1)
	fmt.Print("Digite a segunda string: ")
	fmt.Scanln(&string2)

	if saoAnagramas(string1, string2) {
		fmt.Println("As strings são anagramas.")
	} else {
		fmt.Println("As strings não são anagramas.")
	}
}

func saoAnagramas(string1, string2 string) bool {
	string1 = strings.ReplaceAll(strings.ToLower(string1), " ", "")
	string2 = strings.ReplaceAll(strings.ToLower(string2), " ", "")

	if len(string1) != len(string2) {
		return false
	}

	slice1 := strings.Split(string1, "")
	slice2 := strings.Split(string2, "")

	sort.Strings(slice1)
	sort.Strings(slice2)

	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}
