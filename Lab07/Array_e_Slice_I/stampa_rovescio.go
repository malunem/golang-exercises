package main

import "fmt"

func main() {

    var n int
    fmt.Scan(&n)
    
    var numeri []int 
    numeri = make([]int, n)
    
    fmt.Printf("Inserisci %d numeri:\n", n)    
        
    for i := 0; i < n; i++ {
    
        fmt.Scan(&numeri[i])    
    }
    
    fmt.Println("Numeri in ordine inverso:")
    
    for i := n-1; i >= 0; i-- {
    
        fmt.Printf("%d ", numeri[i])
    }
    
    fmt.Println()
}
