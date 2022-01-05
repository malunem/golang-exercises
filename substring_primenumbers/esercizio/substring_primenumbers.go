//Scrivere un programma che legga da **riga di comando** un numero intero `numero` e stampi tutti i numeri *primi* ottenibili rimuovendo al più `3` cifre consecutive tra quelle che definiscono `numero`.

//In particolare, i numeri primi devono essere stampate in ordine crescente (cioè dal più piccolo al più grande).

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	num := readNumber()
	length := len(num)

	for i := 0; i < 3; i++ {
		for j := 0; j < length-i; j++ {
			newNum := removeDigits(num, i+1, j)
			if isPrime(newNum) {
				fmt.Println(newNum)
			}
		}
	}
}

func readNumber() (input string) {
	input = os.Args[1]
	return input
}

func removeDigits(number string, quantity, position int) (newNumber int) {
	newString := number[:position] + number[position+quantity:]
	newNumber, _ = strconv.Atoi(newString)
	return newNumber
}

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}
