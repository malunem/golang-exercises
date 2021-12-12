// Scrivere un programma che legga da **standard input** un intero `n`.
// Il programma deve stampare a video la sequenza di numeri che vanno da 1 a `n` come mostrato nell'**Esempio d'esecuzione**. In particolare:
//  * ogni numero divisibile per 3 deve essere rimpiazzato dalla parola `"Fizz"`;
//  * ogni numero divisibile per 5 deve essere rimpiazzato dalla parola `"Buzz"`;
//  * ogni numero divisibile sia per 3 sia per 5 deve essere sostituito da `"Fizz Buzz"`.

package main

import (
	"fmt"
)

func main() {

	var num int
	fmt.Print("Inserisci un numero intero: ")
	fmt.Scan(&num)

	var i int
	for i = 1; i <= num; i++ {

		if i%15 == 0 {
			fmt.Println("Fizz Buzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}
}
