// Scrivere un programma che legga da **standard input** due numeri interi che chiameremo `n` e `l` e calcoli l’area di un poligono regolare con `n` lati di lunghezza `l`.

//*Suggerimento:* l'area di un poligono regolare può essere calcolata utilizzando le funzioni `math.Pow` (per il calcolo della potenza) e `math.Tan` (per il calcolo della tangente di un angolo) del package `math` nel seguente modo:

// var area float64 = (n * math.Pow(l, 2)) / (4 * math.Tan(math.Pi/n))

package main

import (
	"fmt"
	"math"
)

func main() {
	var sides_number, sides_length int

	fmt.Print("Inserisci il numero di lati del poligono: ")
	fmt.Scan(&sides_number)
	fmt.Print("Inserisci la lunghezza dei lati del poligono: ")
	fmt.Scan(&sides_length)

	var area float64 = (float64(sides_number) * math.Pow(float64(sides_length), 2)) / (4 * math.Tan(math.Pi/float64(sides_number)))

	fmt.Println("L'area del poligono è:", area)
}
