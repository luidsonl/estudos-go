package main

import "fmt"

func main(){
	slice1 := make([]int16, 10, 11)
	fmt.Println("Slice com tamanho 11")
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// Quando o valor do slice vai estourar, o go dobra o valor do array e
	// libera mais espaço
	slice1 = append(slice1, 1)

	fmt.Println("Slice cheio")
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice1 = append(slice1, 1)

	fmt.Println("Tamanho estourou, o go dobra a memória reservada")
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// quando não passsamos o terceiro argumento ele define a capacidade como
	// o mesmo valor do tamanho
	fmt.Println("Inicialização sem passar capacidade")
	slice2 := make([]int16, 3)
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))
}