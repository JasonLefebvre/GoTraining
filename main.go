package main

import (
	"fmt"
)

func main() {

	var x int // Type après l'identifiant
	x = 15    // Variable assignée après la déclaration

	y := 18 // Variable déclarée et assignée | fonctionne que dans une fonction

	fmt.Printf("Mon nombre est: %v!", x)
	fmt.Printf("Mon nombre est: %v!", y)
}
