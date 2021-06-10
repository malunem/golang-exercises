package main

import "fmt"

func main() {

    var n int
    
    fmt.Print("Inserisci la dimensione: ")
    fmt.Scan(&n)
    
    StampaScacchiera(n)
}

func StampaScacchiera(n int) {
    
    if n <= 0 {
        return
    }
    
    for i := 0; i < n; i++ {
    
        if i%2 == 0 {
            StampaRigaInizioAsterisco(n)
        } else {
            StampaRigaInizioPiù(n)
        }
    }
}

func StampaRigaInizioAsterisco(n int) {

    for i := 0; i <n; i++ {
    
        if i%2 == 0 {
            fmt.Print("*")
        } else {
            fmt.Print("+")
        }
    }
    fmt.Println()
}

func StampaRigaInizioPiù(n int) {

    for i := 0; i <n; i++ {
    
        if i%2 == 0 {
            fmt.Print("+")
        } else {
            fmt.Print("*")
        }
    }
    fmt.Println()
}
