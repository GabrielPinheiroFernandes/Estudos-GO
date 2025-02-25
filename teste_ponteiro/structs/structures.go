package structs

type Locadora struct {
	Idlocadora   int
	Nomelocadora string
}

type Cliente struct {
	IdCliente   int
	NomeCliente string
	carros      []Carro
}

type Carro struct {
	IdCarro   int
	NomeCarro string
}

func (c *Cliente) AddCarro(car Carro) any {
	c.carros = append(c.carros, car)
	return c
}
