package main

import "fmt"
import "math"

func main() {

    var n float64
    
    fmt.Print("Inserisci un numero: ")
    fmt.Scan(&n)
    
    radice, ok := RadiceQuadrata(n)
    
    if ok {
    
        fmt.Println("Radice quadrata: ", radice)
    } else {
        
        fmt.Println("Il valore inserito non appartiene al dominio della funzione.")
    }
}

func RadiceQuadrata(n float64) (float64, bool) {

    if n < 0 {
    
        return 0, false
    } else {
    
        return math.Sqrt(n), true
    }
}
