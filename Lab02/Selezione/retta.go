package main

import "fmt"

func main() {

    var m, q, px, py float64
    
    fmt.Print("Inserisci m e q: ")
    fmt.Scan(&m, &q)
    
    fmt.Print("Inserisci x e y: ")
    fmt.Scan(&px, &py)
    
    if py == (m*px)+q {
    
        fmt.Println("Il punto appartiene alla retta")
    
    } else if (m*px)+q < py {
    
        fmt.Println("Il punto sta sopra la retta")
    
    } else {
    
        fmt.Println("Il punto sta sotto la retta")
    
    }
}
