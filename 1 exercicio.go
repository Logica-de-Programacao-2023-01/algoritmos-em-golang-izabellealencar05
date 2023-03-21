package main

import "fmt"

func main() {
	var nome string
	var idade int
	var peso float64
	fmt.Println(" Qual o seu nome? ")
	fmt.Scan(&nome)
	fmt.Println("Qual a sua idade?")
	fmt.Scan(&idade)
	fmt.Println("Qual o seu peso?")
	fmt.Scan(&peso)
	fmt.Println("Olá,", nome, "sua idade é", idade, "e seu peso é", peso, "obrigado.")
}
