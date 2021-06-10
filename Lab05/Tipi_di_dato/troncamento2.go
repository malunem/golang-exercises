package main

import "fmt"
import "math"

func main() {

    var n float64
    fmt.Print("Inserisci il valore da troncare: ")
    fmt.Scan(&n)
    
    fmt.Println("Valore troncato =", math.Trunc(n*100)/100)
}
