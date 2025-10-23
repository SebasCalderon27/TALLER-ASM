package main

import "fmt"

func main() {
	NOTASCURSO := [6][4]float64{
		{8.2, 9.6, 7.6, 10},
		{6.6, 5.9, 8.5, 7.2},
		{9.2, 9.5, 8.7, 10},
		{5.4, 6.7, 5.9, 4.8},
		{7.2, 8.4, 6.9, 7.2},
		{10, 9.4, 9.2, 10},
	}

	var TOTALCURSO float64

	for i := 0; i < 6; i++ {
		SUMA := 0.0
		MAYORNOTA := NOTASCURSO[i][0]
		MENORNOTA := NOTASCURSO[i][0]

		for j := 0; j < 4; j++ {
			if NOTASCURSO[i][j] > MAYORNOTA {
				MAYORNOTA = NOTASCURSO[i][j]
			}
			if NOTASCURSO[i][j] < MENORNOTA {
				MENORNOTA = NOTASCURSO[i][j]
			}
			SUMA += NOTASCURSO[i][j]
		}

		PROMEDIO := SUMA / 4
		TOTALCURSO += PROMEDIO
		fmt.Println("ESTUDIANTE", i+1, "PROMEDIO:", PROMEDIO, "MAYOR NOTA:", MAYORNOTA, "MENOR NOTA:", MENORNOTA)
	}

	fmt.Println()
	fmt.Println("PROMEDIO GENERAL:", TOTALCURSO/6)
}
