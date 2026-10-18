package urna

import (
	"bufio"
	"fmt"
)

func Votar(urna *[]Voto, scanner *bufio.Scanner){
	var err error
	var numero string
	var voto Voto


	fmt.Println("Insira o voto para escolher o presidente que irá te roubar pelos próximos 4 anos")
	scanner.Scan()
	numero = scanner.Text()
	err= computarVoto(numero, &voto, PRESIDENTE)
	if err != nil{
		fmt.Println("formato inválido")
		return
	}

	fmt.Println("Insira o voto para escolher o governador que irá te roubar pelos próximos 4 anos")
	scanner.Scan()
	numero = scanner.Text()
	err = computarVoto(numero, &voto, GOVERNADOR)
	if err != nil{
		fmt.Println("formato inválido")
		return
	}
	
	fmt.Println("Insira o voto para escolher o senador ladrão que irá te roubar pelos próximos 4 anos")
	scanner.Scan()
	numero = scanner.Text()
	err = computarVoto(numero, &voto, SENADOR)
	if err != nil{
		fmt.Println("formato inválido")
		return
	}

	fmt.Println("Insira o voto para escolher o outro senador ladrão que irá te roubar pelos próximos 4 anos")
	scanner.Scan()
	numero = scanner.Text()
	err = computarVoto(numero, &voto, SENADOR2)
	if err != nil{
		fmt.Println("formato inválido")
		return
	}

	fmt.Println("Insira o voto para escolher o deputado federal ladrão que irá te roubar pelos próximos 4 anos")
	scanner.Scan()
	numero = scanner.Text()
	err = computarVoto(numero, &voto, DEP_FEDERAL)
	if err != nil{
		fmt.Println("formato inválido")
		return
	}

	fmt.Println("Insira o voto para escolher o deputado estadual ladrão que irá te roubar pelos próximos 4 anos")
	scanner.Scan()
	numero = scanner.Text()
	err = computarVoto(numero, &voto, DEP_ESTADUAL)
	if err != nil{
		fmt.Println("formato inválido")
		return
	}

	*urna = append(*urna, voto)
}