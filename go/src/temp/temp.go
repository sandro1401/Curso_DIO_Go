package main

import "fmt"

//declaração da variavel const do ponto de ebuliçao da aua em F
const ebulicaoF = 212.0

func main() {

	tempF := ebulicaoF            //variavel do valor da temp em grau F // uso de operador curto : para, dentro de um bloco nao precisar declarar a Variavel com o var
	tempC := (tempF - 32) * 5 / 9 //conversao de F para C

	fmt.Printf("A temp de ebulição da agua em ºF = %g  e a temp de ebulição da agua em ºC = %g ", tempF, tempC)
}
