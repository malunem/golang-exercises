package main

import "fmt"
import "math"

func main() {

    var n float64
    fmt.Print("Inserisci il valore: ")
    fmt.Scan(&n)
    
    var cifre int
    fmt.Print("Inserisci il numero di cifre dopo la virgola: ")
    fmt.Scan(&cifre)
    
    var potenza float64
    potenza = math.Pow(10, float64(cifre))
    
    fmt.Println("Valore troncato =", math.Trunc(n*potenza)/potenza)
    fmt.Println("Valore arrotondato =", math.Round(n*potenza)/potenza)
}
