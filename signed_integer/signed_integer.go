// Scrivere un programma che legge da **standard input** un numero intero `n` (specificato senza segno se maggiore o uguale a 0) e stampi a video il numero con segno.

package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Print("Inserisci un numero intero: ")
	fmt.Scan(&num)

	if num > 0 {
		fmt.Print("+", num, "\n")
	} else {
		fmt.Println(num)
	}
}
