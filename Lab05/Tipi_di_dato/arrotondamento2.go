package main

import "fmt"
import "math"

func main() {

    var n float64
    fmt.Print("Inserisci il valore da arrotondare: ")
    fmt.Scan(&n)
    
    fmt.Println("Valore arrotondato =", math.Round(n*100)/100)
}
