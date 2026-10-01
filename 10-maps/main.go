package main

import (
	"fmt"
)

func main(){
	luidson := map[string]string{
		"nome": "luidson",
		"poder": "fazer códigos",
		"cor_preferida": "transparente",
	}

	fmt.Println(luidson["cor_preferida"])

	tablet:= map[string]map[string]string{
		"software": {
			"os": "android",
			"os_version": "12",
		},
		"hardware": {
			"ram": "8",
			"processador": "mediatech",
			"tamanho_tela": "12",
		},
	}

	fmt.Println(tablet)
	delete(tablet, "software")
	fmt.Println(tablet)
}
