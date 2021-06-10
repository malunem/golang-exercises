package main 

import (
        "fmt"
        "bufio"
        "os"
)

func main() {

    var testo string
    testo = LeggiTesto()
    testo = Garibaldi(testo)
    
    fmt.Printf("Risultato trasformazione:%s\n", testo)
}

func LeggiTesto() string {

    var testo string
    
    fmt.Println("Inserisci un testo su più righe (termina con riga vuota):")
    
    scanner := bufio.NewScanner(os.Stdin)
    
    for scanner.Scan() {
        
        if scanner.Text() == "" {
            break
        } else {
            testo += "\n" + scanner.Text()
        }
    }

    return testo 
}

func Garibaldi(s string) string {

    //usa TrasformaCarattere
    var garibaldi string
    
    for _, c := range s {
        
        garibaldi += string(TrasformaCarattere(c))
    }
    
    return garibaldi
}

func TrasformaCarattere(c rune) rune {

    //trasforma tutte le vocali in u
    
    switch c {
            
        case 'a', 'e', 'i', 'o':
                
            return 'u'
                
        case 'A', 'E', 'I', 'O':
            
            return 'U'
            
        default:
                
            return c
    }
}
