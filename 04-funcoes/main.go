package main

import (
	"fmt"
	"funcoes/matematica"
)

func main() {

	var num int64 = -3453
	sinal, absoluto := matematica.EhPositivo(num)
	fmt.Println(num)
	fmt.Println(absoluto, "| positivo?", sinal)

	num = 534345
	sinal, absoluto = matematica.EhPositivo(num)
	fmt.Println(num)
	fmt.Println(absoluto, "| positivo?", sinal)

	num = -1
	sinal, _ = matematica.EhPositivo(num)
	fmt.Println(num)
	fmt.Println("positivo?", sinal)
}
