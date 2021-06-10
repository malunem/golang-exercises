package main

import (
        "fmt"
        "unicode"
        "os"
)

func main() {
    
    //Legge sequenza di stringhe da riga di comando
    parole := os.Args[1:]
    
    //Trasforma le parole alternando maiuscole e minuscole
    
    for indice, parola := range parole {
        
        fmt.Printf("%s ", TrasformaParola(parola, indice))
    }
    
    fmt.Printf("\n")
}

func TrasformaParola(parola string, posizione int) (parolaTrasformata string) {
    
    //Crea una slice di rune dalla stringa
    parolaRune := []rune(parola)
    
    //Scandisce i caratteri uno a uno
    for i, ch := range parolaRune {
    
        switch { 
        //Se parola ha indice pari inizia con maiuscola
        case posizione%2 == 0 && i%2 == 0:
            parolaTrasformata += string(unicode.ToUpper(ch))
        
        case posizione%2 == 0 && i%2 != 0:
            parolaTrasformata += string(unicode.ToLower(ch))
        
        //Se parola ha indice dispari inizia con minuscola 
        case posizione%2 != 0 && i%2 == 0:
            parolaTrasformata += string(unicode.ToLower(ch))  
            
        case posizione%2 != 0 && i%2 != 0:
            parolaTrasformata += string(unicode.ToUpper(ch))
        }
    }
    
    return parolaTrasformata 
}
