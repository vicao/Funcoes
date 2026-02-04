package deferpkg

//defer : escalona nossas funções.

import "fmt"

func Dia1(){
	fmt.Println("Domingo")
}

func Dia2(){
	fmt.Println("Segunda-feira")
}
/*
func main(){
	defer Dia2() //adiar a execução da função dia2
	Dia1()
}*/
