package main

import "fmt"

func main() {
    
    var n, sum, min, max, pos, neg, nul int
    
    fmt.Scan(&n)
    
    fmt.Println("Inserisci", n, "numeri:")
    
    for i := 0; i < n; i++ {
        
        var num int
        fmt.Scan(&num)
        
        sum += num
        
        if num < min || i == 0 {
            min = num
        }
        
        if num > max || i == 0 {
            max = num
        }
        
        if num > 0 {
            pos++
        } else if num < 0 {
            neg++
        } else {
            nul++
        }
    }
    
    fmt.Println("somma =", sum)
    fmt.Println("valore minimo =", min)
    fmt.Println("valore massimo =", max)
    fmt.Println("interi > 0 =", pos)
    fmt.Println("interni < 0 =", neg)
    fmt.Println("interi = 0 =", nul)
}
