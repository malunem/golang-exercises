//Scrivere un programma che legga da **standard input** il raggio di un cerchio e ne calcoli circonferenza e area.

//*Suggerimento:* l'area del cerchio si calcola facendo `raggio x raggio x pi_greco`, mentre la circonferenza facendo `2 x raggio x pi_greco`. Il valore numerico di `pi_greco` è memorizzato nella costante `Pi` del package `math`, a cui ci si può riferire con `math.Pi`.

package main

import (
	"fmt"
	"math"
)

func main() {

	var radius float64

	fmt.Print("Inserisci il raggio del cerchio: ")
	fmt.Scan(&radius)

	var circumference float64 = 2 * radius * math.Pi

	var area float64 = math.Pow(radius, 2) * math.Pi

	fmt.Println("La circonferenza è:", circumference)
	fmt.Println("L'area è", area)
}
