// Scrivere un programma che legga da **standard input** un intero `n > 1` e stampi, utilizzando il carattere `*`, il perimetro di due triangoli rettangoli con base e altezza di lunghezza `n`, posizionati come mostrato nell'**Esempio d'esecuzione**.

package main

import (
	"fmt"
)

func main() {
	size := askSize()
	printUpperTriangle(size)
	printLowerTriangle(size)
}

func askSize() (size int) {
	fmt.Print("Inserisci un numero intero: ")
	fmt.Scan(&size)
	return size
}

func printUpperTriangle(size int) {

	//rows
	for i := 0; i < size; i++ {

		//cols
		for j := 0; j < size; j++ {

			isBorder := i == 0 || i == j || j == size-1

			if isBorder {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func printLowerTriangle(size int) {

	//rows
	for i := 0; i < size; i++ {

		//spaces
		for k := 0; k < size; k++ {
			fmt.Print(" ")
		}

		//cols
		for j := 0; j < size; j++ {

			isBorder := i == size-1 || i == j || j == 0

			if isBorder {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
