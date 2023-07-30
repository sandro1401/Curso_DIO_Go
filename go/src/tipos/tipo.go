package main

import "fmt"

var x int
var y float64

func main() {
	x = 10
	y = 10.8
	fmt.Printf("x = %d,%T e y = %g,%T", x, x, y, y) // %d para int %g para float %T informa o tipo da variavel

}
