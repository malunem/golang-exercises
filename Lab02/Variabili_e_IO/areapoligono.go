package main

import "fmt"
import "math"

func main() {

    var n, l int
    
    fmt.Print("Inserisci il numero di lati del poligono: ")
    fmt.Scan(&n)
    
    fmt.Print("Inserisci la lunghezza dei lati del poligono: ")
    fmt.Scan(&l)
    
    var area float64 = (float64(n) * math.Pow(float64(l), 2)) /
                        4 * math.Tan(math.Pi/float64(n))
    fmt.Println("Area calcolata:", area)
}
