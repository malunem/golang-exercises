package main

import (
    "fmt"
    "math/rand"
)

const EPSILON = 1.0e-9

func main() {

    var (
        s int64
        m, q float64
    )
    
    fmt.Print("Inserisci s: ")
    fmt.Scan(&s)
    
    fmt.Print("Inserisci m e q: ")
    fmt.Scan(&m, &q)
    
    rand.Seed(s)
    
    for i := 0; i < 10; i++ {
        px := rand.Float64() * 5.0
        py := rand.Float64() * 5.0
        
        fmt.Print("Punto (", px, ",", py, ") - ")
        
        if èXMaggioreDiY(py, m*px+q) {
                fmt.Print("Il punto sta sopra la retta\n")
            
        } else if èXMinoreDiY(py, m*px+q) {
                fmt.Print("Il punto sta sotto la retta\n")
        
        } else {
                fmt.Print("Il punto appartiene alla retta\n")
        }
    }
}

func èXMaggioreDiY(x, y float64) bool {
    return (x - y) > EPSILON
}

func èXMinoreDiY(x, y float64) bool {
    return (x - y) < -EPSILON
}
