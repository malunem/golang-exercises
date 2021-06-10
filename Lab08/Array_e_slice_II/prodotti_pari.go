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
    
    for i, n1 := range numeri {
        
        for _, n2 := range numeri[i+1:] {
        
            prodotto := n1*n2
            
            if prodotto%2 == 0 {
            
                somma += prodotto
            }
        }
    }
    
    return somma
}
