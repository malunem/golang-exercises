// Scrivere un programma che legga da **standard input** una sequenza di numeri reali strettamente positivi (un numero è strettamente positivo se è maggiore di 0; se un numero è minore o uguale 0 non è strettamente positivo). La lettura termina quando viene letto un numero minore o uguale a 0.

// Il programma deve stampare a video il risultato della media aritmetica dei valori inseriti.

package main

import (
	"fmt"
)

func main() {

	fmt.Println("Inserisci una sequenza di numeri reali positivi:")

	var sum float64
	var counter int

	for {
		var this_number float64
		fmt.Scan(&this_number)

		if this_number <= 0 {
			break
		}

		counter++
		sum += this_number
	}

	var average float64 = sum / float64(counter)

	fmt.Println("Media:", average)
}
