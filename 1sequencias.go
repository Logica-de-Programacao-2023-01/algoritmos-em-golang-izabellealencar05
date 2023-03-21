// Faça um algoritmo que leia três números inteiros e mostre a soma entre eles.
package main

import "fmt"

func main() {
	var x int
	var y int
	var z int

	fmt.Println("insira o primeiro número:")
	fmt.Scan(&x)
	fmt.Println("insira o segundo número:")
	fmt.Scan(&y)
	fmt.Println("insira o terceiro número:")
	fmt.Scan(&z)
	fmt.Println("a soma dos três números é:", (x + y + z))

}
