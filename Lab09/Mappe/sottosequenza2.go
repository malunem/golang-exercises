package main

import (
        "fmt"
        "os"
        "strings"
       // "strconv"
)

func main() {
    //legge da riga di comando stringa di rune
    caratteri := strings.Join(os.Args[1:], " ")
    //una sottosequenza inizia e finisce con lo stesso carattere
    sottosequenze := TrovaSottosequenze(caratteri)
    //stampa le sottosequenze dalla più lunga 
    StampaSottosequenze(sottosequenze)
}

func TrovaSottosequenze(caratteri string) (sottosequenze map[string]int) {
    sottosequenze = make(map[string]int)
    for i, char1 := range caratteri {
        if char1 == ' ' {continue}
        for j, char2 := range caratteri {
            if char1 == char2 && i < j {
                sottosequenze[caratteri[i:j+1]]++
            }
        }
    }
    return
}

func StampaSottosequenze(sottosequenze map[string]int) {
    var maxLen int
    for i := range sottosequenze {
        if len(i) > maxLen {
            maxLen = len(i)
        }
    }
    for i := maxLen; i >= 5; i-- {
        for sottosequenza, occorrenze := range sottosequenze {
            if len(sottosequenza) == i {
                for _, v := range sottosequenza {
                    fmt.Printf("%c", v)
                }
                fmt.Printf(" -> Occorrenze: %d\n", occorrenze)
            }
        }
    }
    return
}
