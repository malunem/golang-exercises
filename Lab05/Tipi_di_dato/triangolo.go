package main

import "fmt"
import "math/rand"

const EPSILON = 1.0e-9

func main() {

    var s int64
    var m1, q1, m2, q2, m3, q3 float64
    
    fmt.Print("Inserisci s: ")
    fmt.Scan(&s)
    fmt.Print("Inserisci m1 e q1: ")
    fmt.Scan(&m1, &q1)
    fmt.Print("Inserisci m2 e q2: ")
    fmt.Scan(&m2, &q2)
    fmt.Print("Inserisci m3 e q3: ")
    fmt.Scan(&m3, &q3)
    
    rand.Seed(s)
    
    for i := 0; i < 10; i++ {
    
        px := rand.Float64() * 10
        py := rand.Float64() * 10
        
        fmt.Print("Punto (", px, ", ", py, ") - ")
        
        if FuoriTriangolo(px, py, m1, q1, m2, q2, m3, q3) {
        
            fmt.Print("Il punto sta all'esterno del triangolo.")
            
        } else {
            
            fmt.Print("Il punto sta all'interno del triangolo.")
            
        }
        
        fmt.Println()
    }
}

func FuoriTriangolo(px, py, m1, q1, m2, q2, m3, q3 float64) bool {

    sotto_r1 := SottoRetta(px, py, m1, q1)
    sopra_r2 := SopraRetta(px, py, m2, q2)
    sopra_r3 := SopraRetta(px, py, m3, q3)
    
    if sotto_r1 || sopra_r2 || sopra_r3 {
        return true
    } else {
        return false
    }
}

func SottoRetta(px, py, m, q float64) bool {

    return (py - (m*px+q)) < -EPSILON
}

func SopraRetta(px, py, m, q float64) bool {

    return (py - (m*px+q)) > EPSILON
}
