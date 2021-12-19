// Scrivere un programma che, dopo aver letto da **standard input** un numero intero `n`, chiede all'utente di inserire `n` numeri interi (sempre da **standard input**).

// Il programma deve stampare gli `n` numeri interi in ordine inverso rispetto a quello di inserimento.

// *Suggerimento:* Per creare dinamicamente una slice si utilizzi la funzione `make()`.

package main

import (
	"fmt"
)

func main() {
	num := askNumber()
	numbers := askNumbers(num)

	printReversedNumbers(numbers)
}

func askNumber() (num int) {
	fmt.Print("Inserisci un numero intero: ")
	fmt.Scan(&num)

	return num
}

func askNumbers(num int) []int {

	numbers := make([]int, num)

	fmt.Printf("Inserisci %d numeri interi: ", num)

	for i := 0; i < num; i++ {
		var n int
		fmt.Scan(&n)

		numbers[i] = n
	}

	return numbers
}

func printReversedNumbers(numbers []int) {
	length := len(numbers)

	fmt.Println("Numeri in ordine inverso:")

	for i := length - 1; i >= 0; i-- {
		fmt.Printf("%d ", numbers[i])
	}

	fmt.Println()
}
