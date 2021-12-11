// Scrivere un programma che legga da **standard input** le misure dell’altezza e della base di un rettangolo e ne calcoli il perimetro e l’area.

package main

import "fmt"

func main() {

	var height, length int

	fmt.Println("Inserisci altezza e lunghezza del rettangolo:")
	fmt.Scan(&height, &length)

	var perimeter int = (height + length) * 2
	var area int = height * length

	fmt.Println("Il perimetro è:", perimeter)
	fmt.Println("L'area è:", area)
}
