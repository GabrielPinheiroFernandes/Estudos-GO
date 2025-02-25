package structs

import "fmt"

// import "fmt"

type Pessoa struct {
	Pai       *Pessoa
	Nome      string
	Sobrenome string
	Idade     int8
	Status    bool
}

func Gerapessoa(nome, sobrenome string, idade int8, status bool) Pessoa {
	// Criando a pessoa
	pai := Pessoa{
		Nome:      "josevaldo",
		Sobrenome: "aaaa",
		Idade:     15,
		Status:    true,
	}
	p := Pessoa{
		Pai:       &pai,
		Nome:      nome,
		Sobrenome: sobrenome,
		Idade:     idade,
		Status:    status,
	}
	p.Pai.Nome = "josevaldinho"

		fmt.Println(pai.Nome)
	// Exibindo os dados da pessoa (para confirmar)
	// fmt.Println("Pessoa criada:", p)
	return p
}

