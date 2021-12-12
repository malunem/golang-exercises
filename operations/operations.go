// Scrivere un programma che, dopo aver letto da **standard input** un numero intero `n`, chiede all'utente di inserire `n` numeri interi (sempre da **standard input**).

//Dopo aver letto gli `n` numeri interi, il programma deve stampare:
// la somma degli `n` numeri letti;
// il minimo tra i numeri letti;
// il massimo tra i numeri letti;
// il numero di interi letti strettamente positivi (maggiori di 0), strettamente negativi (minori di 0), e nulli.

package main

import (
	"fmt"
)

func main() {

	var num int

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&num)

	fmt.Println("Inserisci", num, "numeri: ")

	var i, sum, min, max, positives, negatives, nulls int

	for i = 0; i < num; i++ {
		var this_number int
		fmt.Scan(&this_number)

		if i == 0 {
			sum = this_number
			min = this_number
			max = this_number
		} else {
			sum += this_number

			if this_number < min {
				min = this_number
			} else if this_number > max {
				max = this_number
			}
		}

		if this_number > 0 {
			positives++
		} else if this_number < 0 {
			negatives++
		} else {
			nulls++
		}
	}

	fmt.Println("Somma:", sum)
	fmt.Println("Minimo:", min)
	fmt.Println("Massimo:", max)
	fmt.Println("Positivi:", positives)
	fmt.Println("Negativi:", negatives)
	fmt.Println("Nulli:", nulls)
}
