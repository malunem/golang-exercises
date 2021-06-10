package main

import (
        "fmt"
        "os"
        "strconv"
        "math"
)

func main() {

    numeri := LeggiNumeri() //da riga di comando
    
    fmt.Printf("Minimo: %d\n", Minimo(numeri))
    
    fmt.Printf("Massimo: %d\n", Massimo(numeri))
    
    fmt.Printf("Media: %.2f\n", Media(numeri))
}

func LeggiNumeri() (numeri []int) {

    //crea slice dinamicamente con make e la popola con gli input
    
    numeri = make([]int, len(os.Args)-1)
    
    for i, c := range os.Args[1:] {
    
        numeri[i], _ = strconv.Atoi(c)
    }
    
    return numeri
}

func Minimo(slice []int) int {
    
    //scorre la slice e salva il minore
    
    var minimo int = math.MaxInt64 
    
    for _, c := range slice {
        
        if c < minimo {
            minimo = c
        }
    }
    
    return minimo
}

func Massimo(slice []int) int {
    
    //scorre la slice e salva il maggiore
    var max int = math.MinInt64 
    
    for _, c := range slice {
        
        if c > max {
            max = c
        }
    }
    
    return max
}

func Media(slice []int) float64 {

    //scorre la slice, somma i numeri e poi calcola media
    
    var sum float64
    
    for _, c := range slice {
        
        sum += float64(c)
    }
    
    var media float64 = sum/float64(len(slice))
    
    return media
}
