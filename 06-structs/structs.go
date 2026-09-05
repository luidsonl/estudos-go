package main

import "fmt"

type personagem struct {
	nome         string
	cromossomoY  bool
	idade        uint8
	força        uint8
	velocidade   uint8
	inteligencia uint8
	vitalidade   uint8
	mana         uint8
}

type item struct {
	nome       string
	preço      uint16
	quantidade uint8
}

type time struct {
	nome    string
	itens   []item
	membros []personagem
}

func main() {
	ana := personagem{
		nome:         "Ana",
		cromossomoY:  false,
		idade:        22,
		força:        3,
		velocidade:   6,
		inteligencia: 42,
		vitalidade:   20,
		mana:         40,
	}

	beatriz := personagem{
		nome:         "Beatriz",
		cromossomoY:  false,
		idade:        19,
		força:        21,
		velocidade:   34,
		inteligencia: 9,
		vitalidade:   24,
		mana:         0,
	}

	carlos := personagem{
		nome:         "Carlos",
		cromossomoY:  true,
		idade:        41,
		força:        40,
		velocidade:   29,
		inteligencia: 6,
		vitalidade:   36,
		mana:         0,
	}
	daniel := personagem{
		nome:         "Daniel",
		cromossomoY:  true,
		idade:        18,
		força:        20,
		velocidade:   24,
		inteligencia: 22,
		vitalidade:   23,
		mana:         18,
	}

	pocaoDeVida := item{
		nome:       "Poção de vida",
		preço:      20,
		quantidade: 8,
	}

	flecha := item{
		nome:       "Flecha",
		preço:      10,
		quantidade: 192,
	}

	t := time{
		nome:    "Time ABCD",
		membros: []personagem{ana, beatriz, carlos, daniel},
		itens:   []item{pocaoDeVida, flecha},
	}

	fmt.Println(t)
}
