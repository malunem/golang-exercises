package main

import "fmt"

func main() {
    
    var s1, s2 string
    fmt.Scan(&s1, &s2)
    
    fmt.Println(StringheAlternate(s1, s2))
}

func StringheAlternate(stringa1, stringa2 string) (risultato string) {
    
    var stringaMax string
    
    if len(stringa1) > len(stringa2) {
        stringaMax = stringa1
    } else {
        stringaMax = stringa2
    }
    
    for i, _ := range stringaMax {
        
        if i >= len(stringa1) {
            risultato += "-"
        } else {
            risultato += string(stringa1[i])
        }
        
        if i >= len(stringa2) {
            risultato += "-"
        } else {
            risultato += string(stringa2[i])
        }
    }
    
    return risultato
}
