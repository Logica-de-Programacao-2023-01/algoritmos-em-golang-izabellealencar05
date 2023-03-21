// Faça um algoritmo que leia um número inteiro e mostre o seu sucessor e antecessor.
package main

import "fmt"

func main() {
	var x int

	fmt.Println("Insira um número: ")
	fmt.Scan(&x)
	fmt.Print("O sucessor desse número é ", x+1)
	fmt.Scan(x)
	fmt.Print(" e o antecessor é ", x-1)

}
