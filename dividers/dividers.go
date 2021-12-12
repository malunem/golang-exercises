// Scrivere un programma che legga da **standard input** un numero intero `n` e stampi a video i **divisori propri** del numero letto, ovvero tutti i suoi divisori escluso il numero stesso.
// Ad esempio, i **divisori** del numero 12 sono: 1, 2, 3, 4, 6, 12;
// quindi i **divisori propri** di 12 sono: 1, 2, 3, 4, 6.

package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Print("Inserisci un numero intero: ")
	fmt.Scan(&num)

	fmt.Println("I divisori propri sono:")

	for i := 1; i < num; i++ {
		if num%i == 0 {
			fmt.Println(i)
		}
	}
}
