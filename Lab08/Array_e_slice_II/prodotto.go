package main

import (
        "fmt"
        "os"
        "strconv"
)

func main() {

    var prodotto int = 1
    
    for _, c := range os.Args[1:] {
        
        n, err := strconv.Atoi(c)
        
        if err == nil {
            
            prodotto *= n
        }
    } 
    
    fmt.Printf("Il risultato della moltiplicazione tra i numeri interi è %d\n", prodotto)   
}
