//Faça um algoritmo que leia o preço de um produto e mostre o seu valor com desconto de 10%.

package main

import "fmt"

func main() {
	var preco float64
	var desconto float64

	fmt.Println("qual o preço?")
	fmt.Scan(&preco)
	fmt.Println("o desconto será:"((preco*10)/100) - preco)

}
