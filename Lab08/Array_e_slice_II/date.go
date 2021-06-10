package main

import (
        "fmt"
        "bufio"
        "os"
)

func main() {
    
    date := LeggiTesto() //da standard input
    
    for _, c := range date {
        
        fmt.Println(Formatta(c))
    }
}

func LeggiTesto() []string {
    //interrompere con riga vuota
    //legge una data per ogni riga
    
    var date []string
    
    scanner := bufio.NewScanner(os.Stdin)
    
    for scanner.Scan() {
        
        riga := scanner.Text()
    
        if riga == "" {
            break
        }
        
        date = append(date, riga)
    }
    
    return date
    
}
    
func Formatta(data string) string {

    //restituisce la data scritta aaaa-mm-gg
    var dataFormattata []rune
    dataFormattata = []rune(data)
    
    lunghezza := len(dataFormattata)
    
    if lunghezza == 8 {
        
        dataFormattata = append(dataFormattata[:4], '-', '0', dataFormattata[5], '-', '0', dataFormattata[7])
    
    } else if lunghezza == 9 {
        
        if dataFormattata[6] == '/' {
            
            dataFormattata = append(dataFormattata[:4], '-', '0', dataFormattata[5], '-', dataFormattata[7], dataFormattata[8])
        
        } else {
            
            dataFormattata = append(dataFormattata[:4], '-', dataFormattata[5], dataFormattata[6], '-', '0', dataFormattata[8])
        }
    } else {
    
        dataFormattata[4] = '-'
        dataFormattata[7] = '-'
    }
    
    return string(dataFormattata)
}
