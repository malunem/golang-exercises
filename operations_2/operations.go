// Create un programma che legge da **standard input** due numeri interi, che chiameremo `x` e `y`.
// Letti i due numeri, il programma stampa:
// * il maggiore tra `x` e `y`
// * il minore tra `x` e `y`
// * il risultato della somma tra `x` e `y`
// * il risultato della differenza tra il maggiore e il minore
// * il risultato della divisione `x/y`
// * il risultato del prodotto `x*y`
// * il valore medio tra `x` e `y`
// * il risultato di `x` elevato alla `y` (utilizzando sia un ciclo `for` sia la funzione `math.Pow`)

package main

import (
	"fmt"
	"math"
)

func main() {
	var num1, num2 int
	fmt.Print("Inserisci due numeri interi: ")
	fmt.Scan(&num1, &num2)

	max, min := findMaxMin(num1, num2)
	fmt.Println("Maggiore:", max)
	fmt.Println("Minore:", min)

	fmt.Println("Somma:", num1+num2)
	fmt.Println("Differenza:", max-min)
	fmt.Println("Divisione:", float64(num1)/float64(num2))
	fmt.Println("Moltiplicazione:", num1*num2)
	fmt.Println("Media:", float64(num1+num2)/2)

	pow_math := math.Pow(float64(num1), float64(num2))
	pow_for := calcPowerWithForLoop(num1, num2)

	fmt.Println("Potenza math.Pow:", pow_math)
	fmt.Println("Potenza for loop:", pow_for)
}

func findMaxMin(num1 int, num2 int) (max int, min int) {

	if num1 > num2 {
		return num1, num2
	} else {
		return num2, num1
	}
}

func calcPowerWithForLoop(num1 int, num2 int) (pow int) {

	pow = 1
	for i := 0; i < num2; i++ {
		pow *= num1
	}

	return pow
}
