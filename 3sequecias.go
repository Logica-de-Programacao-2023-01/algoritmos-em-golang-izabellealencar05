// Faça um algoritmo que leia o peso e a altura de uma pessoa e calcule o seu IMC (Índice de Massa Corporal).
package main

import "fmt"

func main() {
	var x float64
	var y float64
	fmt.Println("Qual a sua altura?")
	fmt.Scan(&x)
	fmt.Println("Qual o seu peso?")
	fmt.Scan(&y)
	fmt.Print("Então, seu IMC é ", y/(x*x))

}
