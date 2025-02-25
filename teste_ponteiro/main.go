package main

import (
	"fmt"
	"locacaocarro/structs"
)



func main() {

	
	c := structs.Cliente{
		IdCliente:   1,
		NomeCliente: "Roberto",
	}
	
	carro := structs.Carro{
		IdCarro:   1,
		NomeCarro: "mercedes",
	}
	
	carro2 := structs.Carro{
		IdCarro:   2,
		NomeCarro: "Ferrari",
	}

	carro3 := structs.Carro{
		IdCarro:   3,
		NomeCarro: "Lamborghini",
	}
	
	
	c2 := c.AddCarro(carro)
	c2 = c.AddCarro(carro2)
	c2 = c.AddCarro(carro3)
	
	
	fmt.Println(c2)

}
