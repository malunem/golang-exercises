package main

import (
        "fmt"
        "os"
        "strconv"
)

func main() {
    
    //legge i numeri interi da riga di comando
    numeri := LeggiNumeri()
    
    somma := 0
    //calcola quante volte sono presenti i valori
    for _, v := range numeri {
        
        if Occorrenze(numeri, v) == 1 {
            somma += v
        }
    }
    
    //stampa somma
    fmt.Printf("%d\n", somma)
}

func LeggiNumeri() (numeri []int) {
    
    //scarta i valori non numeri
    for _, v := range os.Args[1:] {
        
        if n, err := strconv.Atoi(v); err == nil {
            numeri = append(numeri, n)
        }
    }
    
    return
}

func Occorrenze(numeri []int, n int) int {
    
    //restituisce quante volte è presente n nella slice numeri
    occorrenze := 0
    for _, v := range numeri {
        
        if n == v {
            occorrenze++
        }
    }
    
    return occorrenze
}
