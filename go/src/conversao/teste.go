package main

import "fmt"

var a int = 20
var b string = "Nome"

func main() {

	var c float64 = float64(a) // Só aceita assim pois c é de um tipo diferente de a,  Tipo(variavel)
	fmt.Printf("O valor de C é:%g e o valor de B é:%s ", c, b)
}
