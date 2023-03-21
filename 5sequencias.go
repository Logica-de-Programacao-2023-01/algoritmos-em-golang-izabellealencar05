// Faça um algoritmo que leia a idade de uma pessoa em anos e mostre a idade em dias.
package main

import "fmt"

func main() {
	var idade int
	fmt.Print("Qual a sua idade? ")
	fmt.Scan(&idade)
	fmt.Print("Logo, sua idade em dias é: ", (idade * 365))

}
