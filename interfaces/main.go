package main

import "fmt"

type Carro interface {
	correr() string
}

type Fiat struct {
	Modelo string
}

func (Fiat) correr() string {
	return "170 km/h"
}

func print(carro Carro) {
	fmt.Printf("Meu carro corre %s\n", carro.correr())
}

func main() {
	fiat := &Fiat{
		Modelo: "Uno",
	}

	print(fiat)
}
