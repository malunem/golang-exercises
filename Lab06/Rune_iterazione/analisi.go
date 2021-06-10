package main

import "fmt"

func main() {

    var stringa string
    fmt.Scan(&stringa)
    
    var maiuscole, minuscole, consonanti, vocali int
    
    for _, c := range stringa {
        
        if !èLetteraValida(c) {
            continue
        }
        
        if èMaiuscola(c) {
            
            maiuscole += 1
            
        } else {
           
           minuscole += 1
            
        }
            
        if èVocale(c) {
            
            vocali += 1
            
        } else {
            
            consonanti += 1
            
        }
    }
    
    fmt.Println("Maiuscole:", maiuscole)
    fmt.Println("Minuscole:", minuscole)
    fmt.Println("Vocali:", vocali)
    fmt.Println("Consonanti:", consonanti)
}

func èLetteraValida(c rune) bool {
    
    return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')
}
func èMaiuscola(c rune) bool {
    
    return c >= 'A' && c <= 'Z'
}
func èVocale(c rune) bool {

    return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
            c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
}
