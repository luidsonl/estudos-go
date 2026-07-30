package main

import (
	"errors"
	"fmt"
	"unsafe"
)

func main() {
	fmt.Printf("int possui %d bytes\n", unsafe.Sizeof(int(0)))
	fmt.Printf("int possui %d bits\n\n", unsafe.Sizeof(int(0))*8)

	// ---------------------------
	// Valor zero (Zero Value)
	// ---------------------------

	var zeroBool bool
	var zeroInt int
	var zeroFloat float64
	var zeroString string
	var zeroRune rune
	var zeroPointer *int
	var zeroSlice []int
	var zeroMap map[string]int
	var zeroError error

	fmt.Println("=== Valores Zero ===")
	fmt.Printf("bool: %v\n", zeroBool)
	fmt.Printf("int: %v\n", zeroInt)
	fmt.Printf("float64: %v\n", zeroFloat)
	fmt.Printf("string: %q\n", zeroString)
	fmt.Printf("rune: %v\n", zeroRune)
	fmt.Printf("ponteiro: %v\n", zeroPointer)
	fmt.Printf("slice: %v\n", zeroSlice)
	fmt.Printf("map: %v\n", zeroMap)
	fmt.Printf("error: %v\n\n", zeroError)

	// Booleano
	var ativo bool = true

	// Inteiros
	var idade int = 25
	var temperatura int8 = -10
	var populacao uint64 = 8000000000

	// Ponto flutuante
	var altura float32 = 1.75
	var pi float64 = 3.1415926535

	// Números complexos
	var c complex64 = 3 + 4i

	// Texto
	var nome string = "Luidson"

	// Rune (caractere Unicode)
	var letra rune = 'A'

	// Byte (uint8)
	var b byte = 255

	// Array
	numeros := [5]int{10, 20, 30, 40, 50}

	// Slice
	frutas := []string{"Maçã", "Banana", "Uva"}

	// Map
	idades := map[string]int{
		"Ana":   20,
		"Bruno": 30,
	}

	// Struct
	type Pessoa struct {
		Nome  string
		Idade int
	}

	pessoa := Pessoa{
		Nome:  "Carlos",
		Idade: 40,
	}

	// Ponteiro
	x := 100
	ptr := &x

	// Interface
	var valor any = 123

	// Error
	var err error = errors.New("Tipo erro")

	fmt.Println("=== Tipos de Dados ===")
	fmt.Println("Bool:", ativo)
	fmt.Println("Int:", idade)
	fmt.Println("Int8:", temperatura)
	fmt.Println("Uint64:", populacao)
	fmt.Println("Float32:", altura)
	fmt.Println("Float64:", pi)
	fmt.Println("Complex64:", c)
	fmt.Println("String:", nome)
	fmt.Println("Rune:", letra)
	fmt.Println("Rune Convertida para ASC:", string(letra))
	fmt.Println("Byte:", b)
	fmt.Println("Array:", numeros)
	fmt.Println("Slice:", frutas)
	fmt.Println("Map:", idades)
	fmt.Println("Struct:", pessoa)
	fmt.Println("Ponteiro:", *ptr)
	fmt.Println("Interface (any):", valor)
	fmt.Println("Error:", err)
}
