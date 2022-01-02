package main

import (
	"fmt"
	"os"
)

func main() {
	inputNum := readNumber()

	for i := 0; i < len(inputNum)-1; i++ {

		//fmt.Println(inputNum[i])

		num1 := string(inputNum[i])
		num2 := string(inputNum[i+1])

		if num1 < num2 {
			fmt.Print(num1)
		}
	}

	fmt.Println()
}

func readNumber() string {
	inputNum := os.Args[1]

	//fmt.Println(inputNum)

	return inputNum
}
