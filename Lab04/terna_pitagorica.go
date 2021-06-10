package main

import "fmt"
import "math"

func main () {

    var soglia int
    fmt.Print("Inserisci la soglia: ")
    fmt.Scan(&soglia)
    
    if soglia <= 0 {
            return
    }
    
    TernePitagoriche(soglia)
}

func TernePitagoriche(soglia int) {

    for i := 1; i < soglia; i++ {
    
        for j := 1; j < soglia; j++ {
        
            for k := 1; k < soglia; k++ {
            
                if èTernaPitagorica(i, j, k) {
                
                    fmt.Print("(",i, ", ", j, ", ", k, ")\n")
                }
            }   
        }
    }
}

func èTernaPitagorica(a int, b int, c int) bool {
    
    var a2 float64 = math.Pow(float64(a), 2)
    var b2 float64 = math.Pow(float64(b), 2)
    var c2 float64 = math.Pow(float64(c), 2)
    
    if c2 == a2 + b2 {
                return true
        } else {
                return false
        }
}
