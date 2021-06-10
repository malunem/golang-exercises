package main 

import (
        "fmt"
        "os"
        "strconv"
        "sort"
)

func main() {
    
    //Legge da riga di comando os.Args un numero int >= 1000
    n := os.Args[1] //stringa
    
    //Salva il numero in input come primo valore della slice
    
    input, err := strconv.Atoi(n)
    
    var numeri []int
    
    if err == nil {
        numeri = append(numeri, input)
    }
    
    //Prova tutte le combinazioni rimuovendo max 3 cifre consecutive in 
        //qualsiasi posizione e salva i numeri ottenuti
    
    for i := 1; i < 4; i++  {
    
        for j := 0; j <= len(n)-i; j++ {
        
            numero := n[:j] + n[j+i:]
            
            num, err := strconv.Atoi(numero)
            if err == nil {
                numeri = append(numeri, num)
            }
        }
    }
        
    //Scorre i numeri ottenuti e salva quelli primi
    var numeriPrimi []int
    
    for _, v := range numeri {
        
        if èPrimo(v) {
            numeriPrimi = append(numeriPrimi, v)
        } 
    }
    
    //Ordina e stampa i numeri primi in ordine crescente (sort.Ints(n))
    sort.Ints(numeriPrimi)
    
    for _, v := range numeriPrimi {
    
        fmt.Printf("%d\n", v)
    }
}

func èPrimo(n int) bool {
    
    //restituisce true se n è primo
    //primo: diviso per tutti i numeri tra 0 e n dà zero solo con 1 e n
    //se è divisibile per qualsiasi altro numero non è primo
    
    for i := 2; i < n; i++ {
        if n%i == 0 {
            return false
        }
    }
    
    return true
}
