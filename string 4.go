package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	var str2 string
	fmt.Println("Digite uma palavra: ")
	fmt.Scan(&str)
	fmt.Println("Digite outra palavra: ")
	fmt.Scan(&str2)

	x := strings.Contains(str, str2)
	if x == true {
		fmt.Println("As palavras são iguais.")
	} else {
		fmt.Println("As palavras são diferentes")
	}
}
