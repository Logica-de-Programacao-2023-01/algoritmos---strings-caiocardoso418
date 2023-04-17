package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	var velho string
	var novo string

	fmt.Print("Digite uma string: ")
	fmt.Scan(&str)

	fmt.Print("Digite o que quer tirar: ")
	fmt.Scan(&velho)

	fmt.Print("Digite o que quer adcionar: ")
	fmt.Scan(&novo)

	newString := strings.ReplaceAll(str, velho, novo)
	fmt.Println("sua nova string:", newString)
}
