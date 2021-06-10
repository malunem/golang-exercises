package main

import (
        "fmt"
        "strconv"
        "os"
        "math/rand"
        "time"
)

func main() {

    soglia, _ := strconv.Atoi(os.Args[1])
    
    //Generare numeri tra 0 e 100 finché uno non è minore di soglia
    numeri := Genera(soglia)
    
    //Stampa tutti i numeri generati
    fmt.Printf("Valori generati %v\n", numeri)
    
    //Stampa tutti i numeri generati superiori a soglia
    fmt.Printf("Valori sopra soglia %v\n", numeri[:len(numeri)-1])
}

func Genera(soglia int) []int {

    var numeri []int
    
    rand.Seed(int64(time.Now().Nanosecond())) 
    
    for {
        
        random := rand.Intn(100) +1
        
        numeri = append(numeri, random)
        
        if random < soglia {
            break
        }
    }
    
    return numeri
}
