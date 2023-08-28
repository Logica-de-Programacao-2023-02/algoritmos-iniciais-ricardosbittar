package main

import "fmt"

func main() {
	var t float64
	var valorInicial float64

	fmt.Println("Qual o valor inicial? ")
	fmt.Scanln(&valorInicial)
	fmt.Println("Quantos meses? ")
	fmt.Scanln(&t)

	juros := 2.3

	jurostempo := t ^ juros
	valorFinal := valorInicial * jurostempo
	fmt.Println(valorFinal)

}
