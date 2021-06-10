package main

import (
        "fmt"
        "os"
        "strconv"
)

func main() {
    
    //legge da riga di comando os.Args
    valori := os.Args[1:]
    
    //scorre i valori dal secondo (1) e controlla:
    for i := range valori {
        
        if !èValido(i, valori) {
            fmt.Printf("Valore in posizione %d non valido.\n", i)
            return
        }
    }
    
    fmt.Printf("Sequenza valida.\n")
}

func èValido(i int, valori []string) bool {
    //che siano tutti numeri
    n1, err1 := strconv.Atoi(valori[i])
    
    if err1 != nil {
        return false
    } 
    
    if i == 0 {
        return true
    }
    n2, err2 := strconv.Atoi(valori[i-1])
    
    if err2 != nil {
        return true
    }
    
    //se è in posizione dispari deve essere minore del precedente
    //se è in posizione pari deve essere maggiore del precedente
    switch {
    
        case i%2 == 0 && n1 > n2: return true
        
        case i%2 != 0 && n1 < n2: return true
        
        default: return false
    }
    
}
