package main

import (
    "fmt"
    "math"
)

func main() {

    var num float64
    
    fmt.Print("Inserisci un numero: ")
    fmt.Scan(&num)
    
    radq := math.Sqrt(num)
    
    if radq*radq == num {
        
        fmt.Println(num, "uguale a", radq, "*", radq)
    } else {
        
        fmt.Println(num, "diverso da", radq, "*", radq)
    }
}
