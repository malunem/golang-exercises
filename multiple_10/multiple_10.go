// Scrivere un programma che legge da **standard input** un numero intero `n` e verifica se il numero è multiplo di 10.

package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Print("Inserisci un numero intero: ")
	fmt.Scan(&num)

	if num%10 == 0 {
		fmt.Print(num, " è divisibile per 10\n")
	} else {
		fmt.Print(num, " non è divisibile per 10\n")
	}
}
