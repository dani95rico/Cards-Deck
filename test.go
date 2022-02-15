package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	b := newDeck()
	if len(b) != 16 {
		t.Errorf("\nERROR: la longitud de la baraja no es correcta.\nLa longitud debería ser 16 y en este caso es: %v.", len(b))
	}

	if b[0] != "1 de Espadas" || b[len(b)-1] != "4 de Tréboles" {
		t.Errorf("\nERROR: el orden de la baraja no es la correcta.")
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	b := newDeck()
	b.saveToFile("_decktesting")

	loadedb := newDeckFromFile("_decktesting")

	if len(loadedb) != 16 {
		t.Errorf("\nERROR: la longitud de la baraja no es correcta.\nLa longitud debería ser 16 y en este caso es: %v.", len(b))
	}

	os.Remove("_decktesting")
}
