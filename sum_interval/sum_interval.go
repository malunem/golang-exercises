// Scrivere un programma che legga da **standard input** due numeri interi `a` e `b` e stampi a video la somma dei numeri pari compresi tra `a` e `b` (estremi esclusi).

package main

import (
	"fmt"
)

func main() {
	var num_1, num_2 int

	fmt.Print("Inserisci due numeri interi: ")
	fmt.Scan(&num_1, &num_2)

	var i, sum int
	for i = (num_1 + 1); i < num_2; i++ {
		if i%2 == 0 {
			sum += i
		}
	}

	fmt.Println("La somma dei numeri pari compresi tra i numeri inseriti è:", sum)
}
