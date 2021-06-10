package main

import (
        "fmt"
        "os"
        "strconv"
)

func main() {
    
    var numeri []int
    
    numeri = make([]int, len(os.Args[1:])) 
    
    for i := 0; i < len(os.Args[1:]); i++ {
        
        numeri[i], _ = strconv.Atoi(os.Args[i+1])
    }
    
    fmt.Println("La somma è", Calcola(numeri)) 
}

func Calcola(numeri []int) int {
    
    var somma int
    
    for i := 0; i < len(numeri); i+=2 {
        
        prodotto := numeri[i]*numeri[i+1]
        somma += prodotto
    }
    
    return somma
}
