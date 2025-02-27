package main

import (
	"fmt"
	"time"
)

func Numeros(done chan<- bool) {
	var n int
	for {
		fmt.Print(n)
		if n == 100 {
			done <- true
		}
		n++
		time.Sleep(time.Millisecond * 150)
	}

}
func Letras(done chan<- bool) {
	// Variável que começa com 'A'
	char := 'A'

	// Loop infinito
	for {
		// Imprime a letra atual
		fmt.Print(string(char))

		// Se chegar em 'Z', reinicia para 'A'
		char++
		if char > 'Z' {
			// char = 'A'
			done <- true
		}
		time.Sleep(time.Millisecond * 250)

	}

}

func main() {

	//chanells
	cN := make(chan bool)
	cL := make(chan bool)

	go Numeros(cN)
	go Letras(cL)

	<-cL
	<-cN

	fmt.Println("Fim da execuçao!!!")

	// time.Sleep(time.Second * 5)
}
