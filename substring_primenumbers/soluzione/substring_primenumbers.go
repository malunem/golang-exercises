package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	number := os.Args[1]

	prime_numbers := getPrimeNumbers(number)

	sort.Ints(prime_numbers)

	// for i := range prime_numbers {
	// 	fmt.Println(prime_numbers[i])
	// }

}

func getPrimeNumbers(number string) (prime_numbers []int) {

	//n, _ := strconv.Atoi(number)

	//if isPrime(n) {
	//fmt.Println(n)
	//	}

	for length := 1; length <= 3; length++ {

		fmt.Println("\nlength: ", length) //test
		fmt.Printf("\nrange su: %s (number[:%d])\n", number[:len(number)-length+1], len(number)-length+1)

		for index := range number[:len(number)-length+1] {

			fmt.Println("\n\tindex: ", index) //test

			substring_number := number[:index] + number[index+length:]

			fmt.Printf("\n\t\t%s := number[:%d] + number[%d:]\n", substring_number, index, index+length)        //test
			fmt.Printf("\n\t\trimosso: %s := number[%d:%d]\n", number[index:index+length], index, index+length) //test

			n, _ := strconv.Atoi(substring_number)

			if isPrime(n) {
				fmt.Printf("\n\n\t\t*** %d ***\n\n", n)
				prime_numbers = append(prime_numbers, n)
			} else {
				fmt.Printf("\n\n\t\txxx %d xxx\n\n", n)
			}
		}
	}
	// fmt.Println()
	return
}

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
