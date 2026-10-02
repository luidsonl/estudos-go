package main

import "fmt"

func main(){
	pais := "brasil"
	estado := "bahia"
	
	var poste_mija_no_cachorro bool
	
	if pais == "brasil"{
		poste_mija_no_cachorro = true
	}else if pais == "argentina"{
		fmt.Println("faz frio")
		poste_mija_no_cachorro = true
	}else{
		poste_mija_no_cachorro = false
	}

	if poste_assalta_voce := (estado == "bahia"); poste_mija_no_cachorro {
		// variável poste_assalta_voce criada com if init fica restrita ao bloco de if
		fmt.Println("Aqui o poste mija no cachorro")
		if poste_assalta_voce {
			fmt.Println("E ainda assalta você")
		}
	}

}