package main

import "fmt"

func main() {

    var n int
    
    fmt.Print("Inserisci un numero: ")
    fmt.Scan(&n)
    
    for i := 1; i <= 10; i++ {
        
        fmt.Println(i, "x", n, "=", i * n)
    
    }

}
