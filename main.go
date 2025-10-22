package main

import (
	"fmt"
)

func main() {
	// Array bidimensional [6][4]float64
	notas := [6][4]float64{
		{8.5, 7.0, 9.0, 6.5},
		{9.0, 8.0, 7.5, 8.5},
		{6.0, 5.5, 7.0, 6.0},
		{10.0, 9.5, 9.0, 10.0},
		{7.0, 6.5, 8.0, 7.5},
		{5.5, 6.0, 5.0, 6.5},
	}

	var promedioGeneral float64
	fmt.Println("=== Análisis de Notas ===")

	for i := 0; i < len(notas); i++ {
		sliceNotas := notas[i][:] // Slice de las notas del estudiante
		var suma float64
		mayor := sliceNotas[0]
		menor := sliceNotas[0]

		for _, n := range sliceNotas {
			suma += n
			if n > mayor {
				mayor = n
			}
			if n < menor {
				menor = n
			}
		}

		promedio := suma / float64(len(sliceNotas))
		promedioGeneral += promedio

		fmt.Printf("Estudiante %d → Promedio: %.2f | Mayor: %.2f | Menor: %.2f\n", i+1, promedio, mayor, menor)
	}

	promedioGeneral /= 6
	fmt.Printf("\nPromedio general de la clase: %.2f\n", promedioGeneral)
}
