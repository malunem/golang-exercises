package main

import "fmt"

func main() {

    var i int
    var somma int
    var media float64
    
    fmt.Print("Inserisci una sequenza di numeri (interrompi con numero <= 0): ")
    
    for {
        
        var n int
        fmt.Scan(&n)
        
        if n <= 0 {
            break
        }
        
        somma += n
        i++
        
    }
    
    media = float64(somma) / float64(i)
    fmt.Println("Media aritmetica: ", media)
}
