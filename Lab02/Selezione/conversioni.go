package main

import "fmt"

func main() {

    var (
        conversione int
        valore float64
        )
    
    fmt.Print("Scegli la conversione:\n" +
              "1) secondi -> ore\n" +
              "2) secondi -> minuti\n" +
              "3) minuti -> ore\n" +
              "4) minuti -> secondi\n" +
              "5) ore -> secondi\n" +
              "6) ore -> minuti\n" +
              "7) minuti -> giorni e ore\n" +
              "8) minuti -> anni e giorni\n")
    
    fmt.Scan(&conversione)
    
    if conversione < 1 || conversione > 8 {
    
        fmt.Println("Scelta errata")
        return
    }
    
    fmt.Print("Inserisci il valore da convertire: ") 
    fmt.Scan(&valore)
  
    switch conversione {
    
        case 1:
            ore := (valore / 60) / 60
            fmt.Print(valore, " secondi corrispondono a ", ore, " ore\n")
        
        case 2:
            minuti := valore / 60
            fmt.Print(valore, " secondi corrispondono a ", minuti, " minuti\n")
        
        case 3:
            ore := valore / 60
            fmt.Print(valore, " minuti corrispondono a ", ore, " ore\n")
            
        case 4:
            secondi := valore * 60
            fmt.Print(valore, " minuti corrispondono a ", secondi, " secondi\n")
        
        case 5:
            secondi := valore * 60 * 60
            fmt.Print(valore, " ore corrispondono a ", secondi, " secondi\n")
        
        case 6:
            minuti := valore * 60
            fmt.Print(valore, " ore corrispondono a ", minuti, " minuti\n")
        
        case 7:
            giorni := (valore / 60) / 24
            //ore := (valore - float64((int(giorni) * 60 * 24))) / 60
            
            ore := (giorni - float64(int(giorni))) * 24
            fmt.Print(valore, " minuti corrispondono a ", int(giorni), " giorni" +
                        " e ", ore, " ore\n")
        
        case 8:
            anni := (((valore / 60) / 24) / 365)
            giorni := (anni - float64(int(anni))) * 365
            fmt.Print(valore, " minuti corrispondono a ", int(anni), " anni" +
                        " e ", giorni, " giorni\n")
        
        default: 
            
    }  
}
