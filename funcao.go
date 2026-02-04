package main

import (
	"fmt"
	"treinamento/deferpkg"
	"treinamento/recursao"
)

func media(lista []float64) float64 {
	total := 0.0
	for _, nota := range lista {
		total += nota
	}
	return total / float64(len(lista))
}

func main() { //programa que calcula a media de uma sala
	lista := []float64{98, 93, 77, 82, 83} //lista de alunos
	fmt.Println("A média da sala é:", media(lista))
	fmt.Println("Fatorial de 2:", recursao.Fatorial(2))
	fmt.Println("==============Defer==============")
	defer deferpkg.Dia2()
	deferpkg.Dia1()
}
