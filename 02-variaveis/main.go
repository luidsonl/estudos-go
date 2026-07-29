package main

import "fmt"

func main() {
	var declaracaoExplicita string = "Declaracao Explicita"
	declaracaoImplicita := "Declaração Implícita"

	var (
		multiExplicita1 bool    = true
		multiExplicita2 int32   = 666
		multiExplicita3 string  = "Iron Maiden"
		multiExplicita4 float32 = 3.14
	)

	multiInferida1, multiInferida2, multiInferida3, multiInferida4 := true, 666, "Iron Maiden", 3.14

	fmt.Println(declaracaoExplicita, declaracaoImplicita)

	fmt.Println(multiExplicita1, multiExplicita2, multiExplicita3, multiExplicita4)

	fmt.Println(multiInferida1, multiInferida2, multiInferida3, multiInferida4)

	const constante string = "constante"
	const (
		constante1 string = "Constante 1"
		constante2 int    = 40028922
	)

	fmt.Println(constante, constante1, constante2)

	// Inverter variáveis
	fmt.Println(declaracaoExplicita, declaracaoImplicita)
	declaracaoImplicita, declaracaoExplicita = declaracaoExplicita, declaracaoImplicita
	fmt.Println(declaracaoExplicita, declaracaoImplicita)

}
