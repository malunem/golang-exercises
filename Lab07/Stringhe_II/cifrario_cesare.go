package main

import (
        "fmt"
        "bufio"
        "os"
        "unicode"
)

func main() {
    var k int //chiave di cifratura (shift lettere)
    fmt.Print("Inserisci un numero: ")
    k = LeggiNumero() //da standard input
    
    var messaggio string
    fmt.Print("Inserisci un testo su più righe (termina con ctrl d):\n")
    messaggio = LeggiTesto() //da standard input su + righe con EOF
    
    var cifrato string
    cifrato = CifraTesto(messaggio, k) //usa CifraCarattere(rune, int) rune
    
    fmt.Println(cifrato)
}

func LeggiNumero() (n int) {
    
    fmt.Scan(&n)
    return
}

func LeggiTesto() (messaggio string) {
    
    scanner := bufio.NewScanner(os.Stdin)
    
    for scanner.Scan() {
        
        messaggio += scanner.Text() + "\n"
    }
    
    return messaggio
}

func CifraTesto(messaggio string, chiave int) (cifrato string) {
    
    for _, c := range messaggio {
        
        cifrato += string(CifraCarattere(c, chiave))
    }
    
    return cifrato
}

func CifraCarattere(carattere rune, chiave int) rune {
    
    const alfabeto int32 = 'Z' - 'A' + 1
    
    var key rune = int32(chiave) %alfabeto
    
    if key < 0 {
        key += alfabeto
    }
    
    var primo rune
    
    if carattere >= 'A' && carattere <= 'Z' || carattere >= 'a' && carattere <= 'z' {
        
        if unicode.IsLower(carattere) {
            primo = 'a'
        } else {
            primo = 'A'
        }
        
        return (((carattere - primo) + key) % alfabeto) + primo
    
    } else {
        
        return carattere
    }
    
    /* NON FUNZIONA BOH
    
    output := int(carattere)+chiave
    var a, z, A, Z int = 'a', 'z', 'A', 'Z'
    
    if output > z && output < A {
        
        output = (int(carattere)-z)+a
        
    } else if output > Z {
        
        output = (int(carattere)-Z)+A
    }
    
    return rune(output)
    */
}
