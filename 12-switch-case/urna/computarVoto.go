package urna

import (
	"strconv"
)

func computarVoto(numeroString string, voto *Voto, cargo string)(error) {
	var candidato Candidato
	numero, err := strconv.Atoi(numeroString)

	if err != nil{
		return err
	}

	switch numero {
	case 13:
		candidato.nome = "Bosta"

	case 22:
		candidato.nome = "Merda"

	default:
		candidato.nome = "Cocô"
	}
	candidato.cargo = cargo
	candidato.numero = numero
	switch cargo {
	case PRESIDENTE:
		voto.presidente = candidato
	case GOVERNADOR:
		voto.governador = candidato
	case SENADOR:
		voto.senador = candidato
	case SENADOR2:
		voto.senador2 = candidato
	case DEP_FEDERAL:
		voto.depFederal = candidato
	case DEP_ESTADUAL:
		voto.depEstadual = candidato
	}
	return nil
}