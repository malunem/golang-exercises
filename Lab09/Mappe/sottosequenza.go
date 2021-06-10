package main

import (
        "fmt"
        "os"
        "strings"
       // "strconv"
)

func main() {
    //legge da riga di comando stringa di int
    numeri := strings.Join(os.Args[1:], " ")
    //una sottosequenza inizia e finisce con lo stesso numero
    sottosequenze := TrovaSottosequenze(numeri)
    //stampa le sottosequenze dalla più corta alla più lunga
    StampaSottosequenze(sottosequenze)
}

func TrovaSottosequenze(numeri string) (sottosequenze map[string]int) {
    sottosequenze = make(map[string]int)
    for i, numero1 := range numeri {
        if numero1 == ' ' {continue}
        for j, numero2 := range numeri {
            if numero1 == numero2 && i < j {
                sottosequenze[numeri[i:j+1]] = len(numeri[i:j+1])
            }
        }
    }
    return
}

func StampaSottosequenze(sottosequenze map[string]int) {
    var maxLen int
    for _, l := range sottosequenze {
        if l > maxLen {
            maxLen = l
        }
    }
    for i := 2; i <= maxLen; i++ {
        for sottosequenza, l := range sottosequenze {
            if l == i {
                for _, v := range sottosequenza {
                    fmt.Printf("%c", v)
                }
                fmt.Println()
            }
        }
    }
    return
}
