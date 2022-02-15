package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type baraja []string

func newDeck() baraja {
	cartas := baraja{}
	tipoCartas := []string{"Espadas", "Diamantes", "Corazones", "Tréboles"}
	numCartas := []string{"1 ", "2 ", "3 ", "4 "}

	for _, tipo := range tipoCartas {
		for _, numero := range numCartas {
			cartas = append(cartas, numero+"de "+tipo)
		}
	}
	return cartas
}

func (b baraja) print() {
	for i, carta := range b {
		fmt.Println(i, carta)
	}
}

func repartir(b baraja, handSize int) (baraja, baraja) {
	return b[:handSize], b[handSize:]
}

func (b baraja) toString() string {
	return strings.Join([]string(b), ",")
}

func (b baraja) saveToFile(nombre string) error {
	return ioutil.WriteFile(nombre, []byte(b.toString()), 0666)
}

func newDeckFromFile(nombre string) baraja {
	bytes, error := ioutil.ReadFile(nombre)
	if error != nil {
		fmt.Print(error)
		os.Exit(1)
	}
	separacion := strings.Split(string(bytes), ",")
	return baraja(separacion)
}

func (b baraja) barajear() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range b {
		newPosition := r.Intn(len(b) - 1)
		b[i], b[newPosition] = b[newPosition], b[i]
	}

}
