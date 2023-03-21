package main

import "fmt"

func main() {
	var base float64
	var altura float64
	var area float64
	fmt.Println("Qual a base do seu retângulo?")
	fmt.Scan(&base)
	fmt.Println("Qual a altura?")
	fmt.Scan(&altura)
	fmt.Println("a área do seu retângulo é:", base*altura)
	fmt.Scan(&area)

}
