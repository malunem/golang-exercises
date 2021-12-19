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

	n, _ := strconv.Atoi(number)

	if isPrime(n) {
		//fmt.Println(n)
	}

	for length := 1; length <= 3; length++ {

		fmt.Println("length: ", length) //test

		for index := range number[:len(number)-length+1] {

			fmt.Println("\tindex: ", index) //test

			substring_number := number[:index] + number[index+length:]

			fmt.Printf("\t\t%s := number[:index] (%s) + number[index+length:] (%s)\n", substring_number, number[:index], number[index+length:]) //test

			n, _ := strconv.Atoi(substring_number)

			if isPrime(n) {
				//fmt.Println(n)
				prime_numbers = append(prime_numbers, n)
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
