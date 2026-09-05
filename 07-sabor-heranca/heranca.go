package main

import "fmt"

type animal struct{
	nome string
	idade uint8
	agressivo bool
	racional bool
}

type humano struct{
	nome string
	idade uint8
}

type pessoa struct{
	humano
	hobby string
}

type programador struct{
	animal
	linguagem []string
	tomaCafe bool
	maluco bool
}

func main(){
	pessoaNormal := pessoa{
		humano: humano{
			nome: "José",
			idade: 30,
		},
		hobby: "Ir ao cinema",
	}

	luidson := programador{
		animal: animal{
			nome: "Luidson",
			idade: 153,
			agressivo: false,
			racional: true,
		},
		linguagem: []string{"PHP", "JS", "TS", "Go", "Python"},
		tomaCafe: false,
		maluco: true,
	}

	fmt.Println(pessoaNormal)
	fmt.Println(luidson)
}