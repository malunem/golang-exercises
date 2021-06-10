package main

import "fmt"
import "math"

func main() {
    var raggio float64
    
    fmt.Printf("Raggio = ")
    fmt.Scan(&raggio)
    
    var circ float64 = 2 * raggio * math.Pi
    fmt.Println("Circonferenza = ", circ)
    
    var area float64 = raggio * raggio * math.Pi
    fmt.Println("Area = ", area)
}
