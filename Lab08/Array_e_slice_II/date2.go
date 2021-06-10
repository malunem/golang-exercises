package main

import (
        "fmt"
        "os"
        "bufio"
        "unicode"
        "sort"
)

func main() {
    
    var date []string
    
    scanner := bufio.NewScanner(os.Stdin)
    
    for scanner.Scan() {
    
        if scanner.Text() == "" {
                break
        } else {
                date = append(date, scanner.Text())
        }
    }
    
    //inverte le date con prima i giorni
    for i, data := range date {
        
        if daInvertire(data) {
            date[i] = Inverti(data)
        }
    }
    
    //formatta le date
    for i, data := range date {
        FormattaData(i, data, date)
    }
    
    //le stampa in ordine cronologico
    sort.Strings(date)
    
    for _, data := range date {
    
        fmt.Printf("%s\n", data)
    }
}

func daInvertire(data string) bool {
    
    for _, c := range data[:4] {
            
        if c == '/' {
            return true
        }
    }
    
    return false
}

func Inverti(data string) string {
    
    var dataInvertita string
    var segmento string
    
    for _, c := range data {
        
        if c != '/' {
            
            segmento += string(c)
            
        } else {
            
            dataInvertita = "/" + segmento + dataInvertita
            
            segmento = ""
        }
    }
    
    dataInvertita = segmento + dataInvertita
    
    return dataInvertita
}

func FormattaData(i int, data string, date []string) {
    
    date[i] = data[:4] + "-"
        
        lunghezza := len(data[5:])
        
        switch lunghezza {
            
            case 3: 
                date[i] += "0" + string(data[5]) + "-0" + string(data[7])
            
            case 4:
                if unicode.IsDigit(rune(data[5])) && unicode.IsDigit(rune(data[6])) {
                    
                    date[i] += string(data[5:7]) + "-" + "0" + string(data[8])
                } else {
                    date[i] += "0" + string(data[5]) + "-" + string(data[7:9])
                }
            
            case 5:
                date[i] += data[5:7] + "-" + data[8:10]
        }
}
