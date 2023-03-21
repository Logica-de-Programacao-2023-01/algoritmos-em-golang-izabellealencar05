package main

import "fmt"

func main() {
	var x float64
	var y float64
	var z float64
	var w float64
	fmt.Println("Informe o 1numero:")
	fmt.Scan(&x)
	fmt.Println("Informe o 2numero:")
	fmt.Scan(&y)
	fmt.Println("Informe o 3numero:")
	fmt.Scan(&z)
	fmt.Println("Informe o 4numero:")
	fmt.Scan(&w)
	fmt.Println("A media aritmetica é:", (x+y+z+w)/4)

}
