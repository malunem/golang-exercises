package main

import "fmt"

func main() {

    var n int
    fmt.Print("Inserisci un numero: ")
    fmt.Scan(&n)
    
    for i := 0; i < n; i++ {
        
        for j := 0; j < n; j++ {
        
            fmt.Print("* ")
        }
        
        fmt.Print("\n")
    
    }
}
