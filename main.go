package main

import "fmt"

const EBULICAO = 373.2

func main() {
	tempC := EBULICAO - 273
	fmt.Printf("A temperatura de ebulição Kelvin é %g e em Celsius é %g", EBULICAO, tempC)
}

// desafio: criar um progama para converter o valor do ponto de ebulição da água de kelvin para graus Celsius.
// Dica : C = K -273