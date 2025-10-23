package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("INGRESA UN TEXTO A INDAGAR: ")
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	TEXTO := strings.ToLower(sc.Text())

	TEXTO = strings.ReplaceAll(TEXTO, ".", "")
	PALABRAS := strings.Fields(TEXTO)

	CONTADOR := map[string]int{}
	for _, p := range PALABRAS {
		CONTADOR[p]++
	}

	for p, c := range CONTADOR {
		fmt.Println(p, ":", c)
	}

	MAXIMO := 0
	var PALABRAMASREPETIDA []string
	for p, c := range CONTADOR {
		if c > MAXIMO {
			MAXIMO = c
			PALABRAMASREPETIDA = []string{p}
		} else if c == MAXIMO {
			PALABRAMASREPETIDA = append(PALABRAMASREPETIDA, p)
		}
	}

	fmt.Println("\nPALABRA MAS REPETIDA:", PALABRAMASREPETIDA, "(", MAXIMO, "VECES )")
}
