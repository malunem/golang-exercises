package main

import (
        "fmt"
        "strconv"
        "os"
        "math/rand"
        "time"
)

func main() {
    //legge da os.Args un int n
    n, _ := strconv.Atoi(os.Args[1])
    
    //crea e inizializza la slice [][]int con le funzioni
    slice := CreaSliceBidimensionale(n)
    InizializzaSliceBidimensionale(slice)
    
    //stampa le coppie ottenute con la funzione
    lunghezza := Coppie(slice)
    for _, v := range lunghezza {
        fmt.Printf("%v\n", v)
    }
    fmt.Println()
}

func CreaSliceBidimensionale(l int) [][]int {
    //crea la prima slice di slice con lunghezza l
    s := make([][]int, l)
    
    //crea le seconde slice con lunghezza l
    for i := range s {
        s[i] = make([]int, l)
    }
    
    //restituisce slice di slice
    return s
}

func InizializzaSliceBidimensionale(s [][]int) {
    //crea un rand.Seed
    rand.Seed(int64(time.Now().Nanosecond()))
    
    //assegna valori casuali compresi tra 0 e 1 alle slice
    for i := range s {
        for j := range s[i] {
            s[i][j] = rand.Intn(2)
        }
    }
}

func Coppie(s [][]int) (coppie [][]int) {
    //coppie = make([][]int, 1)
    for i := range s {
        for j := range s[i] {
            if s[i][j] == 1 {
                slice := []int{i, j}
                coppie = append(coppie, slice)
            }
        }
    }
    return
}
