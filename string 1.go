package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Print("escreva uma palavra: ")
	fmt.Scan(&str)

	str = strings.ToUpper(str)
	fmt.Println(str)
}
