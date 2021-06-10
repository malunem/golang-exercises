package main

import "fmt"
import "math/rand"
import "time"

func main() {

    var giocatori, turni int
    var nomi, vincitori []string
    
    fmt.Print("Inserisci il numero di giocatori: ")
    fmt.Scan(&giocatori)
    
    nomi = make([]string, giocatori)
    
    fmt.Printf("Inserisci i nomi dei %d giocatori:\n", giocatori)
    
    for i := 0; i < giocatori; i++ {
    
        fmt.Scanln(&nomi[i])
    }
    
    fmt.Print("Inserisci il numero di turni: ")
    fmt.Scan(&turni)
    
    vincitori = make([]string, turni)
    
    rand.Seed(int64(time.Now().Nanosecond()))
    
    for i := 0; i < turni; i++ {
    
        Turno(i, nomi, vincitori)
    }
    
    fmt.Println("Vittorie:")
    Vittorie(vincitori)
}

func Turno(turno int, nomi, vincitori []string) {

    fmt.Printf("=== Turno %d ===\n", turno+1)
    
    var somma int
    var vincitore string
    
    for i := range nomi {
        
        var n1, n2 = LanciaDadi()
        fmt.Printf("\t\tGiocatore %s, lanci di valore: %d e %d\n", nomi[i], n1, n2)
        
        if n1+n2 > somma {
            somma = n1+n2
            vincitore = nomi[i]
        }        
    }
    
    fmt.Printf("\t\tTurno %d, vincitore: %s\n", turno+1, vincitore)
    vincitori[turno] = vincitore
}

func LanciaDadi() (int, int) {
    
    var n1, n2 int
    n1 = rand.Intn(6) +1
    n2 = rand.Intn(6) +1
    
    return n1, n2
    
    //return rand.Intn(6), rand.Intn(6)
}

func Vittorie(vincitori []string) {

    for i := range vincitori {
       
        fmt.Printf("Vincitore turno %d: %s\n", i+1, vincitori[i])
    }
}
