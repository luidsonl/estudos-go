package main

import "fmt"

// O go só tem for
func main(){
	
	fmt.Println("Loop com for i < 10")
	i:= 0
	for i < 10{
		fmt.Println(i)
		i++
	}

	fmt.Println("Loop com for j := 0; j < 10; j++")
	for j:= 0; j < 10; j++{
		fmt.Println(j)
	}

	fmt.Println("Loop com for j := 0; j < 20; j += 2")
	for j:= 0; j < 20; j += 2{
		fmt.Println(j)
	}

	fmt.Println("Loop com for k, nome := range nomes")
	nomes := [3]string{"alexandre", "bruno", "carlos"}
	for k, nome := range nomes{
		println(k, ": ", nome)
	}

	fmt.Println("Loop com for _, nome := range nomes")
	for _, nome := range nomes{
		println(nome)
	}

	// Vai pegar os números da tabela asc, pra converter para letra precisa usar a função string
	fmt.Println("Loop com for k, letra := range \"nomes\"")
	for k, letra := range "nomes"{
		println(k, ": ", letra)
		println(k, ": ", string(letra))
	}

	// Para maps
	fmt.Println("Loop com for chave, valor := range user")
	user := map[string]string{
		"nome": "luidson",
		"priffisao": "programador",
		"nacionaliade": "brasileiro",
	}
	for chave, valor := range user{
		fmt.Println(chave, ": ", valor)
	}
}