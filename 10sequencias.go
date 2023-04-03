// Faça um algoritmo que leia o peso de uma pessoa em quilos e converta para libras
package main

import "fmt"

func main() {
	var quilos float64

	fmt.Println("Qual o seu peso em quilos?")
	fmt.Scan(&quilos)
	fmt.Println("O seu peso convertido de quilos para libras é", quilos/0.45)

}
