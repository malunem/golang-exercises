//  axˆ2+bx+c=0
// 5xˆ2+1x+2=0

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	a, b, c, limit, epsilon := readFunction()

	disc := calcDisc(a, b, c)

	resultsNumber := checkDisc(disc)

	results := calcResults(a, b, c, resultsNumber, disc)

	printResults(resultsNumber, results)

	aboveLimit := checkAboveLimit(resultsNumber, results, limit, epsilon)

	printAboveLimit(aboveLimit)

}

func readFunction() (a, b, c int, limit, epsilon float64) {
	function := os.Args[1]
	limit, _ = strconv.ParseFloat(os.Args[2], 64)
	epsilon, _ = strconv.ParseFloat(os.Args[3], 64)
	function = strings.Replace(function, "-", "+-", -1)
	splittedFunction := strings.Split(function, "=")
	abc := splittedFunction[0]

	splittedABC := strings.Split(abc, "+")

	splittedA := strings.Split(splittedABC[0], "x")
	a, _ = strconv.Atoi(splittedA[0])

	splittedB := strings.Split(splittedABC[1], "x")
	b, _ = strconv.Atoi(splittedB[0])

	c, _ = strconv.Atoi(splittedABC[2])

	return a, b, c, limit, epsilon
}

func calcDisc(a, b, c int) (disc float64) {
	//b^2 - 4ac
	b2 := math.Pow(float64(b), 2)
	ac := float64(a) * float64(c)
	disc = b2 - 4*ac

	return disc
}

func checkDisc(disc float64) (resultsNumber int) {

	if disc < 0 {
		resultsNumber = 0
	} else if disc == 0 {
		resultsNumber = 1
	} else {
		resultsNumber = 2
	}

	return resultsNumber
}

func calcResults(a, b, c, resultsNumber int, disc float64) (results []float64) {

	switch resultsNumber {

	case 2:
		// (-b + sqrt(disc)) / (2a)
		result1 := (-float64(b) + math.Sqrt(disc)) / (2 * float64(a))

		// (-b - sqrt(disc)) / (2a)
		result2 := (-float64(b) - math.Sqrt(disc)) / (2 * float64(a))

		results = append(results, result1, result2)

	case 1:
		// -b / 2a
		result := (-float64(b)) / (2 * float64(a))

		results = append(results, result)

	default:
		break
	}

	return results
}

func printResults(resultsNumber int, results []float64) {
	switch resultsNumber {
	case 2:
		fmt.Print("Esistono due soluzioni reali: ")
	case 1:
		fmt.Print("Esiste una soluzione reale: ")
	default:
		fmt.Print("Non esistono soluzioni reali.")
	}

	for i := 0; i < resultsNumber; i++ {
		fmt.Printf("%.3f", results[i])
		if resultsNumber == 2 && i == 0 {
			fmt.Print(" e ")
		}
	}

	fmt.Println()
}

func checkAboveLimit(resultsNumber int, results []float64, limit, epsilon float64) (aboveLimit []float64) {

	for i := 0; i < resultsNumber; i++ {

		if (results[i] - limit) > epsilon {
			aboveLimit = append(aboveLimit, results[i])
		}
	}

	return aboveLimit
}

func printAboveLimit(aboveLimit []float64) {
	for i := 0; i < len(aboveLimit); i++ {
		fmt.Printf("La soluzione %.3f è maggiore della soglia.\n", aboveLimit[i])
	}
}
