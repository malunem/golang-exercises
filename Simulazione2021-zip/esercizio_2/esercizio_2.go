package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"unicode"
)

func main() {
	inputString := readString()

	if !isValidInput(inputString) {
		return
	}

	substrings := Sottostringhe(inputString)

	validSubstrings := checkSubstrings(substrings)

	printSorted(validSubstrings)

}

func readString() string {
	input := os.Args[1]
	return input
}

func isValidInput(input string) bool {

	isValid := true

	asciiInput, _ := strconv.Unquote(strconv.QuoteToASCII(input))

	if input != asciiInput {
		return false
	}

	for _, v := range input {
		if unicode.IsUpper(v) {
			isValid = false
		} else if unicode.IsDigit(v) {
			isValid = false
		} else if unicode.IsSymbol(v) {
			isValid = false
		}
	}

	return isValid
}

func Sottostringhe(s string) map[string]int {

	substrings := make(map[string]int)

	substringMaxLength := len(s)

	for i := 1; i < substringMaxLength; i++ {

		for j := range s[:len(s)-i] {

			substring := getSubstring(s, i+1, j)
			substrings[substring]++
		}
	}

	return substrings
}

func getSubstring(originString string, substringLength, initialPosition int) (substring string) {

	substring = originString[initialPosition : initialPosition+substringLength]

	return substring
}

func checkSubstrings(substrings map[string]int) map[string]int {

	validSubstrings := make(map[string]int)

	for substring, val := range substrings {

		isValid := true

		for j := range substring[:len(substring)-1] {

			if substring[j] >= substring[j+1] {
				isValid = false
			}
		}

		if isValid {
			validSubstrings[substring] = val
		}
	}

	return validSubstrings
}

func printSorted(substrings map[string]int) {

	var keys []string

	for key := range substrings {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	fmt.Println("output:")

	for _, key := range keys {
		fmt.Println(key, substrings[key])
	}
}
