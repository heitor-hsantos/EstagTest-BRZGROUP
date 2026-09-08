package main

import (
	"fmt"
	"time"
)

type IntervaloAno struct {
	Ano        int
	DataInicio time.Time
	DataFim    time.Time
	TotalDias  int
}

func main() {

	anos := []int{2000, 2001, 1900, 1582}

	const mesInicio = time.February
	const diaInicio = 15

	const mesFim = time.October
	const diaFim = 15


	resultados := make([]IntervaloAno, len(anos))

	for i, ano := range anos {

		inicio := time.Date(ano, mesInicio, diaInicio, 0, 0, 0, 0, time.UTC)
		fim := time.Date(ano, mesFim, diaFim, 0, 0, 0, 0, time.UTC)


		dias := int(fim.Sub(inicio).Hours() / 24)

		resultados[i] = IntervaloAno{
			Ano:        ano,
			DataInicio: inicio,
			DataFim:    fim,
			TotalDias:  dias,
		}
	}


	for _, res := range resultados {
		fmt.Printf("Ano %d (%s até %s): %d dias\n",
			res.Ano,
			res.DataInicio.Format("02/01/2006"),
			res.DataFim.Format("02/01/2006"),
			res.TotalDias,
		)
	}
}
