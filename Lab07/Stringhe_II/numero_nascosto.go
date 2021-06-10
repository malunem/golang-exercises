package main

import (
        "fmt"
        "bufio"
        "os"
        "strconv"
        "unicode"
)

func main() {
    
    num, err := NumeroNascosto(LeggiTesto())
    
    if err == nil {
    
        fmt.Printf("Doppio del numero nascosto: %d (%d * 2)\n", num*2, num)
    } else {
        fmt.Println("nessun numero nascosto\n")
    }
}

func LeggiTesto() string {

    var scanner *bufio.Scanner
    var testo string
    
    scanner = bufio.NewScanner(os.Stdin)
        
    if scanner.Scan() {
    
        testo = scanner.Text()
    }
    
    return testo
}

func NumeroNascosto(testo string) (int, error) {

    var num int
    var err error
    var numeri string
    
    for i := range testo {
    
        if unicode.IsDigit(rune(testo[i])) {
            
            numeri += string(testo[i])
        }
    }
    
    num, err = strconv.Atoi(numeri)
            
    return num, err
}
