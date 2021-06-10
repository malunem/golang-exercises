package main

import "fmt"

func main() {

    var n int
    fmt.Print("Inserisci il numero dei gradini: ")
    fmt.Scan(&n)
    
    if n <= 0 {
            fmt.Println("Dimensione non sufficiente.")
            return
    } else {
            StampaScala(n)
    }
}

func StampaScala(gradini int) {

    for i := gradini; i >= 0; i-- {
    
        StampaGradino(i)
    }
}

func StampaGradino(gradino int) {

    if gradino < 0 {
            return
    } else if gradino == 0 {
            fmt.Print("***\n*\n")
    } else {
            
            for j := 0; j < 2; j++ {
                for i := gradino; i > 0; i-- {
                    fmt.Print("  ")                
                }
                
                if j == 0 {
                        fmt.Print("***\n")
                } else {
                        fmt.Print("*\n")
                }
            }
    }
}
