//Scrivere un programma che legga da **riga di comando** un numero intero `numero` e stampi tutti i numeri *primi* ottenibili rimuovendo al più `3` cifre consecutive tra quelle che definiscono `numero`.

//In particolare, i numeri primi devono essere stampate in ordine crescente (cioè dal più piccolo al più grande).

package main

import (
	"os"
)

func main() {
	num := readNumber()
}

func readNumber() (input string) {
	input := os.Args[1]
	return input
}
