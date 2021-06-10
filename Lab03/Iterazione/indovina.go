package main

import (
    "fmt"
    "math/rand"
    "time"
    )

func main() {

    var numeroGenerato, tentativo int
    
    rand.Seed(int64(time.Now().Nanosecond()))
    
    numeroGenerato = rand.Intn(100)
    
    var i int
    for i = 0; tentativo != numeroGenerato; i++ {
    
        fmt.Print("Tentativo n° ", i+1, ": ")
        fmt.Scan(&tentativo)
        
        if tentativo > numeroGenerato {
            
            fmt.Println("Troppo alto! Riprova!")
        
        } else if tentativo < numeroGenerato {
        
            fmt.Println("Troppo basso! Riprova!")   
        }
    }
    
    fmt.Println("Hai indovinato in", i, "tentativi!")
}
