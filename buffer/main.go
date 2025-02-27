package main

import (
	"fmt"
	"time"
)



func Numeros(c chan<- int){
	for i:=0;i<15;i++{
		c <- i
		fmt.Printf("Gerado no Channel: %d \n",i)
		time.Sleep(time.Millisecond * 0)
	}
	fmt.Println("Fim da escrita.")
	close(c)
}

func main(){
	cL:=make(chan int,7) 	
	go Numeros(cL)

	for v := range cL{
		fmt.Printf("Lido no channel:%d \n",v)
		time.Sleep(time.Millisecond * 770)

	}

}