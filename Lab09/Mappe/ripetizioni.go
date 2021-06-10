package main

import (
        "fmt"
        "bufio"
        "os"
        "unicode"
)

func main() {
    //legge da os.Stdin con bufio fino EOF
    testo := LeggiTesto()
    //conta numero di parole 
    parole := SeparaParole(testo)
    
    //conta ripetizioni parole
    ripetizioni := ContaRipetizioni(parole)
    contatore := len(ripetizioni)
    fmt.Printf("Parole distinte: %d\n", contatore)
    
    for i, v := range ripetizioni {
        fmt.Printf("%s: %d\n", i, v)
    }
}

func LeggiTesto() string {
    testo := ""
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        testo += scanner.Text() + "\n"
    }
    return testo
}

func SeparaParole(s string) []string {
    //restituisce slice di parole presenti
    precedente := false
    parola := ""
    var parole []string
    for _, v := range s {
        lettera := unicode.IsLetter(v)
        switch {
            case lettera && !precedente:
                parola += string(v)
                precedente = true
            case lettera && precedente:
                parola += string(v)
            case !lettera && precedente:
                parole = append(parole, parola)
                parola = ""
                precedente = false
            case !lettera && !precedente:
                continue
        }
    }
    return parole
}

func ContaRipetizioni(sl []string) map[string]int {
    ripetizioni := map[string]int{}
    for _, v := range sl {
        ripetizioni[v]++
    }
    return ripetizioni
}    
