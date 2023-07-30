package main

import "fmt"

//declaração da variavel const do ponto de ebuliçao da agua em K
const ebulicaoK = 373.0

func main() {

	tempK := ebulicaoK     //variavel do valor da temp em  grau K // uso de operador curto : para, dentro de um bloco nao precisar declarar a Variavel com o var
	tempC := (tempK - 273) //conversao de C para K (c = k - 273)

	fmt.Printf("A temp de ebulição da agua em ºC = %g  e a temp de ebulição da agua em ºK = %g ", tempC, tempK)
}
