// Faça um algoritmo que leia três números reais e mostre a média ponderada entre eles, com pesos 2, 3 e 5, respectivamente.
package main

import "fmt"

func main() {
	var x int
	var y int
	var z int

	fmt.Print("Insira o primeiro número: ")
	fmt.Scan(&x)
	fmt.Print("Insira o segundo número: ")
	fmt.Scan(&y)
	fmt.Print("Insira o terceiro número: ")
	fmt.Scan(&z)
	fmt.Print("A média ponderada é: ", (x*2+y*3+z*5)/10)

}
