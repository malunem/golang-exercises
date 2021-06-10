package main

import (
        "fmt"
        "os"
        "strconv"
)

func main() {
    //legge da os.Args un int
    v, _ := strconv.Atoi(os.Args[1])
    matrice := CreaTavolaPitagorica(v) //crea tavola pitagorica n*n
    
    
    //stampa tavola pitagorica
    StampaTavolaPitagorica(matrice)
     
}

func CreaTavolaPitagorica(n int) ([][]int) {
    //crea la prima riga e la prima colonna
    matrice := make([][]int, n)
    
    for i := 0; i < n ; i++ {
        matrice[i] = make([]int, n)
        matrice[i][0] = i+1
        matrice[0][i] = i+1
    }
    
    //stampa nei valori dalla riga2col2 in poi la moltiplicazione tra primo e secondo valore della matrice
    for i := 1; i < n; i++ {
        for j := 1; j < n; j++ {
            matrice[i][j] = matrice[i][0] * matrice[0][j]
        }
    }
    //restituisce matrice 
    return matrice   
}

func StampaTavolaPitagorica(s [][]int) {
    //due cicli for annidati che stampano
    for riga := range s {
        for colonna := range s[0] {
            fmt.Printf("%2d ", s[riga][colonna])
        }
        fmt.Println()
    }
    return
}
