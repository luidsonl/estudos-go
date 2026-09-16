package main

import "fmt"

func main(){
	// Array tem tamanho definido
	var array1 [3]string
	// Ao declarar um array antes, para atribuir os valores nas posições é preciso
	// Atribuir cada posição individualmente
	array1[0] = "ana"
	array1[1] = "beatriz"
	array1[2] = "carla"
	// array1[3] = "daniela"
	// Atribuir valores não reservados na memória causa erro

	fmt.Println(array1)

	// É possível assinalar valores em laços de repetição
	var array2 [3]int16

	for i := 0; i < len(array2); i++{
		array2[i] = int16(i * i)
	}

	fmt.Println(array2)

	// Também é possível atribuir valores diretamente na decaração do array com inferência
	// de tipos

	array3 := [3]string{"amanda", "bianca", "capitulina"}
	//array3 = [3]string{"amanda", "bianca", "capitulina", "dionésia"}
	// O tamanho do array é rígido na hora da declaração
	fmt.Println(array3)
	// O tamanho do array pode pode ser inferido com ...
	array4 := [...]int32{1,2,3,4,5,6,7,8}
	fmt.Println(array4)
	//array4[8] = 9
	// O espaço reservado na memória é fixo

	// Slices
	// Não tem tamanho fixo para serem declarados, são fatias de arrays
	// São declarados passando o [] vazio
	var slice1  []string
	//slice1[0] = "anabel"
	//slice1[2] = "bruna"
	// Isso dá erro no slice, para adicionar valores é necessário usar o append

	slice1 = append(slice1, "anabel")
	slice1 = append(slice1, "bruna")
	fmt.Println(slice1)

	// Slice pode ter o valor declarado na inicialização
	slice2 := []int{1,2,3,4,5}
	fmt.Println(slice2)

	// Ele serve como um ponteiro para o array
	slice3 := array3[0:2]
	fmt.Println(slice3)
	array3[1] = "beatriz"
	fmt.Println(slice3)
	slice3[1] = "benedita"
	fmt.Println(slice3)

	// Nesse caso o slice está apontando para os mesmos endereços de memória
	// do array
}