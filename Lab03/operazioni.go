package main

import "fmt"
import "math"

func main() {
    
    var x, y int
    
    fmt.Println("Inserisci due numeri interi:")
    fmt.Scan(&x, &y)
    
    var max, min int
    
    if x >= y {
    
        max = x
        min = y
    } else {
    
        max = y
        min = x 
    }
    
    fmt.Print("Maggiore: ", max, "\n" +
                "Minore: ", min, "\n" +
                "Somma: ", x+y, "\n" +
                "Differenza: ", max-min, "\n" +
                "Prodotto: ", x*y, "\n" +
                "Divisione: ", float64(x)/float64(y), "\n" +
                "Valore medio: ", (float64(x)+float64(y))/2, "\n")
    
    var potenza int = 1
    
    for i := 0; i < y; i++ {
    
        potenza *= x
    }
    
    fmt.Println("Potenza (ciclo for): ", potenza)
    fmt.Println("Potenza (math.Pow): ", math.Pow(float64(x), float64(y)))
}
