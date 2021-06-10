package main

import "fmt"

func main() {

    var a, b int
    fmt.Println("Inserisci due numeri:")
    fmt.Scan(&a, &b)
    
    if b == 0 {
        
        fmt.Println("Impossibile")
        
    } else {
        
        res := float64(a)/float64(b)
        fmt.Println("Quoziente =", res)
    }
}
