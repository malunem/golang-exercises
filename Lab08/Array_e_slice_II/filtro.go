package main

import (
        "fmt"
        "os"
        "strconv"
)

func main() {
    
    //Legge da riga di comando slice di int con func LeggiNumeri
    var voti []int = LeggiNumeri()
    
    var sufficienti, insufficienti []int = FiltraVoti(voti)
    //Stampa slice voti insufficienti
    fmt.Printf("Voti insufficienti: %v\n", insufficienti)
    
    //Stampa slice voti sufficienti
    fmt.Printf("Voti sufficienti: %v\n", sufficienti)
}

func LeggiNumeri() (numeri []int) {

    //Costruisce slice con valori a riga di comando
    numeri = make([]int, len(os.Args[1:]))
    
    for i, v := range os.Args[1:] {
        
        num, err := strconv.Atoi(v)
        
        if err == nil {
        
            numeri[i] = num
        }
    }
    
    return
}

func FiltraVoti(voti []int) (sufficienti, insufficienti []int) {

    //Divide la slice in due slice in base al voto
    
    for _, v := range voti {
        
        if v >= 60 {
            
            sufficienti = append(sufficienti, v)
        } else {
            insufficienti = append(insufficienti, v)
        }
    }
    
    return
}
