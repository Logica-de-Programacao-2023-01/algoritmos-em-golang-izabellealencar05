package main

import "fmt"

func main() {
	var altura float64
	var base float64
	var profundidade float64
	var volume float64
	fmt.Println("Qual a altura da sua caixa?")
	fmt.Scan(&altura)
	fmt.Println("Qual a base da sua caixa?")
	fmt.Scan(&base)
	fmt.Println("Qual a profundidade da sua caixa?")
	fmt.Scan(&profundidade)
	fmt.Println("O volume da sua caixa é de:", altura*base*profundidade)
	fmt.Scan(&volume)

}
