package main

import (
	"fmt"
	"votacao/urna"
	"bufio"
	"os"
)


func main() {
	
	urnaData := make([]urna.Voto, 0)
	running := true
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("Bem vindo a festa da democracia! Vamos votar e delegar seu destino a algum desgraçado ladrão, mentiroso e autoritário! :) Pagar imposto é tão bom! Confie nas instituições ;D")
	
	for running{
		urna.Votar(&urnaData, scanner)

		fmt.Println("Obrigado por participar da democracia! Lembre-se de pagar seus impostos e proteger a pátria quando tiver alguma guerra :D")
		fmt.Println("Votos computados")
		for i, voto := range(urnaData){
			fmt.Printf("Voto %d: -------------------------- \n ", i)
			fmt.Println(voto)
		}

		fmt.Println("Deseja votar novamente? Precione S se sim")

		scanner.Scan()
		running = scanner.Text() == "s" 

	}
	
	
}




