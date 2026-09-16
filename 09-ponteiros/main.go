package main

import(
	"fmt"
)

func main(){
	// Ponteiros referenciam valores de outras variáveis, apontando para o mesmo
	// endereço de memória
	// Em variáveis sem ponteiro
	var var1 int16 = 1
	var var2 int16 = var1

	fmt.Println("var1:", var1, ",", "var2:", var2)
	var2 ++
	fmt.Println("var1:", var1, ",", "var2:", var2)
	
	// Em variáveis com ponteiro
	var var3 int
	var ponteiro1 *int

	// O puteiro inicia como nil
	fmt.Println("var3:", var3, ",", "ponteiro1:", ponteiro1)

	var3 = 1
	ponteiro1 = &var3 // Vai receber o valor do endereço na memória
	fmt.Println("var3:", var3, ",", "ponteiro1:", ponteiro1)

	var3 ++
	// Desreferenciação do ponteiro é feita com *
	fmt.Println("var3:", var3, ",", "ponteiro1:", *ponteiro1)
}
