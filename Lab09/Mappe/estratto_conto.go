package main

import (
        "fmt"
        "bufio"
        "os"
        "strconv"
        "strings"
        "sort"
)

type Operazione struct {
        data string
        tipo bool //true versamento+, false prelievo-
        importo float64
}

func main() {
    //fa indirezione di os.Args (nome file)
    dati := LeggiDati()
    //legge prima riga = saldo iniziale
    saldo, _ := strconv.ParseFloat(dati[0], 64)
    fmt.Printf("SALDO INIZIALE: %11.2f\n", saldo)
    
    operazioni := DividiDati(dati[1:])
    //trasforma data in formato gg-mm-aaaa
    for i := range operazioni {
        operazioni[i].data = FormattaData(operazioni[i].data)
    }
    elencoDate := CreaElencoDate(operazioni)
    
    for _, data := range elencoDate {
        saldo += StampaData(data, operazioni)
    }
    fmt.Printf("\n%-25s%11.2f\n", "SALDO FINALE:", saldo)
    fmt.Printf("%-25s===========\n", "")
}

func LeggiDati() (dati []string){
    scanner := bufio.NewScanner(os.Stdin)
    file := ""
    for scanner.Scan() {
        file += scanner.Text() + "\n"
    }
    riga := ""
    for _, v := range file {
        if v != '\n' {
            riga += string(v)
        } else {
            dati = append(dati, riga)
            riga = ""
        }
    }
    return
}

func DividiDati(righe []string) (operazioni []Operazione){
    for _, riga := range righe {
        tmp := strings.Split(riga, ";")
        var movimento Operazione
        movimento.data = tmp[0]
        tipo := tmp[1] == "V"
        movimento.tipo = tipo
        importo, _ := strconv.ParseFloat(tmp[2], 64)
        movimento.importo = importo
        operazioni = append(operazioni, movimento)
    }
    
    return
}


func FormattaData(data string) (dataFormattata string) {
    //usare strings split con / e poi strings join per riformare data giusta
    tmp := strings.Split(data, "/")
    if len(tmp[0]) != 4 {
        sposta := tmp[0]
        tmp[0] = tmp[2]
        tmp[2] = sposta
    }
    if len(tmp[1]) < 2 {
        tmp[1] = "0" + tmp[1]
    }
    if len(tmp[2]) < 2 {
        tmp[2] = "0" + tmp[2]
    }
    dataFormattata = strings.Join(tmp, "-")
    
    return
}

func CreaElencoDate(operazioni []Operazione) (elencoDate []string) {
    for i := range operazioni {
        esiste := false
        for j := range elencoDate {
            if operazioni[i].data == elencoDate[j] {
                esiste = true
            }
        }
        if !esiste {
            elencoDate = append(elencoDate, operazioni[i].data)
        }
    }
    sort.Strings(elencoDate)
    return
}

func StampaData(data string, operazioni []Operazione) (saldoGiornaliero float64) {
    fmt.Printf("\nDATA: ")
    tmp := strings.Split(data, "-")
    fmt.Printf("%s-%s-%s\n", tmp[2], tmp[1], tmp[0])
    for i := range operazioni {
        if operazioni[i].data != data {continue}
        if operazioni[i].tipo == true {
            fmt.Printf("%-25s%11.2f\n", "Versamento:", operazioni[i].importo)
            saldoGiornaliero += operazioni[i].importo
        } else {
            fmt.Printf("%-25s%11.2f\n", "Prelievo:", operazioni[i].importo)
            saldoGiornaliero -= operazioni[i].importo
        }
    }
    fmt.Printf("%-25s___________\n", "")
    fmt.Printf("%-25s%11.2f\n", "SALDO GIORNALIERO:", saldoGiornaliero)
    fmt.Printf("%-25s===========\n", "")
    
    return
}
