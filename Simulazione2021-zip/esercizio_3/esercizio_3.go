package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Utenza = struct {
	numero_telefono string
	codice_sim      string
}

type RegistroTelefonico = map[string][]string

func main() {

	utenzeInput := LeggiUtenze()

	registro := InizializzaRegistro()

	for _, utenza := range utenzeInput {
		registro = AggiungiUtenza(registro, utenza)
	}

	for telefono := range registro {

		//fmt.Print(telefono[:4], "hop\n")
		if telefono[:3] == "338" {
			sim := Identifica(registro, telefono)

			fmt.Printf("Il numero %s è associato alla sim %s\n", telefono, sim)
		}
	}

}

func LeggiUtenze() (utenze []Utenza) {

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		line := scanner.Text()
		data := strings.Split(line, ";")

		var utenza = Utenza{
			data[0],
			data[1],
		}

		utenze = append(utenze, utenza)
	}

	return utenze
}

func InizializzaRegistro() (registro RegistroTelefonico) {

	registro = make(RegistroTelefonico)

	return registro
}

func AggiungiUtenza(registro RegistroTelefonico, utenza Utenza) (registroAggiornato RegistroTelefonico) {

	registroAggiornato = registro

	registroAggiornato[utenza.numero_telefono] = append(registroAggiornato[utenza.numero_telefono], utenza.codice_sim)

	return registroAggiornato
}

func Identifica(registro RegistroTelefonico, telefono string) (codiceSIM string) {

	numeriPresenti := len(registro[telefono])

	if numeriPresenti == 0 {
		codiceSIM = ""
	} else {
		codiceSIM = registro[telefono][numeriPresenti-1]
	}

	return codiceSIM
}
