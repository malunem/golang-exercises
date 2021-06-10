package main

import "fmt"
import "unicode"

func main() {

    var stringa string
    fmt.Scan(&stringa)
    
    var consonantiMaiuscole, consonantiMinuscole, vocaliMaiuscole, vocaliMinuscole int
    
    for _, c := range stringa {
        
        if !unicode.IsLetter(c) {
            continue
        }
        
        if unicode.IsUpper(c) {
            
            if èVocale(c) {    
                    vocaliMaiuscole += 1
            } else {
                    consonantiMaiuscole += 1
            }
            
        } else {
            if èVocale(c) {
                    vocaliMinuscole += 1
            } else {
                    consonantiMinuscole += 1
            }
           
        }
    }
    
    fmt.Println("Vocali maiuscole:", vocaliMaiuscole)
    fmt.Println("Consonanti maiuscole:", consonantiMaiuscole)
    fmt.Println("Vocali minuscole:", vocaliMinuscole)
    fmt.Println("Consonanti minuscole:", consonantiMinuscole)
}

func èVocale(c rune) bool {

    switch c {
    
        case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
            return true
            
        default:
            return false
    }
}
